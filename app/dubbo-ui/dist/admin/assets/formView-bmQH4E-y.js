import{d as T,k as j,u as x,g as A,a as E,r as F,O as L,f as d,c as t,t as a,F as C,x as h,e as o,h as _,P as M,o as c,H as w,y as l,z as n,V as v,v as D,W as z,_ as G}from"./index-QTFC1BX7.js";import{u as H}from"./index-yN-dmHDD.js";import{c as Y}from"./traffic-84ZA5mgZ.js";import"./request-bLCExP-g.js";const q={class:"__container_app_detail"},J=T({__name:"formView",setup(K){j(e=>({"4cf7a4b8":_(M)}));const $=x(),{appContext:{config:{globalProperties:k}}}=A(),R=H().toClipboard;function f(e){z.success(k.$t("messageDomain.success.copy")),R(e)}const m=E(()=>{const e=s.key.split(":");return e[0]?e[0]:""}),s=F({configVersion:"v3.0",scope:"application",key:"shop-user",enabled:!0,runtime:!0,tags:[{name:"gray",match:[{key:"version",value:{exact:"v1"}}]}]}),V=async()=>{var r;const e=await Y((r=$.params)==null?void 0:r.ruleName);e.code===200&&Object.assign(s,e.data||{})};return L(()=>{V()}),(e,r)=>{const i=o("a-descriptions-item"),u=o("a-typography-paragraph"),N=o("a-descriptions"),g=o("a-card"),O=o("a-flex"),S=o("a-typography-text"),b=o("a-typography-title"),y=o("a-space"),W=o("a-tag");return c(),d("div",q,[t(O,null,{default:a(()=>[t(g,{class:"_detail"},{default:a(()=>[t(N,{column:2,layout:"vertical",title:""},{default:a(()=>[t(i,{label:e.$t("flowControlDomain.ruleName"),labelStyle:{fontWeight:"bold"}},{default:a(()=>[w("p",{onClick:r[0]||(r[0]=p=>f(s.key)),class:"description-item-content with-card"},[l(n(s.key)+" ",1),t(_(v))])]),_:1},8,["label"]),t(i,{label:e.$t("flowControlDomain.ruleGranularity"),labelStyle:{fontWeight:"bold"}},{default:a(()=>[t(u,null,{default:a(()=>[l(n(s.scope),1)]),_:1})]),_:1},8,["label"]),t(i,{label:e.$t("flowControlDomain.actionObject"),labelStyle:{fontWeight:"bold"}},{default:a(()=>[w("p",{onClick:r[1]||(r[1]=p=>f(m.value)),class:"description-item-content with-card"},[l(n(m.value)+" ",1),t(_(v))])]),_:1},8,["label"]),t(i,{label:e.$t("flowControlDomain.enabledState"),labelStyle:{fontWeight:"bold"}},{default:a(()=>[t(u,null,{default:a(()=>[l(n(s.enabled?e.$t("flowControlDomain.enabled"):e.$t("flowControlDomain.disabled")),1)]),_:1})]),_:1},8,["label"]),t(i,{label:e.$t("flowControlDomain.runTimeEffective"),labelStyle:{fontWeight:"bold"}},{default:a(()=>[t(u,null,{default:a(()=>[l(n(s.runtime?e.$t("flowControlDomain.opened"):e.$t("flowControlDomain.closed")),1)]),_:1})]),_:1},8,["label"])]),_:1})]),_:1})]),_:1}),(c(!0),d(C,null,h(s.tags,(p,I)=>(c(),D(g,{title:`标签【${I+1}】`,style:{"margin-top":"10px"},class:"_detail"},{default:a(()=>[t(y,{align:"center"},{default:a(()=>[t(b,{level:5},{default:a(()=>[l(n(e.$t("flowControlDomain.labelName"))+": ",1),t(S,{class:"labelName"},{default:a(()=>[l(n(p.name),1)]),_:2},1024)]),_:2},1024)]),_:2},1024),t(y,{align:"start",style:{width:"100%"}},{default:a(()=>[t(b,{level:5},{default:a(()=>[l(n(e.$t("flowControlDomain.actuatingRange"))+": ",1)]),_:1}),(c(!0),d(C,null,h(p.match,(B,P)=>(c(),D(W,{key:P,color:"#2db7f5"},{default:a(()=>[l(n(B.key),1)]),_:2},1024))),128))]),_:2},1024)]),_:2},1032,["title"]))),256))])}}}),ee=G(J,[["__scopeId","data-v-6a9ecf97"]]);export{ee as default};
