# Picarto.tv API Wrapper

> This wrapper is not the complete Picarto API. Only the public endpoints are represented. If you like the wrapper 
> and wish to help extend it, feel free to submit a PR.

## Use

This wrapper is intended to be used only in a private project and is provided publicly to be pulled as an import. Use
this library at your own risk.

### Setting Client ID and/or Client Secrets

In order to use this package, you MUST set the client ID at a minimum. Please note that some endpoints do require the
use of a client secret. To set these, pass them in when first initializing the wrapper. The wrapper will automatically
pass the client ID and Secret where needed. If one or both is needed and not set or incorrect, the request will return
an error.

```go
api.Rest = api.NewPicarto("CLIENT_ID", "CLIENT_SECRET")
```

## Pull Requests

Pull requests will be accepted on a case-by-case basis to expand upon the library and fix bugs.