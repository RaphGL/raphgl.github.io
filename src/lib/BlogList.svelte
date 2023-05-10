<script lang="ts">
	type Blog = {
		title: string,
		date: string,
		href: string,
	};

	export let blogs: Blog[];
</script>

<div class="blog">
	<h2>Blog</h2>
	<div>
		{#each blogs as blog}
		<a href={blog.href}>
			<div class="blog-title">{blog.title}</div>
			<div class="blog-date">{blog.date}</div>
		</a>
		{/each}
	</div>
	<button>View More</button>
</div>

<style>
	.blog {
		width: 35vw;
	}
	.blog h2 {
		text-align: center;
	}

	.blog button {
		background: none;
		border: none;
		color: var(--fg-color);
		font-size: 14pt;
		width: 100%;
		cursor: pointer;
	}

	.blog button:hover {
		text-decoration: underline;
	}

	.blog button:hover {
		color: var(--hover-color);
	}

	a {
		all: unset;
	}

	.blog > div {
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}

	div a {
		background: rgb(32, 33, 41);
		background: linear-gradient(310deg, rgba(32, 33, 41, 1) 0%, rgba(40, 42, 54, 1) 100%);
		padding: 1em;
		border-radius: var(--radius-size);
		margin-bottom: 1em;
		font-size: 15pt;
		display: flex;
		justify-content: space-between;
	}

	.blog a:hover {
		cursor: pointer;
	}

	.blog-date {
		font-size: 13pt;
	}
</style>
