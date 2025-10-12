<script>
  import { handleEnter } from "../../utils/key";
  import {
    CirclePlus,
    CircleMinus,
    Check,
    ChevronsUpDown,
  } from "@lucide/svelte";

  let {
    isShowTableFilter,
    filters,
    updateQueryTableContents,
    queryTableContents,
    tableColumns,
  } = $props();

  const operators = [
    "=",
    "!=",
    "<",
    ">",
    "<=",
    ">=",
    "IN",
    "NOT IN",
    "LIKE",
    "BETWEEN",
    "IS NULL",
    "IS NOT NULL",
  ];
  const oneColumnOperators = [
    "=",
    "!=",
    "<",
    ">",
    "<=",
    ">=",
    "IN",
    "NOT IN",
    "LIKE",
  ];
  const twoColumnOperators = ["BETWEEN"];

  function updateFilterActiveStatus(index, checked) {
    filters = filters.map((f, i) =>
      i === index ? { ...f, isActive: checked } : f,
    );
  }

  // resizeSelectInput width based on choosed value
  function resizeSelectInput(e) {
    const select = e.target;
    const mirror = select.nextElementSibling; // Assumes mirror <span> right after select
    const selected = select.options[select.selectedIndex];

    mirror.textContent = selected.text;
    select.style.width = mirror.offsetWidth + 32 + "px"; // Add padding/arrow width
  }

  function addFilterItem() {
    filters.push({
      isActive: true,
      operator: "=",
      firstValue: "",
      secondValue: "",
    });
  }

  function closeFilterItem(idx) {
    filters.splice(idx, 1);
  }

  function filterTableRecords() {
    let conditions = [];
    filters.forEach((filter) => {
      if (filter.isActive && filter.field && filter.operator) {
        conditions.push({
          field: filter.field,
          operator: filter.operator,
          first_value: filter.firstValue,
          second_value: filter.secondValue,
        });
      }
    });
    updateQueryTableContents({
      table_name: queryTableContents.table_name,
      sort_by: queryTableContents.sort_by,
      order_by: queryTableContents.order_by,
      offset: queryTableContents.offset,
      conditions: conditions,
    });
  }
</script>

