package cookies

import (
	"errors"
	"strings"
	"syscall/js"
	"time"
)

var ErrNoCookieFound = errors.New("No cookie found")

func SetCookie(key, value string) {
	document := js.Global().Get("document")
	cookieStr := key + "=" + value + "; path=/;"
	document.Set("cookie", cookieStr)
}

func SetCookieWithExpiration(name, value string, tm time.Time) {
	document := js.Global().Get("document")
	expirationStr := tm.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	cookieStr := name + "=" + value + "; " +
		"expires=" + expirationStr + "; path=/;"
	document.Set("cookie", cookieStr)
}

func getCookie(name string) (string, error) {
	document := js.Global().Get("document")
	cookies := document.Get("cookie").String()

	// Cookies are like "key1=value1; key2=value2; ..."
	pairs := strings.Split(cookies, "; ")

	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 && parts[0] == name {
			return parts[1], nil
		}
	}
	// Cookie not found
	return "", ErrNoCookieFound
}
