<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";

  let { value = $bindable([]), ...props }: HTMLInputAttributes = $props();

  let active = $state(false);
  let inputValue = $state("");
  let input: HTMLInputElement;

  export function focus() {
    input?.focus();
  }
</script>

<input
  class="outline-0 field-sizing-content"
  bind:value
  bind:this={input}
  {...props}
  onkeydown={(e) => {
    if (e.key === "Backspace" && inputValue === "") {
      value = value.slice(0, value.length - 1);
    } else if (e.key === "Escape") {
      active = false;
      input?.blur();
    }
  }}
/>
