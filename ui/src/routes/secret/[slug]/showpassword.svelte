<script>
  /** @type {import('$lib/lang.js').Language} */
  export let sl
  /** @type {function} */
  export let handleGetSecret
  let password = ''
  let disables = true

  $: {
    if (password.length == 32){
      disables = false
    } else {
      disables = true
    }
  }
</script>


<h1 class="text-5xl font-bold mb-10">{sl.passwordtitle}</h1>

<div class="join mb-10">
  <input class="input input-bordered join-item mb-10" bind:value={password} placeholder="Password"/>
  <button class="btn join-item btn-primary" disabled={disables}  on:click={() => handleGetSecret(password)}>{sl.passwordbutton}</button>
</div>
