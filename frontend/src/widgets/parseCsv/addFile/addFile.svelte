<script lang="ts">
	import { FilePlus } from '@lucide/svelte';

	let { files = $bindable([]) }: { files?: File[]; 'bind:files'?: File[] } = $props();

	let active: boolean = $state(false);

	$effect(() => {
		if (files.length == 0) {
			active = true;
		}
	});

	function handleInput(e: Event) {
		if (!e.target) return;
		files = [...files, ...Array.from((e.target as HTMLInputElement).files!)];
		active = false;
	}
</script>

<svelte:window
	ondragenter={(e) => {
		e.preventDefault();
		active = true;
	}}
	ondragover={(e) => {
		e.preventDefault();
		// Explicitly tell the browser we want to copy files
		if (e.dataTransfer) e.dataTransfer.dropEffect = 'copy';
		active = true;
	}}
	ondragleave={(e) => {
		e.preventDefault();
		// Only set to false if we are actually leaving the window
		if (!e.relatedTarget) active = false;
	}}
	ondrop={(e) => {
		e.preventDefault();
		active = false;

		if (e.dataTransfer?.files) {
			// Svelte 5 runes check: ensure 'files' is a reactive array
			files = [...files, ...Array.from(e.dataTransfer.files)];
			console.log('Files dropped:', files);
		}
	}}
/>

<label class="h-full p-2" hidden={!active}>
	<div class="h-full border rounded bg-[rgba(60,60,60,1)]">
		<input type="file" multiple hidden onchange={handleInput} />
		<div
			class="flex flex-col gap-2 border rounded h-full w-full items-center justify-center cursor-pointer text-white"
		>
			<FilePlus class="w-10 h-10 " />
			<span>Choose a file or drag it over here</span>
		</div>
	</div>
</label>
