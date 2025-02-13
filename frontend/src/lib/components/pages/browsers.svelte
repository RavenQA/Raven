<script lang="ts">
	import BrowserListItem from "$lib/components/browser-list-item/browser-list-item.svelte";
	import type { BrowserListItemData } from "$lib/components/browser-list-item/types";
	import BrowserList from "$lib/components/browser-list/browser-list.svelte";

	import * as Page from "$lib/components/page/index";

	const data: BrowserListItemData[] = [
		{
			name: "Firefox",
			version: "1.2.3",
			path: "/Users/michael/Local/RavenBrowsers/",
			isAvailable: false,
		},
		{
			name: "Chrome",
			version: "4.5.6",
			path: "/Users/michael/Local/RavenBrowsers/",
			isAvailable: false,
		},
		{
			name: "Safari",
			version: "7.8.9",
			path: "/Users/michael/Local/RavenBrowsers/",
			isAvailable: true,
		},
	];
</script>

<Page.Root title="Browsers">
	<!-- {#each data as data}
		<BrowserListItem {data} />
	{/each} -->
	<BrowserList {data} />
</Page.Root>
