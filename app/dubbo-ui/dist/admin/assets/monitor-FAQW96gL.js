import{g as h}from"./instance-vMiDnq-x.js";import{d as v,n as c,O as b,f as _,H as i,c as t,t as n,h as o,e as f,o as l,y as d,J as g,_ as k}from"./index-tKheeObz.js";import"./request-ovKoEMRQ.js";const w={class:"__container_tabDemo3"},C={class:"option"},x={class:"__container_iframe_container"},I=["src"],N=v({__name:"monitor",setup(V){let a=c(""),e=c(!0);b(async()=>{let s=await h({});a.value=s.data});function m(){e.value=!1,setTimeout(()=>{e.value=!0},200)}function u(){window.open(a.value,"_blank")}return(s,y)=>{const r=f("a-button"),p=f("a-spin");return l(),_("div",w,[i("div",C,[t(r,{class:"btn",onClick:m},{default:n(()=>[d(" refresh ")]),_:1}),t(r,{class:"btn",onClick:u},{default:n(()=>[d(" grafana ")]),_:1})]),t(p,{class:"spin",spinning:!o(e)},{default:n(()=>[i("div",x,[o(e)?(l(),_("iframe",{key:0,id:"grafanaIframe",src:o(a),frameborder:"0"},null,8,I)):g("",!0)])]),_:1},8,["spinning"])])}}}),D=k(N,[["__scopeId","data-v-b18fe3f4"]]);export{D as default};