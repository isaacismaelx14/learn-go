package entry

import (
	"fmt"

	"github.com/isaacismaelx14/learn-go/pkg/calc"
	"github.com/isaacismaelx14/learn-go/pkg/greetings"
)

func Main() {
	greetingMsg := greetings.Get("Michael")
	sum := calc.Sum(33, 40)
	rest := calc.Rest(20, 5)

	fmt.Println(greetingMsg)
	fmt.Println("Sum: ", sum)
	fmt.Println("Rest: ", rest)

}
