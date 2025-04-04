package bootstrap

import "strconv"

const (
	// Inactive Client is created, but not able to exchange messages using Mitras.
	Inactive State = iota
	// Active Client is created, configured, and whitelisted.
	Active
)

// State represents corresponding Mitras Client state. The possible Config States
// as well as description of what that State represents are given in the table:
// | State    | What it means                                                                  |
// |----------+--------------------------------------------------------------------------------|
// | Inactive | Client is created, but isn't able to communicate over Mitras                  |
// | Active   | Client is able to communicate using Mitras                                    |.
type State int

// String returns string representation of State.
func (s State) String() string {
	return strconv.Itoa(int(s))
}