package testing

import "testing"

func TestSimple(t *testing.T) {
	route := InitializeRoute()
	route.RouteUser()
}
