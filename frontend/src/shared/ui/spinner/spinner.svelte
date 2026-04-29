<script lang="ts">
	import { CircleCheck, LoaderCircle } from '@lucide/svelte';
	import type { Snippet } from 'svelte';

	let {
		loading = true,
		children,
		class: className
	}: { loading?: boolean; children?: Snippet; class?: string } = $props();
</script>

{#if loading}
	<LoaderCircle class="animate-spin {className}" />
{:else}
	{@render children?.()}
{/if}
