package interactions

import "fmt"

/* TODO Add unit test
Add a unit test to our Greet function to default the name to World if it is empty.
You should see how simple this is,
and then the business rules are reflected in both applications for "free"
*/
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
