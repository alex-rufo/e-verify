# eVerify

eVerify is an email verifier that can run on HTTP and gRPC. You can also use eVerify as a library.

## Use it HTTP or gRPC server

To generate the binary, you'll need to download the repository and execute `make build`. The output will be located inside `bin` folder. These are the available flags:
```
--disposable-domains string   File where disposable domains are located
--disposable-roles string     File where disposable roles are located
--enable-smtp                 Enable SMTP?
--grpc-port int               Port for gRPC server (defaults 81)
--http-port int               Port for HTTP server (defaults 80)
--xverify-apiKey string       ApiKey of xVerify
--xverify-domain string       Domain to use with xVerify
```

## Use it as a library

```
go get github.com/alex-rufo/e-verify
```

The easiest way is to use eVerify is using the default verifier. This includes `lowercase`, `trim` and `gmail` sanitizers and `syntax`,  `disposable` (default disposable emails can be found [here](https://raw.githubusercontent.com/martenson/disposable-email-domains/master/allowlist.conf)) and `mx` validators. 
```golang
import verifier "github.com/alex-rufo/e-verify"

v := verifier.NewDefault()
valid, err := v.Verify("whatever@email.com")
```

You can also configure the verifier with the sanitizers and validators you want:
```golang
import (
    "time"

    verifier "github.com/alex-rufo/e-verify"
    "github.com/alex-rufo/e-verify/pkg/sanetization"
	"github.com/alex-rufo/e-verify/pkg/validation"
)

sanilizers := []verifier.Sanitizer{
    &sanitizer.Lowercase{},
    &sanitizer.Trim{},
}

validators := []verifier.Validator{
    &validation.Syntax{},
    validation.NewSMTP(&net.Dialer{Timeout: 1*time.Second}),
    validation.NewXVerify("apiKey", "domain", 1*time.Second),
}

v := verifier.New()
valid, err := v.Verify("whatever@email.com")
```

Keep in mind that sanitizers and validators will run in the order they are configured. In case a validator invalidates the email, the chain will stop and the result will be returned immediately. You can find all the available sanitizers inside `pkg/sanetization` and all the validators inside `pkg/validation`.

In case you want to create you own sanitizers or validators you just need to implement the following interfaces.
```golang
type Sanitizer interface {
	Sanitize(email string) (string, error)
}
```

```golang
type Validator interface {
	Validate(email string) (bool, error)
}
```