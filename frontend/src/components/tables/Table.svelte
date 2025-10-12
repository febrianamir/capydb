<script>
  import { onMount } from "svelte";
  import {
    GetTableColumns,
    GetTableRecords,
    GetTables,
  } from "../../../wailsjs/go/usecase/Usecase";
  import TableSidebar from "./TableSidebar.svelte";
  import TableToolbar from "./TableToolbar.svelte";
  import TableFilter from "./TableFilter.svelte";
  import TableRecord from "./TableRecord.svelte";

  let isShowTableList = $state(false);
  let isShowTableFilter = $state(false);
  let tables = $state([]);
  let tableColumns = $state([]);
  let tableRecords = $state([]);
  let tableRecordsCount = $state(0);
  let tableRecordsTotal = $state(0);
  let queryTableContents = $state({
    table_name: "",
    sort_by: "",
    order_by: "",
    offset: 0,
    conditions: [],
  });
  function updateQueryTableContents(query) {
    queryTableContents = query;
  }
  let filters = $state([]);
  let currentPage = $derived.by(() => {
    if (queryTableContents.offset == 0) {
      return 1;
    }
    return queryTableContents.offset / 500 + 1;
  });
  let lastPage = $derived(Math.ceil(tableRecordsTotal / 500));

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
        offset: query.offset,
        sort_by: query.sort_by,
        order_by: query.order_by,
        conditions: query.conditions,
      };
      let tableRecordsRes = await GetTableRecords(getTableRecordsReq);
      tableRecords = tableRecordsRes.Data;
      tableRecordsCount = tableRecordsRes.Data.length;
      tableRecordsTotal = tableRecordsRes.Total;
    } catch (err) {
      console.log("Failed to get table contents:", err);
    }
  }
</script>

<div class="table">
  <TableSidebar
    {isShowTableList}
    {tables}
    {updateQueryTableContents}
    {queryTableContents}
  />
  <div class="table-explorer">
    {#if tableColumns.length > 0}
      <TableToolbar
        {currentPage}
        {lastPage}
        {queryTableContents}
        {updateQueryTableContents}
        {tableColumns}
        {tableRecordsCount}
        {tableRecordsTotal}
        {filters}
        {isShowTableFilter}
      />
      <TableFilter
        {isShowTableFilter}
        {filters}
        {updateQueryTableContents}
        {queryTableContents}
        {tableColumns}
      />
      <TableRecord
        {queryTableContents}
        {updateQueryTableContents}
        {tableColumns}
        {tableRecords}
      />
    {/if}
  </div>
</div>

<style>
  :root {
    --color-almost-black: #1d1d1d;
    --color-dark-grey: #252525;
    --color-dark-grey-2: #3f3f3f;
    --color-carolina-blue: #7591df;
    --color-carolina-blue-1: #6883db;
    --color-carolina-blue-2: #5875d3;
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

  /* Table Explorer */
  .table-explorer {
    background-color: var(--color-almost-black);
    flex: 1;
    overflow: auto;
    text-align: left;
    font-size: 14px;
  }
</style>
