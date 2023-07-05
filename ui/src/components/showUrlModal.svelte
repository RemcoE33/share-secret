<script>

  /** @type {string} */
  export let message
  export let reset
  export let sl

  function copyToClipboard() {
    navigator.clipboard.writeText(message);
    reset()
  }
</script>
<!-- Open the modal using ID.showModal() method -->
<dialog id="dialog" class="modal">
  <form method="dialog" class="modal-box">
    <h2 class="font-bold text-lg py-4">{sl.modaltitle}</h2>

    <textarea class="textarea textarea-primary mb-10 w-full" rows=4 bind:value={message} disabled></textarea>

    <div class="modal-action">
      <!-- if there is a button in form, it will close the modal -->
      <button on:click={() => copyToClipboard()} class="btn">
        {sl.modalbutton}
     </button>

    </div>
  </form>
</dialog>
