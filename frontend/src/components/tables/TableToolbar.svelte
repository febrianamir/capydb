<script>
  import { handleEnter } from "../../utils/key";
  import {
    RefreshCcw,
    ChevronLeft,
    ChevronRight,
    Search,
  } from "@lucide/svelte";

  let {
    currentPage,
    lastPage,
    queryTableContents,
    updateQueryTableContents,
    paginationPrevPage,
    paginationNextPage,
    toggleTableFilter,
    tableColumns,
    tableRecordsCount,
    tableRecordsTotal,
  } = $props();
</script>

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
    font-size: 16px;
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
</style>
