<script lang="ts">
	import Pagination from '$lib/Pagination.svelte';
	import Button from '$lib/Button.svelte';
	import ContentRect from '$lib/ContentRect.svelte';
	import ProjectItem from '$lib/Projects/ProjectItem.svelte';
	import { getProjects } from '../../lib/Projects/projects';
	import { onMount } from 'svelte';

	const itemPerPage = 9;
	let currPage = 0;
	let projects: any = null;

	onMount(() => {
		getProjects().then((projs) => {
			projects = projs;
			console.log(projects);
		});
	});
</script>

<ContentRect>
	<section>
		{#if !projects}
			<div>Retrieving projects...</div>
		{:else}
			{#each projects.slice(currPage * itemPerPage, (currPage + 1) * itemPerPage) as project}
				<ProjectItem href={project.href} desc={project.description} tags={project.tags}>
					{project.name}
				</ProjectItem>
			{/each}
		{/if}
	</section>

	{#if projects}
		<Pagination bind:current={currPage} numOfPages={Math.ceil(++projects.length / itemPerPage)} />
	{/if}
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
