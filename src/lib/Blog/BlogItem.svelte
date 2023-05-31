<script lang="ts">
	import { fly } from 'svelte/transition';
	export let href: string;
	export let date: string;
</script>

<div transition:fly={{ y: 100, duration: 100 }}>
	<a {href}>
		<div class="blog-title"><slot /></div>
		<div class="blog-date">{date}</div>
	</a>
</div>

<style>
	.blog-date {
		font-size: 13pt;
	}

	a {
		width: 35vw;
		color: var(--fg-color);
		text-decoration: unset;
		background: rgb(32, 33, 41);
		background: var(--gradient-fg-color);
		padding: 1em;
		border-radius: var(--radius-size);
		font-size: 15pt;
		display: flex;
		justify-content: space-between;
	}

	a:hover {
		background: var(--gradient-hover-color);
		cursor: pointer;
	}
</style>
