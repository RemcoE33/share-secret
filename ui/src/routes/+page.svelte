<script>
  import ErrorMsg from './../components/errorMsg.svelte'
  import ShowUrlModal from './../components/showUrlModal.svelte' 
  import FlagNL from '$lib/nl.svelte'
  import FlagUS from '$lib/us.svelte'
  import { createSecret } from "./fetch";
  import { page } from '$app/stores';
  /** @type {import('$lib/lang.js').Languages} */
  import { lang } from '$lib/lang.js'
  /** @type {import('$lib/lang.js').Language} */
  let sl = lang.nl
  let secret = ''
  let days = 7
  let message = ''
  let responsestatus = 200
  
  /** @param {string} l */
  function changeLang(l) {
    sl = lang[l]
  }

  async function handleCreateSecret(){
    const { status, id, password } = await createSecret(secret, days)
    responsestatus = status
    const url = $page.url.origin + "/secret/" + id
    message = `URL:\t${url}\nPASSWORD:\t${password}`
    const modal = document.querySelector("dialog")
    modal?.showModal()
  }

  const resetSecret = () => {
    secret = ''
  }

  $: disabled = secret ? false : true
</script>

<div class="hero min-h-screen bg-base-200">
  <div class="hero-content text-center items-center">
    <div class="max-w-md">

      <label class="swap">
              <input type="checkbox" />
        <div class="swap-off" on:click={() => changeLang("en")}><FlagNL />NL</div>
        <div class="swap-on" on:click={() => changeLang("nl")}><FlagUS />EN</div>
      </label>

    <h1 class="text-5xl font-bold mb-10">{sl.title}</h1>
    <ul class="steps mb-10">
      <li class="step step-primary">{sl.steps['A']}</li>
      <li class="step step-primary">{sl.steps['B']}</li>
      <li class="step step-primary">{sl.steps['C']}</li>
        <li class="step step-primary">{sl.steps['D']}</li>
        <li class="step step-secondary">{sl.steps['E'] + days + " " + (days > 1 ? sl.days.plural : sl.days.single)}</li>
      </ul>

      <textarea class="textarea textarea-primary mb-10 w-full" rows=4 maxlength=1500 required bind:value={secret} placeholder={sl.placeholder}></textarea>

      <p class="mb-10">{sl.expires}</p>

      <div class="join mb-10">
        <input class="join-item btn" type="radio" name="days" bind:group={days} value={1} aria-label="1 {sl.days.single}" />
        <input class="join-item btn" type="radio" name="days" bind:group={days} value={3} aria-label="3 {sl.days.plural}" />
        <input class="join-item btn" type="radio" name="days" bind:group={days} value={5} aria-label="5 {sl.days.plural}" />
        <input class="join-item btn" type="radio" name="days" bind:group={days} value={7} aria-label="7 {sl.days.plural}" />
      </div>

      <button disabled={disabled} on:click={() => handleCreateSecret()} class="btn btn-primary mb-10">{sl.button}</button>
      
      {#if responsestatus >= 400}
        <ErrorMsg message={sl.error}/>
      {/if}

      <ShowUrlModal {message} reset={resetSecret} {sl}/>

    </div>
  </div>
</div>
