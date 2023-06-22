import { decrypt } from '$lib/aes'

/**
 * @typedef {Object} Secret 
 * @property {number} status
 * @property {string} secret
* /

/**
 * @param {string} id
 * @param {string} password 
 * @returns {Promise<Secret>}
*/
// export async function getSecret(password, id) {
//   password = password.trim()
//   const res = await fetch(`/api/secret/${id}`)
//   const data = await res.json()
//   let status = res.status
//   let secret = data?.secret
//
//   if (status < 400){
//     try {
//      secret = decrypt(secret, password)
//
//     } catch(e){
//      status = 500
//     }
//   }
//
//   return {
//     status,
//     secret
//   }
// }

export async function getSecret(password, id){
  return {
    status: 200,
    secret: "Some real love!"
  }
}

