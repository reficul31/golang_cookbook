package trees

import "errors"

var (
	// ErrDuplicateNode is returned when an insertion is performed on a duplicate node
	ErrDuplicateNode = errors.New("Duplicate Node insertion")
)