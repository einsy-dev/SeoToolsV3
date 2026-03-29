import type { Action } from 'svelte/action';

export const windowBlur: Action<
	HTMLElement,
	undefined,
	{
		onwindow_blur: (e: CustomEvent) => void;
	}
> = (node) => {
	const handleBlur = (event: FocusEvent) => {
		if (!event.preventDefault) {
			node.dispatchEvent(new CustomEvent('onwindow_blur'));
		}
	};

	window.addEventListener('blur', handleBlur, true);
	return {
		destroy() {
			window.removeEventListener('blur', handleBlur, true);
		}
	};
};
