<script>
  import { onMount } from "svelte";
  import { GetCredentials } from "../../../wailsjs/go/usecase/Usecase";

  let credentials = $state([]);

  onMount(async () => {
    try {
      let req = {};
      let credentialsRes = await GetCredentials(req);
      credentials = credentialsRes.Data;
    } catch (err) {
      console.log("Failed to get list credentials:", err);
    }
  });
</script>

<div class="connection-list">
  <h2>Connection List</h2>
  {#if credentials.length > 0}
    {#each credentials as credential}
      <div class="connection-item">
        {credential.title}
        {credential.db_vendor}
        {credential.hex_color}
        {credential.host}
        {credential.port}
        {credential.user}
        {credential.password}
        {credential.database_name}
      </div>
    {/each}
  {/if}
</div>

<style>
  /* your styles go here */
</style>
