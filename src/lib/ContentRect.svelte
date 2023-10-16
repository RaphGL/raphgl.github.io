<section class="main-content">
	<slot />
</section>

<style>
	.main-content {
		background-color: var(--content-bg-color);
		border-radius: 2em;
	}

	@media only screen and (min-width: 768px) {
		.main-content {
			margin: 0 1em;
		}
	}
</style>
