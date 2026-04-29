<script lang="ts">
	import type { InputI } from '../input';
	import { FilePlus } from '@lucide/svelte';

	let {
		value = $bindable(),
		children,
		class: className,
		multiple = false
	}: InputI['file'] = $props();

	let active: boolean = $state(false);
	let dragCounter = $state(0);

	function handleFiles(fileList?: FileList) {
		if (!fileList) return;

		if (multiple) {
			if (!Array.isArray(value)) value = [];
			value = [...value, ...Array.from(fileList)];
		} else {
			value = fileList[0];
		}
		active = false;
		dragCounter = 0;
	}

	function handleInput(e: Event) {
		let target = e.target as HTMLInputElement;
		if (!target || !target.files) return;
		handleFiles(target.files);
	}
</script>

<svelte:window
	ondragenter={(e) => {
		e.preventDefault();
		dragCounter++;
		active = true;
	}}
	ondragleave={(e) => {
		e.preventDefault();
		dragCounter--;
		if (dragCounter <= 0) active = false;
	}}
	ondragover={(e) => {
		e.preventDefault();
	}}
	ondrop={(e) => {
		e.preventDefault();
		active = false;
		dragCounter = 0;
	}}
/>

<label
	ondrop={(e) => {
		e.preventDefault();
		active = false;
		handleFiles(e.dataTransfer?.files);
	}}
	class="cursor-pointer overflow-hidden text-white {className}"
>
	<div class="relative flex grow">
		<input type="file" multiple hidden onchange={handleInput} />
		{@render children?.()}
		<div
			class="absolute inset-0 z-50 flex items-center justify-center bg-black/50"
			hidden={!active}
		>
			<FilePlus class="h-30 w-30" />
		</div>
	</div>
</label>
