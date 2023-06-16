<script>
	import ProjectItem from './ProjectItem.svelte';
	import { getProjects } from './projects';
</script>

<div class="project">
	<h2>Projects</h2>
	<div>
		{#await getProjects(5)}
			<ProjectItem>Retrieving projects...</ProjectItem>
		{:then projects}
			{#each projects as project}
				<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
					{project.name}
				</ProjectItem>
			{/each}
			{:catch}
			<ProjectItem>Failed to load Github data!</ProjectItem>
		{/await}
	</div>
	<a class="view-more" href="/projects">View More</a>
</div>

<style>
	.project {
		width: 25vw;
	}

	h2 {
		text-align: center;
	}

	.view-more {
		display: flex;
		justify-content: center;
		text-decoration: none;
		background: none;
		border: none;
		color: var(--fg-color);
		font-size: 14pt;
		cursor: pointer;
		height: fit-content;
	}

	.view-more:hover {
		color: var(--hover-color);
	}

	.project > div {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-content: center;
		gap: 1em;
		margin-bottom: 1em;
	}
</style>
