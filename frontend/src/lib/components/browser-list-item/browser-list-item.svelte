<script lang="ts" module>
    import type { types } from "$go/models";
    import type { WithElementRef } from "bits-ui";
    import type { HTMLAttributes } from "svelte/elements";
  
    export type BrowserListItemProps = WithElementRef<HTMLAttributes<HTMLDivElement>> & {
        data: types.BrowserListItem
    }
</script>

<script lang="ts">
	import { FetchFirefox, InstallFirefox, LaunchFirefox } from "$go/raven/App.js";
	import ProgressRing from "$lib/components/progress-ring.svelte";
	import { EventsOn } from "$runtime/runtime.js";
	import { Button } from "$lib/components/ui/button/index";
	import Download from "lucide-svelte/icons/download";
	import Check from "lucide-svelte/icons/check";
	import Archive from "$lib/components/icons/archive/archive.svelte"
	
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
	let isInstalling = $state(false);
	let isInstalled = $state(false); // TODO: set data.Available in the Go func after install (and update DB)
	let startUrl = `http://soikke.li`

	function fetch(): void {
		console.log(`fetching ${data.Version}`);
		let stopListen = EventsOn("fetch-progress", (pct: number) => {
			console.log(`progress: ${pct}`);
			fetchProgress = pct;
		});
		isFetching = true;
		FetchFirefox(data.Version)
			.catch((err: Error) => console.log(err))
			.finally(() => {
				isFetching = false;
				stopListen();
				isInstalling = true;
				InstallFirefox(data.Version)
  		            .catch((err: Error) => console.log(err))
              		.finally(() => {
                 	    isInstalling = false;
                        isInstalled = true;
              		});
			});
	}

	function launch(): void {
		LaunchFirefox(data.Version, startUrl).catch((err: Error) => console.log(err));
	}
</script>

<div>
	<div class="flex items-center gap-4">
		{#if isInstalled}
			<Check class="size-4"/>
			<Button size="sm" variant="outline" onclick={launch}>
				Launch
			</Button>
		{:else if isFetching }
			<ProgressRing class="size-4" percent={fetchProgress} />
		{:else if isInstalling }
			<span class="text-sm">Installing...</span>
		{:else}
			<Button size="sm" variant="outline" onclick={fetch}>
				<Download />
			</Button>
		{/if}
		<!-- <div class="text-foreground text-base">{data.name}</div>
		<div class="text-muted-foreground text-base">{data.version}</div>
		{#if isAvailable}
			<Button size="sm" onclick={run}>Launch Session</Button>
		{/if} -->
	</div>
</div>
