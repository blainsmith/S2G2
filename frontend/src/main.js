import App from './App.svelte';
import { GreeterService, MathService } from './rpc/rpc.client';

window.GreeterService = GreeterService;
window.MathService = MathService;

const app = new App({
	target: document.body,
	props: {
		name: 'world'
	}
});

export default app;