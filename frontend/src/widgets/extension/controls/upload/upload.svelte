<script lang="ts">
	import { Player } from '$go/services/extension';
	import { Input } from '$shared/ui';
	import { ChevronDown } from '@lucide/svelte';
	import { onMount } from 'svelte';

	let { onsubmit }: { onsubmit?: () => void } = $props();

	let value = $state('');

	onMount(() => {
		Player.State().then((state) => {
			value = state.Links.join('\n');
		});
	});

	async function submit() {
		await Player.Upload(value);
		onsubmit?.();
	}

	$effect(() => {
		value = value.replace(/\n{2,}/g, '\n'); // add new line at the end and cleans empty lines
	});
</script>

<div class="absolute inset-0 bg-black/60 p-6">
	<div class="flex gap-2 flex-col grow h-full bg-black text-white p-2">
		<ChevronDown onclick={submit} class="cursor-pointer hover:scale-105 ms-auto" />
		<Input type="textarea" bind:value placeholder="Links" class="" />
	</div>
</div>
