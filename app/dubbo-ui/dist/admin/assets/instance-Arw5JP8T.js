import{d as M,k as Y,u as B,b as $,r as w,O as S,f as h,c as u,t as a,s as F,e as f,h as o,P as g,K as H,o as s,F as k,x as D,v as c,I as v,y as i,z as n,H as j,J as d,Y as q,Z as z,m as J,_ as K}from"./index-QTFC1BX7.js";import{S as U,a as G}from"./SearchUtil-gzPkvecP.js";import{a as X}from"./app-AcX83ZGL.js";import{f as Z}from"./DateUtil-9OgzTl0-.js";import{q as T}from"./promQuery-PGswEixs.js";import{b as Q}from"./ByteUtil-YdHlSEeW.js";import"./request-bLCExP-g.js";const W={class:"__container_app_instance"},ee={class:"statistic-icon-big"},te=M({__name:"instance",setup(ae){var C;Y(p=>({"2af02389":o(g)+"22","35e57e72":o(g)}));const N=B(),E=$();let O=w({info:{},report:{}}),R=(C=N.params)==null?void 0:C.pathId;S(async()=>{});const A=[{title:"instanceDomain.ip",dataIndex:"ip",key:"ip",sorter:!0,width:150,fixed:"left"},{title:"instanceDomain.name",dataIndex:"name",key:"name",sorter:!0,width:180},{title:"instanceDomain.deployState",dataIndex:"deployState",key:"deployState",sorter:!0,width:150},{title:"instanceDomain.deployCluster",dataIndex:"deployClusters",key:"deployClusters",sorter:!0,width:180},{title:"instanceDomain.registerState",dataIndex:"registerState",key:"registerState",sorter:!0,width:150},{title:"instanceDomain.registerClusters",dataIndex:"registerCluster",key:"registerCluster",sorter:!0,width:200},{title:"instanceDomain.cpu",dataIndex:"cpu",key:"cpu",sorter:!0,width:120},{title:"instanceDomain.memory",dataIndex:"memory",key:"memory",sorter:!0,width:120},{title:"instanceDomain.startTime",dataIndex:"startTime",key:"startTime",sorter:!0,width:150},{title:"instanceDomain.registerTime",dataIndex:"registerTime",key:"registerTime",sorter:!0,width:150},{title:"instanceDomain.labels",dataIndex:"labels",key:"labels",sorter:!0,fixed:"left",width:800}];function V(p){return X(p).then(async _=>{var m;let b=(m=_==null?void 0:_.data)==null?void 0:m.list;try{for(let l of b){let I=l.ip.split(":")[0],r=await T(`sum(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{container!=""}) by (pod) * on (pod) group_left(pod_ip)
        kube_pod_info{pod_ip="${I}"}`),e=await T(`sum(container_memory_working_set_bytes{container!=""}) by (pod)
* on (pod) group_left(pod_ip)
kube_pod_info{pod_ip="${I}"}`);l.cpu=F.isNumber(r)?r.toFixed(3)+"u":r,l.memory=Q(e)}}catch(l){console.error(l)}return _})}const y=w(new U([{label:"",param:"type",defaultValue:1,dict:[{label:"ip",value:1},{label:"name",value:2},{label:"label",value:3}],style:{width:"100px"}},{label:"",param:"search",style:{width:"300px"}},{label:"",param:"appName",defaultValue:R,dict:[],dictType:"APPLICATION_NAME"}],V,A,{pageSize:10},!0));S(()=>{y.tableStyle={scrollX:"100",scrollY:"calc(100vh - 400px)"},y.onSearch()});const P=p=>{E.push("/resources/services/detail/"+p)};return H(J.SEARCH_DOMAIN,y),(p,_)=>{const b=f("a-statistic"),m=f("a-flex"),l=f("a-card"),I=f("a-button"),r=f("a-tag");return s(),h("div",W,[u(m,{wrap:"wrap",gap:"small",vertical:!1,justify:"space-around",align:"left"},{default:a(()=>[(s(!0),h(k,null,D(o(O).report,(e,t)=>(s(),c(l,{class:"statistic-card"},{default:a(()=>[u(m,{gap:"middle",vertical:!1,justify:"space-between",align:"center"},{default:a(()=>[u(b,{value:e.value,class:"statistic"},{prefix:a(()=>[u(o(v),{class:"statistic-icon",icon:"solar:target-line-duotone"})]),title:a(()=>[i(n(p.$t(t.toString())),1)]),_:2},1032,["value"]),j("div",ee,[u(o(v),{icon:e.icon},null,8,["icon"])])]),_:2},1024)]),_:2},1024))),256))]),_:1}),u(G,{"search-domain":y},{bodyCell:a(({column:e,text:t})=>[e.dataIndex==="name"?(s(),c(I,{key:0,type:"link",onClick:x=>P(t)},{default:a(()=>[i(n(t),1)]),_:2},1032,["onClick"])):d("",!0),e.dataIndex==="deployState"?(s(),c(r,{key:1,color:o(q)[t.toUpperCase()]},{default:a(()=>[i(n(t),1)]),_:2},1032,["color"])):d("",!0),e.dataIndex==="deployClusters"?(s(),c(r,{key:2},{default:a(()=>[i(n(t),1)]),_:2},1024)):d("",!0),e.dataIndex==="registerState"?(s(),c(r,{key:3,color:o(z)[t.toUpperCase()]},{default:a(()=>[i(n(t),1)]),_:2},1032,["color"])):d("",!0),e.dataIndex==="registerCluster"?(s(),c(r,{key:4},{default:a(()=>[i(n(t),1)]),_:2},1024)):d("",!0),e.dataIndex==="labels"?(s(!0),h(k,{key:5},D(t,(x,L)=>(s(),c(r,{color:o(g)},{default:a(()=>[i(n(L)+" : "+n(x),1)]),_:2},1032,["color"]))),256)):d("",!0),e.dataIndex==="registerTime"?(s(),h(k,{key:6},[i(n(o(Z)(t)),1)],64)):d("",!0)]),_:1},8,["search-domain"])])}}}),de=K(te,[["__scopeId","data-v-2b2cbf59"]]);export{de as default};
