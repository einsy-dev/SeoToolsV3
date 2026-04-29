<script lang="ts">
	import { Clipboard } from '$go/services/clipboard';

	let { domain = '' } = $props();
</script>

<div class="flex items-center justify-between">
	<div
		role="none"
		class="cursor-pointer hover:scale-103"
		onclick={() => {
			Clipboard.CopyText(domain);
		}}
	>
		{domain}
	</div>
</div>
