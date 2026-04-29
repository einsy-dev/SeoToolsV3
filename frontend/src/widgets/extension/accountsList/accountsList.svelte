<script lang="ts">
	import type { AccountI } from '../types';
	import AccountItem from './accountItem/accountItem.svelte';

	let { accounts }: { accounts: Partial<AccountI>[] } = $props();
</script>

<div class="flex flex-col gap-1">
	{#each accounts as item}
		<AccountItem {...item} />
	{/each}
</div>
