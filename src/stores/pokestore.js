import { writable } from 'svelte/store';

export const pokemon = writable([]);

const fetchPokemon = async (num) => {
	const url = `https://pokeapi.co/api/v2/pokemon?limit=${num}`;
	const res = await fetch(url);
	const data = await res.json();

	const loadedpokemon =
		data.results &&
		data.results.map((data, index) => {
			return {
				id: index + 1,
				name: data.name,
				image: `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/${
					index + 1
				}.png`
			};
		});
	pokemon.set(loadedpokemon);
};

fetchPokemon(150);
