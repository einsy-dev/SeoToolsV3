import type { models } from "$lib";

export const defaultAbout = { Type: "", Topic: "", Comment: "", Anchor: "" };
export const defaultRequired = { Email: false, Photo: false, Phone: false, Approve: false, Moderation: false };
export const defaultAccount = { Email: "", Username: "", Password: "" };

export const defaultSite: SiteI = {
  Domain: "",
  About: { ...defaultAbout, Required: defaultRequired } as models.About,
  Ahref: {},
  Majestic: {},
  Semrush: {},
  Links: [],
  Groups: [],
  Accounts: []
};

export const defaultPlayer: PlayerI = {
  Links: [],
  Index: 0,
  Current: "",
  CurrentDomain: ""
};

export const defaultForm: SiteFormI = {
  Site: "",
  About: defaultAbout,
  Required: defaultRequired,
  Account: defaultAccount,
  Group: { Title: "" }
};
