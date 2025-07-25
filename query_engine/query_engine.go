package queryengine

// QueryEngine - interface to all implementations of things that
// want to be the query engine. It defines the contracts and expected
// returns to query engines

// id associates an id, from the 511 api, with an object
// uses typing to enforce that the user knows what
// they're doing, or something like that
type ID string

// TODO: Make a decision on how to better ensure type safety

type QueryEngine interface {
	GetOperatorID() (map[ID]Operator, error)
	// TODO: make oid more type specific
	GetLineID(oid string) (map[ID]Line, error)
	// GetStopID(oid, lid string) (map[ID]Stop, error)
	// GetStopMonintor(sid string)
}

// for the types being returned by the query engine, there are two ids.
// One is ID, and it is for DB indexing
// the other is the specific Id that the 511 API returns as a means for referencing the
// operator, which will help in other query engine resolutions

type Operator struct {
	ID         string `json:"id"`
	OperatorID string `json:"operatorid"`
	Name       string `json:"name"`
}

type Line struct {
	ID     string `json:"id"`
	LineID string `json:"lineid"`
	Name   string `json:"name"`
}

/*
type Stop struct {
	ID     string `json:"id"`
	LineID string `json:"stopid"`
	Name   string `json:"name"`
}
*/
