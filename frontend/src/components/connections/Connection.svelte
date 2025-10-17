<script>
  import ListConnection from "./ListConnection.svelte";
  import InputConnection from "./InputConnection.svelte";
</script>

<div class="connection">
  <ListConnection />
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
