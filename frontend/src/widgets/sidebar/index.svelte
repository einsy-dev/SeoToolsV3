<script lang="ts">
	import { goto } from '$app/navigation';
	import { clickOutside } from '$shared/actions/clickOutside';
	import { sidebarState } from '.';
	import Button from './button/button.svelte';
	import { Blocks, Grid2x2Plus, Settings } from '@lucide/svelte';

	let active = $state(false);

	sidebarState.subscribe((val) => {
		active = val;
	});

	const router = {
		top: [
			{
				path: '/extension',
				icon: Blocks
			},
			{
				path: '/parse-csv',
				icon: Grid2x2Plus
			}
		],
		bottom: [
			{
				path: '/settings',
				icon: Settings
			}
		]
	};
</script>

<div
	class="absolute top-0 z-50 h-full bg-black -translate-x-full transition-all"
	style:transform={active ? `translateX(${100}%)` : ''}
	use:clickOutside
	onclick_outside={() => {
		sidebarState.set(false);
	}}
>
	<nav class="flex flex-col h-full justify-between">
		<ul class="flex flex-col">
			{#each router.top as route}
				<Button onclick={() => goto(route.path)}>
					<route.icon color={'white'} />
				</Button>
			{/each}
		</ul>
		<div class="flex flex-col">
			{#each router.bottom as route}
				<Button onclick={() => goto(route.path)}>
					<route.icon color={'white'} />
				</Button>
			{/each}
		</div>
	</nav>
</div>
