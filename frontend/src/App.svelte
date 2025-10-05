<script>
  import Connection from "./components/connections/Connection.svelte";
  import Table from "./components/tables/Table.svelte";
  import Tab from "./components/Tab.svelte";
  import { getOpenedTab } from "./states/tab.svelte";
</script>

<div class="app">
  <Tab />
  {#if getOpenedTab() === "CONNECTION"}
    <Connection />
  {/if}
  {#if getOpenedTab() === "TABLE"}
    <Table />
  {/if}
</div>

<style>
  .app {
    display: flex;
    flex-direction: column;
    overflow: hidden;
    width: 100%;
    height: 100vh;
  }
</style>
