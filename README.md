### gRPC Users Service

Simple users gRPC service.

### How can I test the code?

First, you would need to clone this repository. Like so: 

```bash
git@github.com:gocanto/users-grpc-service.git
```

After you have done this, you have to position yourself at the project root folder from
you CLI `cd users-grpc-service` from two different tabs or CLI windows.

Once you are in the project, you would have to type the following commands.

- From one of your opened tab/window, you need to type `go run server/server.go` for you to have 
the given gRPC server running. After doing so, you would be able to see the following output
`Server listing at [::]:50051`
- Then, you need to position yourself withing your second tap/window and type the following.
`go run client/client.go` for you to make clients petitions to the running server.

### Proto generation commands

Generate proto files: 
``protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user/user.proto``

### Contributing

Please, feel free to fork this repository to contribute to it by submitting a functionalities/bugs-fixes pull request to enhance it.

### License

Please see [License File](https://github.com/gocanto/users-grpc-service/blob/main/LICENSE) for more information.

## How can I thank you?

There are many ways you would be able to support my open source work. There is not a right one to choose, so the choice is yours.

Nevertheless :grinning:, I would propose the following

- :arrow_up: Follow me on [Twitter](https://twitter.com/gocanto).
- :star: Star the repository.
- :handshake: Open a pull request to fix/improve the codebase.
- :writing_hand: Open a pull request to improve the documentation.
- :coffee: Buy me a [coffee](https://github.com/sponsors/gocanto)?

> Thank you for reading this far. :blush:
