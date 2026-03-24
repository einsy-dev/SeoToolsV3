export function parseDomain(url: string): string | null {
  if (!url || typeof url !== "string") return "";
  const regex = /^(?:https?:\/\/)?(?:www\.)?([^\/\s]+\.[^\/\s?]+)/im;
  let match: RegExpMatchArray | null | string = url.match(regex);
  if (!match) return null;
  match = match[1].toLowerCase();
  return match || null;
}
