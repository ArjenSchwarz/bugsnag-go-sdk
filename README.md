# Bugsnag Go SDK

This is an implementation in Go of the [Bugsnag API](https://bugsnag.com/docs/api).

# Status

Very much in early development. Don't try to use this, as I'm using it to learn Go and will therefore be making a lot of changes to it going forward.

# Usage

Install using `go get`:

```bash
go get github.com/arjenschwarz/bugsnag-go-sdk
```

Include in your project file:

```golang
import (
    bugsnag "github.com/arjenschwarz/bugsnag-go-sdk"
)
```

## Authenticate

You can authenticate with a username/password combination or using an API key:

```golang
creds := bugsnag.UserPass(username, password)
creds := bugsnag.ApiKey(apiKey)
```

Using this you can then open a connection:

```golang
connection := bugsnag.New(creds)
```

## Make API calls

And with the connection you can then make all the calls you wish. For example, to get a list of all your accounts you would use:

```golang
data, err := connection.Accounts(params)
```

Full documentation to follow.