<script lang="ts">
	// import * as Player from "$go/controls/Player";
	import { Button, Textarea } from '$shared/ui';
	import { onMount } from 'svelte';

	let value = $state('');

	onMount(() => {
		// Player.State().then((state) => {
		//   value = state.Links.join("\n");
		// });
	});

	async function submit() {
		// await Player.Upload(value);
	}

	$effect(() => {
		value = value.replace(/\n{2,}/g, '\n'); // add new line at the end and cleans empty lines
	});
</script>

<div class="flex h-full gap-2 p-4 flex-col">
	<Textarea bind:value placeholder="Links" />
	<Button onclick={submit} class="card py-2! hover:scale-105 cursor-pointer">Submit</Button>
</div>
