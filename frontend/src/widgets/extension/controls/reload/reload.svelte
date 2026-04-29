<script lang="ts">
	import { Player } from '$go/services/extension';
	import { Button, Spinner } from '$shared/ui';
	import { RotateCw } from '@lucide/svelte';

	const { loading = false }: { loading?: boolean } = $props();
</script>

<Spinner {loading}>
	<Button
		onclick={() => {
			Player.Reload();
		}}><RotateCw /></Button
	>
</Spinner>
