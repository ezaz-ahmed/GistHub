<script>
	export let pokeman;
</script>

<a
	href={`/pokeman/${pokeman.id}`}
	class=" p-6 bg-gray-100 text-gray-800 rounded-md shadow-sm hover:shadow-md flex flex-col items-center"
>
	<img class=" h-40 w-40" src={pokeman.image} alt={pokeman.name} />

	<h2 class=" uppercase">{pokeman.id}. {pokeman.name}</h2>
</a>
