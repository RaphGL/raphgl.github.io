<script lang="ts">
	export let href: string;
</script>

<a {href}>
	<div class="project-name">Tuckr</div>
	<div class="project-description">A dotfile manager inspired by GNU Stow</div>
	<ul class="project-tags">
		<li>Rust</li>
		<li>Dotfiles</li>
		<li>Automation</li>
	</ul>
</a>

<style>
	a {
		all: unset;
		background: rgb(32, 33, 41);
		background: linear-gradient(310deg, rgba(32, 33, 41, 1) 0%, rgba(40, 42, 54, 1) 100%);
		border-radius: var(--radius-size);
		font-size: 15pt;
		display: flex;
		width: 25vw;
		flex-direction: column;
		text-align: center;
		justify-content: space-between;
		padding: 1em 0;
	}

	a:hover {
		cursor: pointer;
	}

	.project-name {
		font-weight: bold;
		font-size: 18pt;
	}

	.project-tags {
		display: flex;
		gap: 0.5em;
		justify-content: center;
	}

	ul {
		padding: 0;
		list-style: none;
	}

	.project-tags li {
		background: #383a59;
		padding: 0.2em 1em;
		font-size: 12pt;
		border-radius: var(--radius-size);
	}
</style>
