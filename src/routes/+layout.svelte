<script>
	import '../app.css';
	import { page } from '$app/stores';

	const langs = [
		{ id: 'en', name: 'English' },
		{ id: 'es', name: 'Español' },
		{ id: 'pt', name: 'Português' },
	];
	let selectedLang = 'en';
</script>

<nav>
	<a class:active={$page.url.pathname === '/'} href="/">Home</a>
	<a class:active={$page.url.pathname === '/projects'} href="/projects">Projects</a>
	<a class:active={$page.url.pathname === '/blog'} href="/blog">Blog</a>

	<select bind:value={selectedLang}>
		{#each langs as lang}
			<option value={lang.id}>{lang.name}</option>
		{/each}
	</select>
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
		margin: auto;
		margin-top: 0.5em;
		padding: 1em;
		width: fit-content;
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

	select {
		background-color: rgb(40, 42, 54);
		background-color: linear-gradient(310deg, rgba(40, 42, 54, 1) 0%, rgba(32, 33, 41, 1) 100%);
		color: var(--gradient-bg-color);
		border-radius: var(--radius-size);
		border: 0;
		padding: 0.3em 1em;
		font-size: 0.8em;
	}

	select:hover {
		background-color: var(--hover-color);
		cursor: pointer;
	}
</style>
