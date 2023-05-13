<script>
	import '../app.css';
	import {page} from '$app/stores';
	export const ssr = false;
</script>

<nav>
	<a class:active={$page.url.pathname === "/"} href="/">Home</a>
	<a class:active={$page.url.pathname === "/projects"} href="/projects">Projects</a>
	<a class:active={$page.url.pathname === "/blog" } href="/blog">Blog</a>
</nav>

<slot />

<footer>
	<b>RaphGL</b> @2023
</footer>

<style>
	footer {
		text-align: center;
		font-size: 14pt;
		margin: 1.5em;
	}

	nav {
		display: flex;
		justify-content: center;
		gap: 2em;
		padding: 1em 0 1em 0;
	}

	nav a {
		text-decoration: none;
		color: var(--fg-color);
		width: 4em;
		text-align: center;
	}

	nav a:hover {
		color: var(--hover-color);
		font-weight: bold;
	}

	.active {
		color: var(--hover-color);
	}
</style>
