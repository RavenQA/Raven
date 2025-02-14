<script lang="ts">
	import DataTable from "./data-table.svelte";
	import { ScrollArea } from "$lib/components/ui/scroll-area";
	import { columns } from "./columns";
	import { browser } from "$go/models";

	let { data }: { data: browser.Browser[] } = $props();
</script>

<DataTable {data} {columns} />
