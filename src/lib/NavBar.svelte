<script lang="ts">
	import { page } from '$app/stores';
	import { slide } from 'svelte/transition';

	const langs = [
		{ id: 'en', name: 'English' },
		{ id: 'es', name: 'Español' },
		{ id: 'pt', name: 'Português' }
	];

	let selectedLang = 'en';
	let show = false;
</script>

<div class="navbox">
	{#if show}
		{@const pages = [
			['/', 'Home'],
			['/projects', 'Projects'],
			['/blog', 'Blog']
		]}
		<nav transition:slide={{ axis: 'x', duration: 150 }}>
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

	<div
		class="hamburger"
		class:hamburger-open={show}
		on:click={() => (show = !show)}
		on:keydown
		role="button"
		tabindex="0"
	>
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
</div>

<style>
	nav {
		display: flex;
		flex-direction: column;
		width: fit-content;
		background: var(--gradient-bg-color);
		padding: 1em;
		border-top-left-radius: var(--radius-size);
		border-bottom-left-radius: var(--radius-size);
		width: 76vw;
	}

	@media only screen and (min-width: 850px) {
		nav {
			width: 15vw;
		}
	}

	nav a {
		color: var(--fg-color);
		text-decoration: none;
		padding: 0.6em 0;
		padding-left: 1em;
		border-radius: var(--radius-size);
	}

	nav a:hover {
		font-weight: bold;
		background-color: var(--bg-color);
		color: var(--fg-color);
	}

	select {
		font-size: 11pt;
		font-weight: bold;
		padding: 0.5em;
		text-align: center;
		margin-left: 1em;
		margin-top: 1em;
		border-radius: 0.5em;
		color: var(--fg-color);
		background: var(--content-bg-color);
		border-color: transparent;
		cursor: pointer;
	}

	select:hover {
        background: var(--bg-color);
	}

	.active {
		color: var(--hover-color);
		font-weight: bold;
	}

	.navbox {
		display: flex;
		position: fixed;
		right: 1em;
		top: 1em;
	}

	.hamburger {
		padding: 1em;
		cursor: pointer;
	}

	.hamburger-open {
		background-color: var(--bg-color);
		border-top-right-radius: var(--radius-size);
		border-bottom-right-radius: var(--radius-size);
	}
</style>
