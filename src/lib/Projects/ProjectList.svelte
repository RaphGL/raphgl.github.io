<script>
	import ProjectItem from './ProjectItem.svelte';
	import { projects } from '$lib/stores';
	export let size = 4;
</script>

<div class="project">
	<h2>Projects</h2>
	<div>
		{#each $projects.slice(0, size) as project}
			<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
				{project.name}
			</ProjectItem>
		{/each}
		<a class="view-more" href="/projects">View More</a>
	</div>
</div>

<style>
	h2 {
		text-align: center;
		margin-bottom: 1em;
	}

	.view-more {
		display: flex;
		justify-content: center;
		text-decoration: none;
		font-size: calc(var(--font-size) - 2pt);
		cursor: pointer;
		color: var(--fg-color);
		background-color: var(--bg-color);
		padding: 1em;
		border-radius: var(--radius-size);
	}

	.view-more:hover {
		background: var(--gradient-hover-color);
		color: var(--fg-color);
	}

	.project > div {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-content: center;
		gap: 1em;
	}
</style>
