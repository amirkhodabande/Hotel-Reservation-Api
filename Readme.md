# Hotel Reservation API

## Simple rest-api for managing hotel reservation

> **Make sure you have mongo up and running**

- **To run project** \
`cp .env.sample .env` \
`make run`
- **To run all tests** \
`make test`
- **To seed database** \
`make seed`

### Tests are in tests folder

the setup method is responsible for setting up the test environment

- *Database initializing, will read connection info from .env.testing*
- *Loading the Service-Container*

> The setup will inject the service container, for being able to mock any special service implementation \

```go
 mockedServices := container.Services{
  ExampleServicer: &mockedImplementationOfExampleService,
 }

 app, tdb := setup(t, &mockedServices)
```

- *Building the app, and registering routes*
- *Initializing the factory*

> sample usage of factory in: ***tests/auth_handler_test.go***
