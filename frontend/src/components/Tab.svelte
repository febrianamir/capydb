<script>
  import { EllipsisVertical } from '@lucide/svelte'
  import { setOpenedTab } from '../states/tab.svelte'
  import { connection } from '../states/connection.svelte'
  import { handleEnter } from '../utils/key'
</script>

<div class="tab">
  <div
    class="tab-connection"
    role="button"
    tabindex="0"
    onclick={(e) => {
      e.preventDefault()
      setOpenedTab('CONNECTION')
    }}
    onkeydown={(e) => {
      handleEnter(e, () => {
        e.preventDefault()
        setOpenedTab('TABLE')
      })
    }}
  >
    <EllipsisVertical />
  </div>
  <div
    class="tab-explore"
    role="button"
    tabindex="0"
    onclick={(e) => {
      e.preventDefault()
      setOpenedTab('TABLE')
    }}
    onkeydown={(e) => {
      handleEnter(e, () => {
        e.preventDefault()
        setOpenedTab('TABLE')
      })
    }}
  >
    {#if connection.current_connection.connection_id !== ''}
      {connection.current_connection.connection_name}
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
