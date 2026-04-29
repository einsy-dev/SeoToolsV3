<script lang="ts">
	import { Clipboard } from '$go/services/clipboard';
	import { ClipboardCopy } from '@lucide/svelte';

	let { value = '' } = $props();
</script>

<button
	onclick={() => Clipboard.CopyText(value)}
	class="rounded bg-black text-white p-1 cursor-pointer hover:scale-103"
>
	<ClipboardCopy class="h-6 w-6" />
</button>
