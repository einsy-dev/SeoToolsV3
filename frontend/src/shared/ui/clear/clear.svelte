<script lang="ts">
	import { Trash } from '@lucide/svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	let { class: className, ...props }: HTMLButtonAttributes = $props();
</script>

<button
	class="rounded bg-black text-white p-1 cursor-pointer hover:scale-103 {className}"
	{...props}
>
	<Trash class="h-6 w-6" />
</button>
