<section class="main-content">
	<slot />
</section>

<style>
	.main-content {
		background-color: #18191f;
		border-radius: 2em;
	}

	@media only screen and (min-width: 768px) {
		.main-content {
			margin: 1em;
		}
	}
</style>
