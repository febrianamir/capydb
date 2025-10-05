<script>
  import Connection from "./components/connections/Connection.svelte";
  import InputConnection from "./components/connections/InputConnection.svelte";
  import TableList from "./components/tables/TableList.svelte";
  import { EllipsisVertical } from "@lucide/svelte";
  import { dbCredential } from "./states/connection.svelte.js";

  let openedPage = $state("CONNECTION");
</script>

<main>
  <div>
    <div
      onclick={() => {
        openedPage = "CONNECTION";
      }}
    >
      <EllipsisVertical />
    </div>
    <div>
      Active Connection: {dbCredential.connection_name}
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
      <Connection />
      <InputConnection />
    {/if}
    {#if openedPage === "TABLE"}
      <TableList />
    {/if}
  </div>
</main>

<style>
</style>
