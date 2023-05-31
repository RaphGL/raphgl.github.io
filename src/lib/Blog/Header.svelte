<script>
	import { fade } from 'svelte/transition';
</script>

<div transition:fade>
	<h2>
		<slot />
	</h2>
</div>

<style>
	div {
		padding: 1em;
		text-align: center;
	}
</style>
