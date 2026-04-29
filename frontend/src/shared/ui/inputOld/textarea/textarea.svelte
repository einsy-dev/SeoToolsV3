<script lang="ts">
	import type { HTMLTextareaAttributes } from 'svelte/elements';

	let { value = $bindable(''), placeholder = '', ...props }: HTMLTextareaAttributes = $props();
</script>

<textarea
	class="h-full w-full overflow-scroll text-white text-lg resize-none outline-none text-nowrap scrollbar bg-black rounded p-2"
	bind:value
	{placeholder}
	{...props}
></textarea>
