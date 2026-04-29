<script lang="ts">
	import type { InputI } from '../input';
	import Input from '../input/input.svelte';
	import { User } from '@lucide/svelte';

	let { value = $bindable(), placeholder = '' }: InputI['username'] = $props();
</script>

<div class="flex items-center justify-center">
	<User class="icon" />
	<Input type="text" {placeholder} bind:value />
</div>
