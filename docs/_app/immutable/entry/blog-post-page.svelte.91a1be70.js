import{S as M,i as z,s as B,u as ut,e as g,c as f,b as m,f as d,h as p,j as x,k as r,v as dt,w as ct,x as gt,p as j,z as rt,A as P,q as A,t as w,a as S,d as v,g as I,n as ot,C as et,B as ft,X as nt,y as mt,l as T,m as D,o as q,r as W,F as pt}from"../chunks/index.85023978.js";import{a as it,C as yt,f as at}from"../chunks/ContentRect.c0cc4b33.js";function wt(l){let t,n,e,o;const a=l[1].default,i=ut(a,l,l[0],null);return{c(){t=g("div"),n=g("h2"),i&&i.c(),this.h()},l(s){t=f(s,"DIV",{class:!0});var h=m(t);n=f(h,"H2",{});var u=m(n);i&&i.l(u),u.forEach(d),h.forEach(d),this.h()},h(){p(t,"class","svelte-9tlwqu")},m(s,h){x(s,t,h),r(t,n),i&&i.m(n,null),o=!0},p(s,[h]){i&&i.p&&(!o||h&1)&&dt(i,a,s,s[0],o?gt(a,s[0],h,null):ct(s[0]),null)},i(s){o||(j(i,s),rt(()=>{o&&(e||(e=P(t,it,{},!0)),e.run(1))}),o=!0)},o(s){A(i,s),e||(e=P(t,it,{},!1)),e.run(0),o=!1},d(s){s&&d(t),i&&i.d(s),s&&e&&e.end()}}}function vt(l,t,n){let{$$slots:e={},$$scope:o}=t;return l.$$set=a=>{"$$scope"in a&&n(0,o=a.$$scope)},[o,e]}class Ct extends M{constructor(t){super(),z(this,t,vt,wt,B,{})}}function st(l,t,n){const e=l.slice();return e[1]=t[n],e}function ht(l){let t,n=l[0],e=[];for(let o=0;o<n.length;o+=1)e[o]=lt(st(l,n,o));return{c(){for(let o=0;o<e.length;o+=1)e[o].c();t=et()},l(o){for(let a=0;a<e.length;a+=1)e[a].l(o);t=et()},m(o,a){for(let i=0;i<e.length;i+=1)e[i]&&e[i].m(o,a);x(o,t,a)},p(o,a){if(a&1){n=o[0];let i;for(i=0;i<n.length;i+=1){const s=st(o,n,i);e[i]?e[i].p(s,a):(e[i]=lt(s),e[i].c(),e[i].m(t.parentNode,t))}for(;i<e.length;i+=1)e[i].d(1);e.length=n.length}},d(o){ft(e,o),o&&d(t)}}}function lt(l){let t,n,e=l[1].innerHTML+"",o,a,i,s;return{c(){t=g("li"),n=g("a"),o=w(e),i=S(),this.h()},l(h){t=f(h,"LI",{class:!0});var u=m(t);n=f(u,"A",{href:!0,class:!0});var C=m(n);o=v(C,e),C.forEach(d),i=I(u),u.forEach(d),this.h()},h(){p(n,"href",a="#"+l[1].id),p(n,"class","svelte-1cth0oe"),p(t,"class",s=nt(l[1].tagName)+" svelte-1cth0oe")},m(h,u){x(h,t,u),r(t,n),r(n,o),r(t,i)},p(h,u){u&1&&e!==(e=h[1].innerHTML+"")&&mt(o,e),u&1&&a!==(a="#"+h[1].id)&&p(n,"href",a),u&1&&s!==(s=nt(h[1].tagName)+" svelte-1cth0oe")&&p(t,"class",s)},d(h){h&&d(t)}}}function bt(l){let t,n,e,o,a,i,s=l[0]!=null&&ht(l);return{c(){t=g("div"),n=g("h3"),e=w("Table of Contents"),o=S(),a=g("div"),i=g("ul"),s&&s.c(),this.h()},l(h){t=f(h,"DIV",{class:!0});var u=m(t);n=f(u,"H3",{class:!0});var C=m(n);e=v(C,"Table of Contents"),C.forEach(d),o=I(u),a=f(u,"DIV",{class:!0});var b=m(a);i=f(b,"UL",{class:!0});var V=m(i);s&&s.l(V),V.forEach(d),b.forEach(d),u.forEach(d),this.h()},h(){p(n,"class","svelte-1cth0oe"),p(i,"class","svelte-1cth0oe"),p(a,"class","svelte-1cth0oe"),p(t,"class","table-of-contents svelte-1cth0oe")},m(h,u){x(h,t,u),r(t,n),r(n,e),r(t,o),r(t,a),r(a,i),s&&s.m(i,null)},p(h,[u]){h[0]!=null?s?s.p(h,u):(s=ht(h),s.c(),s.m(i,null)):s&&(s.d(1),s=null)},i:ot,o:ot,d(h){h&&d(t),s&&s.d()}}}function St(l,t,n){let{contents:e}=t;return l.$$set=o=>{"contents"in o&&n(0,e=o.contents)},[e]}class It extends M{constructor(t){super(),z(this,t,St,bt,B,{contents:0})}}function _t(l){let t;return{c(){t=w("Setting up NeoVim for Godot")},l(n){t=v(n,"Setting up NeoVim for Godot")},m(n,e){x(n,t,e)},d(n){n&&d(t)}}}function Nt(l){let t,n,e,o,a,i,s,h,u,C,b,V,O,G,Y,R,$,U,J,k,E,H,_,L;return n=new It({props:{contents:l[0]}}),{c(){t=g("div"),T(n.$$.fragment),e=S(),o=g("div"),a=g("h1"),i=w("How to do it"),s=S(),h=g("h2"),u=w("Why I'm doing it"),C=S(),b=g("h2"),V=w("Why I'm doing it 2"),O=S(),G=g("h2"),Y=w("Yeah"),R=S(),$=g("h3"),U=w("Another subsection"),J=S(),k=g("h2"),E=w("Conclusion"),H=w(`
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
			NeoVim). Setting up the editor`),this.h()},l(y){t=f(y,"DIV",{class:!0});var N=m(t);D(n.$$.fragment,N),e=I(N),o=f(N,"DIV",{class:!0});var c=m(o);a=f(c,"H1",{});var F=m(a);i=v(F,"How to do it"),F.forEach(d),s=I(c),h=f(c,"H2",{});var X=m(h);u=v(X,"Why I'm doing it"),X.forEach(d),C=I(c),b=f(c,"H2",{id:!0});var K=m(b);V=v(K,"Why I'm doing it 2"),K.forEach(d),O=I(c),G=f(c,"H2",{});var Q=m(G);Y=v(Q,"Yeah"),Q.forEach(d),R=I(c),$=f(c,"H3",{});var Z=m($);U=v(Z,"Another subsection"),Z.forEach(d),J=I(c),k=f(c,"H2",{id:!0});var tt=m(k);E=v(tt,"Conclusion"),tt.forEach(d),H=v(c,`
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
			NeoVim). Setting up the editor`),c.forEach(d),N.forEach(d),this.h()},h(){p(b,"id","something-else"),p(k,"id","conclusion"),p(o,"class","blog-text svelte-1dl10o2"),p(t,"class","content svelte-1dl10o2")},m(y,N){x(y,t,N),q(n,t,null),r(t,e),r(t,o),r(o,a),r(a,i),r(o,s),r(o,h),r(h,u),r(o,C),r(o,b),r(b,V),r(o,O),r(o,G),r(G,Y),r(o,R),r(o,$),r($,U),r(o,J),r(o,k),r(k,E),r(o,H),L=!0},p(y,N){const c={};N&1&&(c.contents=y[0]),n.$set(c)},i(y){L||(j(n.$$.fragment,y),rt(()=>{L&&(_||(_=P(t,at,{},!0)),_.run(1))}),L=!0)},o(y){A(n.$$.fragment,y),_||(_=P(t,at,{},!1)),_.run(0),L=!1},d(y){y&&d(t),W(n),y&&_&&_.end()}}}function xt(l){let t,n,e,o;return t=new Ct({props:{$$slots:{default:[_t]},$$scope:{ctx:l}}}),e=new yt({props:{$$slots:{default:[Nt]},$$scope:{ctx:l}}}),{c(){T(t.$$.fragment),n=S(),T(e.$$.fragment)},l(a){D(t.$$.fragment,a),n=I(a),D(e.$$.fragment,a)},m(a,i){q(t,a,i),x(a,n,i),q(e,a,i),o=!0},p(a,[i]){const s={};i&2&&(s.$$scope={dirty:i,ctx:a}),t.$set(s);const h={};i&3&&(h.$$scope={dirty:i,ctx:a}),e.$set(h)},i(a){o||(j(t.$$.fragment,a),j(e.$$.fragment,a),o=!0)},o(a){A(t.$$.fragment,a),A(e.$$.fragment,a),o=!1},d(a){W(t,a),a&&d(n),W(e,a)}}}function Vt(l,t,n){let e=null;return pt(()=>{n(0,e=document.querySelectorAll(".blog-text h1, .blog-text h2, .blog-text h3, .blog-text h4, .blog-text h5, .blog-text h6"))}),[e]}class $t extends M{constructor(t){super(),z(this,t,Vt,xt,B,{})}}export{$t as default};
