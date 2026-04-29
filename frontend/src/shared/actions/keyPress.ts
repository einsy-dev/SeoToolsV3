import type { Action } from 'svelte/action';

export const keyPress: Action<
	HTMLElement,
	undefined,
	{
		onbackspace: (e: CustomEvent) => void;
		onescape: (e: CustomEvent) => void;
		// onbackspace: (e: CustomEvent) => void;
	}
> = (node) => {
	const handleKey = (event: KeyboardEvent) => {
		if (!node && event.defaultPrevented) return;

		switch (event.key) {
			case 'Backspace':
				node.dispatchEvent(new CustomEvent('backspace'));
				break;
			case 'Escape':
				node.dispatchEvent(new CustomEvent('escape'));
				break;
			default:
				null;
		}
	};

	window.addEventListener('keydown', handleKey, true);
	return {
		destroy() {
			window.removeEventListener('keydown', handleKey, true);
		}
	};
};
