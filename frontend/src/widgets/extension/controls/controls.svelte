<script lang="ts">
  import * as Player from "$go/extension";
  import { EventsOff, EventsOn } from "$runtime";
  import { RotateCw, SkipBack, SkipForward, StepBack, StepForward } from "@lucide/svelte";
  import { onMount } from "svelte";
  import { store } from "$shared/store";
  import { Button } from "$shared/ui";

  let app = $state({ currentLink: "" });

  onMount(() => {
    EventsOn("currentLink", (url: string) => {
      app.currentLink = url;
    });
    return () => {
      EventsOff("currentLink");
    };
  });

  function proxyFn(cb: () => Promise<void>) {
    return async () => {
      await cb();
      await Player.State().then((res) => {
        store.player.update((state) => {
          state = res;
          return state;
        });
      });
    };
  }
</script>

<div class="sticky bottom-0 mt-auto w-full flex flex-col gap-4 items-center bg-black py-2 px-2 text-white">
  <div class="flex gap-4">
    <Button class="border px-4 py-1 rounded-lg hover:scale-105 hidden sm:block" onclick={proxyFn(Player.First)}
      ><SkipBack /></Button
    >
    <Button class="border px-4 py-1 rounded-lg hover:scale-105" onclick={proxyFn(Player.Prev)}><StepBack /></Button>
    <Button class="border px-4 py-1 rounded-lg hover:scale-105" onclick={proxyFn(Player.Reload)}><RotateCw /></Button>
    <Button class="border px-4 py-1 rounded-lg hover:scale-105" onclick={proxyFn(Player.Next)}><StepForward /></Button>
    <Button class="border px-4 py-1 rounded-lg hover:scale-105 hidden sm:block" onclick={proxyFn(Player.Last)}
      ><SkipForward /></Button
    >
  </div>
</div>
