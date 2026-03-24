type SiteI = Omit<models.Site, "convertValues" | "CreatedAt" | "UpdatedAt" | "DeletedAt" | "PbnID"> & {
  Status?: "loading" | "ready";
};

type PlayerI = controls.State;

interface HeaderI {
  title: string;
  path: string;
}

interface SiteFormI {
  Site: string;
  About: Required<models.AboutProps>;
  Required: models.RequiredProps;
  Account: { Email: string; Username: string; Password: string };
  Group: models.GroupProps;
}
