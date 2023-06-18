<script lang="ts">
	import Button from './Button.svelte';
	export let numOfPages = 1;
	let pages = Array.from(Array(numOfPages).keys());
	export let current = 0;

	function prevPage() {
		if (current > 0) {
			--current;
		}
	}

	function nextPage() {
		if (current < numOfPages - 1) {
			++current;
		}
	}
</script>

<div class="pagination">
	<Button on:click={prevPage}>
		<svg xmlns="http://www.w3.org/2000/svg" height="0.875em" viewBox="0 0 320 512"
			><!--! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><!--! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
			<style>
				svg {
					fill: #deddda;
				}
			</style><path
				d="M41.4 233.4c-12.5 12.5-12.5 32.8 0 45.3l160 160c12.5 12.5 32.8 12.5 45.3 0s12.5-32.8 0-45.3L109.3 256 246.6 118.6c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0l-160 160z"
			/></svg
		>
	</Button>

	{#each pages as page}
		<Button active={current == page} on:click={() => (current = page)}>{page + 1}</Button>
	{/each}

	<Button on:click={nextPage}>
		<svg xmlns="http://www.w3.org/2000/svg" height="0.875em" viewBox="0 0 320 512"
			><!--! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><!--! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
			<style>
				svg {
					fill: #deddda;
				}
			</style><path
				d="M278.6 233.4c12.5 12.5 12.5 32.8 0 45.3l-160 160c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3L210.7 256 73.4 118.6c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0l160 160z"
			/></svg
		>
	</Button>
</div>

<style>
	.pagination {
		display: flex;
		justify-content: center;
		padding-bottom: 2em;
	}
</style>
