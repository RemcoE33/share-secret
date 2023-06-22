import CryptoJS from "crypto-js";

/**
 * creates a 32bit key
 * @returns {string}
 */
function createKey() {
  const input =
    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  let key = "";

  for (let i = 0; i < 32; i++) {
    key += input.charAt(Math.floor(Math.random() * input.length));
  }

  return key;
}

/**
 * Encrypt secret client side
 * @param {string} secret 
*/
export function encrypt(secret) {
  const key = createKey();
  const ciphertext = CryptoJS.AES.encrypt(secret, key).toString();
  return {
    ciphertext,
    key,
  };
}


/**
 * Decript secret with the key that is issued to the user
 * @param {string} key 
 * @param {string} hash 
*/

export function decrypt(hash, key) {
  const bytes = CryptoJS.AES.decrypt(hash, key);
  const secret = bytes.toString(CryptoJS.enc.Utf8);
  return secret;
}
