import{r as a}from"./request-5Xt99UU9.js";const r=t=>a({url:"/application/search",method:"get",params:t}),p=t=>a({url:"/application/detail",method:"get",params:t}),i=t=>a({url:"/application/instance/info",method:"get",params:t}),n=t=>a({url:"/application/service/form",method:"get",params:t}),c=t=>a({url:"/application/metric-dashboard",method:"get",params:t}),s=t=>a({url:"/application/event",method:"get",params:t}),l=t=>a({url:"/application/config/operatorLog",method:"get",params:{appName:t}}),g=(t,o)=>a({url:"/application/config/operatorLog",method:"put",params:{appName:t,operatorLog:o}}),u=t=>a({url:"/application/config/flowWeight",method:"get",params:{appName:t}}),m=(t,o)=>a({url:"/application/config/flowWeight",method:"put",params:{appName:t},data:{flowWeightSets:o}}),h=t=>a({url:"/application/config/gray",method:"get",params:{appName:t}}),d=(t,o)=>a({url:"/application/config/gray",method:"put",params:{appName:t},data:{graySets:o}});export{i as a,n as b,c,m as d,d as e,l as f,p as g,u as h,h as i,s as l,r as s,g as u};
