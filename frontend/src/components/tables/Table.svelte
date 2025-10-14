<script>
  import { getTableTabs } from "../../states/tables.svelte";
  import TableSidebar from "./TableSidebar.svelte";
  import TableTabs from "./TableTabs.svelte";
  import TableTab from "./TableTab.svelte";

  let tableTabs = getTableTabs();
</script>

<div class="table">
  <TableSidebar />
  <div class="table-explorer">
    <TableTabs />
    {#each tableTabs.tabs as tableTab}
      <TableTab tableName={tableTab.tableName} />
    {/each}
  </div>
</div>

<style>
  .table {
    display: flex;
    flex: 1;
    min-height: 0;
  }

  /* Table Explorer */
  .table-explorer {
    background-color: var(--color-almost-black);
    flex: 1;
    overflow: auto;
    text-align: left;
    font-size: 14px;
  }
</style>
