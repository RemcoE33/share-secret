<script>
  import FlagNL from '$lib/nl.svelte'
  import FlagUS from '$lib/us.svelte'
  import ShowResult from './showresult.svelte'
  import ShowPassword from './showpassword.svelte'
  import { getSecret } from './fetch.js'
  import { lang } from '$lib/lang' 
  export let data
  let sl = lang.en
  let showSecret = false
  /** @type {import('./fetch.js').Secret} */
  let secretObj

  /** @param {string} l */
  function changeLang(l) {
    sl = lang[l]
  }

  /** @param {string} password */
  const handleGetSecret = async (password) => {
    secretObj = await getSecret(password, data.slug)
    showSecret = true
  }

</script>

<div class="hero min-h-screen bg-base-200">
  <div class="hero-content text-center items-center">
    <div class="max-w-xl">
      <!-- lang flag switch -->
      <label class="swap mb-10">
        <input type="checkbox" />
        <div class="swap-off" on:click={() => changeLang("en")}><FlagNL />NL</div>
        <div class="swap-on" on:click={() => changeLang("nl")}><FlagUS />EN</div>
      </label>


      {#if !showSecret}
        <ShowPassword {sl} {handleGetSecret}/>
      {:else}
        <ShowResult data={secretObj} {sl}/>
      {/if}

    </div>
  </div>
</div>
