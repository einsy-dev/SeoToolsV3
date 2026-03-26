<script lang="ts">
  import { store } from "../../app/store";

  let { children } = $props();
  store.header.set([
    { title: "Form", path: "/extension" },
    { title: "Links", path: "/extension/links" }
  ]);
</script>

{@render children?.()}
