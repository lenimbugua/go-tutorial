import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v", name)
	return message
}
