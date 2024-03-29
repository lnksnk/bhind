package bhembed

const jsondateexstensions string = `/* JSON Date Extensions, Version 1.2.1 */
(function(n){"use strict";var t,i,r;JSON&&!JSON.dateParser&&(t=/^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2}(?:\.{0,1}\d*))(?:Z|(\+|-)([\d|:]*))?$/,i=/^\/Date\((d|-|.*)\)[\/|\\]$/,JSON.parseMsAjaxDate=!1,JSON.useDateParser=function(t){t!==n?JSON._parseSaved&&(JSON.parse=JSON._parseSaved,JSON._parseSaved=null):JSON._parseSaved||(JSON._parseSaved=JSON.parse,JSON.parse=JSON.parseWithDate)},r=function(r){return function(u,f){var o=f,e,s;return typeof f=="string"&&(e=t.exec(f),e?o=new Date(f):JSON.parseMsAjaxDate&&(e=i.exec(f),e&&(s=e[1].split(/[-+,.]/),o=new Date(s[0]?+s[0]:0-+s[1])))),r!==n?r(u,o):o}},JSON.dateParser=r(),JSON.parseWithDate=function(n,t){var i=JSON._parseSaved?JSON._parseSaved:JSON.parse;try{return i(n,r(t))}catch(u){throw new Error("JSON content could not be parsed");}},JSON.dateStringToDate=function(n,r){var u,f;return(r||(r=null),!n)?r:n.getTime?n:((n[0]==='"'||n[0]==="'")&&(n=n.substr(1,n.length-2)),u=t.exec(n),u)?new Date(n):JSON.parseMsAjaxDate?(u=i.exec(n),u)?(f=u[1].split(/[-,.]/),new Date(+f[0])):r:r})})();
//# sourceMappingURL=json.date-extensions.min.js.map
`
