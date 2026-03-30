<script lang="ts">
	import { Header } from '$widgets';
	import { Events } from '@wailsio/runtime';
	import { onMount, setContext } from 'svelte';

	let { children } = $props();
	let tabs = [
		{ title: 'Metrics', path: '/extension' },
		{ title: 'Player', path: '/extension/player' }
	];

	onMount(() => {
		Events.On('app:update', (event) => {
			console.log('app:update', event.data);
		});
	});
	setContext('ext', { data: { test: 'test' } });
</script>

<Header {tabs} />
{@render children?.()}
