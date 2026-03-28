<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';

	let {
		children,
		onclick,
		button = $bindable(),
		...props
	}: HTMLButtonAttributes & {
		button?: HTMLButtonElement;
		'bind:button'?: HTMLButtonElement;
	} = $props();
</script>

<button
	bind:this={button}
	class={`cursor-pointer card py-2! hover:scale-105 ${props.class}`}
	{onclick}
>
	{@render children?.()}
</button>
