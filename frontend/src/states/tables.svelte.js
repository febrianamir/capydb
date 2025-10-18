let tableTabs = $state({
  tabs: [],
  openedTab: '',
})

export function getTableTabs() {
  return tableTabs
}

export function isTableTabExist(tableName) {
  return tableTabs.tabs.find((t) => {
    return t.tableName === tableName
  })
}

export function addTableTab(tableName) {
  if (!isTableTabExist(tableName)) {
    tableTabs.tabs.push({
      tableName: tableName,
    })
  }
  tableTabs.openedTab = tableName
}

export function closeTableTab(tableName) {
  tableTabs.tabs = tableTabs.tabs.filter((t) => {
    return t.tableName !== tableName
  })
  // If the closed tab was active, switch to another tab or empty
  if (tableTabs.openedTab === tableName) {
    tableTabs.openedTab = tableTabs.tabs.length > 0 ? tableTabs.tabs[0].tableName : ''
  }
}

export function switchActiveTab(tableName) {
  tableTabs.openedTab = tableName
}
