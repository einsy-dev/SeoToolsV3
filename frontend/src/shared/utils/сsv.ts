import { parseDomain } from "./parseDomain";

export function parseCsv(text: string | undefined): string[][] | undefined {
  if (!text) return;
  let arr = text
    .replaceAll("\r", "")
    .split("\n")
    .map((item: any) => item.split(/,(?=(?:[^"]*"[^"]*")*[^"]*$)/).map((el: any) => el.replaceAll('"', "")));

  return arr;
}

let keys = ["Item", "Target"];

type CsvMap = { [key: string]: number | number[] };

export class Csv {
  data: string[][];
  rows: CsvMap = {};
  cols: CsvMap = {};
  key: number | undefined;

  constructor(data?: string) {
    this.data = parseCsv(data) || [];
    this.getCols();

    Object.keys(this.cols).find((c, i) => {
      if (keys.includes(c)) {
        this.key = Array.isArray(this.cols[c]) ? this.cols[c][0] : this.cols[c];
      }
    });

    this.getRows(this.key);
  }

  private getRows(keyCol: number | undefined) {
    if (typeof keyCol == "undefined" || !this.data.length) return;
    for (let i = 1; i < this.data.length; i++) {
      let val = parseDomain(this.data[i][keyCol]);
      if (!val) continue;
      this.rows[val] = i;
    }
  }

  private getCols() {
    if (!this.data[0]?.length) return;
    for (let i = 0; i < this.data[0].length; i++) {
      let val = this.data[0][i];
      if (this.cols[val]) {
        if (Array.isArray(this.cols[val])) {
          this.cols[val].push(i);
        } else {
          this.cols[val] = [this.cols[val], i];
        }
      } else {
        this.cols[val] = i;
      }
    }
  }

  join(csv: Csv) {
    if (typeof this.key == "undefined") return;
    for (let domain in csv.rows) {
      for (let col in csv.cols) {
        if (Array.isArray(this.rows[domain]) || Array.isArray(this.cols[col])) return;
        if (Array.isArray(csv.rows[domain]) || Array.isArray(csv.cols[col])) return;
        if (csv.cols[col] == csv.key) continue;

        if (!this.cols[col]) {
          let i = Object.keys(this.cols).length;
          this.cols[col] = Object.keys(this.cols).length;
          this.data[0][i] = col;
        }

        if (!this.rows[domain]) {
          this.rows[domain] = Object.keys(this.rows).length;
          this.data[this.rows[domain]] = Array.from({ length: Object.keys(this.cols).length });
          this.data[this.rows[domain]][this.key] = domain;
        }

        let newVal = csv.data[csv.rows[domain]][csv.cols[col]];
        this.data[this.rows[domain]][this.cols[col]] = newVal;
      }
    }

    let l = Object.keys(this.cols).length;
    for (let i = 1; i < this.data.length; i++) {
      let start = this.data[i].length;
      this.data[i].length = l;
      this.data[i].fill("", start);
    }
  }
}
