<script lang="ts">
	// import BrowserListItem from "$lib/components/browser-list-item/browser-list-item.svelte";
	// import type { BrowserListItemData } from "$lib/components/browser-list-item/types";
	import { types } from "$go/models";
	import BrowserList from "$lib/components/browser-list/browser-list.svelte";
	import { FetchVersions } from "$go/raven/App.js";

	import * as Page from "$lib/components/page/index";
	import { onMount } from "svelte";

	let data: types.BrowserListItem[] = [];
	let error: Error | null = null;
	let loading = true;

	onMount(async () => {
		try {
			data = await FetchVersions();
		} catch (err) {
			error = err as Error;
			console.error(err);
		} finally {
			loading = false;
		}
	});
</script>

<Page.Root title="Browsers">
	<BrowserList {data} />
</Page.Root>
