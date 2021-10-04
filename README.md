# gowallet

eWallet REST API service



### Wallet types:
* anonymous - balance limit 10,000 TJS
* identified - balance limit 100,000 TJS


### API auth:
* X-UserId - user ID
* X-Digest - request body hmac-sha1


### API methods:
* /check - check existing account
* /topup - top up balance
* /stats - monthly stats. operations count and sums
* /balance - check wallet balance
