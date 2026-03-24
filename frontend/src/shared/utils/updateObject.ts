export function updateObject(t: { [key: string]: any }, s: { [key: string]: any }): any {
  for (const key in t) {
    if (Object.prototype.hasOwnProperty.call(t, key) && Object.prototype.hasOwnProperty.call(s, key)) {
      t[key] = s[key];
    }
  }
  return t;
}
