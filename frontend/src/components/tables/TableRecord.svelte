<script>
  import { ArrowDownNarrowWide, ArrowUpWideNarrow } from "@lucide/svelte";

  let {
    queryTableContents,
    updateQueryTableContents,
    tableColumns,
    tableRecords,
    columnsVisibility,
  } = $props();

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
</script>

<table class="table-record">
  <thead class="table-record-head">
    <tr>
      {#each tableColumns as tableColumn}
        {#if columnsVisibility.length == 0 || (columnsVisibility.length > 0 && columnsVisibility.includes(tableColumn.column_name))}
          <th
            class="table-record-cell"
            onclick={(e) => {
              e.preventDefault();
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
            {#if columnsVisibility.length == 0 || (columnsVisibility.length > 0 && columnsVisibility.includes(tableColumn.column_name))}
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

<style>
  /* Table Record */
  .table-record {
    width: max-content;
    border-spacing: 0;
    padding: 0.75rem;
  }

  .table-record-head {
    background-color: var(--color-almost-black);
    border: 2px solid var(--color-dark-grey);
  }

  .table-record-body {
    background-color: var(--color-dark-grey);
  }

  .table-record-cell {
    padding: 0.2rem 0.4rem;
    text-wrap: nowrap;
  }

  .table-record-head .table-record-cell {
    font-weight: 600;
    border-top: 2px solid var(--color-dark-grey);
    border-bottom: 2px solid var(--color-dark-grey);
    border-left: 2px solid var(--color-dark-grey);
    cursor: pointer;
    padding-right: 2rem;
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
</style>
