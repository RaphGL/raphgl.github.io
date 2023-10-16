<script lang="ts">
	import Pagination from '$lib/Pagination.svelte';
	import ContentRect from '$lib/ContentRect.svelte';
	import ProjectItem from '$lib/Projects/ProjectItem.svelte';
	import { projects } from '$lib/stores';

	const itemPerPage = 9;
	let currPage = 0;
	let projectsPages: any = [];
	for (let i = 0; i < $projects.length; i++) {
		projectsPages.push($projects.slice(i * itemPerPage, (i + 1) * itemPerPage));
	}
</script>

<svelte:head>
	<title>RaphGL - Projects</title>
</svelte:head>

<ContentRect>
	<section>
		{#each projectsPages[currPage] as project}
			<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
				{project.name}
			</ProjectItem>
		{/each}
	</section>

	{#if projectsPages}
		<Pagination
			bind:current={currPage}
			numOfPages={Math.ceil(++projectsPages.length / itemPerPage)}
		/>
	{/if}
</ContentRect>

<style>
	section {
		display: flex;
		padding: 2em;
		flex-flow: row wrap;
		justify-content: center;
		gap: 1em;
	}

	h2 {
		text-align: center;
		font-weight: normal;
	}
</style>
