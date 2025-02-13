import type { ColumnDef } from "@tanstack/table-core";
import { renderComponent } from "$lib/components/ui/data-table/index.js";
import DataTableNameButton from "./data-table-name-button.svelte";
import DataTableVersionButton from "./data-table-version-button.svelte";
import type { BrowserListItemData } from "$lib/components/browser-list-item/types";
import { BrowserListItem } from "../browser-list-item";

export const columns: ColumnDef<BrowserListItemData>[] = [
  {
    accessorKey: "isAvailable",
    header: "",
    cell: ({ row }) =>
      renderComponent(BrowserListItem, {
        data: { isAvailable: row.getValue("isAvailable") },
      }),
  },
  {
    accessorKey: "name",
    header: ({ column }) =>
      renderComponent(DataTableNameButton, {
        onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }),
  },
  {
    accessorKey: "version",
    header: ({ column }) =>
      renderComponent(DataTableVersionButton, {
        onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }),
  },
];
