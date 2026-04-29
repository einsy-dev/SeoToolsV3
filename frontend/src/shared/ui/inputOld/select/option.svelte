<script lang="ts">
	let { onclick, value } = $props();
</script>

<div class="">
	<button class="w-full cursor-pointer" {onclick}>
		<div class=" text-start">
			{value}
		</div>
	</button>
</div>
