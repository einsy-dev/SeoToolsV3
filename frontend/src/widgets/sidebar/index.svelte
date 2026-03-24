<script lang="ts">
  import { Blocks, Database, Grid2x2Plus, Settings, Table, Wrench } from "@lucide/svelte";
  import { goto } from "$app/navigation";
  import Button from "./button/button.svelte";
  let active = $state(false);
  let width = $state<number>();

  let sidebar: Node;
  function handleClick(e: MouseEvent) {
    if (!sidebar || !e.target) return;
    if (!sidebar.contains(e.target as Node)) {
      active = false;
    }
  }
</script>

<svelte:window bind:innerWidth={width} onclick={handleClick} />

<div
  class="absolute top-0 z-50 h-full bg-black -translate-x-full transition-all"
  style:transform={active ? `translateX(${100}%)` : ""}
  bind:this={sidebar}
>
  <nav class="flex flex-col h-full justify-between">
    <ul class="flex flex-col">
      <Button onclick={() => goto("/extension")}>
        <Blocks color={"white"} />
      </Button>
      <Button onclick={() => goto("/parse-csv")}>
        <Grid2x2Plus color={"white"} />
      </Button>
    </ul>
    <div class="flex flex-col">
      <Button onclick={() => goto("/settings")}>
        <Settings color={"white"} />
      </Button>
    </div>
  </nav>
  <button
    class="bg-black text-white absolute top-20 right-0 translate-x-full w-3 h-20 rounded-[0_10px_10px_0] cursor-pointer"
    onclick={() => (active = !active)}
  >
    |
  </button>
</div>
