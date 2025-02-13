<script lang="ts">
	import { cn } from "$lib/utils.js";
	import type { WithElementRef } from "bits-ui";
	import type { HTMLAttributes } from "svelte/elements";
	import { useSidebar } from "./context.svelte.js";

	let {
		ref = $bindable(null),
		side = "left",
		variant = "sidebar",
		class: className,
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLDivElement>> & {
		side?: "left" | "right";
		variant?: "sidebar" | "floating" | "inset";
	} = $props();

	const sidebar = useSidebar();
</script>

<div
	class={cn(
		"bg-sidebar text-sidebar-foreground flex w-[--sidebar-width] flex-col",
		className
	)}
	bind:this={ref}
	{...restProps}
>
	{@render children?.()}
</div>
