<script>
  import ListConnection from './ListConnection.svelte'
  import InputConnection from './InputConnection.svelte'
  import { onMount } from 'svelte'
  import { GetCredentials } from '../../../wailsjs/go/usecase/Usecase'

  let credentials = $state([])

  onMount(async () => {
    try {
      let req = {}
      let credentialsRes = await GetCredentials(req)
      credentials = credentialsRes.Data
    } catch (err) {
      console.log('Failed to get list credentials:', err)
    }
  })
</script>

<div class="connection">
  <ListConnection credentials={credentials} />
  <InputConnection credentials={credentials} />
</div>

<style>
  .connection {
    display: flex;
    padding: 1rem;
  }
</style>
