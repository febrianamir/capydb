<script>
  import { handleEnter } from "../../utils/key";
  import { Table } from "@lucide/svelte";
  import { onMount, onDestroy } from "svelte";
  import { GetTables } from "../../../wailsjs/go/usecase/Usecase";
  import { addTableTab } from "../../states/tables.svelte";
  import { connection } from "../../states/connection.svelte";

  let tables = $state([]);
  let isShowTableList = $state(false);
  let isTableSidebarResizing = $state(false);
  let tableSidebarWidth = $state(250); // Default width

  function startResize(e) {
    e.preventDefault();
    isTableSidebarResizing = true;
  }

  function stopResize() {
    isTableSidebarResizing = false;
  }

  function onMouseMove(e) {
    if (isTableSidebarResizing) {
      tableSidebarWidth = Math.max(200, e.clientX); // Minimum width = 200px
    }
  }

  onMount(async () => {
    try {
      let req = {
        connection_id: connection.current_connection.connection_id,
      };
      tables = await GetTables(req);
      isShowTableList = true;
    } catch (err) {
      console.log("Failed to get table list:", err);
    }
  });

  onMount(() => {
    window.addEventListener("mousemove", onMouseMove);
    window.addEventListener("mouseup", stopResize);
  });

  onDestroy(() => {
    window.removeEventListener("mousemove", onMouseMove);
    window.removeEventListener("mouseup", stopResize);
  });
</script>

<div class="table-sidebar">
  <div class="table-sidebar-content" style="width: {tableSidebarWidth}px">
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
                addTableTab(table);
              }}
              onkeydown={(e) => {
                handleEnter(e, () => {
                  e.preventDefault();
                  addTableTab(table);
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
  <div
    class="table-sidebar-resizer"
    role="button"
    tabindex="0"
    class:active={isTableSidebarResizing}
    onmousedown={startResize}
  ></div>
</div>

<style>
  /* Table Sidebar */
  .table-sidebar {
    display: flex;
  }

  .table-sidebar-content {
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

  /* Table Sidebar Resizer */
  .table-sidebar-resizer {
    width: 2px;
    cursor: ew-resize;
    height: 100%;
    flex-shrink: 0;
    background-color: #1e1e1e;
  }

  .table-sidebar-resizer:hover,
  .table-sidebar-resizer.active {
    background-color: #9d9d9d;
  }
</style>
