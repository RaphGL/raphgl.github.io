<script lang="ts">
	export let href: string = "";
</script>

<div on:click>
	<a {href}>
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
	}

    a:hover {
        background: var(--hover-color);
        background: var(--gradient-hover-color);
    }
</style>
