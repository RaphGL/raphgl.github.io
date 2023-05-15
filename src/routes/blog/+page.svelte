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

<style>
	section {
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 2em;
        gap: 1em;
	}
</style>
