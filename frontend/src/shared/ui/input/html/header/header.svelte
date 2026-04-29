<script lang="ts">
	import { Button, Input } from '$shared/ui';

	let { input = $bindable(), value = $bindable() } = $props();

	function wrap(tag: string) {
		const selection = window.getSelection();
		if (!selection || selection.rangeCount === 0 || selection.isCollapsed) return;

		const range = selection.getRangeAt(0);

		// 1. Find the top-level block (the direct child of our editor)
		let currentBlock: Node | null = range.commonAncestorContainer;
		if (currentBlock.nodeType === Node.TEXT_NODE) {
			currentBlock = currentBlock.parentElement;
		}

		// Climb until we hit the direct child of the input container
		while (currentBlock && currentBlock.parentNode !== input) {
			currentBlock = currentBlock.parentNode;
		}

		// 2. Create the new wrapper
		const newWrapper = document.createElement(tag);

		if (currentBlock && currentBlock !== input) {
			// SCENARIO A: Replace existing block (prevents h3 > h2 > h1)
			// Move children from old block to new block
			while (currentBlock.firstChild) {
				newWrapper.appendChild(currentBlock.firstChild);
			}
			(currentBlock as Element).replaceWith(newWrapper);
		} else {
			// SCENARIO B: Fresh wrap (no existing block)
			newWrapper.appendChild(range.extractContents());
			range.insertNode(newWrapper);
		}

		// 3. Cleanup and sync
		value = input.innerHTML;
		selection.removeAllRanges();
	}

	function removeWrap() {
		const selection = window.getSelection();
		// 1. Check if selection exists and isn't just a cursor blink (isCollapsed)
		// Remove 'selection.isCollapsed' if you want to allow unwrapping empty tags
		if (!selection || selection.rangeCount === 0 || selection.isCollapsed) return;

		const range = selection.getRangeAt(0);

		// 2. Start from where the user began their selection
		let node: Node | null = range.startContainer;

		// 3. Climb up the DOM tree
		// We stop when the parent of our current node is the 'input' container
		// This identifies the "wrapper" in a flat structure.
		while (node && node.parentNode !== input) {
			node = node.parentNode;

			// Safety check: if we hit the body or null, the selection is outside the editor
			if (!node || node === document.body) return;
		}

		// 4. 'node' is now the direct child (e.g., <h3>) of the editor
		if (node && node !== input) {
			const parent = node.parentNode!; // This is 'input'

			// Move all children (text, spans, etc.) out of the wrapper
			while (node.firstChild) {
				parent.insertBefore(node.firstChild, node);
			}

			// Delete the empty wrapper tag
			parent.removeChild(node);

			// 5. Sync the Svelte state
			value = input.innerHTML;
		}
	}
</script>

<div class="flex gap-2 border-b">
	<Input
		value="Heading"
		type="select"
		data={[
			{ title: 'Heading 1', callback: () => wrap('h1') },
			{ title: 'Heading 2', callback: () => wrap('h2') },
			{ title: 'Heading 3', callback: () => wrap('h3') },
			{ title: 'Paragraph', callback: () => removeWrap() }
		]}
	/>
	<Button
		class="bg-white px-2"
		onclick={(e) => {
			e.preventDefault();
			removeWrap();
			wrap('code');
		}}>Code</Button
	>
</div>
