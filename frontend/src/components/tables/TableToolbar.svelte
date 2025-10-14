<script>
  import { handleEnter } from "../../utils/key";
  import {
    RefreshCcw,
    ChevronLeft,
    ChevronRight,
    Search,
    SlidersHorizontal,
  } from "@lucide/svelte";

  let {
    currentPage,
    lastPage,
    queryTableContents,
    updateQueryTableContents,
    toggleTableFilter,
    tableColumns,
    tableRecordsCount,
    tableRecordsTotal,
    setColumnsVisibility,
  } = $props();

  let isShowColumnVisibilityInput = $state(false);
  let columnsVisibilityInputRel;

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

  function refreshTableRecords() {
    updateQueryTableContents({
      table_name: queryTableContents.table_name,
      sort_by: queryTableContents.sort_by,
      order_by: queryTableContents.order_by,
      offset: queryTableContents.offset,
      conditions: queryTableContents.conditions,
    });
  }

  function toggleColumnVisibility() {
    if (isShowColumnVisibilityInput) {
      let columns = columnsVisibilityInputRel.value;
      setColumnsVisibility(columns);
    }
    isShowColumnVisibilityInput = !isShowColumnVisibilityInput;
  }
</script>

<div class="table-toolbar">
  <div
    role="button"
    tabindex="0"
    class="table-refresh"
    onclick={(e) => {
      e.preventDefault();
      refreshTableRecords();
    }}
    onkeydown={(e) => {
      handleEnter(e, () => {
        e.preventDefault();
        refreshTableRecords();
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
      class="table-pagination-next {currentPage == lastPage ? 'disable' : ''}"
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
  <div
    role="button"
    tabindex="0"
    class="table-toggle-column"
    onclick={(e) => {
      e.preventDefault();
      toggleColumnVisibility();
    }}
    onkeydown={(e) => {
      handleEnter(e, () => {
        e.preventDefault();
        toggleColumnVisibility();
      });
    }}
  >
    <div class="table-toggle-column-icon">
      <SlidersHorizontal size="19" strokeWidth="2.5" />
    </div>
    <div class="table-toggle-column-text">Columns</div>
    <div
      role="button"
      tabindex="0"
      class="table-column-visibility {isShowColumnVisibilityInput
        ? 'active'
        : ''}"
      onclick={(e) => {
        e.preventDefault();
        e.stopPropagation();
      }}
      onkeydown={(e) => {
        handleEnter(e, () => {
          e.preventDefault();
          e.stopPropagation();
        });
      }}
    >
      <div class="table-column-visibility-info">
        Select columns to show, separate by comma
      </div>
      <textarea
        class="table-column-visibility-input"
        bind:this={columnsVisibilityInputRel}
        onkeydown={(e) => {
          handleEnter(e, () => {
            e.preventDefault();
            toggleColumnVisibility();
          });
        }}
      ></textarea>
    </div>
  </div>
  <div class="table-toolbar-separator"></div>
  <div class="table-result-info">
    <div class="table-result-info-count">
      Results: {tableRecordsCount}
    </div>
    <div class="table-result-info-total">Total: {tableRecordsTotal}</div>
  </div>
</div>

<style>
  /* Table Toolbar  */
  .table-toolbar {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 0.75rem 0 0.75rem;
  }

  /* Table Refresh */
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

  /* Table Column */
  .table-toggle-column {
    position: relative;
    display: flex;
    align-items: center;
    padding: 0.35rem 0.6rem;
    border-radius: 0.5rem;
    background-color: var(--color-dark-grey-2);
    border: 1px solid var(--color-dark-grey-2);
    cursor: pointer;
    gap: 0.5rem;
    font-size: 15px;
    transition: 0.2s ease;
  }

  .table-toggle-column-icon {
    display: flex;
    align-items: center;
  }

  .table-column-visibility {
    position: absolute;
    opacity: 0;
    top: 0;
    left: 0;
    padding: 1rem;
    padding: 0.35rem 0.6rem 0.6rem 0.6rem;
    border-radius: 0.5rem;
    pointer-events: none;
    transition: 0.2s ease;
    background-color: var(--color-dark-grey);
    border: 1px solid var(--color-dark-grey-2);
    width: 300px;
    font-size: 13px;
    z-index: 1000;
  }

  .table-column-visibility.active {
    top: 37px;
    opacity: 1;
    pointer-events: all;
  }

  .table-column-visibility-input {
    outline: none;
    width: 100%;
    min-height: 5rem;
    border: 1px solid var(--color-dark-grey-2);
    border-radius: 0.2rem;
    background-color: var(--color-bg);
    color: var(--color-text);
    margin-top: 0.5rem;
    padding: 0.2rem;
    font-size: 12px;
    font-family: "JetBrains Mono", monospace;
  }

  /* Table Pagination */
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

  /* Table Result Info */
  .table-result-info {
    display: flex;
    gap: 1rem;
    font-size: 15px;
    padding-right: 1rem;
  }

  /* Table Toggle Filter */
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

  /* Table Toolbar Separator  */
  .table-toolbar-separator {
    flex: 1;
  }
</style>
