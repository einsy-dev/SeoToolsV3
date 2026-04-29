import { animate } from 'animejs';
import type { Action } from 'svelte/action';

export const animHeight: Action<HTMLElement, { trigger: boolean }> = (node, params) => {
	node.style.overflow = 'hidden';
	node.classList.add('h-0');

	function toggle(active: boolean) {
		animate(node, {
			height: active ? node.scrollHeight : 0,
			easing: 'easeInOut',
			duration: 300
		});
	}

	toggle(params.trigger);

	return {
		update(newParams) {
			toggle(newParams.trigger);
		}
	};
};
