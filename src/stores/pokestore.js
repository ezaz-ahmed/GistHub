import { writable } from 'svelte/store';

export const pokemon = writable([]);
const pokemonDetails = {};
let loaded = false;


export const fetchPokemon = async (num) => {
	if (loaded) return;
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
	loaded = true;
};


export const getPokemonById = async (id) => {
	if (pokemonDetails[id]) return pokemonDetails[id];

	try {
		const url = `https://pokeapi.co/api/v2/pokemon/${id}`;
		const res = await fetch(url);
		const data = await res.json();
		pokemonDetails[id] = data;
		return data;
	} catch (err) {
		console.error(err);
		return null;
	}
};