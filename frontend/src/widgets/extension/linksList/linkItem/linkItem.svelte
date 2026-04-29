<script lang="ts">
	import { animHeight } from '$shared/anim';
	import type { LinkI } from '../../types';
	import { animate } from 'animejs';

	let { linkPage, targetUrl, anchor, rel, reviewedAt, stetus }: Partial<LinkI> = $props();

	let active = $state(false);
</script>

<div class="flex flex-col gap-1 bg-black text-white rounded pt-1 px-2">
	<div class="" role="none" onclick={() => active = !active}>{linkPage}</div>
	<div class="" use:animHeight={{ trigger: active }}>
		{targetUrl}
	</div>
</div>
