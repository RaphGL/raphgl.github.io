<script>
	import ProjectList from '$lib/Projects/ProjectItem.svelte';
	import ContentRect from '$lib/ContentRect.svelte';
	import ProjectItem from '$lib/Projects/ProjectItem.svelte';
	import { getProjects } from '../../lib/Projects/projects';
</script>

<ContentRect>
	<section>
		{#await getProjects()}
			<div>Retrieving projects...</div>
		{:then projects}
			{#each projects as project}
				<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
					{project.name}
				</ProjectItem>
			{/each}
		{/await}
	</section>
</ContentRect>

<style>
	section {
		display: flex;
		padding: 2em;
		flex-flow: row wrap;
		align-content: center;
		justify-content: center;
		gap: 1em;
	}
</style>
