<script>
  import { EllipsisVertical } from "@lucide/svelte";
  import { setOpenedTab } from "../states/tab.svelte";
  import { dbCredential } from "../states/connection.svelte";
</script>

<div class="tab">
  <div
    class="tab-connection"
    onclick={() => {
      setOpenedTab("CONNECTION");
    }}
  >
    <EllipsisVertical />
  </div>
  <div
    class="tab-explore"
    onclick={() => {
      setOpenedTab("TABLE");
    }}
  >
    {#if dbCredential.has_active_connection}
      {dbCredential.connection_name}
    {:else}
      No Active Connection
    {/if}
  </div>
</div>

<style>
  .tab {
    display: flex;
    align-items: center;
    padding: 0 0.25rem;
    height: 35px;
    border-bottom: 2px solid #1a1a1a;
  }

  .tab-connection {
    display: flex;
    align-items: center;
    padding: 0.25rem 0.25rem 0.25rem 0;
    min-height: 100%;
    cursor: pointer;
  }

  .tab-explore {
    display: flex;
    align-items: center;
    padding: 0.25rem 0.75rem;
    min-height: 100%;
    background-color: #333333;
    cursor: pointer;
  }
</style>
