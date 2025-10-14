<script>
  import { onMount } from "svelte";
  import {
    GetTableColumns,
    GetTableRecords,
  } from "../../../wailsjs/go/usecase/Usecase";
  import TableToolbar from "./TableToolbar.svelte";
  import TableFilter from "./TableFilter.svelte";
  import TableRecord from "./TableRecord.svelte";
  import { getTableTabs } from "../../states/tables.svelte";

  let { tableName } = $props();
  let tableTabs = getTableTabs();
  let isShowTableFilter = $state(false);
  let tableColumns = $state([]);
  let tableRecords = $state([]);
  let tableRecordsCount = $state(0);
  let tableRecordsTotal = $state(0);
  let queryTableContents = $state({
    table_name: tableName,
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
  // Empty columnsVisibility means every columns is visible
  let columnsVisibility = $state([]);

  $effect(() => {
    getTableContents(queryTableContents);
  });

  onMount(() => {
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
        limit: 500, // Default limit per page
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

  function toggleTableFilter(defaultColumn) {
    isShowTableFilter = !isShowTableFilter;
    if (filters.length == 0) {
      filters.push({
        isActive: true,
        field: defaultColumn,
        operator: "=",
        firstValue: "",
        secondValue: "",
      });
    }
  }

  function setColumnsVisibility(columns) {
    let trimmedColumns = columns.trim();
    if (trimmedColumns.length > 0) {
      columnsVisibility = trimmedColumns.split(",").map((s) => s.trim());
    } else {
      columnsVisibility = [];
    }
  }
</script>

<div class="table-tab {tableTabs.openedTab == tableName ? 'active' : ''}">
  {#if tableColumns.length > 0}
    <TableToolbar
      {currentPage}
      {lastPage}
      {queryTableContents}
      {updateQueryTableContents}
      {toggleTableFilter}
      {tableColumns}
      {tableRecordsCount}
      {tableRecordsTotal}
      {setColumnsVisibility}
    />
    <TableFilter
      {isShowTableFilter}
      {filters}
      {updateQueryTableContents}
      {queryTableContents}
      {tableColumns}
    />
    <TableRecord
      {columnsVisibility}
      {queryTableContents}
      {updateQueryTableContents}
      {tableColumns}
      {tableRecords}
    />
  {/if}
</div>

<style>
  .table-tab {
    display: none;
  }

  .table-tab.active {
    display: flex;
    flex-direction: column;
    flex: 1;
    min-height: 0;
  }
</style>
