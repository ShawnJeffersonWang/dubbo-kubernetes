import{d,u as f,n as c,O as p,f as _,c as h,t as v,h as r,e as g,o as y,H as i,U as w,X as S}from"./index-7CzEMIL1.js";import{c as q}from"./app-ZMUzdKgx.js";import"./request-SYZV5A5R.js";const I={class:"__container_tabDemo3"},D={class:"__container_iframe_container"},b=["src"],N=d({__name:"monitor",setup(k){var s;let l=(s=f().params)==null?void 0:s.pathId,n=c(""),t=c(!0);p(async()=>{var a;let e=await q({});t.value=!1,n.value=`${window.location.origin}/grafana/d/${(a=e.data)==null?void 0:a.baseURL.split("/d/")[1].split("?")[0]}?var-application=${l}&kiosk=tv`,console.log(n)});function o(e){try{e()}catch(a){console.log(a)}}function m(){console.log("The iframe has been loaded."),setTimeout(()=>{try{let e=document.querySelector("#grafanaIframe").contentDocument;o(()=>{e.querySelector("header").remove()}),o(()=>{e.querySelector("[data-testid*='controls']").remove()}),setTimeout(()=>{o(()=>{e.querySelector("[data-testid*='navigation mega-menu']").remove()}),o(()=>{for(let a of e.querySelectorAll("[data-testid*='Panel menu']"))console.log(a),a.remove()})},1e3)}catch{}t.value=!0},1e3)}return(e,a)=>{const u=g("a-spin");return y(),_("div",I,[h(u,{class:"spin",spinning:!r(t)},{default:v(()=>[i("div",D,[w(i("iframe",{onload:m,id:"grafanaIframe",src:r(n),frameborder:"0"},null,8,b),[[S,r(t)]])])]),_:1},8,["spinning"])])}}});export{N as default};
