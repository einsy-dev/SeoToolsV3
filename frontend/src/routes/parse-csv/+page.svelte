<script lang="ts">
  import { Select } from "$shared/ui";
  import { AddFile, Table } from "$widgets/parseCsv";
  import { ClipboardCopy } from "@lucide/svelte";
</script>

<div class="h-full flex flex-col">
  <AddFile />
  <!-- <div class="flex gap-4">
    <Select />
  </div>
  <div class=" bg-black overflow-scroll"><Table /></div>
  <div class="mt-auto w-full flex justify-center p-2 gap-2">
    <div class="border rounded flex items-center justify-center">
      <button class="cursor-pointer p-2">
        <ClipboardCopy class="w-8 h-8" />
      </button>
    </div>
  </div> -->
</div>
