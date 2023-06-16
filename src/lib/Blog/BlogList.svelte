<script lang="ts">
	import BlogItem from '../Blog/BlogItem.svelte';

	type Blog = {
		title: string;
		date: string;
		href: string;
	};

	export let blogs: Blog[];
</script>

<div class="blog">
	<h2>Blog</h2>
	<div class="blog-posts">
		{#each blogs as blog}
			<BlogItem date={blog.date} href={blog.href}>{blog.title}</BlogItem>
		{/each}
	</div>
	<a href="/blog">View More</a>
</div>

<style>
	.blog h2 {
		text-align: center;
	}

	a {
		background: none;
		border: none;
		color: var(--fg-color);
		font-size: 14pt;
		display: flex;
		justify-content: center;
		cursor: pointer;
		text-decoration: unset;
	}

	a:hover {
		color: var(--hover-color);
	}

	.blog-posts {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		gap: 1em;
		margin-bottom: 1em;
	}
</style>
