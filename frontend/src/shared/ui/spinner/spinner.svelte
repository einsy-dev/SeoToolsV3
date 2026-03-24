<script lang="ts">
  import { CircleCheck, LoaderCircle } from "@lucide/svelte";
  let { active = true } = $props<{ active?: boolean }>();
</script>

{#if active}
  <LoaderCircle class="animate-spin h-5" />
{:else}
  <CircleCheck class="h-5" />
{/if}
