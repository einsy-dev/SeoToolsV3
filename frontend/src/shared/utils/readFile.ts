export function readFile(file: File): Promise<string> | undefined {
	if (!file) return;
	return new Promise((res, rej) => {
		const reader = new FileReader();
		reader.onload = async function (event) {
			let result = event.target?.result as string;
			if (result) {
				res(result);
			} else {
				rej('Unable to read file');
			}
		};
		reader.readAsText(file);
	});
}
