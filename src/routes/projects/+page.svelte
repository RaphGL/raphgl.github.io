<script>
	import ProjectList from '../../lib/Projects/ProjectItem.svelte';
	import ContentRect from '../../lib/ContentRect.svelte';
	import ProjectItem from '../../lib/Projects/ProjectItem.svelte';
</script>

<ContentRect>
	<section>
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
		<ProjectItem href="/" />
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
