import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import wails from '@wailsio/runtime/plugins/vite';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [wails('./src/lib'), tailwindcss(), sveltekit()]
});
