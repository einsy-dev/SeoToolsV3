// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    // interface Locals {}
    // interface PageData {}
    // interface PageState {}
    // interface Platform {}
  }
}

export {};

// add protected types here
type Site = {
  Domain: string;
  Comment: string;
  Language: string;
  Location: string;
  Topic: string[];
  Type: string; // Blog | News | Corporate | E-commerce | Forum | Wiki | Portfolio | Personal | Other
  SiteLink: {
    Anchor: string;
    PageAccess: boolean;
    Rel: string[]; // "dofollow" | "nofollow" | "ugc" | "sponsored" | noopener | noreferrer | external
    Target: string; // "_blank" | "_self"
  };

  // Referring tables

  Pbn: {
    Price: string;
    Pbn: {
      Name: string;
      Email: string;
      Phone: string;
      Telegram: string;
      Whatsapp: string;
      Sites: Site[];
    };
  };
  Links: {
    Links: {
      RefferingPage: string;
      TargetUrl: string;
      Anchor: string;
      Type: string;
      Lost: boolean;
      Article: number;
    }[];
  }[];
  Required: {
    Required: {
      Email: boolean;
      BusinessEmail: boolean;
      Phone: boolean;
      Photo: boolean;
      Instagramm: boolean;
      Facebook: boolean;
      SberId: boolean;
      Yandex: boolean;
      BankCard: boolean;
    };
  };
  Ahrefs: {
    Ahrefs: {
      DomainRating: string;
      Traffic: string;
      Trend: string;
    };
  };
  Semrush: { Semrush: { AuthorityScore: string; Backlinks: string; Traffic: string; Farm: string } };
  Majestic: {
    Majestic: {
      TrustFlow: string;
      CitationFlow: string;
      Topic: string;
      Backlinks: string;
    };
  };
  Groups: {
    Group: {
      Name: string;
      Comment: string;
      Sites: Site[];
    }[];
  }[];
  Network: {
    Network: {
      Name: string; // hosting platforms with multiple domains exact match to domain (parent site)
      Main: string;
      Sites: Site[];
    };
  };
  Accounts: {
    Accounts: {
      Email: string;
      Username: string;
      Password: string;
      Default: boolean;
    }[];
  }[];
};
