<script lang="ts">
	import type { LinkI } from '../types';
	import LinkItem from './linkItem/linkItem.svelte';

	let { links }: { links?: LinkI[] } = $props();
</script>

<div class="flex flex-col gap-1">
	{#each links as link}
		<LinkItem {...link} />
	{/each}
</div>
