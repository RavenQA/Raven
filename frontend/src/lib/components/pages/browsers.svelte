<script lang="ts">
	// import BrowserListItem from "$lib/components/browser-list-item/browser-list-item.svelte";
	// import type { BrowserListItemData } from "$lib/components/browser-list-item/types";
	import { browser } from "$go/models";
	import BrowserList from "$lib/components/browser-list/browser-list.svelte";
	import { SyncBrowsers } from "$go/raven/App.js";

	import * as Page from "$lib/components/page/index";
	import { onMount } from "svelte";

	let data: browser.Browser[] = [];
	let error: Error | null = null;
	let loading = true;

	onMount(async () => {
		try {
			data = await SyncBrowsers();
		} catch (err) {
			console.log(`failed to sync browsers: ${err}`);
		} finally {
			loading = false;
		}
	});
</script>

<Page.Root title="Browsers">
	<BrowserList {data} />
</Page.Root>
