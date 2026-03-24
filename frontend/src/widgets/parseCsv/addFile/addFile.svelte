<script lang="ts">
  import { FilePlus } from "@lucide/svelte";

  let { files = $bindable([]) }: { files?: File[]; "bind:files"?: File[] } = $props();

  let active: boolean = $state(false);

  $effect(() => {
    if (files.length == 0) {
      active = true;
    }
    console.log(files);
  });

  function handleInput(e: Event) {
    if (!e.target) return;
    files.push(...Array.from((e.target as HTMLInputElement).files!));
    active = false;
  }
</script>

<svelte:window
  ondragenter={(e) => {
    e.preventDefault();
    active = true;
  }}
  ondragover={(e) => {
    e.preventDefault();
  }}
  ondragleave={(e) => {
    e.preventDefault();
    if (e.relatedTarget === null) active = false;
  }}
  ondrop={(e) => {
    e.preventDefault();
    active = false;
    if (!e.dataTransfer) return;
    files.push(...Array.from(e.dataTransfer.files));
  }}
/>

<label class="h-full p-2" hidden={!active}>
  <div class="h-full border rounded bg-[rgba(60,60,60,1)]">
    <input type="file" multiple hidden onchange={handleInput} />
    <div
      class="flex flex-col gap-2 border rounded h-full w-full  items-center justify-center cursor-pointer text-white"
    >
      <FilePlus class="w-10 h-10 " />
      <span>Choose a file or drag it over here</span>
    </div>
  </div>
</label>
