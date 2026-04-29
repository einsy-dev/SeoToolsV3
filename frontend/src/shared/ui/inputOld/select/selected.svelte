<script lang="ts">
	import { X } from '@lucide/svelte';

	let { value, ondelete }: { value: string; ondelete: () => void } = $props();
</script>

<div
	role="none"
	class="rounded border text-sm flex items-center justify-center gap-1"
	onclick={(e) => e.stopPropagation()}
>
	<div class="ps-2">
		{value}
	</div>
	<button onclick={ondelete}>
		<X class="w-3 h-3" />
	</button>
</div>
