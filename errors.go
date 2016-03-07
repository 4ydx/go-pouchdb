// +build js

package pouchdb

import (
	"github.com/gopherjs/gopherjs/js"
)

// PouchError records an error returned by the PouchDB library
type PouchError struct {
	e       *js.Error
	status  int
	message string
	name    string
	isError bool
}

// Error satisfies the error interface for the PouchError type
func (e *PouchError) Error() string {
	return e.message
}

// ErrorStatus returns the status of a PouchError, or 0 for other errors
func ErrorStatus(err error) int {
	switch pe := err.(type) {
	case *PouchError:
		return pe.status
	}
	return 0
}

// ErrorMessage returns the message portion of a PouchError, or "" for other errors
func ErrorMessage(err error) string {
	switch pe := err.(type) {
	case *PouchError:
		return pe.message
	}
	return ""
}

// ErrorName returns the name portion of a PouchError, or "" for other errors
func ErrorName(err error) string {
	switch pe := err.(type) {
	case *PouchError:
		return pe.name
	}
	return ""
}

// IsNotExist returns true if the passed error represents a PouchError with
// a status of 404 (not found)
func IsNotExist(err error) bool {
	switch pe := err.(type) {
	case *PouchError:
		return pe.status == 404
	}
	return false
}

// IsConflict returns true if the passed error is a PouchError with a status
// of 409 (conflict)
func IsConflict(err error) bool {
	switch pe := err.(type) {
	case *PouchError:
		return pe.status == 409
	}
	return false
}

// IsPouchError returns true if the passed error is a PouchError, false
// if it is any other type of error.
func IsPouchError(err error) bool {
	switch err.(type) {
	case *PouchError:
		return true
	}
	return false
}

// NewPouchError creates a new PouchError from a js.Error object returned from the PouchDB library
func NewPouchError(err *js.Error) error {
	if err == nil {
		return nil
	}
	return &PouchError{
		e:       err,
		status:  err.Get("status").Int(),
		message: err.Get("message").String(),
		name:    err.Get("name").String(),
		isError: err.Get("error").Bool(),
	}
}

// Underlying returns the underlying js.Error object, as returned from the PouchDB library
func (e *PouchError) Underlying() *js.Error {
	return e.e
}
