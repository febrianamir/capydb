let openedTab = $state("CONNECTION");

export function getOpenedTab() {
  return openedTab
}

export function setOpenedTab(value) {
  openedTab = value;
}
