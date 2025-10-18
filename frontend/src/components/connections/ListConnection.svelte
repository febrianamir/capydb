<script>
  import { DeleteCredential } from "../../../wailsjs/go/usecase/Usecase";
  import { Server, User, Database, Trash, Plus } from "@lucide/svelte";
  import { connection } from "../../states/connection.svelte";
  import { handleEnter } from "../../utils/key";

  let { credentials } = $props();

  function fillCredential(credential, credentialIdx) {
    connection.credential.credential_idx = credentialIdx;
    connection.credential.credential_id = credential.id;
    connection.credential.connection_name = credential.title;
    connection.credential.db_vendor = credential.db_vendor;
    connection.credential.host = credential.host;
    connection.credential.port = credential.port;
    connection.credential.user = credential.user;
    connection.credential.password = credential.password;
    connection.credential.database_name = credential.database_name;
  }

  async function deleteCredential(credential, credentialIdx) {
    try {
      let req = {
        credential_id: credential.id,
      };
      await DeleteCredential(req);
      credentials.splice(credentialIdx, 1);
    } catch (err) {
      console.log("Failed to delete credential:", err);
    }
  }

  function addCredential() {
    connection.credential.credential_idx = 0;
    connection.credential.credential_id = 0;
  }
</script>

<div class="connection-list-container">
  <h2 class="connection-header">
    <div class="connection-header-text">ALL CONNECTIONS</div>
    <div
      class="connection-header-btn-add"
      role="button"
      tabindex="0"
      onclick={(e) => {
        e.preventDefault();
        addCredential();
      }}
      onkeydown={(e) => {
        handleEnter(e, () => {
          e.preventDefault();
          addCredential();
        });
      }}
    >
      <Plus size="18" strokeWidth="3" />
    </div>
  </h2>
  <div class="connection-list">
    {#if credentials.length > 0}
      {#each credentials as credential, i (credential.id)}
        <div
          class="connection-item"
          role="button"
          tabindex="0"
          onclick={(e) => {
            e.preventDefault();
            fillCredential(credential, i);
          }}
          onkeydown={(e) => {
            handleEnter(e, () => {
              e.preventDefault();
              fillCredential(credential, i);
            });
          }}
        >
          <div
            class="connection-btn-delete"
            role="button"
            tabindex="0"
            onclick={(e) => {
              e.preventDefault();
              e.stopPropagation();
              deleteCredential(credential, i);
            }}
            onkeydown={(e) => {
              handleEnter(e, () => {
                e.preventDefault();
                e.stopPropagation();
                deleteCredential(credential, i);
              });
            }}
          >
            <Trash size="14" />
          </div>
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
    width: 700px;
  }

  .connection-header {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .connection-header-text {
    font-size: 1.15rem;
    font-weight: 600;
  }

  .connection-header-btn-add {
    display: flex;
    align-items: center;
    padding: 0.5rem;
    background-color: var(--color-dark-grey-2);
    border-radius: 1rem;
    color: var(--color-text);
    cursor: pointer;
    transition: 0.2s ease;
  }

  .connection-header-btn-add:hover {
    background-color: var(--color-dark-grey-3);
  }

  .connection-list {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    gap: 1rem;
    padding: 1rem 0;
  }

  .connection-item {
    position: relative;
    padding: 1rem;
    min-width: 200px;
    border-radius: 1rem;
    transition: 0.2s ease;
  }

  .connection-btn-delete {
    opacity: 0;
    position: absolute;
    display: flex;
    align-items: center;
    top: 1rem;
    right: 1rem;
    padding: 0.25rem;
    cursor: pointer;
    border-radius: 0.25rem;
    transition: 0.2s ease;
  }

  .connection-item:hover .connection-btn-delete {
    opacity: 1;
  }

  .connection-btn-delete:hover {
    background-color: var(--color-bg);
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
