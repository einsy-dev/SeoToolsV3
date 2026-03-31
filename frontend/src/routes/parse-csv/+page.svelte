<script lang="ts">
	import { CSV } from '$go/services';
	import { Clear, Copy, Select } from '$shared/ui';
	import { readFile } from '$shared/utils/readFile';
	import { AddFile, Table } from '$widgets/parseCsv';

	let files = $state([]);
	let parsed: string[][] = $state([]);
	let value = $state([]);

	let defValue = [
		'Domain',
		'Domain Rating',
		'Organic / Traffic',
		'Organic / Top Countries',
		'RefDomains',
		'Outgoing domains / All time',
		'Authority Score',
		'TrustFlow',
		'CitationFlow',
		'TopicalTrustFlow_Topic_0'
	];

	let filtered = $derived(filterCSV(parsed, value.length ? value : defValue));

	function filterCSV(csv: string[][], columns: string[]): string[][] {
		if (!columns.length) return csv;
		let res: string[][] = [];
		for (let i = 0; i < csv.length; i++) {
			for (let j = 0; j < csv[0].length; j++) {
				if (columns.includes(csv[0][j])) {
					if (!Array.isArray(res[i])) res[i] = [];
					res[i].push(csv[i][j]);
				}
			}
		}

		return res;
	}

	function copyFormatCSV(csv: string[][]) {
		return csv.map((el) => el.join('\t')).join('\r\n');
	}

	$effect(() => {
		if (!files.length) return;
		Promise.all(files.map((f) => readFile(f))).then(async (res) => {
			if (!res.length) {
				parsed = [];
				return;
			}
			parsed = await CSV.Parse(res as string[]);
		});
	});
</script>

<div class="flex flex-col h-screen overflow-hidden">
	<div class="flex justify-between bg-black">
		<Clear
			onclick={() => {
				files = [];
				value = [];
				parsed = [];
			}}
		/>
		<Select bind:value options={parsed[0]} class="p-1" />
		<Copy value={copyFormatCSV(filtered)} />
	</div>

	<div class="relative flex-1 min-h-0 p-4 flex flex-col">
		<AddFile bind:files />
		<div class="flex-1 overflow-auto">
			<Table data={filtered} />
		</div>
	</div>
</div>
