<script lang="ts">
	import { Header } from '$widgets';
	import { store } from '../../app/store';

	let { children } = $props();
	let tabs = [
		{ title: 'Form', path: '/extension' },
		{ title: 'Links', path: '/extension/links' }
	];
</script>

<Header {tabs} />
{@render children?.()}
