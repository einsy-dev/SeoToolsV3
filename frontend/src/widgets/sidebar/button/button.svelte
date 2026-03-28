<script lang="ts">
	let button: HTMLButtonElement;
	let { onclick, children } = $props();
</script>

<li>
	<button bind:this={button} {onclick} class="p-4 cursor-pointer hover:scale-105">
		{@render children?.()}
	</button>
</li>
