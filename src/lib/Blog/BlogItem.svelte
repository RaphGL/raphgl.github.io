<script lang="ts">
	import { fade } from 'svelte/transition';

	export let href: string;
	export let date: string;
</script>

<div in:fade={{ duration: 150 }}>
	<a {href}>
		<div class="blog-title"><slot /></div>
		<div class="blog-date">{date}</div>
	</a>
</div>

<style>
	.blog-date {
		font-size: calc(var(--font-size) - 2pt);
	}

	a {
		color: var(--fg-color);
		text-decoration: unset;
		background: rgb(32, 33, 41);
		background: var(--gradient-bg-color);
		padding: 1em;
		margin: auto;
		border-radius: var(--radius-size);
		display: flex;
		justify-content: space-between;
	}

	@media only screen and (min-width: 768px) {
		a {
			width: 80vw;
		}
	}

	@media only screen and (min-width: 850px) {
		a {
			width: 35vw;
		}
	}

	a:hover {
		background: var(--gradient-hover-color);
		cursor: pointer;
	}
</style>
