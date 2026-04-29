<script lang="ts">
	import { Button } from '$shared/ui';
	import type { Snippet } from 'svelte';

	let { onclick, children }: { onclick?: () => void; children?: Snippet } = $props();
</script>

<Button class="px-1 py-0! rounded-lg hover:scale-105" {onclick}>{@render children?.()}</Button>
