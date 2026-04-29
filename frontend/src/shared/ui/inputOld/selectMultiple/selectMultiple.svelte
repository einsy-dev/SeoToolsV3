<script lang="ts">
	import Input from './input/input.svelte';
	import { onMount } from 'svelte';
	import type { HTMLSelectAttributes } from 'svelte/elements';

	let {
		options = [],
		onselect = () => null,
		value = $bindable([]),
		edit = $bindable(false)
	}: HTMLSelectAttributes & {
		options?: string[];
		onselect?: (value: string[]) => void;
		edit?: boolean;
		'bind:edit'?: boolean;
	} = $props();

	let active = $state(false);
	let input: { focus: () => void };
	let inputValue = $state('');

	//filer search use regexp
	let filtered = $derived(
		options.filter((el: string) => !value.includes(el) && el.includes(inputValue))
	);

	let select: HTMLElement;

	onMount(() => {
		select.addEventListener('click', () => {
			input.focus();
			edit = true;
			active = true;
		});
	});

	function handleClick(e: MouseEvent) {
		if (!select || !e.target) return;
		if (!select.contains(e.target as Node)) {
			active = false;
		}
	}
</script>

<svelte:window onclick={handleClick} />

<div class="bg-black rounded px-3 py-1" bind:this={select}>
	<div class="flex flex-wrap gap-1">
		{#each value as v}
			<div class="px-1 bg-gray-500 rounded">{v}</div>
		{/each}
		<Input bind:this={input} />
	</div>

	{#if active}
		<div class="flex flex-col *:text-start">
			{#each filtered as opt}
				<button
					class="w-full hover:bg-gray-500 rounded px-2"
					onclick={() => {
						value.push(opt);
						inputValue = '';
					}}
				>
					{opt}
				</button>
			{/each}
		</div>
	{/if}
</div>
<!-- for form event -->
<select bind:value class="hidden" multiple onchange={() => onselect(value)}>
	{#each options as opt}
		<option value={opt} class="">{opt}</option>
	{/each}
</select>
