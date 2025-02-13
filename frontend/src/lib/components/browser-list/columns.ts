import type { ColumnDef } from "@tanstack/table-core";
import { renderComponent } from "$lib/components/ui/data-table/index.js";
import DataTableNameButton from "./data-table-name-button.svelte";
import DataTableVersionButton from "./data-table-version-button.svelte";
import DataTableReleaseDateButton from "./data-table-release-date-button.svelte";
// import type { BrowserListItemData } from "$lib/components/browser-list-item/types";
import { types } from "$go/models";
import { BrowserListItem } from "../browser-list-item";

export const columns: ColumnDef<types.BrowserListItem>[] = [
  {
    accessorKey: "Available",
    header: "",
    cell: ({ row }) =>
      renderComponent(BrowserListItem, {
        data: { isAvailable: row.getValue("Available") },
      }),
  },
  {
    accessorKey: "Name",
    header: ({ column }) =>
      renderComponent(DataTableNameButton, {
        onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }),
  },
  {
    accessorKey: "Version",
    header: ({ column }) =>
      renderComponent(DataTableVersionButton, {
        onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }),
  },
  {
    accessorKey: "ReleaseDate",
    header: ({ column }) =>
      renderComponent(DataTableReleaseDateButton, {
        onclick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }),
  },
];
