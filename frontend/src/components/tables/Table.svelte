<script>
  import { onMount } from "svelte";
  import {
    GetTableColumns,
    GetTableRecords,
    GetTables,
  } from "../../../wailsjs/go/usecase/Usecase";
  import { Table } from "@lucide/svelte";

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
        <ul class="table-list">
          {#each tables as table}
            <li
              class="table-item"
              onclick={(e) => {
                e.preventDefault();
                getTableContents(table);
              }}
            >
              <div class="table-item-icon">
                <Table size="15" />
              </div>
              <div class="table-item-text">
                {table}
              </div>
            </li>
          {/each}
        </ul>
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
    flex: 1;
    min-height: 0;
  }

  .table-content {
    width: 100%;
    overflow: auto;
  }

  .table-sidebar {
    flex: 0 0 250px;
    width: 250px;
    min-width: 0;
    overflow: auto;
  }

  .table-list {
    width: max-content;
    font-size: 14px;
    list-style: none;
    padding: 0.75rem;
    margin: 0;
  }

  .table-item {
    display: flex;
    align-items: center;
    padding: 0.2rem 0.4rem;
    gap: 0.2rem;
    transition: 0.2s ease;
    border-radius: 0.3rem;
  }

  .table-item:hover {
    background-color: #333333;
  }

  .table-item-icon {
    display: flex;
    align-items: center;
  }
</style>
