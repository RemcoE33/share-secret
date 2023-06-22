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


