<script>
	import BlogItem from '$lib/Blog/BlogItem.svelte';
	import ContentRect from '$lib/ContentRect.svelte';
</script>

<ContentRect>
	<section>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
		<BlogItem href="/" date="10 Nov 2022">asdfadsfadsf</BlogItem>
	</section>
</ContentRect>

<svelte:head>
	<title>RaphGL - Blog</title>
</svelte:head>

<style>
	section {
		display: flex;
		flex-direction: column;
		padding: 2em;
		gap: 1em;
	}

	@media only screen and (min-width: 768px) {
		section {
			align-items: center;
		}
	}
</style>
