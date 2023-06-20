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
		background: var(--gradient-fg-color);
		border-radius: var(--radius-size);
		font-size: 15pt;
		display: flex;
		flex-direction: column;
		text-align: center;
		justify-content: center;
		padding: 1.5em;
	}

	@media only screen and (min-width: 768px) {
		a {
			width: 25vw;
		}
	}

	a:hover {
		cursor: pointer;
		background: var(--gradient-hover-color);
	}

	a:hover .project-tags li {
		background: var(--gradient-fg-color);
	}

	.project-name {
		font-weight: bold;
		font-size: 18pt;
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
		background: #383a59;
		padding: 0.2em 1em;
		font-size: 12pt;
		border-radius: var(--radius-size);
	}
</style>
