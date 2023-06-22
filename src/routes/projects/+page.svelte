<script lang="ts">
	import Pagination from '$lib/Pagination.svelte';
	import ContentRect from '$lib/ContentRect.svelte';
	import ProjectItem from '$lib/Projects/ProjectItem.svelte';
	import { getProjects } from '$lib/Projects/projects';
	import { onMount } from 'svelte';

	const itemPerPage = 9;
	let currPage = 0;
	let projectsPages: any = null;

	onMount(() => {
		getProjects().then((projs) => {
			projectsPages = [];
			for (let i = 0; i < projs.length; i++) {
				projectsPages.push(projs.slice(i * itemPerPage, (i + 1) * itemPerPage));
			}
		});
	});
</script>

<svelte:head>
	<title>RaphGL - Projects</title>
</svelte:head>

<ContentRect>
	<section>
		{#if !projectsPages}
			<div>Retrieving projects...</div>
		{:else}
			{#each projectsPages[currPage] as project}
				<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
					{project.name}
				</ProjectItem>
			{/each}
		{/if}
	</section>

	{#if projectsPages}
		<Pagination bind:current={currPage} numOfPages={Math.ceil(++projectsPages.length / itemPerPage)} />
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
</style>
