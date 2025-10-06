<script>
  import { onMount } from "svelte";
  import { handleEnter } from "../../utils/key";
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

  function getColumnTypeColorClass(type, data) {
    if (data === null) {
      return "table-data-cell-null";
    }

    switch (type) {
      case "integer":
      case "numeric":
      case "double precision":
        return "table-data-cell-blue";
      case "character varying":
      case "text":
        return "table-data-cell-green";
      case "timestamp with time zone":
      case "timestamp without time zone":
        return "table-data-cell-orange";
      case "boolean":
        return "table-data-cell-yellow";
    }
  }
</script>

<div class="table">
  <div class="table-sidebar">
    {#if isShowTableList}
      {#if tables.length > 0}
        <div class="table-list">
          {#each tables as table}
            <div
              role="button"
              tabindex="0"
              class="table-item"
              onclick={(e) => {
                e.preventDefault();
                getTableContents(table);
              }}
              onkeydown={(e) => {
                handleEnter(e, () => {
                  e.preventDefault();
                  getTableContents(table);
                });
              }}
            >
              <div class="table-item-icon">
                <Table size="15" />
              </div>
              <div class="table-item-text">
                {table}
              </div>
            </div>
          {/each}
        </div>
      {/if}
    {/if}
  </div>
  <div class="table-content">
    {#if tableColumns.length > 0}
      <table class="table-data">
        <thead class="table-data-head">
          <tr>
            {#each tableColumns as tableColumn}
              <th class="table-data-cell">{tableColumn.column_name}</th>
            {/each}
          </tr>
        </thead>
        <tbody class="table-data-body">
          {#if tableRecords.length > 0}
            {#each tableRecords as tableRecord}
              <tr>
                {#each tableColumns as tableColumn}
                  <td
                    class="table-data-cell {getColumnTypeColorClass(
                      tableColumn.data_type,
                      tableRecord[tableColumn.column_name],
                    )}"
                    >{tableRecord[tableColumn.column_name] === null
                      ? "NULL"
                      : tableRecord[tableColumn.column_name]}</td
                  >
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
  :root {
    --color-almost-black: #1d1d1d;
    --color-dark-grey: #252525;
    --color-carolina-blue: #7591df;
    --color-dark-sage: #52895c;
    --color-dark-peach: #e6865c;
    --color-medium-gray: #808080;
    --color-light-mustard: #cca642;
  }

  .table {
    display: flex;
    flex: 1;
    min-height: 0;
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

  .table-content {
    background-color: var(--color-almost-black);
    flex: 1;
    overflow: auto;
    text-align: left;
    font-size: 14px;
  }

  .table-data {
    width: max-content;
    border-spacing: 0;
    padding: 0.75rem;
  }

  .table-data-head {
    background-color: var(--color-almost-black);
    border: 2px solid var(--color-dark-grey);
  }

  .table-data-body {
    background-color: var(--color-dark-grey);
  }

  .table-data-cell {
    padding: 0.2rem 0.4rem;
    text-wrap: nowrap;
  }

  .table-data-head .table-data-cell {
    font-weight: 600;
    border-top: 2px solid var(--color-dark-grey);
    border-bottom: 2px solid var(--color-dark-grey);
    border-left: 2px solid var(--color-dark-grey);
  }

  .table-data-body .table-data-cell {
    border-bottom: 2px solid var(--color-almost-black);
    border-left: 2px solid var(--color-almost-black);
  }

  .table-data-cell-null {
    color: var(--color-medium-gray);
  }

  .table-data-cell-blue {
    color: var(--color-carolina-blue);
  }

  .table-data-cell-green {
    color: var(--color-dark-sage);
  }

  .table-data-cell-orange {
    color: var(--color-dark-peach);
  }

  .table-data-cell-yellow {
    color: var(--color-light-mustard);
  }
</style>
