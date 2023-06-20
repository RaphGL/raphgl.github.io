<script lang="ts">
	import BlogList from '../lib/Blog/BlogList.svelte';
	import ProjectList from '../lib/Projects/ProjectList.svelte';
	import ContentRect from '../lib/ContentRect.svelte';
</script>

<section class="info">
	<div>
		<img
			src="https://avatars.githubusercontent.com/u/28673457?s=400&u=bc6d2bccf7d9bc6b8d403c0d462a4bc49b0fc678&v=4"
			alt="Girl staring at the camera"
		/>
		<div class="info-text">
			<div class="info-title">Hi, I'm <b>RaphGL</b></div>

			<ul>
				<li>
					<a href="https://github.com/RaphGL">
						<svg xmlns="http://www.w3.org/2000/svg" height="0.875em" viewBox="0 0 496 512">
							<style>
								svg {
									fill: #deddda;
								}
							</style><path
								d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"
							/></svg
						>
						Github
					</a>
				</li>
				<li>
					<a href="https://gitlab.com/RaphGL">
						<svg xmlns="http://www.w3.org/2000/svg" height="0.875em" viewBox="0 0 512 512">
							<style>
								svg {
									fill: #deddda;
								}
							</style><path
								d="M105.2 24.9c-3.1-8.9-15.7-8.9-18.9 0L29.8 199.7h132c-.1 0-56.6-174.8-56.6-174.8zM.9 287.7c-2.6 8 .3 16.9 7.1 22l247.9 184-226.2-294zm160.8-88l94.3 294 94.3-294zm349.4 88l-28.8-88-226.3 294 247.9-184c6.9-5.1 9.7-14 7.2-22zM425.7 24.9c-3.1-8.9-15.7-8.9-18.9 0l-56.6 174.8h132z"
							/></svg
						>
						Gitlab
					</a>
				</li>
				<li>
					<a href="https://codeberg.org/RaphGL">
						<svg xmlns="http://www.w3.org/2000/svg" height="0.875em" viewBox="0 0 512 512">
							<style>
								svg {
									fill: #deddda;
								}
							</style><path
								d="M511.4 37.9C515.1 18.2 500 0 480 0H32C10.6 0-4.8 20.7 1.4 41.2l87.1 273.4c2.5 7.2 12.7 7.2 15.1 0L140 190.5l44.2 187.3c1.9 8.3 13.7 8.3 15.6 0l46.5-196.9 34.1 133.4c2.3 7.6 13 7.6 15.3 0l45.8-172.5 66.7 363.8c1.7 8.6 14 8.6 15.7 0l87.5-467.7z"
							/></svg
						>
						Codeberg
					</a>
				</li>
			</ul>
		</div>
	</div>
</section>

<div class="scroll">
	<svg xmlns="http://www.w3.org/2000/svg" height="1.25em" viewBox="0 0 320 512">
		<style>
			svg {
				fill: #ffffff;
			}
		</style><path
			d="M143 352.3L7 216.3c-9.4-9.4-9.4-24.6 0-33.9l22.6-22.6c9.4-9.4 24.6-9.4 33.9 0l96.4 96.4 96.4-96.4c9.4-9.4 24.6-9.4 33.9 0l22.6 22.6c9.4 9.4 9.4 24.6 0 33.9l-136 136c-9.2 9.4-24.4 9.4-33.8 0z"
		/></svg
	>
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
		flex-direction: column;
		justify-content: center;
		padding: 3em 0;
		margin: 1em;
		gap: 3em;
	}

	@media only screen and (min-width: 768px) {
	.main-content {
		flex-direction: row;
	}
	}

	.info {
		height: 90h;
		width: 90vw;
		display: flex;
		justify-items: center;
		min-height: 84vh;
		margin:auto;
	}

	.info img {
		margin-left: auto;
		margin-right: auto;
		width: 80vw;
		margin-bottom: 1em;
		border-radius: var(--radius-size);
	}

	@media only screen and (min-width: 768px) {
		.info img {
			width: auto;
			height: 12em;
			flex-direction: row;
		}
	}

	@media only screen and (min-width: 850px) {
		.info img {
			margin-bottom: 0;
		}
	}

	.info > div {
		display: flex;
		flex-direction: column;
		background-color: rgb(40, 42, 54);
		background-color: linear-gradient(310deg, rgba(40, 42, 54, 1) 0%, rgba(32, 33, 41, 1) 100%);
		border-radius: var(--radius-size);
		margin: auto;
		padding: 1em;
	}

	@media only screen and (min-width: 850px) {
		.info > div {
			flex-direction: row;
		}
	}

	.info ul {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
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
		gap: 1em;
		padding: 0 1em;
	}

	.info-text ul li {
		list-style: none;
	}

	.info-text li a {
		display: flex;
		color: var(--fg-color);
		text-decoration: none;
		gap: 0.5em;
		padding: 0.6em;
		border-radius: var(--radius-size);
		align-items: center;
		background-color: var(--bg-color);
	}

	.info-text li a:hover {
		color: var(--hover-color);
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
