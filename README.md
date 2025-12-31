# Local Password Manager (CLI)

## Goals

Build a password manager that can run on local (Mac & Windows), that manage user password. Build from scratch is because
dont want to play any $$$ and not trusting the third party source (data store at server side). Didnt use any DB because
this project is too small to have a DB.

## How to use
### Developer
- Run `go run main.go` or build as a executable file, then `./execute-file`

### Local (if Go installed)
- run `go install` (This installs to $GOPATH/bin (usually ~/go/bin))
- run `wk-local-pwd-manager` in terminal

### TODO

- [X] Main
  - [X] CRUD for key and pwd value
  - [X] Encrypt and decrypt pwd kv file with secret key.
  - [X] Use KDF & add salt
- [ ] Optional
  - [ ] Import/Output key file
  - [ ] Verify Keyword Signature
  - [ ] Profile
      - [ ] Each profile has its own KV
  - [ ] UI version