import{C as L}from"./ConfigPage-mTUr_FdD.js";import{d as N,r as V,O as $,f as h,c as a,t as e,h as v,e as n,o as d,F as C,x as D,v as f,y as s,z as m,H as F,I as k,_ as W}from"./index-7CzEMIL1.js";const E={class:"__container_app_config"},H={style:{float:"right"}},M={style:{float:"right"}},O=N({__name:"config",setup(S){let z=V({list:[{title:"applicationDomain.operatorLog",key:"log",form:{logFlag:!1},submit:c=>new Promise(l=>{setTimeout(()=>{l(1)},1e3)}),reset(c){c.logFlag=!1}},{title:"applicationDomain.flowWeight",key:"flow",ext:{title:"添加权重配置",fun(){}},form:{rules:[{weight:"100",scope:{label:"key1",condition:"=",value:"value1"}}]},submit(c){return new Promise(l=>{setTimeout(()=>{l(1)},1e3)})}},{title:"applicationDomain.gray",key:"gray",ext:{title:"添加灰度环境",fun(){}},form:{rules:[{name:"100",scope:{label:"key1",condition:"=",value:"value1"}}]},submit(c){return new Promise(l=>{setTimeout(()=>{l(1)},1e3)})}}],current:[0]});return $(()=>{console.log(333)}),(c,l)=>{const B=n("a-switch"),i=n("a-form-item"),y=n("a-button"),g=n("a-space"),P=n("a-input-number"),r=n("a-select-option"),w=n("a-select"),b=n("a-input"),U=n("a-table"),x=n("a-card");return d(),h("div",E,[a(L,{options:v(z)},{form_log:e(({current:_})=>[a(i,{label:c.$t("applicationDomain.operatorLog"),name:"logFlag"},{default:e(()=>[a(B,{checked:_.form.logFlag,"onUpdate:checked":t=>_.form.logFlag=t},null,8,["checked","onUpdate:checked"])]),_:2},1032,["label"])]),form_flow:e(({current:_})=>[a(g,null,{default:e(()=>[(d(!0),h(C,null,D(_.form.rules,(t,u)=>(d(),f(x,null,{title:e(()=>[s(m(c.$t("applicationDomain.flowWeight"))+" "+m(u+1)+" ",1),F("div",H,[a(g,null,{default:e(()=>[a(y,{type:"dashed",onClick:l[0]||(l[0]=()=>{})},{default:e(()=>[a(v(k),{style:{"font-size":"20px"},icon:"material-symbols-light:contract-edit"})]),_:1}),a(y,{danger:"",type:"dashed",onClick:l[1]||(l[1]=()=>{})},{default:e(()=>[a(v(k),{style:{"font-size":"20px"},icon:"fluent:delete-12-filled"})]),_:1})]),_:1})])]),default:e(()=>[a(i,{name:"rules["+u+"].weight",label:"权重"},{default:e(()=>[a(P,{value:t.weight,"onUpdate:value":o=>t.weight=o},null,8,["value","onUpdate:value"])]),_:2},1032,["name"]),a(i,{label:"作用范围"},{default:e(()=>[a(U,{style:{width:"40vw"},pagination:!1,columns:[{key:"label",title:"label"},{key:"condition",title:"condition"},{key:"value",title:"value"}],"data-source":[t.scope]},{bodyCell:e(({column:o,record:T,index:I})=>[o.key==="condition"?(d(),f(i,{key:0,name:"rules["+u+"].scope.condition",label:""},{default:e(()=>[a(w,{value:t.scope.condition,"onUpdate:value":p=>t.scope.condition=p},{default:e(()=>[a(r,{value:"="},{default:e(()=>[s("=")]),_:1}),a(r,{value:"!="},{default:e(()=>[s("!=")]),_:1}),a(r,{value:">"},{default:e(()=>[s(">")]),_:1}),a(r,{value:"<"},{default:e(()=>[s(m("<"))]),_:1})]),_:2},1032,["value","onUpdate:value"])]),_:2},1032,["name"])):(d(),f(i,{key:1,name:"rules["+u+"].scope.condition."+o.key,label:""},{default:e(()=>[a(b,{value:t.scope[o.key],"onUpdate:value":p=>t.scope[o.key]=p},null,8,["value","onUpdate:value"])]),_:2},1032,["name"]))]),_:2},1032,["data-source"])]),_:2},1024)]),_:2},1024))),256))]),_:2},1024)]),form_gray:e(({current:_})=>[a(g,null,{default:e(()=>[(d(!0),h(C,null,D(_.form.rules,(t,u)=>(d(),f(x,null,{title:e(()=>[s(m(c.$t("applicationDomain.gray"))+" "+m(u+1)+" ",1),F("div",M,[a(g,null,{default:e(()=>[a(y,{type:"dashed",onClick:l[2]||(l[2]=()=>{})},{default:e(()=>[a(v(k),{style:{"font-size":"20px"},icon:"material-symbols-light:contract-edit"})]),_:1}),a(y,{danger:"",type:"dashed",onClick:l[3]||(l[3]=()=>{})},{default:e(()=>[a(v(k),{style:{"font-size":"20px"},icon:"fluent:delete-12-filled"})]),_:1})]),_:1})])]),default:e(()=>[a(i,{name:"rules["+u+"].name",label:"环境名称"},{default:e(()=>[a(b,{value:t.name,"onUpdate:value":o=>t.name=o},null,8,["value","onUpdate:value"])]),_:2},1032,["name"]),a(i,{label:"作用范围"},{default:e(()=>[a(U,{style:{width:"40vw"},pagination:!1,columns:[{key:"label",title:"label"},{key:"condition",title:"condition"},{key:"value",title:"value"}],"data-source":[t.scope]},{bodyCell:e(({column:o,record:T,index:I})=>[o.key==="condition"?(d(),f(i,{key:0,name:"rules["+u+"].scope.condition",label:""},{default:e(()=>[a(w,{value:t.scope.condition,"onUpdate:value":p=>t.scope.condition=p},{default:e(()=>[a(r,{value:"="},{default:e(()=>[s("=")]),_:1}),a(r,{value:"!="},{default:e(()=>[s("!=")]),_:1}),a(r,{value:">"},{default:e(()=>[s(">")]),_:1}),a(r,{value:"<"},{default:e(()=>[s(m("<"))]),_:1})]),_:2},1032,["value","onUpdate:value"])]),_:2},1032,["name"])):(d(),f(i,{key:1,name:"rules["+u+"].scope.condition."+o.key,label:""},{default:e(()=>[a(b,{value:t.scope[o.key],"onUpdate:value":p=>t.scope[o.key]=p},null,8,["value","onUpdate:value"])]),_:2},1032,["name"]))]),_:2},1032,["data-source"])]),_:2},1024)]),_:2},1024))),256))]),_:2},1024)]),_:1},8,["options"])])}}}),A=W(O,[["__scopeId","data-v-917fc341"]]);export{A as default};
