<script>
  import { generateDBTime } from "../../utils/db.js";
  import {
    CreateConnection,
    SaveCredential,
  } from "../../../wailsjs/go/usecase/Usecase.js";
  import { connection } from "../../states/connection.svelte.js";
  import { Cable, Save } from "@lucide/svelte";

  let { credentials } = $props();

  async function createConnection(e) {
    e.preventDefault();

    try {
      let res = await CreateConnection(connection.credential);

      // Append new active connections
      connection.active_connections.push({
        connection_id: res.ConnectionId,
        connection_name: connection.credential.connection_name,
      });

      // Set current connection
      connection.current_connection.connection_id = res.ConnectionId;
      connection.current_connection.connection_name =
        connection.credential.connection_name;
    } catch (err) {
      console.log("Failed to create connection:", err);
    }
  }

  async function saveConnection(e) {
    e.preventDefault();

    try {
      let timeNow = generateDBTime();
      let req = {
        title: connection.credential.connection_name,
        db_vendor: connection.credential.db_vendor,
        host: connection.credential.host,
        port: connection.credential.port,
        user: connection.credential.user,
        password: connection.credential.password,
        database_name: connection.credential.database_name,
        created_at: timeNow,
        updated_at: timeNow,
      };
      if (connection.credential.credential_id > 0) {
        // Add credential id to update
        req.id = connection.credential.credential_id;
      }

      await SaveCredential(req);

      if (connection.credential.credential_id > 0) {
        // Update credential list
        credentials[connection.credential.credential_idx] =
          connection.credential;
        credentials[connection.credential.credential_idx].id =
          connection.credential.credential_id;
        credentials[connection.credential.credential_idx].title =
          connection.credential.connection_name;
        credentials[connection.credential.credential_idx].db_vendor =
          connection.credential.db_vendor;
        credentials[connection.credential.credential_idx].host =
          connection.credential.host;
        credentials[connection.credential.credential_idx].port =
          connection.credential.port;
        credentials[connection.credential.credential_idx].user =
          connection.credential.user;
        credentials[connection.credential.credential_idx].password =
          connection.credential.password;
        credentials[connection.credential.credential_idx].database_name =
          connection.credential.database_name;
      }
    } catch (err) {
      console.log("Failed to save connection:", err);
    }
  }
</script>

<div class="connection-input-container">
  <h3 class="connection-input-header">
    {#if connection.credential.credential_id > 0}
      CONNECT TO DB
    {:else}
      ADD NEW CONNECTION
    {/if}
  </h3>
  <form onsubmit={createConnection} class="connection-form">
    <div class="connection-form-item">
      <label for="connection_name" class="connection-form-label">
        Connection Name
      </label>
      <input
        class="connection-form-input"
        type="text"
        name="connection_name"
        id="connection_name"
        bind:value={connection.credential.connection_name}
      />
    </div>
    <div class="connection-form-item">
      <label for="db_vendor" class="connection-form-label">Client</label>
      <select
        class="connection-form-input"
        name="db_vendor"
        id="db_vendor"
        bind:value={connection.credential.db_vendor}
      >
        <option value="PostgreSQL">PostgreSQL</option>
      </select>
    </div>
    <div class="connection-form-item">
      <label for="host" class="connection-form-label">Host</label>
      <input
        class="connection-form-input"
        type="text"
        name="host"
        id="host"
        bind:value={connection.credential.host}
      />
    </div>
    <div class="connection-form-item">
      <label for="port" class="connection-form-label">Port</label>
      <input
        class="connection-form-input"
        type="text"
        name="port"
        id="port"
        bind:value={connection.credential.port}
      />
    </div>
    <div class="connection-form-item">
      <label for="user" class="connection-form-label">User</label>
      <input
        class="connection-form-input"
        type="text"
        name="user"
        id="user"
        bind:value={connection.credential.user}
      />
    </div>
    <div class="connection-form-item">
      <label for="password" class="connection-form-label">Password</label>
      <input
        class="connection-form-input"
        type="password"
        name="password"
        id="password"
        bind:value={connection.credential.password}
      />
    </div>
    <div class="connection-form-item">
      <label for="database_name" class="connection-form-label">
        Database Name
      </label>
      <input
        class="connection-form-input"
        type="text"
        name="database_name"
        id="database_name"
        bind:value={connection.credential.database_name}
      />
    </div>
    <div class="connection-form-footer">
      <button class="connection-btn-save" onclick={saveConnection}>
        <div class="connection-btn-icon">
          <Save size="14" />
        </div>
        <div class="connection-btn-text">Save</div>
      </button>
      <button type="submit" class="connection-btn-connect">
        <div class="connection-btn-icon">
          <Cable size="14" />
        </div>
        <div class="connection-btn-text">Connect</div>
      </button>
    </div>
  </form>
</div>

<style>
  .connection-input-container {
    flex: 1;
  }

  .connection-input-header {
    font-size: 1.15rem;
    font-weight: 600;
  }

  .connection-form {
    margin-top: 1rem;
    padding: 1rem;
    border-radius: 1rem;
    max-width: 520px;
    background-color: var(--color-dark-grey-3);
  }

  .connection-form-item {
    display: flex;
    align-items: center;
    padding-bottom: 0.5rem;
  }

  .connection-form-label {
    display: block;
    width: 175px;
    font-size: 1rem;
  }

  .connection-form-input {
    outline: none;
    flex: 1;
    color: var(--color-text);
    background-color: var(--color-dark-grey);
    border-radius: 0.25rem;
    border: 1px solid var(--color-dark-grey-2);
    padding: 0.4rem 0.8rem;
  }

  .connection-form-footer {
    display: flex;
    justify-content: flex-end;
    padding: 0.75rem 0 0 0;
    gap: 0.75rem;
  }

  .connection-btn-save,
  .connection-btn-connect {
    display: flex;
    align-items: center;
    padding: 0.3rem 0.6rem;
    gap: 0.4rem;
    color: var(--color-text);
    border-radius: 0.25rem;
    outline: none;
    border: none;
    cursor: pointer;
  }

  .connection-btn-save {
    background-color: var(--color-dark-peach-2);
  }

  .connection-btn-connect {
    background-color: var(--color-dark-sage);
  }

  .connection-btn-icon {
    display: flex;
    align-items: center;
  }
</style>
