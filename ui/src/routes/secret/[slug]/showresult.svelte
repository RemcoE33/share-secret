<script>
  import CheckSvg from './../../../components/checksvg.svelte'
  /** @type {import('./fetch.js').Secret} */
  export let data
  /** @type {import('$lib/lang.js').Language} */
  export let sl
  let showcheck = false

  function copyToClipboard() {
    navigator.clipboard.writeText(data.secret);
    showcheck = true
  }
</script>


<h1 class="text-5xl font-bold mb-10">{sl.secrettitle}</h1>

<!-- error message or secret-->
<div class="mb-10">
  {#if data.status < 400}
    <textarea class="textarea textarea-primary mb-10 w-full" rows=4 bind:value={data.secret}></textarea>
  {:else}
    <div class="alert alert-warning">
      <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
      <span>{sl.errormsg}</span>
    </div>
  {/if}
</div>

<!-- buttons -->
<div class="flex gap-4">
  {#if data.status < 400}
    <button class="btn btn-primary" on:click={() => copyToClipboard()}>
      {sl.modalbutton}
      {#if showcheck}
        <CheckSvg />
      {/if}
    </button>
  {/if}

  <a href="/">
    <button class="btn btn-outline btn-primary">{sl.errorbutton}</button>
  </a>
</div>

