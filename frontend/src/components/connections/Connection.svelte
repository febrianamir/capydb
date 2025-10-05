<script>
  import ListConnection from "./ListConnection.svelte";

  // your script goes here
</script>

<div class="connection">
  <h2 class="connection-header">ALL CONNECTIONS</h2>
  <ListConnection />
</div>

<style>
  .connection {
    padding: 1rem;
  }

  .connection-header {
    font-size: 1.15rem;
    font-weight: 600;
  }
</style>
