<script lang="ts">
	import { Player } from '$go/services/extension';
	import Button from './button/Button.svelte';
	import Domain from './domain/domain.svelte';
	import Reload from './reload/reload.svelte';
	import Upload from './upload/upload.svelte';
	import { ClipboardList, SkipBack, SkipForward, StepBack, StepForward } from '@lucide/svelte';

	let data = $state({ domain: 'google.com' });
	let uploadActive = $state(false);
</script>

{#if uploadActive}
	<Upload onsubmit={() => (uploadActive = false)} />
{/if}

<div
	class="sticky bottom-0 mt-auto w-full flex flex-col items-center bg-black text-white py-1 z-50"
>
	<Domain domain={data.domain} />

	<div class="flex items-center justify-between w-full">
		<div class="w-1/7 flex items-center justify-center">
			<Button
				onclick={() => {
					uploadActive = !uploadActive;
				}}><ClipboardList /></Button
			>
		</div>
		<div class="flex gap-2">
			<Button
				onclick={() => {
					Player.First();
				}}><SkipBack /></Button
			>
			<Button
				onclick={() => {
					Player.Prev();
				}}><StepBack /></Button
			>
			<Reload />
			<Button
				onclick={() => {
					Player.Next();
				}}><StepForward /></Button
			>
			<Button
				onclick={() => {
					Player.Last();
				}}><SkipForward /></Button
			>
		</div>
		<div class="w-1/7"></div>
	</div>
</div>
