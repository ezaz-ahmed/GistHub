<script>
	import Header from '$lib/header/Header.svelte';
</script>

<main class="p-8 max-w-6xl mx-auto">
	<Header />
	<slot />
</main>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
