<script>
  import CreateConnection from "./components/connections/CreateConnection.svelte";
  import TableList from "./components/tables/TableList.svelte";
  import { dbCredential } from "./states/connection.svelte.js";

  let openedPage = $state("CONNECTION");
</script>

<main>
  <div>
    Active Connection: {dbCredential.connection_name}
    <div
      onclick={() => {
        openedPage = "CONNECTION";
      }}
    >
      Open Connection Tab
    </div>
    <div
      onclick={() => {
        openedPage = "TABLE";
      }}
    >
      Open Table List Tab
    </div>
  </div>

  <div>
    {#if openedPage === "CONNECTION"}
      <CreateConnection />
    {/if}
    {#if openedPage === "TABLE"}
      <TableList />
    {/if}
  </div>
</main>

<style>
</style>
