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
	<Button on:click={prevPage}>&lt;</Button>
	{#each pages as page}
		<Button active={current == page} on:click={() => (current = page)}>{page + 1}</Button>
	{/each}
	<Button on:click={nextPage}>&gt;</Button>
</div>

<style>
	.pagination {
		display: flex;
		justify-content: center;
		padding: 1em;
	}
</style>
