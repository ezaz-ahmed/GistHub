<script>
	import { pokemon } from '../stores/pokestore';
</script>

<svelte:head>
	<title>Pokedex</title>
</svelte:head>

<h1 class="text-4xl text-center my-8 uppercase">Pokedex</h1>

{#each $pokemon as pokeman}
	<img src={pokeman.image} alt={pokeman.name} />
{/each}
