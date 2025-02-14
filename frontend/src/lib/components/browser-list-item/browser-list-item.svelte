<script lang="ts" module>
    import type { browser } from "$go/models";
    import type { WithElementRef } from "bits-ui";
    import type { HTMLAttributes } from "svelte/elements";
  
    export type BrowserListItemProps = WithElementRef<HTMLAttributes<HTMLDivElement>> & {
        data: browser.Browser
    }
</script>

<script lang="ts">
	import { FetchFirefox, InstallFirefox, LaunchFirefox } from "$go/raven/App.js";
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
	let isInstalling = $state(false);
	let isInstalled = $state(false);
	let startUrl = `http://soikke.li`
	
	function fetch(): void {
		console.log(`fetching ${data.Version}`);
		let stopListen = EventsOn("fetch-progress", (pct: number) => {
			fetchProgress = pct;
		});
		isFetching = true;
		FetchFirefox(data.Version)
    		.catch((err: Error) => console.log(err))
    		.finally(() => {
     			isFetching = false;
       			stopListen();
    		})
		    .then(() => {
                isInstalling = true;
                InstallFirefox(data.Version)
                    .then(() => {
                        isInstalled = true;
                    })
                    .catch((err: Error) => {
                        console.log(err);
                        return
                    })
                    .finally(() => {
                        isInstalling = false;
                    });
			});
	}

	function launch(): void {
		LaunchFirefox(data.Version, startUrl).catch((err: Error) => console.log(err));
	}
</script>

<div>
	<div class="flex items-center gap-4">
		{#if (data.InstallPath != "" || isInstalled)}
			<Check class="size-4 ml-3"/>
			<Button size="sm" variant="outline" onclick={launch}>
				Launch
			</Button>
		{:else if isFetching }
			<ProgressRing class="size-4" percent={fetchProgress} />
		{:else if isInstalling }
			<div class="size-4 text-sm">Installing...</div>
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
