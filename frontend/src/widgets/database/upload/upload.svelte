<script lang="ts">
	import { Input, Button } from '$shared/ui';
	import { parseCSV } from '$shared/utils';

	let value: string = $state('');
	let data = $state({});

	async function ondrop(e: any) {
		e.preventDefault();
		let parsed = await parseCSV(e.dataTransfer.files[0]);
		value = JSON.stringify(parsed)
			.replaceAll('},{', '},\n{')
			.replaceAll(/[\[\]]*/g, '');
	}
</script>

<div class="text-red-500 text-center">! use subdomains / domains !</div>
<Input type="textarea" {ondrop} bind:value />
<Button class="bg-black text-white rounded hover:scale-105 hover:cursor-pointer">Save</Button>
