package queryengine

// QueryEngine - interface to all implementations of things that
// want to be the query engine. It defines the contracts and expected
// returns to query engines

// id associates an id, from the 511 api, with an object
// uses typing to enforce that the user knows what
// they're doing, or something like that
type ID string

type QueryEngine interface {
	GetOperatorID() (map[ID]Operator, error)
	// GetLineID(oid ID)
	// GetStopID(oid, lid string)
	// GetStopMonintor(sid string)
}

type Operator struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
