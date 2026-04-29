<script lang="ts">
	import { browser } from '$app/environment';
	import Konva from 'konva';
	import { onDestroy, onMount, tick } from 'svelte';
	import { Stage, Layer, Circle, Line } from 'svelte-konva';

	let list = $state<any[]>([]);
	let lines = $state<any[]>([]); // This will hold our pairs
	let config = $state({ width: 0, height: 0, dots: 50, maxDist: 150 ** 2 });
	let mainAnim: Konva.Animation | null = null;

	onMount(async () => {
		config.width = window.innerWidth;
		config.height = window.innerHeight;

		if (config.width <= 440) {
			config.maxDist = 90 ** 2;
		}

		list = Array.from({ length: config.dots }, (_, n) => ({
			id: n.toString(),
			x: Math.random() * window.innerWidth,
			y: Math.random() * window.innerHeight,
			vx: Math.random() - 0.5,
			vy: Math.random() - 0.5,
			component: null as any
		}));

		await tick();

		const layer = list[0].component.node.getLayer();

		mainAnim = new Konva.Animation((frame) => {
			const activeLines = [];

			// 1. Update Dot Positions
			for (let i = 0; i < list.length; i++) {
				const dot = list[i];
				const node = dot.component.node;

				let nx = node.x() + dot.vx;
				let ny = node.y() + dot.vy;

				if (nx < 0 || nx > config.width) dot.vx *= -1;
				if (ny < 0 || ny > config.height) dot.vy *= -1;

				node.x(nx);
				node.y(ny);
			}

			// 2. Find Close Pairs
			for (let i = 0; i < list.length; i++) {
				for (let j = i + 1; j < list.length; j++) {
					const p1 = list[i].component.node;
					const p2 = list[j].component.node;

					const dx = p1.x() - p2.x();
					const dy = p1.y() - p2.y();
					const distSq = dx * dx + dy * dy;

					if (distSq < config.maxDist) {
						// Push the pair coordinates
						activeLines.push([p1.x(), p1.y(), p2.x(), p2.y()]);
					}
				}
			}

			// 3. Update the lines state
			lines = activeLines;
		}, layer);

		mainAnim.start();
	});

	onDestroy(() => mainAnim?.stop());
</script>

<svelte:window bind:innerWidth={config.width} bind:innerHeight={config.height} />

<div class="fixed inset-0 -z-50 bg-primary">
	{#if browser}
		<Stage {...config}>
			<Layer>
				{#each lines as points, i}
					<Line {points} stroke="rgba(255, 255, 255, 0.15)" strokeWidth={1} listening={false} />
				{/each}

				{#each list as item (item.id)}
					<Circle
						bind:this={item.component}
						x={item.x}
						y={item.y}
						radius={2}
						fill="rgba(28, 58, 106, 0.5)"
					/>
				{/each}
			</Layer>
		</Stage>
	{/if}
</div>
