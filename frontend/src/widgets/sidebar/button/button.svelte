<script lang="ts">
  let { onclick, children } = $props();
  let button: HTMLButtonElement;
</script>

<li>
  <button bind:this={button} {onclick} class="p-4 cursor-pointer hover:scale-105">
    {@render children?.()}
  </button>
</li>
