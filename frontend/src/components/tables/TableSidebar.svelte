<script>
  import { handleEnter } from "../../utils/key";
  import { Table } from "@lucide/svelte";

  let {
    isShowTableList,
    tables,
    updateQueryTableContents,
    queryTableContents,
  } = $props();
</script>

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
                offset: queryTableContents.offset,
              });
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.preventDefault();
                updateQueryTableContents({
                  table_name: table,
                  sort_by: "",
                  order_by: "",
                  offset: queryTableContents.offset,
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

<style>
  /* Table Sidebar */
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
</style>
