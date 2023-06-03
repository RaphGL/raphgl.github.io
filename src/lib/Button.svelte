<script lang="ts">
	export let href: string = '';
	export let active = false;
</script>

<div>
	<a {href} on:click class:active>
		<slot />
	</a>
</div>

<style>
	a {
		text-decoration: none;
		color: var(--fg-color);
		background-color: var(--bg-color);
		border-radius: var(--radius-size);
		padding: 0.5em 1em;
		margin: 0.1em;
	}

	a:hover {
		background: var(--hover-color);
		background: var(--gradient-hover-color);
	}

	.active {
		background: var(--hover-color);
		background: var(--gradient-hover-color);
	}
</style>
