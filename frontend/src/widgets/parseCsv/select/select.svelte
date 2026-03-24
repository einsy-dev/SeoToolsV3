<script lang="ts">
  import { clickOutside } from "$shared/actions/clickOutside";
  import { Hand } from "@lucide/svelte";

  let { value = $bindable(["Test"]), options = [] } = $props<{
    value?: string[];
    "bind:value"?: string[];
    options?: string[];
  }>();
  let active = $state(false);

  function handleBackspace(e: KeyboardEvent) {
    if (e.key == "Backspace") {
      value = value.slice(0, -1);
    }
  }
</script>

<svelte:window onkeydown={handleBackspace} />
<div class="relative bg-black text-white">
  <button onclick={() => (active = !active)}>
    <div class="px-1">{value.length ? value.join(", ") : "Select"}</div>
  </button>
  <div class="" use:clickOutside onclick_outside={() => (active = false)}>
    <div class="absolute overflow-scroll max-h-[50vh] flex flex-col bg-black">
      <div class="flex flex-col px-2">
        {#if active}
          {#each options as el}
            {#if !value.includes(el)}
              <button onclick={() => value.push(el)}>
                <div class="text-nowrap text-left">{el}</div>
              </button>
            {/if}
          {/each}
        {/if}
      </div>
    </div>
  </div>
</div>
