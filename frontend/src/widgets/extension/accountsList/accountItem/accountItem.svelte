<script lang="ts">
	import { Clipboard } from '$go/services/clipboard';
	import { animHeight } from '$shared/anim';

	let { username, password, note }: { username?: string; password?: string; note?: string } =
		$props();

	let active = $state(false);
</script>

<div class="flex flex-col gap-1 bg-black text-white rounded pt-1 px-2">
	<div class="" role="none" onclick={() => (active = !active)}>
		{username}
	</div>
	<div class="flex flex-col gap-1 *:px-4" use:animHeight={{ trigger: active }}>
		<div class="flex w-full gap-2">
			<div
				class="w-1/2 cursor-pointer"
				role="none"
				onclick={() => {
					if (username) Clipboard.CopyText(username);
				}}
			>
				{username}
			</div>
			<div
				class="w-1/2 cursor-pointer"
				role="none"
				onclick={() => {
					if (password) Clipboard.CopyText(password);
				}}
			>
				{password}
			</div>
		</div>
		<div class="">{note}</div>
	</div>
</div>
