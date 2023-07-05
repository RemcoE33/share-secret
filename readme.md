# Share Secret

Share secret is used to share a string once. A user go's to the website to create an new secret,
when saved the user gets an url back. This can be shared and opend once. If the url is not opend within configed time the secret will be lost.

## Tech stack

- Go backend
  - google uuid
  - interfaces for the db
    - sample with firestore
- SvelteKit frontend
  - [DaisyUi](DaisyUi.com) components
  - TailwindCSS
  - CryptoJS

## How it works

With the help of CryptoJS the secret will be encrypted with AES encryption. The password to decrypt the secret will be generated client side and not stored on the server.

When the password is enterd the secret will be pulled from the server, when the delivery is succefully made to the client the secret will be deleted, even if the password to decypt the secret is incorrect. This is all done client side.
