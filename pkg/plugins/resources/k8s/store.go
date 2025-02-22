/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package k8s

import (
	"context"
	"regexp"
	"strings"
	"time"
)

import (
	"github.com/pkg/errors"

	"golang.org/x/exp/maps"

	kube_apierrs "k8s.io/apimachinery/pkg/api/errors"
	kube_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	kube_runtime "k8s.io/apimachinery/pkg/runtime"

	kube_client "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

import (
	"github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	"github.com/apache/dubbo-kubernetes/pkg/core/logger"
	core_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/registry"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/store"
	k8s_common "github.com/apache/dubbo-kubernetes/pkg/plugins/common/k8s"
	k8s_model "github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/model"
	k8s_registry "github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/registry"
	util_store "github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/util"
	util_k8s "github.com/apache/dubbo-kubernetes/pkg/util/k8s"
)

func typeIsUnregistered(err error) bool {
	var typeErr *k8s_registry.UnknownTypeError
	return errors.As(err, &typeErr)
}

var _ store.ResourceStore = &KubernetesStore{}

type KubernetesStore struct {
	Client    kube_client.Client
	Converter k8s_common.Converter
	Scheme    *kube_runtime.Scheme
}

func NewStore(client kube_client.Client, scheme *kube_runtime.Scheme, converter k8s_common.Converter) (store.ResourceStore, error) {
	return &KubernetesStore{
		Client:    client,
		Converter: converter,
		Scheme:    scheme,
	}, nil
}

func (s *KubernetesStore) Create(ctx context.Context, r core_model.Resource, fs ...store.CreateOptionsFunc) error {
	opts := store.NewCreateOptions(fs...)

	obj, err := s.Converter.ToKubernetesObject(r)
	if err != nil {
		if typeIsUnregistered(err) {
			return errors.Errorf("cannot create instance of unregistered type %q", r.Descriptor().Name)
		}
		return errors.Wrap(err, "failed to convert core model into k8s counterpart")
	}
	name, namespace, err := k8sNameNamespace(opts.Name, obj.Scope())
	if err != nil {
		return err
	}

	// for k8s metadata.name require
	if name, err = EncodeK8sResName(name); err != nil {
		return err
	}

	obj.GetObjectMeta().SetLabels(opts.Labels)
	obj.SetMesh(opts.Mesh)
	obj.GetObjectMeta().SetName(name)
	obj.GetObjectMeta().SetNamespace(namespace)

	if opts.Owner != nil {
		k8sOwner, err := s.Converter.ToKubernetesObject(opts.Owner)
		if err != nil {
			return errors.Wrap(err, "failed to convert core model into k8s counterpart")
		}
		if err := controllerutil.SetOwnerReference(k8sOwner, obj, s.Scheme); err != nil {
			return errors.Wrap(err, "failed to set owner reference for object")
		}
	}

	if err := s.Client.Create(ctx, obj); err != nil {
		if kube_apierrs.IsAlreadyExists(err) {
			// 如果资源已经存在了就直接返回空即可
			logger.Sugar().Warn("资源已经存在了")
			return nil
		}
		return errors.Wrap(err, "failed to create k8s resource")
	}

	if name, err = DecodeK8sResName(obj.GetName()); err != nil {
		return err
	} else {
		obj.SetName(name)
	}

	err = s.Converter.ToCoreResource(obj, r)
	if err != nil {
		return errors.Wrap(err, "failed to convert k8s model into core counterpart")
	}
	return nil
}

func (s *KubernetesStore) Update(ctx context.Context, r core_model.Resource, fs ...store.UpdateOptionsFunc) error {
	opts := store.NewUpdateOptions(fs...)

	obj, err := s.Converter.ToKubernetesObject(r)
	if err != nil {
		if typeIsUnregistered(err) {
			return errors.Errorf("cannot update instance of unregistered type %q", r.Descriptor().Name)
		}
		return errors.Wrapf(err, "failed to convert core model of type %s into k8s counterpart", r.Descriptor().Name)
	}

	obj.GetObjectMeta().SetLabels(opts.Labels)

	name, namespace, err := k8sNameNamespace(opts.Name, obj.Scope())
	if err != nil {
		return err
	}
	// for k8s metadata.name require
	if name, err = EncodeK8sResName(name); err != nil {
		return err
	}

	if r.GetMeta() == nil {
		if err = s.Client.Get(ctx, kube_client.ObjectKey{Namespace: namespace, Name: name}, obj); err != nil {
			if kube_apierrs.IsNotFound(err) {
				return store.ErrorResourceNotFound(r.Descriptor().Name, opts.Name, opts.Mesh)
			}
			return errors.Wrap(err, "failed to get k8s resource")
		}
		obj.SetSpec(r.GetSpec())
	} else {
		obj.SetName(name)
		obj.SetMesh(r.GetMeta().GetMesh())
	}

	if err = s.Client.Update(ctx, obj); err != nil {
		if kube_apierrs.IsConflict(err) {
			return store.ErrorResourceConflict(r.Descriptor().Name, r.GetMeta().GetName(), r.GetMeta().GetMesh())
		}
		return errors.Wrap(err, "failed to update k8s resource")
	}

	if name, err = DecodeK8sResName(obj.GetName()); err != nil {
		return err
	} else {
		obj.SetName(name)
	}

	err = s.Converter.ToCoreResource(obj, r)
	if err != nil {
		return errors.Wrap(err, "failed to convert k8s model into core counterpart")
	}
	return nil
}

func (s *KubernetesStore) Delete(ctx context.Context, r core_model.Resource, fs ...store.DeleteOptionsFunc) error {
	opts := store.NewDeleteOptions(fs...)

	// get object and validate mesh
	if err := s.Get(ctx, r, store.GetByKey(opts.Name, opts.Mesh)); err != nil {
		return err
	}

	obj, err := s.Converter.ToKubernetesObject(r)
	if err != nil {
		// Unregistered types can't exist in the first place, so deletion would automatically succeed.
		if typeIsUnregistered(err) {
			return nil
		}
		return errors.Wrapf(err, "failed to convert core model of type %s into k8s counterpart", r.Descriptor().Name)
	}

	name, namespace, err := k8sNameNamespace(opts.Name, obj.Scope())
	if err != nil {
		return err
	}

	// for k8s metadata.name require
	if name, err = EncodeK8sResName(name); err != nil {
		return err
	}

	obj.GetObjectMeta().SetName(name)
	obj.GetObjectMeta().SetNamespace(namespace)

	if name, err = DecodeK8sResName(obj.GetName()); err != nil {
		return err
	} else {
		obj.SetName(name)
	}

	if err := s.Client.Delete(ctx, obj); err != nil {
		if kube_apierrs.IsNotFound(err) {
			return nil
		}
		return errors.Wrap(err, "failed to delete k8s resource")
	}
	return nil
}

func (s *KubernetesStore) Get(ctx context.Context, r core_model.Resource, fs ...store.GetOptionsFunc) error {
	opts := store.NewGetOptions(fs...)

	obj, err := s.Converter.ToKubernetesObject(r)
	if err != nil {
		if typeIsUnregistered(err) {
			return store.ErrorResourceNotFound(r.Descriptor().Name, opts.Name, opts.Mesh)
		}
		return errors.Wrapf(err, "failed to convert core model of type %s into k8s counterpart", r.Descriptor().Name)
	}

	name, namespace, err := k8sNameNamespace(opts.Name, obj.Scope())
	if err != nil {
		return err
	}

	// for k8s metadata.name require
	if name, err = EncodeK8sResName(name); err != nil {
		return err
	}

	if err := s.Client.Get(ctx, kube_client.ObjectKey{Namespace: namespace, Name: name}, obj); err != nil {
		if kube_apierrs.IsNotFound(err) {
			return store.ErrorResourceNotFound(r.Descriptor().Name, opts.Name, opts.Mesh)
		}
		return errors.Wrap(err, "failed to get k8s resource")
	}

	if name, err = DecodeK8sResName(obj.GetName()); err != nil {
		return err
	} else {
		obj.SetName(name)
	}

	if err := s.Converter.ToCoreResource(obj, r); err != nil {
		return errors.Wrap(err, "failed to convert k8s model into core counterpart")
	}
	if opts.Version != "" && r.GetMeta().GetVersion() != opts.Version {
		return store.ErrorResourceConflict(r.Descriptor().Name, opts.Name, opts.Mesh)
	}
	if r.GetMeta().GetMesh() != opts.Mesh {
		return store.ErrorResourceNotFound(r.Descriptor().Name, opts.Name, opts.Mesh)
	}
	return nil
}

func (s *KubernetesStore) List(ctx context.Context, rs core_model.ResourceList, fs ...store.ListOptionsFunc) error {
	opts := store.NewListOptions(fs...)

	obj, err := s.Converter.ToKubernetesList(rs)
	if err != nil {
		if typeIsUnregistered(err) {
			return nil
		}
		return errors.Wrapf(err, "failed to convert core list model of type %s into k8s counterpart", rs.GetItemType())
	}

	if err := s.Client.List(ctx, obj); err != nil {
		return errors.Wrap(err, "failed to list k8s resources")
	}

	predicate := func(r core_model.Resource) bool {
		if opts.Mesh != "" && r.GetMeta().GetMesh() != opts.Mesh {
			return false
		}
		if opts.NameContains != "" && !strings.Contains(r.GetMeta().GetName(), opts.NameContains) {
			return false
		}
		return true
	}

	fullList, err := registry.Global().NewList(rs.GetItemType())
	if err != nil {
		return err
	}

	for _, object := range obj.GetItems() {
		name, _ := DecodeK8sResName(object.GetName())
		object.SetName(name)
	}

	if err := s.Converter.ToCoreList(obj, fullList, predicate); err != nil {
		return errors.Wrap(err, "failed to convert k8s model into core counterpart")
	}

	for _, item := range fullList.GetItems() {
		_ = rs.AddItem(item)
	}

	rs.GetPagination().SetTotal(uint32(len(fullList.GetItems())))
	return nil
}

func k8sNameNamespace(coreName string, scope k8s_model.Scope) (string, string, error) {
	if coreName == "" {
		return "", "", store.PreconditionFormatError("name can't be empty")
	}
	switch scope {
	case k8s_model.ScopeCluster:
		return coreName, "", nil
	case k8s_model.ScopeNamespace:
		name, ns, err := util_k8s.CoreNameToK8sName(coreName)
		if err != nil {
			return "", "", store.PreconditionFormatError(err.Error())
		}
		return name, ns, nil
	default:
		return "", "", errors.Errorf("unknown scope %s", scope)
	}
}

var _ core_model.ResourceMeta = &KubernetesMetaAdapter{}

type KubernetesMetaAdapter struct {
	kube_meta.ObjectMeta
	Mesh string
}

func (m *KubernetesMetaAdapter) GetName() string {
	if m.Namespace == "" { // it's cluster scoped object
		return m.ObjectMeta.Name
	}
	return util_k8s.K8sNamespacedNameToCoreName(m.ObjectMeta.Name, m.ObjectMeta.Namespace)
}

func (m *KubernetesMetaAdapter) GetNameExtensions() core_model.ResourceNameExtensions {
	return k8s_common.ResourceNameExtensions(m.ObjectMeta.Namespace, m.ObjectMeta.Name)
}

func (m *KubernetesMetaAdapter) GetVersion() string {
	return m.ObjectMeta.GetResourceVersion()
}

func (m *KubernetesMetaAdapter) GetMesh() string {
	return m.Mesh
}

func (m *KubernetesMetaAdapter) GetCreationTime() time.Time {
	return m.GetObjectMeta().GetCreationTimestamp().Time
}

func (m *KubernetesMetaAdapter) GetModificationTime() time.Time {
	return m.GetObjectMeta().GetCreationTimestamp().Time
}

func (m *KubernetesMetaAdapter) GetLabels() map[string]string {
	labels := maps.Clone(m.GetObjectMeta().GetLabels())
	if labels == nil {
		labels = map[string]string{}
	}
	if _, ok := labels[v1alpha1.DisplayName]; !ok {
		labels[v1alpha1.DisplayName] = m.GetObjectMeta().GetName()
	}
	if m.Namespace != "" {
		labels[v1alpha1.KubeNamespaceTag] = m.Namespace
	}
	return labels
}

type KubeFactory interface {
	NewObject(r core_model.Resource) (k8s_model.KubernetesObject, error)
	NewList(rl core_model.ResourceList) (k8s_model.KubernetesList, error)
}

var _ KubeFactory = &SimpleKubeFactory{}

type SimpleKubeFactory struct {
	KubeTypes k8s_registry.TypeRegistry
}

func (f *SimpleKubeFactory) NewObject(r core_model.Resource) (k8s_model.KubernetesObject, error) {
	return f.KubeTypes.NewObject(r.GetSpec())
}

func (f *SimpleKubeFactory) NewList(rl core_model.ResourceList) (k8s_model.KubernetesList, error) {
	return f.KubeTypes.NewList(rl.NewItem().GetSpec())
}

// Define the regex pattern for a valid RFC 1123 subdomain
// for k8s [metadata.name] require.
var k8sNameCheck = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)

