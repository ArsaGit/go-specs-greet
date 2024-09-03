package interactions_test

import (
	"testing"

	"github.com/ArsaGit/go-specs-greet/domain/interactions"
	"github.com/ArsaGit/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
