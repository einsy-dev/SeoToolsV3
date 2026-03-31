<script lang="ts">
	let { label = '', data = {} }: { label?: string; data?: { [key: string]: string | number } } =
		$props();
	let keys = $derived(Object.keys(data));

	function format(n: any) {
		if (typeof n != 'number') return n;
		if (n > 1e9) {
			return Math.round(n / 1e9) + 'B';
		}
		if (n > 1e6) {
			return Math.round(n / 1e6) + 'M';
		}
		if (n > 1e3) {
			return Math.round(n / 1e3) + 'K';
		}
		return n;
	}
</script>

{#if keys.length}
	<div class="rounded bg-black flex">
		<div class="flex flex-1 p-1 px-2 items-center justify-between">
			<div class="grid gap-1 grow grid-cols-3 grid-rows-2 sm:grid-cols-5 md:grid-cols-6">
				{#each keys as key}
					<div class="overflow-scroll px-4">
						<div class="relative pt-3 text-sm">
							{format(data[key])}
							<div class="absolute top-0 -left-3 text-xs">{key}:</div>
						</div>
					</div>
				{/each}
			</div>
		</div>

		<div class="flex items-center justify-center">
			<div class="[writing-mode:vertical-lr] px-2">
				<span>{label}</span>
			</div>
		</div>
	</div>
{/if}
