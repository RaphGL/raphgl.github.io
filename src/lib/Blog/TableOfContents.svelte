<script lang="ts">
	import {fly} from 'svelte/transition'
	export let contents: NodeListOf<Element> | null;
	let show = false;
</script>

<div class={show ? 'table-of-contents' : 'hidden-table'}>
	<div class="expand-btn">
		<div class="btn" on:click={() => (show = !show)} on:keydown>
			<svg xmlns="http://www.w3.org/2000/svg" height="1.25em" viewBox="0 0 448 512">
				<style>
					svg {
						fill: #deddda;
					}
				</style><path
					d="M16 132h416c8.837 0 16-7.163 16-16V76c0-8.837-7.163-16-16-16H16C7.163 60 0 67.163 0 76v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16z"
				/>
			</svg>
		</div>

		{#if show}
			<span>Table of Contents</span>
		{/if}
	</div>

	{#if show}
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
	{/if}
</div>

<style>
	.expand-btn {
		display: flex;
		gap: 2em;
		padding: 0.5em;
	}

	.expand-btn span {
		font-weight: bold;
		font-size: 20pt;
		cursor: default;
	}

	.expand-btn .btn {
		cursor: pointer;
	}

	.hidden-table {
		background-color: #141519;
		padding: 1em;
		width: 30vw;
		border-top-left-radius: 2em;
		border-bottom-left-radius: 2em;
	}

	.table-of-contents {
		background-color: #141519;
		padding: 1em;
		width: 100vw;
		border-top-left-radius: 2em;
		border-bottom-left-radius: 2em;
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
		font-weight: bold;
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
