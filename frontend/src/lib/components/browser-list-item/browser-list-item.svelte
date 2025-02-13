<script lang="ts" module>
    import type { BrowserListItemData } from "./types";
    import type { WithElementRef } from "bits-ui";
    import type { HTMLAttributes } from "svelte/elements";
  
    export type BrowserListItemProps = WithElementRef<HTMLAttributes<HTMLDivElement>> & {
        data: BrowserListItemData
    }
</script>

<script lang="ts">
	import { Fetch, Run } from "$go/raven/App.js";
	import ProgressRing from "$lib/components/progress-ring.svelte";
	import { EventsOn } from "$runtime/runtime.js";
	import { Button } from "$lib/components/ui/button/index";
	import Download from "lucide-svelte/icons/download";
	import Check from "lucide-svelte/icons/check";
	
	let {
       ref = $bindable(null),
       class: className,
       title,
       children,
       data,
       ...restProps
	}: BrowserListItemProps = $props();
	
	let fetchProgress = $state(0);
	let isFetching = $state(false);

	function fetch(): void {
		console.log("fetching");
		let stopListen = EventsOn("fetchProgress", (pct: number) => {
			console.log(`progress: ${pct}`);
			fetchProgress = pct;
		});
		isFetching = true;
		Fetch()
			.catch((err: Error) => console.log(err))
			.finally(() => {
				isFetching = false;
				stopListen();
			});
	}

	function run(): void {
		Run().catch((err: Error) => console.log(err));
	}
</script>

<div>
	<div class="flex items-center gap-4">
		{#if data.isAvailable}
			<Check class="size-4"/>
		{:else if isFetching }
			<ProgressRing class="size-4" percent={fetchProgress} />
		{:else}
			<Button size="sm" variant="outline" onclick={fetch}
				><Download /></Button
			>
		{/if}
		<!-- <div class="text-foreground text-base">{data.name}</div>
		<div class="text-muted-foreground text-base">{data.version}</div>
		{#if isAvailable}
			<Button size="sm" onclick={run}>Launch Session</Button>
		{/if} -->
	</div>
</div>
