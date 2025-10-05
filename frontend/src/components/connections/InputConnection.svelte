<script>
  import {
    CreateConnection,
    SaveCredential,
  } from "../../../wailsjs/go/usecase/Usecase.js";
  import { dbCredential } from "../../states/connection.svelte.js";

  async function createConnection(e) {
    e.preventDefault();

    try {
      await CreateConnection(dbCredential);
      dbCredential.has_active_connection = true;
    } catch (err) {
      console.log("Failed to create connection:", err);
    }
  }

  async function saveConnection(e) {
    e.preventDefault();

    try {
      let req = {
        title: dbCredential.connection_name,
        host: dbCredential.host,
        port: dbCredential.port,
        user: dbCredential.user,
        password: dbCredential.password,
        database_name: dbCredential.database_name,
      };
      await SaveCredential(req);
    } catch (err) {
      console.log("Failed to save connection:", err);
    }
  }
</script>

<div>
  <h3>Connect to DB</h3>
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
        bind:value={dbCredential.connection_name}
      />
    </div>
    <div class="connection-form-item">
      <label for="host" class="connection-form-label">Host</label>
      <input
        class="connection-form-input"
        type="text"
        name="host"
        id="host"
        bind:value={dbCredential.host}
      />
    </div>
    <div class="connection-form-item">
      <label for="port" class="connection-form-label">Port</label>
      <input
        class="connection-form-input"
        type="text"
        name="port"
        id="port"
        bind:value={dbCredential.port}
      />
    </div>
    <div class="connection-form-item">
      <label for="user" class="connection-form-label">User</label>
      <input
        class="connection-form-input"
        type="text"
        name="user"
        id="user"
        bind:value={dbCredential.user}
      />
    </div>
    <div class="connection-form-item">
      <label for="password" class="connection-form-label">Password</label>
      <input
        class="connection-form-input"
        type="password"
        name="password"
        id="password"
        bind:value={dbCredential.password}
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
        bind:value={dbCredential.database_name}
      />
    </div>
    <div class="connection-form-footer">
      <button type="submit" class="connection-btn-connect">Connect</button>
      <button class="connection-btn-save" onclick={saveConnection}>Save</button>
    </div>
  </form>
</div>
