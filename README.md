# Result

**Inspired by Result from Rust.**
This is a lightweight package for error handling, returning a Result type to indicate the success or failure of an operation instead of only using exceptions.

### Usage

```go
func main() {
	var url = "http://github.com/golang/go"

	makeRequest(url).AndThen(readBody).AndThen(print).OrElse(print)
}
```
