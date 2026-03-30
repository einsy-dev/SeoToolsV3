<script lang="ts">
	import { clickOutside } from '$shared/actions/clickOutside';
	import Input from './input.svelte';
	import Option from './option.svelte';
	import Selected from './selected.svelte';
	import { ChevronDown } from '@lucide/svelte';
	import type { HTMLSelectAttributes } from 'svelte/elements';

	let {
		options = [],
		placeholder = '',
		value = $bindable([]),
		class: className
	}: HTMLSelectAttributes & {
		options?: string[];
		placeholder?: string;
		value?: string[];
		'bind:value'?: string[];
	} = $props();

	let active = $state(false);
	let search = $state('');
	let filtered = $derived(
		options.filter((el) => !value.includes(el) && el.toLowerCase().includes(search.toLowerCase()))
	);

	// svelte-ignore non_reactive_update
	let input: HTMLInputElement;

	function onkeydown(e: KeyboardEvent) {
		if (!active) return;
		input.focus();
		if (e.key == 'Backspace' && search == '') {
			value = value.slice(0, -1);
		} else if (e.key == 'Enter') {
			active = false;
			input.blur();
			search = '';
		}
	}
</script>

<svelte:window {onkeydown} />

<div
	class="relative text-white *px-2 *py-1 h-full flex-1 *:bg-black {className}"
	use:clickOutside
	onclick_outside={() => {
		active = false;
	}}
>
	<div
		role="none"
		class="flex items-center h-full w-full justify-between"
		onclick={() => {
			active = true;
		}}
	>
		<div class="flex gap-1 flex-wrap max-h-19 overflow-scroll">
			{#each value as v}
				<Selected
					value={v}
					ondelete={() => {
						value = (value as string[]).filter((el) => el != v);
					}}
				/>
			{/each}
			<Input bind:value={search} bind:input />
		</div>
		<button
			onclick={(e) => {
				e.stopPropagation();
				active = !active;
			}}
		>
			<ChevronDown />
		</button>
	</div>
	<div class="absolute top-full left-0 w-full z-50 max-h-50 overflow-scroll">
		{#if active}
			{#each filtered as opt}
				<Option
					value={opt}
					onclick={() => {
						value.push(opt);
					}}
				/>
			{/each}
		{/if}
	</div>
</div>
