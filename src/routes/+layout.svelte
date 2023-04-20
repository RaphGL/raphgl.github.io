<script>
	import '../app.css';
</script>

<nav>
	<a href="/">Home</a>
	<a href="/projects">Projects</a>
	<a href="/blog">Blog</a>
</nav>

<slot />

<style>
	nav {
		display: flex;
		justify-content: center;
		gap: 2em;
		padding: 1em 0 1em 0;
	}

	nav a {
		text-decoration: none;
		color: var(--fg-color);
        width: 4em;
        text-align: center;
		transition: font-weight 0.1s ease-out;
	}

	nav a:hover {
		font-weight: bold;
	}
</style>
