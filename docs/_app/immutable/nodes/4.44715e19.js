import{s as P,f as X,u as F,g as K,h as Q,d as Z,n as O,r as ee,i as R,e as te,o as oe}from"../chunks/scheduler.3c61ddd7.js";import{S as Y,i as A,g as m,h as p,j as _,f as u,k as c,a as w,d as $,t as k,s as b,x as T,c as S,l as N,z as y,A as x,e as U,D as ne,m as W,n as B,o as ie,r as V,u as G,v as j,w as L}from"../chunks/index.0ac3503a.js";import{e as J}from"../chunks/each.e59479a4.js";import{C as se}from"../chunks/ContentRect.040ac44e.js";function ae(a){let e,i;const o=a[1].default,t=X(o,a,a[0],null);return{c(){e=m("h2"),t&&t.c(),this.h()},l(n){e=p(n,"H2",{class:!0});var s=_(e);t&&t.l(s),s.forEach(u),this.h()},h(){c(e,"class","svelte-1lmv1b5")},m(n,s){w(n,e,s),t&&t.m(e,null),i=!0},p(n,[s]){t&&t.p&&(!i||s&1)&&F(t,o,n,n[0],i?Q(o,n[0],s,null):K(n[0]),null)},i(n){i||($(t,n),i=!0)},o(n){k(t,n),i=!1},d(n){n&&u(e),t&&t.d(n)}}}function le(a,e,i){let{$$slots:o={},$$scope:t}=e;return a.$$set=n=>{"$$scope"in n&&i(0,t=n.$$scope)},[t,o]}class re extends Y{constructor(e){super(),A(this,e,le,ae,P,{})}}function E(a,e,i){const o=a.slice();return o[6]=e[i],o}function M(a){let e,i="Table of Contents";return{c(){e=m("span"),e.textContent=i,this.h()},l(o){e=p(o,"SPAN",{class:!0,"data-svelte-h":!0}),T(e)!=="svelte-1ycwk3l"&&(e.textContent=i),this.h()},h(){c(e,"class","svelte-mlj1ph")},m(o,t){w(o,e,t)},d(o){o&&u(e)}}}function D(a){let e,i,o=a[0]!=null&&z(a);return{c(){e=m("div"),i=m("ul"),o&&o.c(),this.h()},l(t){e=p(t,"DIV",{class:!0});var n=_(e);i=p(n,"UL",{class:!0});var s=_(i);o&&o.l(s),s.forEach(u),n.forEach(u),this.h()},h(){c(i,"class","svelte-mlj1ph"),c(e,"class","svelte-mlj1ph"),N(e,"transform",`translateY(${a[2]}px)`)},m(t,n){w(t,e,n),y(e,i),o&&o.m(i,null)},p(t,n){t[0]!=null?o?o.p(t,n):(o=z(t),o.c(),o.m(i,null)):o&&(o.d(1),o=null),n&4&&N(e,"transform",`translateY(${t[2]}px)`)},d(t){t&&u(e),o&&o.d()}}}function z(a){let e,i=J(a[0]),o=[];for(let t=0;t<i.length;t+=1)o[t]=q(E(a,i,t));return{c(){for(let t=0;t<o.length;t+=1)o[t].c();e=U()},l(t){for(let n=0;n<o.length;n+=1)o[n].l(t);e=U()},m(t,n){for(let s=0;s<o.length;s+=1)o[s]&&o[s].m(t,n);w(t,e,n)},p(t,n){if(n&1){i=J(t[0]);let s;for(s=0;s<i.length;s+=1){const l=E(t,i,s);o[s]?o[s].p(l,n):(o[s]=q(l),o[s].c(),o[s].m(e.parentNode,e))}for(;s<o.length;s+=1)o[s].d(1);o.length=i.length}},d(t){t&&u(e),ne(o,t)}}}function q(a){let e,i,o=a[6].innerHTML+"",t,n,s,l;return{c(){e=m("li"),i=m("a"),t=W(o),s=b(),this.h()},l(r){e=p(r,"LI",{class:!0});var h=_(e);i=p(h,"A",{href:!0,class:!0});var C=_(i);t=B(C,o),C.forEach(u),s=S(h),h.forEach(u),this.h()},h(){c(i,"href",n="#"+a[6].id),c(i,"class","svelte-mlj1ph"),c(e,"class",l=R(a[6].tagName)+" svelte-mlj1ph")},m(r,h){w(r,e,h),y(e,i),y(i,t),y(e,s)},p(r,h){h&1&&o!==(o=r[6].innerHTML+"")&&ie(t,o),h&1&&n!==(n="#"+r[6].id)&&c(i,"href",n),h&1&&l!==(l=R(r[6].tagName)+" svelte-mlj1ph")&&c(e,"class",l)},d(r){r&&u(e)}}}function he(a){let e=!1,i=()=>{e=!1},o,t,n,s,l=`<svg xmlns="http://www.w3.org/2000/svg" height="1.25em" viewBox="0 0 448 512"><style>svg {
						fill: #deddda;
					}
				</style><path d="M16 132h416c8.837 0 16-7.163 16-16V76c0-8.837-7.163-16-16-16H16C7.163 60 0 67.163 0 76v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16z"></path></svg>`,r,h,C,H;Z(a[4]);let f=a[1]&&M(),d=a[1]&&D(a);return{c(){t=m("div"),n=m("div"),s=m("div"),s.innerHTML=l,r=b(),f&&f.c(),h=b(),d&&d.c(),this.h()},l(g){t=p(g,"DIV",{class:!0});var v=_(t);n=p(v,"DIV",{class:!0});var I=_(n);s=p(I,"DIV",{class:!0,role:!0,tabindex:!0,"data-svelte-h":!0}),T(s)!=="svelte-1ep160e"&&(s.innerHTML=l),r=S(I),f&&f.l(I),I.forEach(u),h=S(v),d&&d.l(v),v.forEach(u),this.h()},h(){c(s,"class","btn svelte-mlj1ph"),c(s,"role","button"),c(s,"tabindex","0"),c(n,"class","expand-btn svelte-mlj1ph"),N(n,"transform",`translateY(${a[2]}px)`),c(t,"class","table-of-contents svelte-mlj1ph")},m(g,v){w(g,t,v),y(t,n),y(n,s),y(n,r),f&&f.m(n,null),y(t,h),d&&d.m(t,null),C||(H=[x(window,"scroll",()=>{e=!0,clearTimeout(o),o=setTimeout(i,100),a[4]()}),x(s,"click",a[5]),x(s,"keydown",a[3])],C=!0)},p(g,[v]){v&4&&!e&&(e=!0,clearTimeout(o),scrollTo(window.pageXOffset,g[2]),o=setTimeout(i,100)),g[1]?f||(f=M(),f.c(),f.m(n,null)):f&&(f.d(1),f=null),v&4&&N(n,"transform",`translateY(${g[2]}px)`),g[1]?d?d.p(g,v):(d=D(g),d.c(),d.m(t,null)):d&&(d.d(1),d=null)},i:O,o:O,d(g){g&&u(t),f&&f.d(),d&&d.d(),C=!1,ee(H)}}}function ue(a,e,i){let{contents:o}=e,t=!1,n;function s(h){te.call(this,a,h)}function l(){i(2,n=window.pageYOffset)}const r=()=>i(1,t=!t);return a.$$set=h=>{"contents"in h&&i(0,o=h.contents)},[o,t,n,s,l,r]}class de extends Y{constructor(e){super(),A(this,e,ue,he,P,{contents:0})}}function ce(a){let e;return{c(){e=W("Setting up NeoVim for Godot")},l(i){e=B(i,"Setting up NeoVim for Godot")},m(i,o){w(i,e,o)},d(i){i&&u(e)}}}function fe(a){let e,i,o,t,n=`<h1>How to do it</h1> <h2>Why I&#39;m doing it</h2> <h2 id="something-else">Why I&#39;m doing it 2</h2> <h2>Yeah</h2> <h3>Another subsection</h3> <h2 id="conclusion">Conclusion</h2>
			Recently I’ve been using Godot a bit to mess around with game development and I wanted to have
			autocompletion working with CoC. I looked around and there were no extensions for the Godot LSP.
			So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up If you have CoC working
			correctly already, open NeoVim and run :CocConfig, it’ll open the configuration file for your CoC,
			then just add this to the JSON config: Additional steps You probably want syntax highlighting,
			for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor`,s;return i=new de({props:{contents:a[0]}}),{c(){e=m("div"),V(i.$$.fragment),o=b(),t=m("div"),t.innerHTML=n,this.h()},l(l){e=p(l,"DIV",{class:!0});var r=_(e);G(i.$$.fragment,r),o=S(r),t=p(r,"DIV",{class:!0,"data-svelte-h":!0}),T(t)!=="svelte-yi5dv"&&(t.innerHTML=n),r.forEach(u),this.h()},h(){c(t,"class","blog-text svelte-1dl10o2"),c(e,"class","content svelte-1dl10o2")},m(l,r){w(l,e,r),j(i,e,null),y(e,o),y(e,t),s=!0},p(l,r){const h={};r&1&&(h.contents=l[0]),i.$set(h)},i(l){s||($(i.$$.fragment,l),s=!0)},o(l){k(i.$$.fragment,l),s=!1},d(l){l&&u(e),L(i)}}}function ge(a){let e,i,o,t;return e=new re({props:{$$slots:{default:[ce]},$$scope:{ctx:a}}}),o=new se({props:{$$slots:{default:[fe]},$$scope:{ctx:a}}}),{c(){V(e.$$.fragment),i=b(),V(o.$$.fragment)},l(n){G(e.$$.fragment,n),i=S(n),G(o.$$.fragment,n)},m(n,s){j(e,n,s),w(n,i,s),j(o,n,s),t=!0},p(n,[s]){const l={};s&2&&(l.$$scope={dirty:s,ctx:n}),e.$set(l);const r={};s&3&&(r.$$scope={dirty:s,ctx:n}),o.$set(r)},i(n){t||($(e.$$.fragment,n),$(o.$$.fragment,n),t=!0)},o(n){k(e.$$.fragment,n),k(o.$$.fragment,n),t=!1},d(n){n&&u(i),L(e,n),L(o,n)}}}function me(a,e,i){let o=null;return oe(()=>{i(0,o=document.querySelectorAll(".blog-text h1, .blog-text h2, .blog-text h3, .blog-text h4, .blog-text h5, .blog-text h6"))}),[o]}class _e extends Y{constructor(e){super(),A(this,e,me,ge,P,{})}}export{_e as component};
