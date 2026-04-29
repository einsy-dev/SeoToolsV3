<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { sidebarState } from '$widgets/sidebar';
	import { Menu } from '@lucide/svelte';

	let { tabs }: { tabs: { title: string; path: string }[] } = $props();

	let index = $derived(tabs.indexOf(tabs.find((el) => el.path === page.url.pathname)!));

	const prev = 'rounded-br-xl';
	const active = '!bg-[rgba(60,60,60,1)] rounded-t-xl';
	const next = 'rounded-bl-xl';
	const def = 'bg-black';
</script>

<nav class="sticky top-0 z-20 border-y-3 border-t-black border-b-[rgba(60,60,60,1)]">
	<ul class="flex text-white">
		<div class="bg-black px-1" role="none" onclick={() => sidebarState.set(true)}>
			<Menu class="text-white p-0.5" />
		</div>
		{#each tabs as tab, i}
			<li class={`${i === index ? 'bg-black' : 'bg-[rgba(60,60,60,1)]'}`}>
				<div
					class={`${i === index ? active : i - 1 == index ? next : i + 1 === index ? prev : ''} ${def}`}
				>
					<button onclick={() => goto(tab.path)} class="px-4">{tab.title}</button>
				</div>
			</li>
		{/each}
		<div class="flex-1 bg-[rgba(60,60,60,1)]">
			<div class={`w-full h-full bg-black ${index === tabs.length - 1 ? next : ''}`}></div>
		</div>
	</ul>
</nav>
