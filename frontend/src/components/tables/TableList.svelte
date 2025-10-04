<script>
  import { onMount } from "svelte";
  import { GetTables } from "../../../wailsjs/go/usecase/Usecase";

  let isShowTableList = $state(false);
  let tables = $state([]);

  onMount(async () => {
    try {
      tables = await GetTables();
      isShowTableList = true;
    } catch (err) {
      console.log("Failed to get table list:", err);
    }
  });
</script>

<h1>Table List</h1>
{#if isShowTableList}
  {#if tables}
    <div class="ul">
      {#each tables as table}
        <li>{table}</li>
      {/each}
    </div>
  {/if}
{/if}
