<script lang="ts">
	import Globe from "lucide-svelte/icons/globe";
	import Rocket from "lucide-svelte/icons/rocket";
	import Bird from "lucide-svelte/icons/bird";
	import Raven from "$assets/images/logo.png";
	
	import * as Sidebar from "$lib/components/ui/sidebar/index";
	import NavUser from "$lib/components/nav-user.svelte"

	import { Routes } from "$lib/routes";

	// This is sample data.
	const data = {
		sidebarItems: [
			{
				title: "Browsers",
				route: Routes.Browsers,
				url: "#",
				icon: Globe,
				isActive: true,
			},
			{
				title: "Sessions",
				route: Routes.Sessions,
				url: "#",
				icon: Rocket,
				isActive: true,
			},
		],
		user: {
		    name: "Michael Soikkeli",
    		email: "mike@soikke.li",
    		avatar: "/assets/images/logo.png"
		}
	};

	let {
		ref = $bindable(null),
		activeRoute = $bindable(Routes.Browsers),
		...restProps
	} = $props();

	let activeItem = $state(data.sidebarItems[0]);
</script>

<Sidebar.Root bind:ref {...restProps} class="border-r">
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton size="lg" class="md:h-8 md:p-0">
					{#snippet child({ props })}
						<a href="##" {...props}>
							<div
								class="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg"
							>
								<Bird />
							</div>
						</a>
					{/snippet}
				</Sidebar.MenuButton>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Header>
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupContent class="px-1.5 md:px-0">
				<Sidebar.Menu>
					{#each data.sidebarItems as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton 
								tooltipContentProps={{
									hidden: false,
								}}
								class="px-2.5 md:px-2"
								onclick={() => {
								    activeItem = item;
									activeRoute = item.route;
								}}
								isActive={ item.title == activeItem.title }
							>
								{#snippet tooltipContent()}
									{item.title}
								{/snippet}
								<item.icon />
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
	<Sidebar.Footer>
		<NavUser user={data.user} />
	</Sidebar.Footer>
</Sidebar.Root>
