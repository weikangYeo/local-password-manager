# Local Password Manager (CLI)

## Goals

Build a password manager that can run on local (Mac & Windows), that manage user password. Build from scratch is because dont want to play any $$$ and not trusting the third party source (data store at server side). Didnt use any DB because this project is too small to have a DB.

### Main PLAN

- a. Use key to create signature of secret word and store signature, this could use as a soft verfication to remind user that he/she might entered wrong secret key, do they still want to proceed.
  - context: intend to encrypt the password kv as secret file. This file can output and share to other device, hence the password is saved, later can be re-imported to other device via secret key to reopen it.
- b. Since it is only local, shall i just use asymc key to encrypt/decrypt file?
  - implication is that this password kv cannot be "share" to other device.
  - it can still "share", but only the device has the same private key can read it.

### TODO

- [ ] CRUD for key and pwd value
- [ ] Encrypt and decrypt pwd kv file with secret key.
- [ ] Import/Output key file
- [ ] Verify Keyword Signature
- [ ] Profile (future epic)
  - [ ] Each profile has its own KV