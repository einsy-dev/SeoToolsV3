<script lang="ts">
	import { Button } from '$shared/ui';
	import { RotateCw, SkipBack, SkipForward, StepBack, StepForward } from '@lucide/svelte';
</script>

<div
	class="sticky bottom-0 mt-auto w-full flex flex-col gap-4 items-center bg-black py-2 px-2 text-white"
>
	<div class="flex gap-4">
		<Button class="border px-4 py-1 rounded-lg hover:scale-105 hidden sm:block" onclick={() => {}}
			><SkipBack /></Button
		>
		<Button class="border px-4 py-1 rounded-lg hover:scale-105" onclick={() => {}}
			><StepBack /></Button
		>
		<Button class="border px-4 py-1 rounded-lg hover:scale-105" onclick={() => {}}
			><RotateCw /></Button
		>
		<Button class="border px-4 py-1 rounded-lg hover:scale-105" onclick={() => {}}
			><StepForward /></Button
		>
		<Button class="border px-4 py-1 rounded-lg hover:scale-105 hidden sm:block" onclick={() => {}}
			><SkipForward /></Button
		>
	</div>
</div>
