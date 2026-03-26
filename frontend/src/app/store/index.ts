import { writable } from "svelte/store";
import { defaultPlayer, defaultSite } from "./default";

class Store {
  header;
  site;
  player;
  constructor() {
    this.header = writable<HeaderI[]>([]);
    this.site = writable<SiteI>(defaultSite);
    this.player = writable<PlayerI>(defaultPlayer);
  }
}

export const store = new Store();
