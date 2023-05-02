<div class="blog">
	<h2>Blog</h2>
	<ul>
		<li>
			<div class="blog-title">Poetry is a saner package manager</div>
			<div class="blog-date">10 Oct 2022</div>
		</li>
		<li>
			<div class="blog-title">Setting up NeoVim for Godot</div>
			<div class="blog-date">10 Oct 2023</div>
		</li>
	</ul>
	<button>View More</button>
</div>

<style>
	.blog {
		width: 30vw;;
	}
	.blog h2 {
		text-align: center;
	}

	.blog button {
		background: none;
		border: none;
		color: var(--fg-color);
		font-size: 14pt;
		text-decoration: underline;
		width: 100%;
		cursor: pointer;
	}

	.blog button:hover {
		color: var(--hover-color);
	}

	.blog ul {
		padding: 0;
		list-style: none;
	}

	.blog li {
		background: rgb(32, 33, 41);
		background: linear-gradient(310deg, rgba(32, 33, 41, 1) 0%, rgba(40, 42, 54, 1) 100%);
		padding: 1em;
		border-radius: var(--radius-size);
		margin-top: 1em;
		font-size: 15pt;
		display: flex;
		justify-content: space-between;
		gap: 4em;
	}

	.blog-date {
		font-size: 13pt;
	}
</style>