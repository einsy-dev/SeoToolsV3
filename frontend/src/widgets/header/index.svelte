<script lang="ts">
  import { page } from "$app/state";
  import { goto } from "$app/navigation";
  import { store } from "../../app/store";
  import { Button } from "$shared/ui";

  let tabs = $state<{ title: string; path: string }[]>([]);
  let index = $derived(tabs.indexOf(tabs.find((el) => el.path === page.url.pathname)!));

  store.header.subscribe((val) => (tabs = val));

  const prev = "rounded-br-xl";
  const active = "!bg-[rgba(60,60,60,1)] rounded-t-xl";
  const next = "rounded-bl-xl";
  const def = "bg-black";
</script>

<nav class="sticky top-0 z-20 border-y-3 border-t-black border-b-[rgba(60,60,60,1)]">
  <ul class="flex text-white">
    {#each tabs as tab, i}
      <li class={`${i === index ? "bg-black" : "bg-[rgba(60,60,60,1)]"}`}>
        <div class={`${i === index ? active : i - 1 == index ? next : i + 1 === index ? prev : ""} ${def}`}>
          <button onclick={() => goto(tab.path)} class="px-4">{tab.title}</button>
        </div>
      </li>
    {/each}
    <div class="flex-1 bg-[rgba(60,60,60,1)]">
      <div class={`w-full h-full bg-black ${index === tabs.length - 1 ? next : ""}`}></div>
    </div>
  </ul>
</nav>
