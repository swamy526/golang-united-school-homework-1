package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	return "Hello :world_map:"
}

func main() {
	fmt.Println(emoji.Sprint(GetMessage()))
}
