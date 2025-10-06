<script>
  import { onMount } from "svelte";
  import { handleEnter } from "../../utils/key";
  import {
    GetTableColumns,
    GetTableRecords,
    GetTables,
  } from "../../../wailsjs/go/usecase/Usecase";
  import {
    Table,
    ArrowDownNarrowWide,
    ArrowUpWideNarrow,
  } from "@lucide/svelte";

  let isShowTableList = $state(false);
  let tables = $state([]);
  let tableColumns = $state([]);
  let tableRecords = $state([]);
  let queryTableContents = $state({
    table_name: "",
    sort_by: "",
    order_by: "",
  });
  function updateQueryTableContents(query) {
    queryTableContents = query;
  }

  onMount(async () => {
    try {
      tables = await GetTables();
      isShowTableList = true;
    } catch (err) {
      console.log("Failed to get table list:", err);
    }
  });

  $effect(() => {
    getTableContents(queryTableContents);
  });

  async function getTableContents(query) {
    if (query.table_name == "") {
      return;
    }

    try {
      tableColumns = await GetTableColumns(query.table_name);

      let getTableRecordsReq = {
        table_name: query.table_name,
        limit: 500,
        sort_by: query.sort_by,
        order_by: query.order_by,
      };
      let tableRecordsRes = await GetTableRecords(getTableRecordsReq);
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

  function toggleColumnSort(tableColumn) {
    // If sort_by changed, set order_by back to ASC
    // If sort_by unchanged, toggle order_by
    let sortBy = tableColumn.column_name;
    let orderBy = "";
    if (tableColumn.column_name != queryTableContents.sort_by) {
      orderBy = "ASC";
    } else {
      if (queryTableContents.order_by == "ASC") {
        orderBy = "DESC";
      } else {
        sortBy = "";
        orderBy = "";
      }
    }
    updateQueryTableContents({
      table_name: queryTableContents.table_name,
      sort_by: sortBy,
      order_by: orderBy,
    });
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
                updateQueryTableContents({
                  table_name: table,
                  sort_by: "",
                  order_by: "",
                });
              }}
              onkeydown={(e) => {
                handleEnter(e, () => {
                  e.preventDefault();
                  updateQueryTableContents({
                    table_name: table,
                    sort_by: "",
                    order_by: "",
                  });
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
              <th
                class="table-data-cell"
                onclick={(e) => {
                  e.preventDefault();
                  toggleColumnSort(tableColumn);
                }}
              >
                {tableColumn.column_name}
                <div
                  class="table-data-cell-sort {tableColumn.column_name ==
                  queryTableContents.sort_by
                    ? 'active'
                    : ''}"
                >
                  {#if queryTableContents.order_by == "ASC"}
                    <div class="table-data-cell-icon">
                      <ArrowDownNarrowWide size="16" />
                    </div>
                  {:else}
                    <div class="table-data-cell-icon">
                      <ArrowUpWideNarrow size="16" />
                    </div>
                  {/if}
                </div>
              </th>
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
    cursor: pointer;
    padding-right: 2rem;
    position: relative;
  }

  .table-data-cell-sort {
    position: absolute;
    top: 0.2rem;
    right: 0.4rem;
    visibility: hidden;
  }

  .table-data-cell-sort.active {
    visibility: visible;
  }

  .table-data-cell-icon {
    display: flex;
    align-items: center;
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
