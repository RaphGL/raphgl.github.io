<script>
	import BlogList from '../lib/BlogList.svelte';
	import ProjectList from '../lib/ProjectList.svelte';
	import ContentRect from '../lib/ContentRect.svelte';
</script>

<section class="info">
	<div>
		<!-- svelte-ignore a11y-img-redundant-alt -->
		<img
			src="https://avatars.githubusercontent.com/u/28673457?s=400&u=bc6d2bccf7d9bc6b8d403c0d462a4bc49b0fc678&v=4"
			alt="Profile Picture"
		/>
		<div class="info-text">
			<div class="info-title">Hi, I'm <b>RaphGL</b></div>

			<ul>
				<li>
					<div style="background-color:white;height: 1em; width: 1em;border-radius:0.2em;" />
					<a href="https://github.com/RaphGL">Github</a>
				</li>
				<li>
					<div style="background-color:white;height: 1em; width: 1em;border-radius:0.2em;" />
					<a href="https://gitlab.com/RaphGL">Gitlab</a>
				</li>
				<li>
					<div style="background-color:white;height: 1em; width: 1em;border-radius:0.2em;" />
					<a href="https://codeberg.org/RaphGL">Codeberg</a>
				</li>
			</ul>
		</div>
	</div>
</section>

<div class="scroll">
	<div style="background-color:white;height: 1em; width: 1em;border-radius:0.2em;" />
	Scroll for more
</div>

<ContentRect>
	<div class="main-content">
		<BlogList
			blogs={[
				{ title: 'Poetry is a saner package manager for python', date: '10 Oct 2022', href: '/' },
				{ title: 'Setting up NeoVim for Godot', date: '11 Nov 2022', href: '/' }
			]}
		/>
		<ProjectList />
	</div>
</ContentRect>

<style>
	.main-content {
		display: flex;
		justify-content: center;
		padding: 3em 0;
		margin-top: 1em;
		gap: 5em;
	}

	.info {
		height: 100%;
		width: 100%;
		display: flex;
		justify-items: center;
		min-height: 80vh;
	}

	.info img {
		height: 30vh;
		border-radius: var(--radius-size);
	}

	.info > div {
		background-color: rgb(40, 42, 54);
		background-color: linear-gradient(310deg, rgba(40, 42, 54, 1) 0%, rgba(32, 33, 41, 1) 100%);
		border-radius: var(--radius-size);
		margin: auto;
		padding: 1.3em;
		display: flex;
	}

	.info ul {
		display: flex;
		gap: 1em;
	}

	.info-title {
		font-size: 30pt;
	}

	.info-text {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 0 1em;
		width: 100%;
	}

	.info-text a {
		color: var(--fg-color);
		text-decoration: none;
	}

	.info-text a:hover {
		color: var(--hover-color);
	}

	.info-text ul li {
		list-style: none;
		display: flex;
		gap: 0.5em;
		padding: 0.6em;
		border-radius: var(--radius-size);
		align-items: center;
	}

	.info-text li img {
		height: 3vh;
	}

	.scroll {
		display: flex;
		justify-content: center;
		flex-direction: column;
		align-items: center;
		font-size: 15pt;
		gap: 1em;
		cursor: default;
	}
</style>
