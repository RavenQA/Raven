<script lang="ts">
    import logo from "../assets/images/logo.png";
    import { Fetch, Run } from "../../wailsjs/go/main/App";
    import ProgressRing from "./../components/ProgressRing.svelte";
    import { EventsOn } from "../../wailsjs/runtime/runtime";

    let name: string;
    let fetchProgress = $state(0);
    let isFetching = $state(false);

    function fetch(): void {
        console.log("fetching");
        let stopListen = EventsOn("fetchProgress", (pct) => {
            console.log(`progress: ${pct}`);
            fetchProgress = pct;
        });
        isFetching = true;
        Fetch()
            .catch((err) => console.log(err))
            .finally(() => {
                isFetching = false;
                stopListen();
            });
    }

    function run(): void {
        Run().catch((err) => console.log(err));
    }
</script>

<main>
    <img alt="Raven logo" id="logo" src={logo} />
    <div class="flex items-center justify-center">
        <ProgressRing bind:percent={fetchProgress} />
        <div>Firefox - Latest</div>
        <button
            class="disabled:opacity-50
                disabled:cursor-not-allowed"
            disabled={isFetching}
            onclick={fetch}>Install</button
        >
        <button class="" onclick={run}>Launch</button>
    </div>
</main>

<style>
    #logo {
        display: block;
        width: 50%;
        height: 50%;
        margin: auto;
        padding: 10% 0 0;
        background-position: center;
        background-repeat: no-repeat;
        background-size: 100% 100%;
        background-origin: content-box;
        opacity: 0.3;
    }
</style>
