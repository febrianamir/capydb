<script>
  import { handleEnter } from '../../utils/key'
  import { closeTableTab, switchActiveTab, getTableTabs } from '../../states/tables.svelte'
  import { X } from '@lucide/svelte'

  let tableTabs = getTableTabs()
</script>

<div class="table-tabs">
  {#each tableTabs.tabs as tableTab}
    <div
      class="table-tab-item {tableTabs.openedTab === tableTab.tableName ? 'active' : ''}"
      role="button"
      tabindex="0"
      onclick={(e) => {
        e.preventDefault()
        switchActiveTab(tableTab.tableName)
      }}
      onkeydown={(e) => {
        handleEnter(e, () => {
          e.preventDefault()
          switchActiveTab(tableTab.tableName)
        })
      }}
    >
      <div class="table-tab-item-text">{tableTab.tableName}</div>
      <div
        class="table-tab-item-close"
        role="button"
        tabindex="0"
        onclick={(e) => {
          e.preventDefault()
          e.stopPropagation()
          closeTableTab(tableTab.tableName)
        }}
        onkeydown={(e) => {
          handleEnter(e, () => {
            e.preventDefault()
            e.stopPropagation()
            closeTableTab(tableTab.tableName)
          })
        }}
      >
        <X size="14" />
      </div>
      <div class="table-tab-item-active-indicator"></div>
    </div>
  {/each}
</div>

<style>
  .table-tabs {
    display: flex;
    margin-bottom: 0.5rem;
  }

  .table-tab-item {
    position: relative;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.4rem 0.7rem 0.6rem 0.8rem;
    background-color: var(--color-dark-grey-2);
  }

  .table-tab-item.active {
    background-color: var(--color-dark-grey);
  }

  .table-tab-item-close {
    display: flex;
    align-items: center;
    cursor: pointer;
  }

  .table-tab-item-active-indicator {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 2px;
    background-color: red;
    background-color: var(--color-dark-grey-2);
  }

  .table-tab-item.active .table-tab-item-active-indicator {
    background-color: var(--color-carolina-blue);
  }
</style>
