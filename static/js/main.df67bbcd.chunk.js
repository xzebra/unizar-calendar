(this.webpackJsonpuzcalendar=this.webpackJsonpuzcalendar||[]).push([[1],{107:function(e){e.exports=JSON.parse('{"form":{"subjects":"Subjects","schedules":"Schedules","semester":"Semester","export_type":"Export Type","export_types":{"ics":"Google Calendar ICS (Recommended)","csv":"Google Calendar CSV","org":"Org Mode"},"submit":"Generate"},"semester":{"first_semester":"First semester","second_semester":"Second semester"},"tables":{"tooltip":{"mobile":"Double tap on a cell to edit the content. Hold tap on a row to delete it or insert another one.","web":"Double click on a cell to edit the content. Right click on a row to delete it or insert another one."},"subjects":{"id":"Subject ID","name":"Subject Name","desc":"Subject Description","id_tooltip":"Subject ID is a shortname of the subject to make it easier to be referenced in Schedules table"},"schedules":{"weekday":"Weekday","subject_id":"Subject ID","subject_id_tooltip":"ID specified in Subjects table","start_hour":"Start Hour","end_hour":"End Hour","is_practical":"Is practical?","is_practical_tooltip":"Some weeks have classes but not practical ones"}},"guides":{"title":"How to import?","csv":"It is recommended to create a **new calendar** before importing the data.\\nAlso, make sure you are actually selecting the new calendar as the **destination** of the import.\\n\\n[Click here](https://support.google.com/calendar/answer/37118?hl=en#import_to_gcal) for a guide about importing *.csv* files.","ics":"It is recommended to create a **new calendar** before importing the data.\\nAlso, make sure you are actually selecting the new calendar as the **destination** of the import.\\n\\n[Click here](https://support.google.com/calendar/answer/37118?hl=en#import_to_gcal) for a guide about importing *.ics* files."},"warning":{"title":"Warning","inconsistent":"Calendar published in the official EINA page is inconsistent with the internal Google Calendar which this application uses (but it is created by the University). Day changes and non practical days may be incorrect."}}')},238:function(e,t,n){},239:function(e,t,n){"use strict";n.r(t);var a=n(7),s=n.n(a),r=n(20),c=n(57),o=n(0),i=n.n(o),l=n(13),u=n.n(l),d=n(244),b=n(21),j=n(66),p=n(11),h=n(85),m=n(38),f=n(104),x=n(105),O={en:function(){return Promise.resolve().then(n.t.bind(null,107,3))},es:function(){return n.e(0).then(n.t.bind(null,246,3))}},v=Object.keys(O),w=O,g=n(107);var y={loadPath:"{{lng}}|{{ns}}",request:function(e,t,n,a){try{var s=t.split("|"),r=Object(p.a)(s,1)[0];(c=r,"en"!==c?w[c]().then((function(e){return e.default})).catch((function(e){console.error(e)})):Promise.resolve(g)).then((function(e){a(null,{data:JSON.stringify(e),status:200})}))}catch(o){console.error(o),a(null,{status:500})}var c}};h.a.use(f.a).use(x.a).use(m.e).init({fallbackLng:"en",debug:!0,whitelist:v,detection:{order:["cookie","localStorage","navigator"],lookupCookie:"locale",lookupLocalStorage:"locale",caches:["localStorage","cookie"],excludeCacheFor:["cimode"],checkWhitelist:!0},interpolation:{escapeValue:!1},backend:y});var _=h.a,k=n(18),S=n(110),C=n(111),I=n(245),N=(n(121),n(1));function R(e){var t=e.title,n=e.children;return Object(N.jsxs)("div",{className:"row g-2",children:[Object(N.jsxs)("div",{className:"d-flex justify-content-between col-12",children:[Object(N.jsx)("div",{children:Object(N.jsx)("h2",{className:"popup-title",children:t})}),Object(N.jsx)("div",{className:"d-flex",children:Object(N.jsx)("button",{type:"button",className:"btn-close my-auto","aria-label":"Close",onClick:b.PopupboxManager.close})})]}),Object(N.jsx)("div",{className:"col-12",children:n})]})}function D(e){var t=e.err;return Object(N.jsx)(R,{title:"Error",children:t.split("\n").map((function(e,t){return Object(N.jsx)("p",{children:e},t)}))})}var M=n(59),T=n.n(M);function E(e){e.result;var t=e.exportType,n=Object(I.a)().t;return Object(N.jsx)(R,{title:n("guides.title"),children:Object(N.jsx)("div",{className:"col-12",children:Object(N.jsx)(T.a,{className:"line-break",children:n("guides."+t)})})})}var A=n(112),L=n.n(A),F=n(9),P=n(19),B=n(26),G=n(23),z=n(22),H=n(80),J=n.n(H),W=n(113),V=n.n(W),U=function(e){Object(G.a)(n,e);var t=Object(z.a)(n);function n(e){var a;return Object(P.a)(this,n),(a=t.call(this,e)).handleChangeComplete=function(e){a.setState({hour:e.formatted24},(function(){return a.props.onCommit()}))},a.state={hour:e.value},a.label=e.label,a}return Object(B.a)(n,[{key:"getValue",value:function(){return Object(F.a)({},this.label,this.state.hour)}},{key:"getInputNode",value:function(){return u.a.findDOMNode(this).getElementsByTagName("input")[0]}},{key:"render",value:function(){return Object(N.jsx)(V.a,{time:this.state.hour,onChange:this.handleChangeComplete})}}]),n}(i.a.Component),q=n(64),K=n(67),Q=n(63),X=n(65),Y=n(31),Z=n(5),$=Y.Menu.ContextMenu,ee=Y.Menu.MenuItem,te=Y.Menu.SubMenu;function ne(e){var t=e.idx,n=e.id,a=e.rowIdx,s=e.onRowDelete,r=e.onRowInsertAbove,c=e.onRowInsertBelow;return Object(N.jsxs)($,{id:n,children:[Object(N.jsx)(ee,{data:{rowIdx:a,idx:t},onClick:s,children:"Delete Row"}),Object(N.jsxs)(te,{title:"Insert Row",children:[Object(N.jsx)(ee,{data:{rowIdx:a,idx:t},onClick:r,children:"Above"}),Object(N.jsx)(ee,{data:{rowIdx:a,idx:t},onClick:c,children:"Below"})]})]})}var ae,se=function(e,t){return function(n){var a=t,s=Object(Z.a)(n);return s.splice(e,0,a),s}},re=Y.Menu.ContextMenuTrigger;function ce(e){var t=e.title,n=e.tooltip,a=e.startingRows,s=e.defaultRow,r=e.cols,c=e.onChange,i=Object(o.useState)(a),l=Object(p.a)(i,2),u=l[0],d=l[1];return Object(N.jsxs)("div",{className:"col-12",children:[Object(N.jsx)("label",{htmlFor:t.toLowerCase(),className:"form-label",children:t}),Object(N.jsx)(K.a,{placement:"auto",overlay:Object(N.jsx)(q.a,{children:n}),children:Object(N.jsx)(Q.a,{className:"ms-1",icon:X.a})}),Object(N.jsx)(J.a,{name:t.toLowerCase(),columns:r,rowGetter:function(e){return u[e]},rowsCount:u.length,onGridRowsUpdated:function(e){for(var t=e.fromRow,n=e.toRow,a=e.updated,s=u.slice(),r=t;r<=n;r++)s[r]=Object(k.a)(Object(k.a)({},s[r]),a);d(s),c(s)},enableCellSelect:!0,minHeight:150,contextMenu:Object(N.jsx)(ne,{id:t.toLowerCase()+"ContextMenu",onRowDelete:function(e,t){var n=t.rowIdx;return d(function(e,t){return function(n){var a=Object(Z.a)(n);if(1===a.length){var s=t;a.splice(e,1,s)}else a.splice(e,1);return a}}(n,s))},onRowInsertAbove:function(e,t){var n=t.rowIdx;return d(se(n,s))},onRowInsertBelow:function(e,t){var n=t.rowIdx;return d(se(n+1,s))}}),RowsContainer:re})]})}var oe=Y.Editors.DropDownEditor,ie=j.a.form(ae||(ae=Object(c.a)(["\n  justify-content: center;\n  width: 80%;\n"]))),le={editable:!0},ue=function(e){return Object(N.jsxs)("span",{children:[e.column.name,Object(N.jsx)(K.a,{placement:"auto",overlay:Object(N.jsx)(q.a,{children:e.column.tooltip}),children:Object(N.jsx)(Q.a,{className:"ms-1",icon:X.a})})]})},de=Object(N.jsx)(oe,{options:[{id:"true",value:"True"},{id:"false",value:"False"}]}),be=[{weekday:"Lx",class_id:"ia",start_hour:"18:00",end_hour:"18:50",is_practical:"True"},{weekday:"La",class_id:"ia",start_hour:"19:00",end_hour:"19:50",is_practical:"False"}];function je(e,t){return e.map((function(e){return e.key})).join(";")+"\n"+t.map((function(e){return Object.entries(e).map((function(e){return e[1]})).join(";")})).join("\n")}function pe(){var e=Object(I.a)().t,t=[{key:"class_id",name:e("tables.subjects.id"),tooltip:e("tables.subjects.id_tooltip"),headerRenderer:ue},{key:"class_name",name:e("tables.subjects.name")},{key:"class_desc",name:e("tables.subjects.desc")}].map((function(e){return Object(k.a)(Object(k.a)({},e),le)})),n=[{class_id:"ia",class_name:"Inteligencia Artificial",class_desc:"algo"},{class_id:"ssdd",class_name:"Sistemas Distribuidos",class_desc:"otro"}],a=C.isMobile?e("tables.tooltip.mobile"):e("tables.tooltip.web"),s=[{key:"weekday",name:e("tables.schedules.weekday")},{key:"class_id",name:e("tables.schedules.subject_id"),tooltip:e("tables.schedules.subject_id_tooltip"),headerRenderer:ue},{key:"start_hour",name:e("tables.schedules.start_hour"),editor:Object(N.jsx)(U,{label:"start_hour"})},{key:"end_hour",name:e("tables.schedules.end_hour"),editor:Object(N.jsx)(U,{label:"end_hour"})},{key:"is_practical",name:e("tables.schedules.is_practical"),editor:de,tooltip:e("tables.schedules.is_practical_tooltip"),headerRenderer:ue}].map((function(e){return Object(k.a)(Object(k.a)({},e),le)})),r=Object(S.a)(),c=r.register,i=r.handleSubmit,l=Object(o.useState)(n),u=Object(p.a)(l,2),d=u[0],j=u[1],h=Object(o.useState)(be),m=Object(p.a)(h,2),f=m[0],x=m[1],O=Object(o.useState)(""),v=Object(p.a)(O,2),w=v[0],g=v[1],y=Object(o.useState)(""),_=Object(p.a)(y,2),R=_[0],M=_[1],T=Object(o.useState)(""),A=Object(p.a)(T,2),F=(A[0],A[1]);Object(o.useEffect)((function(){fetch("/unizar-calendar/data/semester1.json").then((function(e){return e.text()})).then((function(e){return g(e)})),fetch("/unizar-calendar/data/semester2.json").then((function(e){return e.text()})).then((function(e){return M(e)}))}),[]);return Object(N.jsxs)(ie,{className:"row g-3",onSubmit:i((function(e){var n=je(t,d),a=je(s,f);e.semester=parseInt(e.semester);var r=window.calendar(e.semester,n,a,e.exportType,1===e.semester?w:R);if(r instanceof Error)!function(e){var t=Object(N.jsx)(D,{err:e.message});b.PopupboxManager.open({content:t,config:{fadeIn:!0,fadeInSpeed:400}})}(r);else{F(r);var c=new Blob([r],{type:"text/plain"});L()(c,"calendar."+e.exportType),"ics"!=e.exportType&&"csv"!=e.exportType||function(e,t){var n=Object(N.jsx)(E,{result:e,exportType:t});b.PopupboxManager.open({content:n,config:{fadeIn:!0,fadeInSpeed:400}})}(r,e.exportType)}})),children:[Object(N.jsx)(ce,{title:e("form.subjects"),tooltip:a,startingRows:n,defaultRow:{class_id:"",class_name:"",class_desc:""},cols:t,onChange:j}),Object(N.jsx)(ce,{title:e("form.schedules"),tooltip:a,startingRows:be,defaultRow:{weekday:"Lx",class_id:"",start_hour:"00:00",end_hour:"00:00",is_practical:"False"},cols:s,onChange:x}),Object(N.jsxs)("div",{className:"col-md-6",children:[Object(N.jsx)("label",{htmlFor:"semester",className:"form-label",children:e("form.semester")}),Object(N.jsxs)("select",{type:"number",name:"semester",className:"form-select",ref:c,children:[Object(N.jsx)("option",{value:1,children:e("semester.first_semester")}),Object(N.jsx)("option",{value:2,children:e("semester.second_semester")})]})]}),Object(N.jsxs)("div",{className:"col-md-6",children:[Object(N.jsx)("label",{htmlFor:"exportType",className:"form-label",children:e("form.export_type")}),Object(N.jsxs)("select",{name:"exportType",className:"form-select",ref:c,children:[Object(N.jsx)("option",{value:"ics",children:e("form.export_types.ics")}),Object(N.jsx)("option",{value:"csv",children:e("form.export_types.csv")}),Object(N.jsx)("option",{value:"org",children:e("form.export_types.org")})]})]}),Object(N.jsx)("div",{className:"col-md-4",children:Object(N.jsx)("input",{type:"submit",className:"w-100 btn btn-primary btn-lg",value:e("form.submit")})})]})}var he;n(238);function me(e){var t=e.warningID,n=Object(I.a)().t;return Object(N.jsx)(R,{title:n("warning.title"),children:Object(N.jsx)("div",{className:"col-12",children:Object(N.jsx)(T.a,{className:"line-break",children:n("warning."+t)})})})}var fe=j.a.div(he||(he=Object(c.a)(["\n  height: 100%;\n  display: flex;\n  justify-content: center;\n  width: 100%;\n"])));function xe(){return Object(o.useEffect)((function(){function e(){return(e=Object(r.a)(s.a.mark((function e(){var t,n,a,r,c;return s.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=new window.Go,e.next=3,fetch("/unizar-calendar/calendar.wasm");case 3:return n=e.sent,e.next=6,n.arrayBuffer();case 6:return a=e.sent,e.next=9,WebAssembly.instantiate(a,t.importObject);case 9:return r=e.sent,c=r.instance,e.next=13,t.run(c);case 13:case"end":return e.stop()}}),e)})))).apply(this,arguments)}!function(){e.apply(this,arguments)}(),function(e){var t=Object(N.jsx)(me,{warningID:e});b.PopupboxManager.open({content:t,config:{fadeIn:!0,fadeInSpeed:400}})}("inconsistent")}),[]),Object(N.jsx)(o.Suspense,{fallback:Object(N.jsx)(Oe,{}),children:Object(N.jsxs)(fe,{children:[Object(N.jsx)(b.PopupboxContainer,{}),Object(N.jsx)(pe,{})]})})}var Oe=function(){return Object(N.jsx)("div",{className:"App",children:Object(N.jsx)("div",{children:"Loading..."})})};u.a.render(Object(N.jsx)(d.a,{i18n:_,children:Object(N.jsx)(xe,{})}),document.getElementById("root"))}},[[239,2,3]]]);
//# sourceMappingURL=main.df67bbcd.chunk.js.map