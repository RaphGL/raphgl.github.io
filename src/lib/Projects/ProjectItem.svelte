<script lang="ts">
	import { fly } from 'svelte/transition';
	export let href: string = '';
	export let desc: string = '';
	export let tags: string[] = [];
</script>

<a {href} transition:fly={{ y: 100, duration: 100 }}>
	<div class="project-name"><slot /></div>

	{#if desc}
		<div class="project-description">{desc}</div>
	{/if}

	{#if tags.length}
		<ul class="project-tags">
			{#each tags as tag}
				<li>{tag}</li>
			{/each}
		</ul>
	{/if}
</a>

<style>
	a {
		all: unset;
		background: rgb(32, 33, 41);
		background: linear-gradient(310deg, rgba(32, 33, 41, 1) 0%, rgba(40, 42, 54, 1) 100%);
		border-radius: var(--radius-size);
		font-size: 15pt;
		display: flex;
		width: 25vw;
		flex-direction: column;
		text-align: center;
		padding: 1em;
	}

	a:hover {
		cursor: pointer;
	}

	.project-name {
		font-weight: bold;
		font-size: 18pt;
	}

	.project-tags {
		display: flex;
		gap: 0.5em;
		justify-content: center;
		flex-wrap: wrap;
	}

	ul {
		padding: 0;
		list-style: none;
	}

	.project-tags li {
		background: #383a59;
		padding: 0.2em 1em;
		font-size: 12pt;
		border-radius: var(--radius-size);
	}
</style>
