<script lang="ts">
	import { type InputI } from '../input';
	import Input from '../input/input.svelte';
	import { AtSign } from '@lucide/svelte';

	let { value = $bindable(), placeholder = '' }: InputI['email'] = $props();
</script>

<div class="flex items-center justify-center">
	<AtSign class="icon" />
	<Input {placeholder} type="email" bind:value />
</div>
