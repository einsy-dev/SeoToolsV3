<script lang="ts">
	import { Textarea, Button } from '$shared/ui';
	import { extract } from '$shared/utils/extract';
	import { parseCSV } from '$shared/utils/parseCSV';
	import { updateObject } from '$shared/utils/updateObject';

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
<Textarea {ondrop} bind:value />
<Button class="bg-black text-white rounded hover:scale-105 hover:cursor-pointer">Save</Button>
