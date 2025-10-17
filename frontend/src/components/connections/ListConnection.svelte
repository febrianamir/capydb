<script>
  import { onMount } from "svelte";
  import { GetCredentials } from "../../../wailsjs/go/usecase/Usecase";
  import { Server, User, Database } from "@lucide/svelte";
  import { dbCredential } from "../../states/connection.svelte";
  import { handleEnter } from "../../utils/key";

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

  function fillCredential(credential) {
    dbCredential.has_active_connection = false;
    dbCredential.connection_name = credential.title;
    dbCredential.db_vendor = credential.db_vendor;
    dbCredential.host = credential.host;
    dbCredential.port = credential.port;
    dbCredential.user = credential.user;
    dbCredential.password = credential.password;
    dbCredential.database_name = credential.database_name;
  }
</script>

<div class="connection-list-container">
  <h2 class="connection-header">ALL CONNECTIONS</h2>
  <div class="connection-list">
    {#if credentials.length > 0}
      {#each credentials as credential}
        <div
          class="connection-item"
          role="button"
          tabindex="0"
          onclick={(e) => {
            e.preventDefault();
            fillCredential(credential);
          }}
          onkeydown={(e) => {
            handleEnter(e, () => {
              e.preventDefault();
              fillCredential(credential);
            });
          }}
        >
          <div class="connection-db-vendor">
            {credential.db_vendor}
          </div>
          <div class="connection-title">
            {credential.title}
          </div>
          <div class="connection-server">
            <div class="connection-icon">
              <Server size="16" />
            </div>
            <div class="conection-text">
              {credential.host}:{credential.port}
            </div>
          </div>
          <div class="connection-user">
            <div class="connection-icon">
              <User size="16" />
            </div>
            <div class="conection-text">
              {credential.user}
            </div>
          </div>
          <div class="connection-db-name">
            <div class="connection-icon">
              <Database size="16" />
            </div>
            <div class="conection-text">
              {credential.database_name}
            </div>
          </div>
        </div>
      {/each}
    {/if}
  </div>
</div>

<style>
  .connection-list-container {
    flex: 1;
    max-width: 700px;
  }

  .connection-header {
    font-size: 1.15rem;
    font-weight: 600;
  }

  .connection-list {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    gap: 1rem;
    padding: 1rem 0;
  }

  .connection-item {
    padding: 1rem;
    min-width: 200px;
    border-radius: 1rem;
    transition: 0.2s ease;
  }

  .connection-item:hover {
    background-color: #272727;
  }

  .connection-db-vendor {
    padding: 0.25rem 0.5rem;
    border-radius: 0.5rem;
    display: inline-block;
    background-color: #505050;
    margin-bottom: 0.5rem;
  }

  .connection-title {
    font-weight: bold;
    margin-bottom: 0.75rem;
  }

  .connection-server,
  .connection-user,
  .connection-db-name {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    margin: 0.2rem 0;
  }

  .connection-icon {
    display: flex;
    align-items: center;
  }
</style>
