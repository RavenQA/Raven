<script lang="ts">
	import AppSidebar from "$lib/components/app-sidebar.svelte";
	import * as Sidebar from "$lib/components/ui/sidebar/index";
	import SidebarRail from "$lib/components/ui/sidebar/sidebar-rail.svelte";
	import { Routes } from "$lib/routes";
	import type { Component } from "svelte";
	import { fade } from "svelte/transition";

	let activeRoute: Component = $state(Routes.Browsers);
</script>

<Sidebar.Provider>
	<AppSidebar bind:activeRoute />
	<Sidebar.Inset>
		{@const Page = activeRoute}
		<Page />
	</Sidebar.Inset>
	<SidebarRail />
</Sidebar.Provider>
