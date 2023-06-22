import { encrypt } from '$lib/aes'
/**
 * @typedef {Object} ID
 * @property {string} id the secret id
 * @property {number} status http status code
 * @property {string} password the password to open the secret
* /

/**
 * @param {string} secret 
 * @param {number} days 
 * @returns {Promise<ID>}
 * */
export async function createSecret(secret, days) {
  const { key, ciphertext } = encrypt(secret)
  const payload = {
    ciphertext,
    days
  };
  const resp = await fetch("/api/secret", {
    method: "post",
    body: JSON.stringify(payload),
  });

  const data = await resp.json()
  const status = resp.status
  return { status, id: data?.id, password: key };
}


