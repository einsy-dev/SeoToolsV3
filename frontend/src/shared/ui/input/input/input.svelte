<script lang="ts">
	import type { InputIDef } from '../input';

	let { value = $bindable(), placeholder = '', type = 'text' }: InputIDef = $props();
</script>

<input class="outline-0 border-0 px-1 w-full" {type} {placeholder} bind:value />
