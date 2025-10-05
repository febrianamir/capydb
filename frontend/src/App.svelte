<script>
  import Connection from "./components/connections/Connection.svelte";
  import InputConnection from "./components/connections/InputConnection.svelte";
  import TableList from "./components/tables/TableList.svelte";
  import Tab from "./components/Tab.svelte";
  import { getOpenedTab } from "./states/tab.svelte";
</script>

<div class="app">
  <Tab />
  <div>
    {#if getOpenedTab() === "CONNECTION"}
      <Connection />
      <InputConnection />
    {/if}
    {#if getOpenedTab() === "TABLE"}
      <TableList />
    {/if}
  </div>
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
