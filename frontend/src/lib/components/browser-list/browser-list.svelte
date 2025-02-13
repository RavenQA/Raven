<script lang="ts">
	import DataTable from "./data-table.svelte";
	import { columns } from "./columns";
	// import type { BrowserListItemData } from "$lib/components/browser-list-item/types";
	import { types } from "$go/models";

	let { data }: { data: types.BrowserListItem[] } = $props();
</script>

<DataTable {data} {columns} />
