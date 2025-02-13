<script lang="ts">
	import type { WithElementRef } from "bits-ui";
	import type { HTMLAttributes } from "svelte/elements";
	const circumference: number = 8 * 2 * Math.PI;

	let { ref = $bindable(null), percent, ...restProps } = $props();
</script>

<div
	class="inline-flex items-center justify-center overflow-hidden rounded-full"
>
	<svg class="w-8 h-8">
		<circle
			class="text-gray-300"
			stroke-width="3"
			stroke="currentColor"
			fill="transparent"
			r="8"
			cx="16"
			cy="16"
		/>
		<circle
			class="text-blue-600"
			stroke-width="3"
			stroke-dasharray={circumference}
			stroke-dashoffset={circumference - (percent / 100) * circumference}
			stroke-linecap="round"
			stroke="currentColor"
			fill="transparent"
			r="8"
			cx="16"
			cy="16"
		/>
	</svg>
</div>
