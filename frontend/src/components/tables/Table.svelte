<script>
  import { onMount } from "svelte";
  import { handleEnter } from "../../utils/key";
  import {
    GetTableColumns,
    GetTableRecords,
    GetTables,
  } from "../../../wailsjs/go/usecase/Usecase";
  import {
    RefreshCcw,
    ArrowDownNarrowWide,
    ArrowUpWideNarrow,
    ChevronLeft,
    ChevronRight,
    Search,
    CirclePlus,
    CircleMinus,
    Check,
    ChevronsUpDown,
  } from "@lucide/svelte";
  import TableSidebar from "./TableSidebar.svelte";

  const operators = [
    "=",
    "!=",
    "<",
    ">",
    "<=",
    ">=",
    "IN",
    "NOT IN",
    "LIKE",
    "BETWEEN",
    "IS NULL",
    "IS NOT NULL",
  ];
  const oneColumnOperators = [
    "=",
    "!=",
    "<",
    ">",
    "<=",
    ">=",
    "IN",
    "NOT IN",
    "LIKE",
  ];
  const twoColumnOperators = ["BETWEEN"];

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
      offset: queryTableContents.offset,
    });
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

  function updateFilterActiveStatus(index, checked) {
    filters = filters.map((f, i) =>
      i === index ? { ...f, isActive: checked } : f,
    );
  }

  function paginationPrevPage() {
    let newOffset = queryTableContents.offset - 500;
    if (newOffset < 0) {
      newOffset = 0;
    }

    updateQueryTableContents({
      table_name: queryTableContents.table_name,
      sort_by: queryTableContents.sort_by,
      order_by: queryTableContents.order_by,
      offset: newOffset,
    });
  }

  function paginationNextPage() {
    let prevOffset = queryTableContents.offset;
    let newOffset = queryTableContents.offset + 500;
    if (newOffset >= tableRecordsTotal) {
      newOffset = prevOffset;
    }

    updateQueryTableContents({
      table_name: queryTableContents.table_name,
      sort_by: queryTableContents.sort_by,
      order_by: queryTableContents.order_by,
      offset: newOffset,
    });
  }

  // resizeSelectInput width based on choosed value
  function resizeSelectInput(e) {
    const select = e.target;
    const mirror = select.nextElementSibling; // Assumes mirror <span> right after select
    const selected = select.options[select.selectedIndex];

    mirror.textContent = selected.text;
    select.style.width = mirror.offsetWidth + 32 + "px"; // Add padding/arrow width
  }

  function addFilterItem() {
    filters.push({
      isActive: true,
      operator: "=",
      firstValue: "",
      secondValue: "",
    });
  }

  function closeFilterItem(idx) {
    filters.splice(idx, 1);
  }

  function filterTableRecords() {
    let conditions = [];
    filters.forEach((filter) => {
      if (filter.isActive && filter.field && filter.operator) {
        conditions.push({
          field: filter.field,
          operator: filter.operator,
          first_value: filter.firstValue,
          second_value: filter.secondValue,
        });
      }
    });
    updateQueryTableContents({
      table_name: queryTableContents.table_name,
      sort_by: queryTableContents.sort_by,
      order_by: queryTableContents.order_by,
      offset: queryTableContents.offset,
      conditions: conditions,
    });
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
      <div class="table-toolbar">
        <div
          role="button"
          tabindex="0"
          class="table-refresh"
          onclick={(e) => {
            e.preventDefault();
            updateQueryTableContents({
              table_name: queryTableContents.table_name,
              sort_by: queryTableContents.sort_by,
              order_by: queryTableContents.order_by,
              offset: queryTableContents.offset,
            });
          }}
          onkeydown={(e) => {
            handleEnter(e, () => {
              e.preventDefault();
              updateQueryTableContents({
                table_name: queryTableContents.table_name,
                sort_by: queryTableContents.sort_by,
                order_by: queryTableContents.order_by,
                offset: queryTableContents.offset,
              });
            });
          }}
        >
          <div class="table-refresh-icon">
            <RefreshCcw size="19" strokeWidth="2.5" />
          </div>
        </div>
        <div class="table-pagination">
          <div
            class="table-pagination-prev {currentPage == 1 ? 'disable' : ''}"
            role="button"
            tabindex="0"
            onclick={(e) => {
              e.preventDefault();
              paginationPrevPage();
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.preventDefault();
                paginationPrevPage();
              });
            }}
          >
            <div class="table-pagination-icon">
              <ChevronLeft size="19" />
            </div>
          </div>
          <div class="table-pagination-page">
            {currentPage}
          </div>
          <div
            class="table-pagination-next {currentPage == lastPage
              ? 'disable'
              : ''}"
            role="button"
            tabindex="0"
            onclick={(e) => {
              e.preventDefault();
              paginationNextPage();
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.preventDefault();
                paginationNextPage();
              });
            }}
          >
            <div class="table-pagination-icon">
              <ChevronRight size="19" />
            </div>
          </div>
        </div>
        <div
          role="button"
          tabindex="0"
          class="table-toggle-filter"
          onclick={(e) => {
            e.preventDefault();
            toggleTableFilter(tableColumns[0].column_name);
          }}
          onkeydown={(e) => {
            handleEnter(e, () => {
              e.preventDefault();
              toggleTableFilter(tableColumns[0].column_name);
            });
          }}
        >
          <div class="table-toggle-filter-icon">
            <Search size="19" strokeWidth="2.5" />
          </div>
        </div>
        <div class="table-result-info">
          <div class="table-result-info-count">
            Results: {tableRecordsCount}
          </div>
          <div class="table-result-info-total">Total: {tableRecordsTotal}</div>
        </div>
      </div>
      <div class="table-filter {isShowTableFilter ? 'active' : ''}">
        <div class="table-filter-list">
          {#each filters as filter, i}
            <div class="table-filter-item">
              <div class="table-filter-item-check">
                <input
                  type="checkbox"
                  name="active"
                  onchange={(e) =>
                    updateFilterActiveStatus(i, e.target.checked)}
                  class="table-filter-check-input"
                />
                <div
                  class="table-filter-check-icon {filter.isActive
                    ? 'active'
                    : ''}"
                >
                  <Check size="14" strokeWidth="3" />
                </div>
              </div>
              <div class="table-filter-item-inputs">
                <div class="table-filter-field">
                  <select
                    name="field"
                    class="table-filter-field-select"
                    bind:value={filter.field}
                    onchange={(e) => resizeSelectInput(e)}
                  >
                    {#each tableColumns as tableColumn}
                      <option value={tableColumn.column_name}>
                        {tableColumn.column_name}
                      </option>
                    {/each}
                  </select>
                  <!-- To mirror choosed option width -->
                  <span
                    style="position:absolute;visibility:hidden;white-space:pre;font:inherit;"
                  ></span>
                  <div class="table-filter-field-icon">
                    <ChevronsUpDown size="15" strokeWidth="3" />
                  </div>
                </div>
                <div class="table-filter-operator">
                  <select
                    name="operator"
                    class="table-filter-operator-select"
                    bind:value={filter.operator}
                    onchange={(e) => resizeSelectInput(e)}
                  >
                    {#each operators as operator}
                      <option value={operator}>{operator}</option>
                    {/each}
                  </select>
                  <!-- To mirror choosed option width -->
                  <span
                    style="position:absolute;visibility:hidden;white-space:pre;font:inherit;"
                  ></span>
                  <div class="table-filter-operator-icon">
                    <ChevronsUpDown size="15" strokeWidth="3" />
                  </div>
                </div>
                {#if oneColumnOperators.includes(filter.operator) || twoColumnOperators.includes(filter.operator)}
                  <input
                    type="text"
                    name="value"
                    class="table-filter-value"
                    bind:value={filter.firstValue}
                  />
                {/if}
                {#if twoColumnOperators.includes(filter.operator)}
                  <input
                    type="text"
                    name="value"
                    class="table-filter-value"
                    bind:value={filter.secondValue}
                  />
                {/if}
              </div>
              <div
                class="table-filter-item-close"
                role="button"
                tabindex="0"
                onclick={(e) => {
                  e.preventDefault();
                  closeFilterItem(i);
                }}
                onkeydown={(e) => {
                  handleEnter(e, () => {
                    e.preventDefault();
                    closeFilterItem(i);
                  });
                }}
              >
                <CircleMinus size="15" />
              </div>
            </div>
          {/each}
        </div>
        <div class="table-filter-buttons">
          <div
            class="table-filter-button-add"
            role="button"
            tabindex="0"
            onclick={(e) => {
              e.preventDefault();
              addFilterItem();
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.preventDefault();
                addFilterItem();
              });
            }}
          >
            <CirclePlus size="15" />
          </div>
          <div
            class="table-filter-button-submit"
            role="button"
            tabindex="0"
            onclick={(e) => {
              e.preventDefault();
              filterTableRecords();
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.preventDefault();
                filterTableRecords();
              });
            }}
          >
            Filter
          </div>
        </div>
      </div>
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

  .table-toolbar {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 0.75rem 0 0.75rem;
  }

  .table-refresh {
    padding: 0.4rem;
    border-radius: 0.5rem;
    background-color: var(--color-dark-grey-2);
    border: 1px solid var(--color-dark-grey-2);
    cursor: pointer;
    transition: 0.2s ease;
  }

  .table-refresh:hover {
    background-color: var(--color-dark-grey);
  }

  .table-refresh-icon {
    display: flex;
    align-items: center;
  }

  .table-pagination {
    display: flex;
    align-items: center;
    font-size: 16px;
  }

  .table-pagination-prev,
  .table-pagination-page,
  .table-pagination-next {
    padding: 0.4rem;
    cursor: pointer;
    width: 34px;
    height: 34px;
    background-color: var(--color-dark-grey-2);
    transition: 0.2s ease;
    border: 1px solid var(--color-dark-grey-2);
    text-align: center;
  }

  .table-pagination-prev:hover,
  .table-pagination-page:hover,
  .table-pagination-next:hover {
    background-color: var(--color-dark-grey);
  }

  .table-pagination-prev.disable,
  .table-pagination-next.disable {
    background-color: var(--color-dark-grey);
    color: var(--color-medium-gray);
  }

  .table-pagination-prev {
    border-top-left-radius: 0.5rem;
    border-bottom-left-radius: 0.5rem;
  }

  .table-pagination-next {
    border-top-right-radius: 0.5rem;
    border-bottom-right-radius: 0.5rem;
  }

  .table-pagination .table-pagination-icon {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  /* Table Filter */
  .table-toggle-filter {
    padding: 0.4rem;
    border-radius: 0.5rem;
    background-color: var(--color-dark-grey-2);
    border: 1px solid var(--color-dark-grey-2);
    cursor: pointer;
    transition: 0.2s ease;
  }

  .table-toggle-filter:hover {
    background-color: var(--color-dark-grey);
  }

  .table-toggle-filter-icon {
    display: flex;
    align-items: center;
  }

  .table-filter {
    display: none;
    margin-bottom: 0.5rem;
  }

  .table-filter.active {
    display: block;
  }

  .table-filter-item {
    display: flex;
    padding: 1rem 1rem 0 1rem;
    align-items: center;
    gap: 1rem;
  }

  .table-filter-item-inputs {
    display: flex;
    gap: 0.75rem;
    align-items: center;
    width: 100%;
  }

  .table-filter-item-check {
    position: relative;
  }

  .table-filter-check-input {
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    width: 1.2rem;
    height: 1.2rem;
    margin: 0;
    border: 1px solid var(--color-dark-grey-2);
    border-radius: 0.4rem;
    cursor: pointer;
    transition: 0.2s ease;
  }

  .table-filter-check-icon {
    display: flex;
    opacity: 0;
    position: absolute;
    top: 0;
    left: 0;
    width: 1.2rem;
    height: 1.2rem;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    pointer-events: none;
    transition: 0.1s ease;
  }

  .table-filter-check-icon.active {
    opacity: 1;
    border-radius: 0.4rem;
    color: #fff;
    border: 1px solid var(--color-carolina-blue-1);
    background-color: var(--color-carolina-blue-1);
  }

  .table-filter-field,
  .table-filter-operator {
    position: relative;
  }

  .table-filter-field-select,
  .table-filter-operator-select {
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    background-color: var(--color-almost-black);
    color: var(--color-text);
    border: 1px solid var(--color-dark-grey-2);
    padding: 0.3rem 1.5rem 0.4rem 0.45rem;
    border-radius: 0.4rem;
  }

  .table-filter-field-icon,
  .table-filter-operator-icon {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.4rem 0.3rem;
    pointer-events: none;
    background-color: transparent;
  }

  .table-filter-value {
    background-color: var(--color-almost-black);
    color: var(--color-text);
    border: 1px solid var(--color-dark-grey-2);
    border-radius: 0.4rem;
    transition: 0.2s ease;
    padding: 0.3rem 0.4rem;
    outline: none;
    flex: 1;
  }

  .table-filter-value:focus {
    border: 1px solid var(--color-carolina-blue);
  }

  .table-filter-item-close,
  .table-filter-button-add {
    display: flex;
    align-items: center;
    padding: 0.3rem;
    border-radius: 0.3rem;
    background-color: var(--color-dark-grey-2);
    border: 1px solid var(--color-dark-grey-2);
    cursor: pointer;
    transition: 0.2s ease;
  }

  .table-filter-item-close:hover,
  .table-filter-button-add:hover {
    background-color: var(--color-dark-grey);
  }

  .table-filter-buttons {
    display: flex;
    justify-content: flex-end;
    padding: 0 1rem;
    margin-top: 1rem;
    gap: 0.75rem;
    align-items: center;
  }

  .table-filter-button-submit {
    padding: 0.4rem 0.6rem;
    cursor: pointer;
    color: #fff;
    background-color: var(--color-carolina-blue-1);
    border-radius: 0.3rem;
    transition: 0.2s ease;
  }

  .table-filter-button-submit:hover {
    background-color: var(--color-carolina-blue-2);
  }

  .table-result-info {
    display: flex;
    gap: 1rem;
    font-size: 16px;
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
