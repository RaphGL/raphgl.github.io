<script>
	import ProjectItem from './ProjectItem.svelte';
</script>

<div class="project">
	<h2>Projects</h2>
	<div>
		<ProjectItem href="/" />
		<ProjectItem href="/" />
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
		text-decoration: underline;
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
