import{d as G,k as w,b as x,r as N,f as u,c as _,t as d,h as m,P as C,K as R,e as h,o as n,H as T,I as q,y as f,z as y,v as b,F as V,x as A,J as P,m as D,_ as E}from"./index-7CzEMIL1.js";import{s as $}from"./service-2gqozivS.js";import{S as B,a as O}from"./SearchUtil-Od3Elmwo.js";import{q as p}from"./promQuery-8dWPhEaC.js";import"./request-SYZV5A5R.js";const Q={class:"__container_services_index"},M=["onClick"],F=G({__name:"search",setup(H){w(a=>({"51ee2c8e":m(C)}));const g=x(),S=[{title:"service",key:"service",dataIndex:"serviceName",sorter:!0,width:"30%"},{title:"versionGroup",key:"versionGroup",dataIndex:"versionGroupSelect",width:"25%"},{title:"avgQPS",key:"avgQPS",dataIndex:"avgQPS",sorter:!0,width:"15%"},{title:"avgRT",key:"avgRT",dataIndex:"avgRT",sorter:!0,width:"15%"},{title:"requestTotal",key:"requestTotal",dataIndex:"requestTotal",sorter:!0,width:"15%"}],v=a=>a.map(e=>(e.versionGroupSelect={},e.versionGroupSelect.versionGroupArr=e.versionGroups.map(t=>t.versionGroup=(t.version?"version: "+t.version+", ":"")+(t.group?"group: "+t.group:"")||"无"),e.versionGroupSelect.versionGroupValue=e.versionGroupSelect.versionGroupArr[0],e));function k(a){return $(a).then(async e=>{var i;let t=(i=e==null?void 0:e.data)==null?void 0:i.list;try{for(let r of t){let o=await p(`sum (dubbo_provider_qps_total{interface='${r.serviceName}'}) by (interface)`),s=await p(`avg(dubbo_consumer_rt_avg_milliseconds_aggregate{interface="${r.serviceName}",method=~"$method"}>0)`),l=await p(`sum (increase(dubbo_provider_requests_total{interface="${r.serviceName}"}[1m]))`);r.avgQPS=o,r.avgRT=s,r.requestTotal=l}}catch(r){console.error(r)}return e})}const c=N(new B([{label:"serviceName",param:"keywords",placeholder:"typeAppName",style:{width:"200px"}}],k,S,void 0,void 0,v));c.onSearch(v);const I=a=>{g.push({name:"distribution",params:{pathId:a}})};return R(D.SEARCH_DOMAIN,c),(a,e)=>{const t=h("a-select-option"),i=h("a-select");return n(),u("div",Q,[_(O,{"search-domain":c},{bodyCell:d(({column:r,text:o})=>[r.dataIndex==="serviceName"?(n(),u("span",{key:0,class:"service-link",onClick:s=>I(o)},[T("b",null,[_(m(q),{style:{"margin-bottom":"-2px"},icon:"material-symbols:attach-file-rounded"}),f(" "+y(o),1)])],8,M)):r.dataIndex==="versionGroupSelect"?(n(),b(i,{key:1,value:o.versionGroupValue,bordered:!1,style:{width:"80%"}},{default:d(()=>[(n(!0),u(V,null,A(o.versionGroupArr,(s,l)=>(n(),b(t,{value:s,key:l},{default:d(()=>[f(y(s),1)]),_:2},1032,["value"]))),128))]),_:2},1032,["value"])):P("",!0)]),_:1},8,["search-domain"])])}}}),j=E(F,[["__scopeId","data-v-49b763bb"]]);export{j as default};
