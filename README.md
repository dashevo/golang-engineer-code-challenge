<p align="center">
	<a href="https://git.io/col">
		<img src="https://img.shields.io/badge/%E2%9C%93-collaborative_etiquette-brightgreen.svg" alt="Collaborative Etiquette">
	</a>
	<a href="https://twitter.com/intent/follow?screen_name=dashpay">
		<img src="https://img.shields.io/twitter/follow/dashpay.svg?style=social&logo=twitter" alt="follow on Twitter">
	</a>
	<a href="#">
		<img src="https://travis-ci.com/dashevo/go-engineer-code-challenge.svg?branch=main" alt="travis-ci">
	</a>
	<a href="#">
		<img src="https://img.shields.io/dub/l/vibe-d.svg" alt="MIT">
	</a>
</p>

<p>&nbsp;</p>

<p align="center">
	<a href="https://dash.org">
		<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/f0/Dash_digital-cash_logo_2018_rgb_for_screens.svg/1600px-Dash_digital-cash_logo_2018_rgb_for_screens.svg.png" width="600">
	</a>
</p>

<p>&nbsp;</p>

# Go Engineer Code Challenge

> Code challenge for Dash Core Team candidates

The goal of this challenge is to create an application which receives,
validates and displays data from a user.

This repository is designed as a mono-repository for your application and external services. We've prepared an application skeleton for you and assume that your solution will be a light client. In the provided skeleton, please implement a [service](internal/app/usecase/service.go) which validates and persists [sample data](assets/data.json) in external service(s). It should also retrieve the data and and verify its integrity.

Your solution must implement at least one of the following types of external services:

  1. **Peer-to-peer service**: This service runs on the user's host.
     Networking and storage will be cheap for you - **0.0001 DASH per byte**, but you can’t trust
     this service because a malicious user may spoof (modify) your data.
     Code for this service should be placed [here](internal/p2p) and the entrypoint of the service is [here](cmd/p2p/main.go)
  2. **Self-hosted service**: This service runs on your server.
     Networking and storage will be much more expensive for you - **0.001 DASH per byte**, but the data is
     located on your server, so you can trust it.
     Code for this service should be placed [here](internal/selfhosted) and the entrypoint of the service is [here](cmd/selfhosted/main.go).

For communication with your external service(s) we provide a [service](internal/app/usecase/service.go)
which are used http client for communication between services.
The middleware helps us calculate inbound and outbound traffic.

## Your mission

- Use this repository as a template for the solution.
- Implement as many external services as you need to store sample data from the application.
- Implement [the store method](internal/app/usecase/service.go#L27). Validate and persist sample data
   in the external service(s).
- Implement [the fetch method](internal/app/usecase/service.go#L23). Fetch sample data back and ensure its
   integrity. **When you fetch data back from the untrusted service, you should verify it for spoofing protection**
- Write beautiful code. Code design (SOLID, Clean Architecture, 12factor) is important to us.
- Run the application and review the results. **Try to spend as little money as possible**. Cost depends on the size
   of the request / response and elapsed time. The exact formula is found in
   [the application skeleton code](internal/app/metric/calculator.go).
- Share a link to your private solution repository with us or send us an archive containing your solution.

## Requirements

### External services

- Services should be written in Go
- Data should be permanently persisted (i.e. available after a service restart)

### Application

- You should validate sample data in the store method. Return an error if any data is invalid.
- You should check data integrity in the fetch method to avoid spoofing. Keep in mind that the fetch method will not have access to the original input data provided to the store method.
- Make sure the data returned by the fetch method matches the input data from the store action.
- In order to specify which type of network you are going to use, follow the instruction in the [.env.dist](.env.dist) file
- You cannot store any data on the application side

### Sample data validation rules

The provided [sample data](assets/data.json) contains a collection of various objects.
Each type of object has its own validation rules as defined below:

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
   Not equal or less than `0`
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
