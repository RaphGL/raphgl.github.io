<script lang="ts">
	export let contents: NodeListOf<Element> | null;
</script>

<div class="table-of-contents">
	<h3>Table of Contents</h3>
	<div>
		<ul>
			{#if contents != null}
				{#each contents as content}
					<li class={content.tagName}>
						<a href="#{content.id}">{content.innerHTML}</a>
					</li>
				{/each}
			{/if}
		</ul>
	</div>
</div>

<style>
	.table-of-contents {
		background-color: #141519;
		padding: 1em;
		width: 100vw;
		border-top-left-radius: 2em;
		border-bottom-left-radius: 2em;
	}

	.table-of-contents h3 {
		text-align: center;
	}

	div ul {
		min-width: 20vw;
		max-width: 30vw;
	}

	ul li {
		list-style: none;
		font-size: 0.9em;
		margin-bottom: 0.2em;
	}

	li a {
		color: var(--fg-color);
		text-decoration: none;
	}

	li a:hover {
		color: var(--hover-color);
	}

	.H2 {
		margin-left: 2em;
	}

	.H3 {
		margin-left: 4em;
	}

	.H4 {
		margin-left: 6em;
	}

	.H5 {
		margin-left: 8em;
	}

	.H6 {
		margin-left: 1em;
	}
</style>
