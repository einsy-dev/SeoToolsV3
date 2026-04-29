<script lang="ts">
	let { children, title } = $props();
</script>

{#if title}
	<label class="flex flex-col">
		<span>{title}</span>
		{@render children()}
	</label>
{:else}
	{@render children()}
{/if}
