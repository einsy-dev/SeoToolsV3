<script lang="ts">
	import { CsvParser } from '$go/services/csvParser';
	import { Clear, Copy, Input } from '$shared/ui';
	import { readFile } from '$shared/utils/readFile';
	import { AddFile, Table } from '$widgets/parseCsv';
	import { sidebarState } from '$widgets/sidebar';
	import { Menu } from '@lucide/svelte';

	let files = $state([]);
	let parsed: string[][] = $state([]);
	let value = $state([]);

	let defValue = [
		'Domain',
		'Domain Rating',
		'Organic / Traffic',
		'Organic / Top Countries',
		'Ref. domains / All',
		'Outgoing domains / All time',
		'Authority Score',
		'TrustFlow',
		'CitationFlow',
		'TopicalTrustFlow_Topic_0'
	];

	let filtered = $derived(filterCSV(parsed, value.length ? value : defValue));

	function filterCSV(csv: string[][], columns: string[]): string[][] {
		if (!columns.length || !csv.length) return csv;
		let res: string[][] = [];
		for (let i = 0; i < csv.length; i++) {
			for (let j = 0; j < csv[0].length; j++) {
				if (columns.includes(csv[0][j])) {
					if (!Array.isArray(res[i])) res[i] = [];
					res[i].push(csv[i][j]);
				}
			}
		}

		const currentHeader = res[0];

		// 2. Create a map of the target indices
		// This tells us: "Column 0 should actually be at index X"
		const sortedIndices = currentHeader
			.map((colName, index) => index)
			.sort((a, b) => {
				return defValue.indexOf(currentHeader[a]) - defValue.indexOf(currentHeader[b]);
			});

		// 3. Map every row to the new index order
		const sortedRes = res.map((row) => sortedIndices.map((i) => row[i]));

		return sortedRes;
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
			parsed = await CsvParser.Parse(res as string[]);
		});
	});
</script>

<div class="flex flex-col h-screen overflow-hidden">
	<div class="flex justify-between bg-black text-white">
		<Menu
			class="h-full"
			onclick={() => {
				sidebarState.set(true);
			}}
		/>
		<Clear
			onclick={() => {
				files = [];
				value = [];
				parsed = [];
			}}
		/>
		<Input
			type="select"
			bind:value
			options={parsed[0]?.map((el) => ({ title: el, value: el }))}
			multiple
			class="p-1 bg-amber-700 w-full"
		/>
		<Copy value={copyFormatCSV(filtered)} />
	</div>

	<div class="relative flex-1 min-h-0 p-4 flex flex-col">
		<AddFile bind:files />
		<div class="flex-1 overflow-auto">
			<Table data={filtered} />
		</div>
	</div>
</div>