var (
	_Aa = []byte{'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i', 'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r', 'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z', ':': '0', '_': '1', '@': '2'}
	_aA = []byte{'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D', 'e': 'E', 'f': 'F', 'g': 'G', 'h': 'H', 'i': 'I', 'j': 'J', 'k': 'K', 'l': 'L', 'm': 'M', 'n': 'N', 'o': 'O', 'p': 'P', 'q': 'Q', 'r': 'R', 's': 'S', 't': 'T', 'u': 'U', 'v': 'V', 'w': 'W', 'x': 'X', 'y': 'Y', 'z': 'Z', '0': ':', '1': '_', '2': '@'}
)

func EncodeK8sResName(name string) (string, error) {
	// if match success, return
	if k8sNameCheck.MatchString(name) {
		return name, nil
	}
	bs := util_store.NewBitset()
	bf := []byte(name)
	for index, char := range bf {
		if char < uint8(len(_Aa)) && _Aa[char] != 0 {
			bf[index] = _Aa[char]
			bs.Set(int32(index))
		}
	}
	// return transString.base32toTransIndex.re
	return string(bf) + "." + bs.Encode() + ".re", nil
}

func DecodeK8sResName(name string) (string, error) {
	// check is encoded name
	raw, ok := strings.CutSuffix(name, ".re")
	if !ok {
		return name, nil
	}
	i := strings.LastIndex(raw, ".")
	if i == -1 {
		return "", store.PreconditionFormatError("DecodeK8sResName fail")
	}
	bs := util_store.NewBitset()
	err := bs.Decode(raw[i+1:])
	if err != nil {
		return "", store.PreconditionFormatError(err.Error())
	}
	bt := []byte(raw[:i])
	bs.Range(func(idx int, val bool) bool {
		if idx >= len(bt) {
			return false
		}
		if val && len(_aA) > int(bt[idx]) && _aA[bt[idx]] != 0 {
			bt[idx] = _aA[bt[idx]]
		}
		return true
	})
	return string(bt), nil
}
