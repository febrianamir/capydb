<script>
  import { onMount } from "svelte";
  import {
    GetTableColumns,
    GetTableRecords,
    GetTables,
  } from "../../../wailsjs/go/usecase/Usecase";

  let isShowTableList = $state(false);
  let tables = $state([]);
  let tableColumns = $state([]);
  let tableRecords = $state([]);

  onMount(async () => {
    try {
      tables = await GetTables();
      isShowTableList = true;
    } catch (err) {
      console.log("Failed to get table list:", err);
    }
  });

  async function getTableContents(tableName) {
    try {
      tableColumns = await GetTableColumns(tableName);

      let req = {
        table_name: tableName,
      };
      let tableRecordsRes = await GetTableRecords(req);
      tableRecords = tableRecordsRes.Data;
    } catch (err) {
      console.log("Failed to get table contents:", err);
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
                getTableContents(table);
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
        <tbody>
          {#if tableRecords.length > 0}
            {#each tableRecords as tableRecord}
              <tr>
                {#each tableColumns as tableColumn}
                  <td>{tableRecord[tableColumn.column_name]}</td>
                {/each}
              </tr>
            {/each}
          {/if}
        </tbody>
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
