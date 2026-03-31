<script lang="ts">
	import * as Domain from '$go/services/domain';
	import { Spinner } from '$shared/ui';
	import Controls from '$widgets/extension/controls/controls.svelte';
	import SiteForm from '$widgets/extension/siteForm/siteForm.svelte';
	import Card from './card/card.svelte';

	let data: { [key: string]: any } = $state({});

	Domain.Fetch('google2.com').then((res) => (data = JSON.parse(res)));

	function ahrefsFormat(data: { [key: string]: any }) {
		return {
			DR: data?.dr || 'n/a',
			Traffic: data?.traffic || 'n/a',
			Geo: data?.geo || 'n/a',
			RefDom: data?.RefDomains || 'n/a',
			OutDom: data?.OutDomains || 'n/a',
			Age: data?.age || 'n/a'
		};
	}
	function semrushFormat(data: { [key: string]: any }) {
		return {
			AS: data?.as || 'n/a',
			Traffic: data?.traffic || 'n/a',
			RefDom: data?.RefDomains || 'n/a',
			OutDom: data?.OutDomains || 'n/a',
			Farm: data?.linkFarm || 'n/a'
		};
	}
	function majesticFormat(data: { [key: string]: any }) {
		return {
			TF: data?.tf || 'n/a',
			CF: data?.cf || 'n/a',
			Topic: data?.topic || 'n/a'
		};
	}
</script>

<div class="flex grow flex-col text-white px-2 py-1 gap-1">
	<div class="rounded bg-black flex px-4 py-1 items-center justify-between">
		<div class="">https://google.com</div>
		<div class=""><Spinner /></div>
	</div>

	<Card label="Ahrefs" data={ahrefsFormat(data.ahrefs)} />
	<Card label="Semrush" data={semrushFormat(data.semrush)} />
	<Card label="Majestic" data={majesticFormat(data.majestic)} />
	<textarea class="bg-black rounded outline-none resize-none px-2 py-1" bind:value={data.comment}></textarea>
</div>
<Controls />
