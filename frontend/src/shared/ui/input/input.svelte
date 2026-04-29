<script lang="ts" generics="T extends keyof InputI">
	import Checkbox from './checkbox/checkbox.svelte';
	import Email from './email/email.svelte';
	import File from './file/file.svelte';
	import Html from './html/html.svelte';
	import type { InputI, InputIDef } from './input';
	import { default as DefInput } from './input/input.svelte';
	import Label from './label/label.svelte';
	import Password from './password/password.svelte';
	import Select from './select/select.svelte';
	import Textarea from './textarea/textarea.svelte';
	import Username from './username/username.svelte';
	import type { Component } from 'svelte';
	import type { HTMLInputAttributes } from 'svelte/elements';

	let { value = $bindable(''), ...props }: InputI[T] | InputIDef = $props();

	let InputComponent = $derived(switchInput(props.type));

	function switchInput(type: HTMLInputAttributes['type']): Component<any> {
		switch (type) {
			case 'password':
				return Password;
			case 'email':
				return Email;
			case 'username':
				return Username;
			case 'html':
				return Html;
			case 'file':
				return File;
			case 'select':
				return Select;
			case 'checkbox':
				return Checkbox;
			case 'textarea':
				return Textarea;
			default:
				return DefInput;
		}
	}
</script>

<Label title={props.label}>
	<InputComponent bind:value {...props} />
</Label>
