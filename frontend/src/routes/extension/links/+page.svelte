<script lang="ts">
	import { Input } from '$shared/ui';
	import { LinksList } from '$widgets/extension';

	let files = $state([]);
</script>

<Input type="file" class="flex grow" bind:value={files} multiple>
	<div class="relative flex grow flex-col text-white px-2 py-1 gap-1">
		<LinksList links={[{ linkPage: 'majesticFormat', targetUrl: 'google.com' }]} />
	</div>
</Input>
