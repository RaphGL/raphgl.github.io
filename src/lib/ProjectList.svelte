<div class="project">
	<h2>Projects</h2>
	<div>
		<a href="/">
			<div class="project-name">Tuckr</div>
			<div class="project-description">A dotfile manager inspired by GNU Stow</div>
			<ul class="project-tags">
				<li>Rust</li>
				<li>Dotfiles</li>
				<li>Automation</li>
			</ul>
		</a>
		<a href="/">
			<div class="project-name">NeoSent</div>
			<div class="project-description">Suckful Sent - Minimal presentation tool</div>
			<ul class="project-tags">
				<li>Rust</li>
				<li>Dotfiles</li>
				<li>Automation</li>
			</ul>
		</a>
	</div>
	<button>View More</button>
</div>

<style>
	.project {
		width: 25vw;
	}

	.project-name {
		font-weight: bold;
		font-size: 18pt;
	}

	.project h2 {
		text-align: center;
	}

	.project button {
		background: none;
		border: none;
		color: var(--fg-color);
		font-size: 14pt;
		width: 100%;
		cursor: pointer;
	}

	.project button:hover {
		text-decoration: underline;
	}

	.project button:hover {
		color: var(--hover-color);
	}

	.project ul {
		padding: 0;
		list-style: none;
	}

	.project > div {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-content: center;
		text-align: center;
	}

	.project a {
		all: unset;
		background: rgb(32, 33, 41);
		background: linear-gradient(310deg, rgba(32, 33, 41, 1) 0%, rgba(40, 42, 54, 1) 100%);
		border-radius: var(--radius-size);
		font-size: 15pt;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
		padding: 1em 0;
		margin-bottom: 1em;
	}

	.project a:hover {
		cursor: pointer;
	}

	.project-tags {
		display: flex;
		gap: 0.5em;
		justify-content: center;
	}

	.project-tags li {
		background: #383a59;
		padding: 0.2em 1em;
		font-size: 12pt;
		border-radius: var(--radius-size);
	}
</style>
