<script>
  import ListConnection from "./ListConnection.svelte";
  import InputConnection from "./InputConnection.svelte";
</script>

<div class="connection">
  <div>
    <h2 class="connection-header">ALL CONNECTIONS</h2>
    <ListConnection />
  </div>
  <InputConnection />
</div>

<style>
  .connection {
    display: flex;
    padding: 1rem;
  }

  .connection-header {
    font-size: 1.15rem;
    font-weight: 600;
  }
</style>
