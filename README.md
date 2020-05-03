# Summary ![Production](https://github.com/navalta3030/go-jwt-generator/workflows/Production/badge.svg?branch=master) ![Staging](https://github.com/navalta3030/go-jwt-generator/workflows/Staging/badge.svg)

# Routes
  - /token (POST)
    - Requires
      - Name
      - Email
    - shall provide the expiration, access_token, and refresh_token.
    - curl --header "Content-Type: application/json" \
        --request POST \
        --data '{"Name":"Mark","Email":"xyz@gmail.com"}' \
        http://localhost:8020/token
  - /static
    - will provide the jwk entrypoint for validation and signing of the api gateway.
