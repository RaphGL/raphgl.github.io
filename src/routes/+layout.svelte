<script>
	import '../app.css';
	import { page } from '$app/stores';

	const langs = [
		{ id: 'en', name: 'English' },
		{ id: 'es', name: 'Español' },
		{ id: 'pt', name: 'Português' }
	];

	let selectedLang = 'en';
	let show = false;
</script>

<svelte:head>
	<title>RaphGL</title>
</svelte:head>

<div class="hamburger" on:click={() => (show = !show)} on:keydown>
	<svg xmlns="http://www.w3.org/2000/svg" height="1.25em" viewBox="0 0 448 512">
		<style>
			svg {
				fill: #deddda;
			}
		</style><path
			d="M16 132h416c8.837 0 16-7.163 16-16V76c0-8.837-7.163-16-16-16H16C7.163 60 0 67.163 0 76v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16z"
		/>
	</svg>
</div>

{#if show}
	{@const pages = [
		['/', 'Home'],
		['/projects', 'Projects'],
		['/blog', 'Blog']
	]}
	<nav>
		{#each pages as [pageUrl, pageName]}
			<!-- hides menu when a link is clicked -->
			<a
				on:click={() => (show = false)}
				class:active={$page.url.pathname === pageUrl}
				href={pageUrl}>{pageName}</a
			>
		{/each}

		<select bind:value={selectedLang}>
			{#each langs as lang}
				<option value={lang.id}>{lang.name}</option>
			{/each}
		</select>
	</nav>
{/if}

<slot />

<footer>
	<b>RaphGL</b> @2023
</footer>

<style>
	.hamburger {
		position: fixed;
		top: 1.5em;
		right: 2em;
		cursor: pointer;
		z-index: 100;
	}

	nav {
		position: fixed;
		top: 0.5em;
		right: 1em;
		display: flex;
		flex-direction: column;
		gap: 2em;
		padding: 1em;
		padding-right: 5vw;
		background-color: var(--content-bg-color);
		border-radius: var(--radius-size);
	}

	@media only screen and (min-width: 768px) {
		nav {
			flex-direction: row;
		}
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
		font-weight: bold;
	}

	select {
		background-color: rgb(40, 42, 54);
		background-color: linear-gradient(310deg, rgba(40, 42, 54, 1) 0%, rgba(32, 33, 41, 1) 100%);
		color: var(--fg-color);
		border-radius: var(--radius-size);
		border: 0;
		padding: 0.3em 1em;
		text-align: center;
		font-size: 0.8em;
	}

	select:hover {
		background-color: var(--hover-color);
		cursor: pointer;
	}

	option {
		background-color: red;
		color: white;
	}

	footer {
		text-align: center;
		font-size: 14pt;
		margin: 1.5em;
	}
</style>
