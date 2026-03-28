<script lang="ts">
	import { onMount } from 'svelte';
	import type { HTMLSelectAttributes } from 'svelte/elements';

	let {
		options,
		onselect = () => null,
		placeholder = '',
		value = $bindable('')
	}: HTMLSelectAttributes & {
		options?: string[];
		onselect?: (value: string) => void;
	} = $props();
</script>

<select
	bind:value
	class="text-white bg-black px-2 rounded py-1 h-fit flex-1"
	{placeholder}
	onchange={() => onselect(value)}
>
	{#each options as opt}
		<option value={opt} class="">{opt}</option>
	{/each}
</select>
