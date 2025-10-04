<script>
  import { onMount } from "svelte";
  import {
    GetTableColumns,
    GetTables,
  } from "../../../wailsjs/go/usecase/Usecase";

  let isShowTableList = $state(false);
  let tables = $state([]);
  let tableColumns = $state([]);

  onMount(async () => {
    try {
      tables = await GetTables();
      isShowTableList = true;
    } catch (err) {
      console.log("Failed to get table list:", err);
    }
  });

  async function getTableColumns(tableName) {
    try {
      tableColumns = await GetTableColumns(tableName);
      console.log(tableColumns);
    } catch (err) {
      console.log("Failed to get table columns:", err);
    }
  }
</script>

<div class="table">
  <div class="table-sidebar">
    {#if isShowTableList}
      {#if tables.length > 0}
        <div class="ul">
          {#each tables as table}
            <li
              onclick={(e) => {
                e.preventDefault();
                getTableColumns(table);
              }}
            >
              {table}
            </li>
          {/each}
        </div>
      {/if}
    {/if}
  </div>
  <div class="table-content">
    {#if tableColumns.length > 0}
      <table class="table-data">
        <thead>
          <tr>
            {#each tableColumns as tableColumn}
              <th>{tableColumn.column_name}</th>
            {/each}
          </tr>
        </thead>
      </table>
    {/if}
  </div>
</div>

<style>
  .table {
    display: flex;
  }

  .table-content {
    width: 100%;
    overflow: auto;
  }

  .table-sidebar {
    width: 250px;
    overflow: auto;
  }
</style>
