<script>
	import Button from '$lib/Button.svelte';
	import ProjectList from '$lib/Projects/ProjectItem.svelte';
	import ContentRect from '$lib/ContentRect.svelte';
	import ProjectItem from '$lib/Projects/ProjectItem.svelte';
	import { getProjects } from '../../lib/Projects/projects';

	const itemPerPage = 6;
	let currPage = 1;
</script>

<ContentRect>
	<section>
		{#await getProjects()}
			<div>Retrieving projects...</div>
		{:then projects}
			{#each projects.slice(currPage * itemPerPage, (currPage + 1) * itemPerPage) as project}
				<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
					{project.name}
				</ProjectItem>
			{/each}
		{/await}
	</section>
		<Button on:click={() => console.log("clicked")}>hello</Button>
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
