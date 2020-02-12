# Result

**Inspired by Result from Rust.**
This is a lightweight package for error handling, returning a Result type to indicate the success or failure of an operation instead of only using exceptions.

### Usage

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alissonbrunosa/result"
)

func main() {
	var url = "http://github.com/golang/go"

	makeRequest(url).AndThen(readBody).AndThen(print).OrElse(print)
}

func print(v interface{}) result.Result {
	fmt.Println(v)

	return result.Ok(v)
}

func readBody(v interface{}) result.Result {
	resp, ok := v.(*http.Response)

	if !ok {
		return result.Err("Not a http.Response")
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return result.Err(err)
	}

	return result.Ok(string(bytes))
}

func makeRequest(url string) result.Result {
	resp, err := http.Get(url)
	if err != nil {
		return result.Err(err)
	}

	return result.Ok(resp)
}

```
