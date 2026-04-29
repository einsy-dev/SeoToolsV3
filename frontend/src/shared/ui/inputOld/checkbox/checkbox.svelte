<script lang="ts">
	let { checked = $bindable(), title = '' } = $props<{
		checked?: boolean;
		'bind:checked'?: boolean;
		title?: string;
	}>();
</script>

<label class="">
	<input type="checkbox" bind:checked />
	<span>{title}</span>
</label>
