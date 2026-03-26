<script lang="ts">
  import "../app.css";
  import Sidebar from "$widgets/sidebar/index.svelte";
  import Header from "$widgets/header/index.svelte";
  let { children } = $props();
</script>

<div class="relative h-screen flex flex-col bg-[rgba(43,43,43,1)]">
  <Header />
  <Sidebar />
  <div class="relative flex-1 overflow-y-scroll">
    {@render children?.()}
  </div>
</div>
