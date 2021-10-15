<script>
	import Header from '$lib/header/Header.svelte';
</script>

<Header />

<main>
	<slot />
</main>

<footer>
	<p>Created With Love 💓 By <a href="https://github.com/ezaz-ahmed">Ezaz Ahmed</a></p>
</footer>

<style>
	/* ./styles/globals.css */
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
