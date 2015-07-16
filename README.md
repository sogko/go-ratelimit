# go-ratelimit
Golang package to ratelimit messages/connections per unit time

## Usage
```go
import "github.com/sogko/go-ratelimit"

func main() {
  rate := 500
  per := 8
  rl := ratelimit.NewRateLimiter(rate, per)
  ....
  if rl.Limit() {
    drop_connection(conn)
  }
}
```

## Notes
- Modified implementation by [Sudhi Herle](https://code.google.com/p/go-wiki/wiki/RateLimiting)
- Used per ```second``` unit time.
  Previously it erroneously was in ```nanosecond```.
- Allowed user to specify per unit time (for eg: 500 messages per 8 seconds)

## Credits:
- Sudhi Herle <sudhi-dot-herle-at-gmail-com>
  - Reference: https://code.google.com/p/go-wiki/wiki/RateLimiting
- Anti Huimaa
  - Reference: http://stackoverflow.com/questions/667508/whats-a-good-rate-limiting-algorithm

## Links
- [wehavefaces.net](http://wehavefaces.net)
- [twitter.com/sogko](http://twitter.com/sogko)
- [github.com/sogko](http://github.com/sogko)
- [github.com/sogko](http://github.com/sogko)

## License
Copyright (c) 2015 Hafiz Ismail. This software is licensed under the GPLv2.
