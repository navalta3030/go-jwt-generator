# Summary
- This Rest Api endpoint aims to generate payload for the jwt which the api-gateway will sign.

# Routes
  - /token
    - shall provide the expiration, access_token, and refresh_token.
  - /static
    - will provide the jwk entrypoint for validation and signing of the api gateway.