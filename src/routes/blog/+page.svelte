<script>
	import Header from '../../lib/Blog/Header.svelte';
	import TableOfContents from '../../lib/Blog/TableOfContents.svelte';
	import ContentRect from '../../lib/ContentRect.svelte';
</script>

<Header>Setting up NeoVim for Godot</Header>
<ContentRect>
	<div class="content">
		<TableOfContents>
			<ul>
				<li>1. asdfasdfasdfasdf</li>
				<li>2. asdfasdfasdfasdf</li>
				<li>3. asdfasdfasdfasdf</li>
				<li>4. asdfasdfasdfasdf</li>
			</ul>
		</TableOfContents>
		<div class="blog-text">
			Recently I’ve been using Godot a bit to mess around with game development and I wanted to have
			autocompletion working with CoC. I looked around and there were no extensions for the Godot
			LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up If you have
			CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want
			syntax highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot
			or use habamax/vim-godot (which provides syntax highlighting as well as a set of commands
			allows you to run scenes through NeoVim). Setting up the editor
		</div>
	</div>
</ContentRect>

<style>
	.content {
		display: flex;
		gap: 2em;
	}

	.blog-text {
		padding: 3em;
	}
</style>
