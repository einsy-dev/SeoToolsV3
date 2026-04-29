<script lang="ts">
	import type { InputI } from '../input';
	import Header from './header/header.svelte';

	let { value = $bindable('') }: InputI['html'] = $props();

	let input: HTMLDivElement | undefined = $state();
</script>

<div class=" rounded-xl flex flex-col flex-1 p-4">
	<Header {input} bind:value />

	<div
		bind:this={input}
		bind:innerHTML={value}
		class=" p-4 border-none outline-none prose prose-slate [&>h1]:text-2xl [&>h2]:text-xl"
		contenteditable
	>
		{value}
	</div>
</div>
