<script lang="ts">
	import { fade } from 'svelte/transition';
	export let href: string = '';
	export let desc: string = '';
	export let tags: string[] = [];
</script>

<a {href} in:fade|global={{ duration: 150 }}>
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
		background: var(--gradient-bg-color);
		border-radius: var(--radius-size);
		font-size: calc(var(--font-size) - 2pt);
		display: flex;
		flex-direction: column;
		gap: 0.2em;
		text-align: center;
		justify-content: center;
		padding: 1.5em;
	}

	a:hover {
		cursor: pointer;
		background: var(--gradient-hover-color);
	}

	a:hover .project-tags li {
		background: var(--gradient-bg-color);
	}

	.project-name {
		font-weight: bold;
		font-size: calc(var(--font-size) + 2pt);
	}

	.project-tags {
		margin-top: 0.5em;
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
		background: var(--gradient-hover-color);
		padding: 0.2em 1em;
		font-size: calc(var(--font-size) - 4pt);
		border-radius: var(--radius-size);
	}
</style>
