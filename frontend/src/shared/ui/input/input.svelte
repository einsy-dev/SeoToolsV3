<script lang="ts">
  import { Files } from "@lucide/svelte";
  import type { HTMLInputAttributes } from "svelte/elements";
  let {
    value = $bindable(""),
    edit = $bindable(false),
    placeholder = ""
  }: HTMLInputAttributes & { edit?: boolean; "bind:edit"?: boolean } = $props();
</script>

<div class="relative flex card py-2!">
  {#if edit}
    <input class="flex w-full outline-0" type="text" {placeholder} bind:value />
  {:else}
    <button class="flex w-full outline-0 cursor-pointer" onclick={() => (edit = true)}>{value || placeholder}</button>
    <button
      class="cursor-pointer absolute right-2 ps-8 bg-linear-to-r from-transparent via-black/80 to-black"
      onclick={() => null}><Files onclick={() => {}} /></button
    >
  {/if}
</div>
