export function parseCSV(file: File): Promise<{ [key: string]: string | number }[]> {
  return new Promise((res, rel) => {
    let reader = new FileReader();
    reader.onload = (e) => {
      let str = e.target?.result;
      if (!str || typeof str !== "string") return;
      let arr = str
        .trim()
        .replaceAll("\r", "")
        .split("\n")
        .map((el) => el.split(","));
      let keys: string[] = [];
      let data = [];
      for (let i = 0; i < arr.length; i++) {
        let obj: { [key: string]: string | number } = {};
        for (let j = 0; j < arr[i].length; j++) {
          if (i === 0) {
            keys[j] = arr[i][j].replaceAll('"', "");
          } else {
            let v = arr[i][j].replaceAll('"', "");
            obj[keys[j]] = isNaN(+v) ? v : +v;
          }
        }
        if (i !== 0) {
          data.push(obj);
        }
      }
      res(data);
    };
    reader.readAsText(file);
  });
}
