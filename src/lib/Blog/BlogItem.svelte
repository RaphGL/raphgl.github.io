<script lang="ts">
	export let href: string;
	export let date: string;
</script>

<div>
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
		background: linear-gradient(310deg, rgba(32, 33, 41, 1) 0%, rgba(40, 42, 54, 1) 100%);
		padding: 1em;
		border-radius: var(--radius-size);
		font-size: 15pt;
		display: flex;
		justify-content: space-between;
	}

	a:hover {
		cursor: pointer;
	}
</style>
