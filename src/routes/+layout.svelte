<script>
	import '../app.css';
	import NavBar from '$lib/NavBar.svelte';
</script>

<svelte:head>
	<title>RaphGL</title>
</svelte:head>

<NavBar />

<slot />

<footer>
	<b>RaphGL</b> @2023
</footer>

<style>
	footer {
		text-align: center;
		font-size: calc(var(--font-size) - 3pt);
		margin: 1em;
	}
</style>