<div class="table-filter {isShowTableFilter ? 'active' : ''}">
  <div class="table-filter-list">
    {#each filters as filter, i}
      <div class="table-filter-item">
        <div class="table-filter-item-check">
          <input
            type="checkbox"
            name="active"
            onchange={(e) => updateFilterActiveStatus(i, e.target.checked)}
            class="table-filter-check-input"
          />
          <div
            class="table-filter-check-icon {filter.isActive ? 'active' : ''}"
          >
            <Check size="14" strokeWidth="3" />
          </div>
        </div>
        <div class="table-filter-item-inputs">
          <div class="table-filter-field">
            <select
              name="field"
              class="table-filter-field-select"
              bind:value={filter.field}
              onchange={(e) => resizeSelectInput(e)}
            >
              {#each tableColumns as tableColumn}
                <option value={tableColumn.column_name}>
                  {tableColumn.column_name}
                </option>
              {/each}
            </select>
            <!-- To mirror choosed option width -->
            <span
              style="position:absolute;visibility:hidden;white-space:pre;font:inherit;"
            ></span>
            <div class="table-filter-field-icon">
              <ChevronsUpDown size="15" strokeWidth="3" />
            </div>
          </div>
          <div class="table-filter-operator">
            <select
              name="operator"
              class="table-filter-operator-select"
              bind:value={filter.operator}
              onchange={(e) => resizeSelectInput(e)}
            >
              {#each operators as operator}
                <option value={operator}>{operator}</option>
              {/each}
            </select>
            <!-- To mirror choosed option width -->
            <span
              style="position:absolute;visibility:hidden;white-space:pre;font:inherit;"
            ></span>
            <div class="table-filter-operator-icon">
              <ChevronsUpDown size="15" strokeWidth="3" />
            </div>
          </div>
          {#if oneColumnOperators.includes(filter.operator) || twoColumnOperators.includes(filter.operator)}
            <input
              type="text"
              name="value"
              class="table-filter-value"
              bind:value={filter.firstValue}
            />
          {/if}
          {#if twoColumnOperators.includes(filter.operator)}
            <input
              type="text"
              name="value"
              class="table-filter-value"
              bind:value={filter.secondValue}
            />
          {/if}
        </div>
        <div
          class="table-filter-item-close"
          role="button"
          tabindex="0"
          onclick={(e) => {
            e.preventDefault();
            closeFilterItem(i);
          }}
          onkeydown={(e) => {
            handleEnter(e, () => {
              e.preventDefault();
              closeFilterItem(i);
            });
          }}
        >
          <CircleMinus size="15" />
        </div>
      </div>
    {/each}
  </div>
  <div class="table-filter-buttons">
    <div
      class="table-filter-button-add"
      role="button"
      tabindex="0"
      onclick={(e) => {
        e.preventDefault();
        addFilterItem();
      }}
      onkeydown={(e) => {
        handleEnter(e, () => {
          e.preventDefault();
          addFilterItem();
        });
      }}
    >
      <CirclePlus size="15" />
    </div>
    <div
      class="table-filter-button-submit"
      role="button"
      tabindex="0"
      onclick={(e) => {
        e.preventDefault();
        filterTableRecords();
      }}
      onkeydown={(e) => {
        handleEnter(e, () => {
          e.preventDefault();
          filterTableRecords();
        });
      }}
    >
      Filter
    </div>
  </div>
</div>

<style>
  /* Table Filter */
  .table-filter {
    display: none;
    margin-bottom: 0.5rem;
  }

  .table-filter.active {
    display: block;
  }

  .table-filter-item {
    display: flex;
    padding: 1rem 1rem 0 1rem;
    align-items: center;
    gap: 1rem;
  }

  .table-filter-item-inputs {
    display: flex;
    gap: 0.75rem;
    align-items: center;
    width: 100%;
  }

  .table-filter-item-check {
    position: relative;
  }

  .table-filter-check-input {
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    width: 1.2rem;
    height: 1.2rem;
    margin: 0;
    border: 1px solid var(--color-dark-grey-2);
    border-radius: 0.4rem;
    cursor: pointer;
    transition: 0.2s ease;
  }

  .table-filter-check-icon {
    display: flex;
    opacity: 0;
    position: absolute;
    top: 0;
    left: 0;
    width: 1.2rem;
    height: 1.2rem;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    pointer-events: none;
    transition: 0.1s ease;
  }

  .table-filter-check-icon.active {
    opacity: 1;
    border-radius: 0.4rem;
    color: #fff;
    border: 1px solid var(--color-carolina-blue-1);
    background-color: var(--color-carolina-blue-1);
  }

  .table-filter-field,
  .table-filter-operator {
    position: relative;
  }

  .table-filter-field-select,
  .table-filter-operator-select {
    appearance: none;
    -webkit-appearance: none;
    -moz-appearance: none;
    background-color: var(--color-almost-black);
    color: var(--color-text);
    border: 1px solid var(--color-dark-grey-2);
    padding: 0.3rem 1.5rem 0.4rem 0.45rem;
    border-radius: 0.4rem;
  }

  .table-filter-field-icon,
  .table-filter-operator-icon {
    position: absolute;
    top: 0;
    right: 0;
    padding: 0.4rem 0.3rem;
    pointer-events: none;
    background-color: transparent;
  }

  .table-filter-value {
    background-color: var(--color-almost-black);
    color: var(--color-text);
    border: 1px solid var(--color-dark-grey-2);
    border-radius: 0.4rem;
    transition: 0.2s ease;
    padding: 0.3rem 0.4rem;
    outline: none;
    flex: 1;
  }

  .table-filter-value:focus {
    border: 1px solid var(--color-carolina-blue);
  }

  .table-filter-item-close,
  .table-filter-button-add {
    display: flex;
    align-items: center;
    padding: 0.3rem;
    border-radius: 0.3rem;
    background-color: var(--color-dark-grey-2);
    border: 1px solid var(--color-dark-grey-2);
    cursor: pointer;
    transition: 0.2s ease;
  }

  .table-filter-item-close:hover,
  .table-filter-button-add:hover {
    background-color: var(--color-dark-grey);
  }

  .table-filter-buttons {
    display: flex;
    justify-content: flex-end;
    padding: 0 1rem;
    margin-top: 1rem;
    gap: 0.75rem;
    align-items: center;
  }

  .table-filter-button-submit {
    padding: 0.4rem 0.6rem;
    cursor: pointer;
    color: #fff;
    background-color: var(--color-carolina-blue-1);
    border-radius: 0.3rem;
    transition: 0.2s ease;
  }

  .table-filter-button-submit:hover {
    background-color: var(--color-carolina-blue-2);
  }
</style>
