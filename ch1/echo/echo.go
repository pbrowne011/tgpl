// Exercise 1.1

package echo

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
