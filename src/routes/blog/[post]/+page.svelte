<script lang="ts">
	import { onMount } from 'svelte';
	import { fly } from 'svelte/transition';
	import Header from '$lib/Blog/Header.svelte';
	import TableOfContents from '$lib/Blog/TableOfContents.svelte';
	import ContentRect from '$lib/ContentRect.svelte';

	let contents: NodeListOf<Element> | null = null;
	onMount(() => {
		contents = document.querySelectorAll(
			'.blog-text h1, .blog-text h2, .blog-text h3, .blog-text h4, .blog-text h5, .blog-text h6'
		);
	});
</script>

<Header>Setting up NeoVim for Godot</Header>
<ContentRect>
	<div class="content" transition:fly>
		<TableOfContents {contents} />
		<div class="blog-text">
			<h1>How to do it</h1>
			<h2>Why I'm doing it</h2>
			<h2 id="something-else">Why I'm doing it 2</h2>
			<h2>Yeah</h2>
			<h3>Another subsection</h3>
			<h2 id="conclusion">Conclusion</h2>
			Recently I’ve been using Godot a bit to mess around with game development and I wanted to have
			autocompletion working with CoC. I looked around and there were no extensions for the Godot LSP.
			So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up If you have CoC working
			correctly already, open NeoVim and run :CocConfig, it’ll open the configuration file for your CoC,
			then just add this to the JSON config: Additional steps You probably want syntax highlighting,
			for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor Recently I’ve been using Godot a bit to mess around with game development
			and I wanted to have autocompletion working with CoC. I looked around and there were no extensions
			for the Godot LSP. So, to use Godot with NeoVim I just had to connect to their LSP. Setting Up
			If you have CoC working correctly already, open NeoVim and run :CocConfig, it’ll open the configuration
			file for your CoC, then just add this to the JSON config: Additional steps You probably want syntax
			highlighting, for that you have 2 options: you could install the sheerun/vim-polyglot or use habamax/vim-godot
			(which provides syntax highlighting as well as a set of commands allows you to run scenes through
			NeoVim). Setting up the editor
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
