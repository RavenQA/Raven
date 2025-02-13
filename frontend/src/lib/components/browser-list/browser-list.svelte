<script lang="ts">
	import DataTable from "./data-table.svelte";
	import { columns } from "./columns";
	import type { BrowserListItemData } from "$lib/components/browser-list-item/types";

	let { data }: { data: BrowserListItemData[] } = $props();
</script>

<DataTable {data} {columns} />
