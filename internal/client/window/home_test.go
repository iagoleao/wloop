package window

import (
	"os"
	"regexp"
	"testing"
)

func TestHomeRouteNamesTheCockpitsHome(t *testing.T) {
	src, err := os.ReadFile("../../web/ui/src/lib/routes.ts")
	if err != nil {
		t.Fatal(err)
	}
	m := regexp.MustCompile(`export const HOME = "([^"]+)";`).FindSubmatch(src)
	if m == nil {
		t.Fatal("routes.ts declares no HOME route")
	}
	if want := "view=" + string(m[1]); HomeRoute != want {
		t.Fatalf("HomeRoute is %q, the cockpit's home is %q", HomeRoute, want)
	}
}
