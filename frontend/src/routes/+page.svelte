<script lang="ts">
  import { onMount } from "svelte";
  var devices = [{Data: 0, Id: "A123"}]
  onMount(() => {
    const ws = new WebSocket("ws://localhost:3000/ws")
  
    ws.onmessage = (msg) => {
      let data = JSON.parse(msg.data)
      devices = data
    }
  })
</script>
<h1>
  Smart Home Tracker
</h1>
{#each devices as device}
<div>
  {device.Id} : {device.Data}
</div>
{/each}
