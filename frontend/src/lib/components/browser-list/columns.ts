import type { ColumnDef } from "@tanstack/table-core";
import { renderComponent } from "$lib/components/ui/data-table/index.js";
import DataTableNameButton from "./data-table-name-button.svelte";
import DataTableVersionButton from "./data-table-version-button.svelte";
import DataTableReleaseDateButton from "./data-table-release-date-button.svelte";
import { types } from "$go/models";
import BrowserListItem from "$lib/components/browser-list-item/browser-list-item.svelte";

export const columns: ColumnDef<types.BrowserListItem>[] = [
  {
    accessorKey: "Available",
    header: "",
    cell: ({ row }) =>
      renderComponent(BrowserListItem, {
        data: {
          Available: row.getValue("Available"),
          Version: row.getValue("Version"),
        },
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
