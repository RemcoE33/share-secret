/**
 * @typedef {Object} Languages
 * @property {Language} nl
 * @property {Language} en
*/

/**
 * @typedef {Object} Language
 * @property {string} error general error message
 * @property {string} title 
 * @property {string} passwordtitle
 * @property {string} secrettitle 
 * @property {Days} days
 * @property {Steps} steps 
 * @property {string} button
 * @property {string} placeholder
 * @property {string} expires
 * @property {string} modaltitle
 * @property {string} modalbutton
 * @property {string} errormsg message if secret is not found
 * @property {ErrorTitle} errortitle
 * @property {string} errorbutton
 * @property {string} passwordbutton
 * @property {string} password
 *
*/

/**
 * @typedef {Object} Days
 * @property {string} plural
 * @property {string} single
*/

/**
 * @typedef {Object} Steps
  * @property {string} A
  * @property {string} B
  * @property {string} C
  * @property {string} D
  * @property {string} E
*/

/**
 * @typedef {Object} ErrorTitle
 * @property {string} ok
 * @property {string} nok
*/

/** @type {Languages} */
export const lang = {
  nl: {
    error: 'Oeps.. er is iets niet goed gegaan',
    title: 'Maak geheim',
    passwordtitle: 'Bijna in je bezit',
    secrettitle: 'Jouw geheim',
    days:{ 
      plural: 'dagen',
      single: 'dag'
    },
    steps: {
      'A': 'Geheim invoeren',
      'B': 'Krijg uniek url',
      'C': 'Deel',
      'D': 'Open eenmalig',
      'E': 'Verwijderd in '
    },
    button: 'Aanmaken',
    placeholder: 'Voeg geheim in van maximaal 1500 tekens...',
    expires: 'Vervalt in:',
    modaltitle: 'Kopieer en deel url:',
    modalbutton: 'Kopieer',
    errormsg: 'Geheim is al geopend of verlopen',
    errortitle: {
      ok: 'Jouw geheim:',
      nok: 'Oeps....'
    },
    errorbutton: 'Maak geheim aan',
    passwordbutton: 'Indienen',
    password: "Wachtwoord"
  },
  en: {
    error: 'Oeps.. something went wrong',
    title: 'Create secret',
    passwordtitle: "Almost in your possession",
    secrettitle: "You're secret",
    days:{ 
      plural: 'days',
      single: 'day'
    },
    steps: {
      'A': 'Enter secret',
      'B': 'Get unique url',
      'C': 'Share',
      'D': 'Open once',
      'E': 'Deleted in '
    },
    button: 'Create secret',
    placeholder: 'Enter secret of maximum 1500 characters...',
    expires: 'Expires in:',
    modaltitle: 'Copy and share url:',
    modalbutton: 'Copy',
    errormsg: 'This secret is already seen or expired',
    errortitle: {
      ok: "You're secret:",
      nok: 'Oeps....'
    },
    errorbutton: 'Create secret',
    passwordbutton: 'Submit',
    password: 'password'
  }
}
