<script lang="ts">
  import { readFile } from "$shared/utils/readFile";
  import { AddFile, Table } from "$widgets/parseCsv";
  import { CSV } from "$go/services";

  let files = $state([]);
  let parsed: string[][] = $state([]);

  // let parsed: string[][] = $derived(
  //   readFile(files[0])!.then(async (res) => {
  //     return await CSV.Parse(res);
  //   })
  //   return [] as string[][];
  // );

  $effect(() => {
    if (!files.length) return;
    Promise.all(files.map((f) => readFile(f))).then(async (res) => {
      if (!res.length) {
        parsed = [];
        return;
      }
      parsed = await CSV.Parse(res as string[]);
    });
  });
</script>

<div class="h-full flex flex-col">
  <AddFile bind:files />

  <!-- <div class="flex gap-4"> -->
  <!-- <div class=" bg-black overflow-scroll"> -->
  <Table data={parsed} />
  <!-- </div> -->
  <!-- <Select /> -->
  <!-- </div> -->
  <!--<div class="mt-auto w-full flex justify-center p-2 gap-2">
    <div class="border rounded flex items-center justify-center">
      <button class="cursor-pointer p-2">
        <ClipboardCopy class="w-8 h-8" />
      </button>
    </div>
  </div> -->
</div>
