<script lang="ts">
	let { value = $bindable(''), input = $bindable() } = $props();
</script>

<div class="max-w-30 overflow-scroll">
	<input
		type="text"
		class="outline-none border-none field-sizing-content"
		bind:this={input}
		bind:value
	/>
</div>
