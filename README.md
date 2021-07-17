
## Usage
```go
import (
	"fmt"
	"log"

	token "github.com/getstreamai/go-sdk"
)

func main() {
	request := &token.RequestData{
		Name:   "Harkirat",
		Room:   "kirat-room",
		Type:   "producer",
		Record: true,
	}
	requestBody := token.RequestBody{
		Data:         *request,
		AccessKey:    "<your-access-key>",
		AccessSecret: "<your-access-secret>",
	}
	token, err := token.GenerateToken(&requestBody)
	if err != nil {
		log.Fatal("Err response from generate token")
	}
	fmt.Println(token)
}
```
