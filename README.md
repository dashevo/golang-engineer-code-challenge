# Golang Engineer's Code Challenge

> Code challenge for Dash Core Team candidates

The goal of this challenge is to create an application which receives,
validates and displays data from a user.

Since it takes a while, we've prepared the application skeleton for you.
It's not really application, but we'll pretend that it is. ;)

In the provided skeleton, you should implement a [service](internal/usecase/service.go), which should validate, persist [sample data](assets/data.json) in external service(s), retrieve and ensure data its integrity 

Also, your solution must implement at least one of the following types of external services:
  1. **Peer-to-peer service** which hypothetically runs on user's hosts. Let’s call it "P2P" service.
     Networking and storage will be cheap for you - **0.0001 DASH per byte**, but you can’t trust
     this service because a malicious user may spoof (modify) your data.
  2. **Self-hosted service** which runs on your server. Let’s call it "hosted" service.
     Networking and storage will be much more expensive for you - **0.001 DASH per byte**, but the data is
     located on your server, so you can trust it.

For communication with your external service(s) we provide two functions
which are used http client for communication between services
The middleware helps us to calculate incoming and outgoing traffics

## Your mission

 - Use this repository as a template for the solution 
 - Implement as many external services as you need to store sample data from the application
 - Implement [the store method](internal/usecase/service.go#L27). Validate and persist sample data
   in the external service(s)
 - Implement [the fetch method](internal/usecase/service.go#L32). Fetch sample data back and ensure its 
   integrity. **When you fetch data back from the untrusted service, you should verify it for spoofing protection**
 - Write beautiful code. Code design (SOLID, Clean Architecture, 12factor) is important to us.
 - Run application and see results. **Try to spend as little money as possible**. Cost depends on the size
   of request / response and elapsed time. You may find the exact formula in
   [the application skeleton code](internal/metric/calculator.go)
 - Share with us the link on a private repository or send us an archive with your solution
 
## Requirements

### External services
 - Services should be written in Golang
 - Data should be permanently persisted (i.e. available after a service restart)

### Application
 - You should validate sample data in the store method. Return an error if a data is invalid
 - You should check data integrity in the fetch method to avoid spoofing. Keep in mind that you don't have access to the original input data which you received in the store method
 - Make sure the data returned by the fetch method matches the input data from the store action
 - In order to specify which type of network you are going to use, just following instruction at .env.dist file
 - You cannot store any data on the application side

### Sample data validation rules

[Sample data](assets/data.json) represents a collection of various objects
Each type of object has its own validation rules.

#### User

- `id`
   - Format: `a-zA-Z0-9`
   - Length: `256`
   - Required
- `type`
   - Value: `user`
   - Required
- `userName`
   - Format: `a-zA-Z0-9_.`
   - Max length: `20`
   - Required
- `firstName`
   - Max length: `100`
- `lastName`
   - Max length: `100`
- `email`
   - According to RFC

#### Payment

- `id`
   - Format: `a-zA-Z0-9`
   - Length: `256`
   - Required
- `type`
   - Value: `payment`
   - Required 
- `fromUserId`
   - Format: `a-zA-Z0-9`
   - Length: `256`
   - Required
- `toMerchantId` or `toUserId`
   - Format: `a-zA-Z0-9`
   - Length: `256`
   - Required
- `amount`
   - Format: float number
   - Not equal or less than `0`
   - Required
- `createdAt`
   - Format: Date ISO 8601
   - Required

#### Merchant

- `id`
   - Format: `a-zA-Z0-9`
   - Length: `256`
   - Required
- `type`
   - Value: `merchant`
   - Required
- `name`
   - Format: `a-zA-Z0-9`
   - Max length: `20`
   - Required

## Summary

Follow the [challenge mission](#your-mission) according to the [provided requirements](#requirements) and do your 
best. Good luck!

## License

[MIT](LICENSE) © 2021 Dash Core Team 
 