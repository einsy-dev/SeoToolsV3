<script lang="ts">
	import { clickOutside, keyPress, windowBlur } from '$actions';
	import { animHeight } from '$shared/anim/animHeight';
	import type { InputI } from '../input';

	let {
		value = $bindable([]),
		multiple = false,
		options = [],
		class: className
	}: InputI['select'] = $props();

	let filtered = $derived(options.filter((opt) => !value.includes(opt.value)));

	let active = $state(false);

	function defCallback(v?: string | number) {
		if (!v) return;
		if (multiple) {
			if (!Array.isArray(value)) value = [];
			value.push(v);
			return;
		}
		value = v;
	}
</script>

<div
	class="flex flex-col grow cursor-pointer text-white"
	use:keyPress
	onbackspace={() => {
		value = value.slice(0, value.length - 1);
	}}
	onescape={() => {
		active = false;
	}}
	use:clickOutside
	onclick_outside={() => (active = false)}
	use:windowBlur
	onwindow_blur={() => {
		active = false;
	}}
>
	<div role="none" class="flex grow" onclick={() => (active = !active)}>{value}</div>
	<div class="relative">
		<div
			class="absolute inset-0 bg-black z-50 max-h-50 overflow-scroll"
			use:animHeight={{ trigger: active }}
		>
			<div class="flex flex-col gap-1 px-2 py-1">
				{#each filtered as opt}
					<div
						role="none"
						onclick={() => {
							(opt.callback || defCallback)(opt.value);
							if (multiple) return;
							active = false;
						}}
						class="text-nowrap text-sm"
					>
						{opt.title}
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>
