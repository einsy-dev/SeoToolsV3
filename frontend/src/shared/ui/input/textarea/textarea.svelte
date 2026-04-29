<script lang="ts">
	import type { HTMLTextareaAttributes } from 'svelte/elements';

	let { value = $bindable(''), class: className, ...props }: HTMLTextareaAttributes = $props();
</script>

<textarea
	class="flex grow overflow-scroll resize-none outline-none scrollbar {className}"
	bind:value
	{...props}
></textarea>
