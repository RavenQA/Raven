<script lang="ts">
	import { cn } from "$lib/utils.js";
	import type { WithElementRef } from "bits-ui";
	import type { HTMLAttributes } from "svelte/elements";

	let {
		ref = $bindable(null),
		class: className,
		title,
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLDivElement>> = $props();
</script>

<div
	class={cn(
		"flex flex-col",
		className
	)}
	bind:this={ref}
	{...restProps}
>
	<header
		class="flex h-12 shrink-0 items-center gap-2"
	>
		<div class="flex items-center gap-2 px-4">
			<div class="text-foreground text-base font-medium">{title}</div>
		</div>
	</header>
	<div class="flex flex-1 flex-col gap-4 p-4 pt-0">
		{@render children?.()}
	</div>
</div>
