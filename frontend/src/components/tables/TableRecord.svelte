<script>
  import { ArrowDownNarrowWide, ArrowUpWideNarrow } from "@lucide/svelte";
  import { onMount, onDestroy } from "svelte";

  let {
    queryTableContents,
    updateQueryTableContents,
    tableColumns,
    tableRecords,
    columnsVisibility,
  } = $props();

  let cols = [];
  let isColResizeActive = $state(false);
  let colResizeActiveIndex = $state(0);
  let startX = 0;
  let startWidth = 0;

  function getColumnTypeColorClass(type, data) {
    if (data === null) {
      return "table-record-cell-null";
    }

    switch (type) {
      case "integer":
      case "numeric":
      case "double precision":
        return "table-record-cell-blue";
      case "character varying":
      case "text":
        return "table-record-cell-green";
      case "timestamp with time zone":
      case "timestamp without time zone":
        return "table-record-cell-orange";
      case "boolean":
        return "table-record-cell-yellow";
    }
  }

  function getColumnInitialWidth(columnName, type) {
    let baseWidth = 0;
    switch (type) {
      case "integer":
      case "numeric":
      case "double precision":
        baseWidth = 30;
        break;
      case "character varying":
      case "text":
        baseWidth = 270;
        break;
      case "timestamp with time zone":
      case "timestamp without time zone":
        baseWidth = 270;
        break;
      case "boolean":
        baseWidth = 40;
        break;
    }
    return measureHeaderWidth(columnName, baseWidth);
  }

  function measureHeaderWidth(name, baseWidth) {
    const estimated = name.length * 10 + 20; // Roughly 10px per char, add 20px because header cell has additional padding-right to show sorting icon
    return Math.max(baseWidth, Math.min(estimated, 400));
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

  function isColumnVisible(columnName) {
    return (
      columnsVisibility.length == 0 ||
      (columnsVisibility.length > 0 && columnsVisibility.includes(columnName))
    );
  }

  function startResize(e, colIdx) {
    e.preventDefault();
    isColResizeActive = true;
    colResizeActiveIndex = colIdx;
    startX = e.clientX;
    startWidth = parseInt(getComputedStyle(cols[colIdx]).width, 10);
  }

  function stopResize(e) {
    e.preventDefault();
    if (!isColResizeActive) return;
    // Set small timeout/sleep to prevent header click event triggered
    // See .table-record-cell onclick event
    setTimeout(() => {
      isColResizeActive = false;
      colResizeActiveIndex = 0;
    }, 10);
  }

  function onMouseMove(e) {
    e.preventDefault();
    if (!isColResizeActive) return;
    const diff = e.clientX - startX;
    const newWidth = Math.max(60, startWidth + diff);
    cols[colResizeActiveIndex].style.width = `${newWidth}px`;
  }

  onMount(() => {
    window.addEventListener("mousemove", onMouseMove);
    window.addEventListener("mouseup", stopResize);
  });

  onDestroy(() => {
    window.removeEventListener("mousemove", onMouseMove);
    window.removeEventListener("mouseup", stopResize);
  });
</script>

<div class="table-record-container">
  <table class="table-record">
    <colgroup>
      {#each tableColumns as tableColumn, i}
        {#if isColumnVisible(tableColumn.column_name)}
          <col
            bind:this={cols[i]}
            style="width: {getColumnInitialWidth(
              tableColumn.column_name,
              tableColumn.data_type,
            )}px;"
          />
        {/if}
      {/each}
    </colgroup>
    <thead class="table-record-head">
      <tr>
        {#each tableColumns as tableColumn, i}
          {#if isColumnVisible(tableColumn.column_name)}
            <th
              class="table-record-cell"
              onclick={(e) => {
                e.preventDefault();
                if (isColResizeActive) return;
                toggleColumnSort(tableColumn);
              }}
            >
              {tableColumn.column_name}
              <div
                class="table-record-cell-sort {tableColumn.column_name ==
                queryTableContents.sort_by
                  ? 'active'
                  : ''}"
              >
                {#if queryTableContents.order_by == "ASC"}
                  <div class="table-record-cell-icon">
                    <ArrowDownNarrowWide size="16" />
                  </div>
                {:else}
                  <div class="table-record-cell-icon">
                    <ArrowUpWideNarrow size="16" />
                  </div>
                {/if}
              </div>
              <div
                role="button"
                tabindex="0"
                class="table-record-cell-resizer"
                onmousedown={(e) => startResize(e, i)}
              ></div>
            </th>
          {/if}
        {/each}
      </tr>
    </thead>
    <tbody class="table-record-body">
      {#if tableRecords.length > 0}
        {#each tableRecords as tableRecord}
          <tr>
            {#each tableColumns as tableColumn}
              {#if isColumnVisible(tableColumn.column_name)}
                <td
                  class="table-record-cell {getColumnTypeColorClass(
                    tableColumn.data_type,
                    tableRecord[tableColumn.column_name],
                  )}"
                >
                  {tableRecord[tableColumn.column_name] === null
                    ? "NULL"
                    : tableRecord[tableColumn.column_name]}
                </td>
              {/if}
            {/each}
          </tr>
        {/each}
      {/if}
    </tbody>
  </table>
</div>

<style>
  /* Table Record */
  .table-record-container {
    flex: 1;
    overflow: auto;
  }

  .table-record {
    width: 100%;
    table-layout: fixed;
    border-spacing: 0;
    padding: 0 0.75rem 1.5rem 0.75rem;
  }

  .table-record-head {
    background-color: var(--color-almost-black);
    border: 2px solid var(--color-dark-grey);
  }

  .table-record-body {
    background-color: var(--color-dark-grey);
  }

  .table-record-cell {
    position: relative;
    padding: 0.2rem 0.4rem;
    text-wrap: nowrap;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
  }

  .table-record-head .table-record-cell {
    font-weight: 600;
    border-top: 2px solid var(--color-dark-grey);
    border-bottom: 2px solid var(--color-dark-grey);
    border-left: 2px solid var(--color-dark-grey);
    cursor: pointer;
    padding-right: 1.5rem;
    position: relative;
  }

  .table-record-cell-sort {
    position: absolute;
    top: 0.2rem;
    right: 0.4rem;
    visibility: hidden;
  }

  .table-record-cell-sort.active {
    visibility: visible;
  }

  .table-record-cell-icon {
    display: flex;
    align-items: center;
  }

  .table-record-body .table-record-cell {
    border-bottom: 2px solid var(--color-almost-black);
    border-left: 2px solid var(--color-almost-black);
  }

  .table-record-cell-null {
    color: var(--color-medium-gray);
  }

  .table-record-cell-blue {
    color: var(--color-carolina-blue);
  }

  .table-record-cell-green {
    color: var(--color-dark-sage);
  }

  .table-record-cell-orange {
    color: var(--color-dark-peach);
  }

  .table-record-cell-yellow {
    color: var(--color-light-mustard);
  }

  /* Table Column Resizer */
  .table-record-cell-resizer {
    position: absolute;
    top: 0;
    right: 0;
    width: 5px;
    cursor: ew-resize;
    height: 100%;
    flex-shrink: 0;
    background-color: #1e1e1e;
  }
</style>
