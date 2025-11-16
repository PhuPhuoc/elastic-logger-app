package common

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// AppError represents a structured error for the application.
// It includes fields for client response, logging, and debugging.
type AppError struct {
	// Code: HTTP status code to be returned to the client.
	Code int `json:"code"`
	// Message: A user-friendly message intended for the client.
	// Should not contain sensitive information.
	Message string `json:"message"`
	// ReasonField: A more specific, potentially technical description of the error.
	// Can be used for logging and debugging.
	ReasonField string `json:"reason_field,omitempty"`
	// Details: A map for storing arbitrary detailed information.
	// Useful for logging context (e.g., IDs, parameters) without exposing them to the client directly.
	// Consider if you want this in the JSON response based on security needs.
	Details map[string]any `json:"details,omitempty"`
	// File: The file path where the error originated (from runtime.Caller).
	File string `json:"file,omitempty"`
	// Line: The line number in the file where the error originated.
	Line int `json:"line,omitempty"`
	// Function: The function name where the error originated.
	Function string `json:"function,omitempty"`
	// Inner: The underlying error (if any) that caused this AppError.
	// Useful for debugging the root cause.
	Inner error `json:"-"` // Not serialized to JSON to prevent exposing internal errors
	// Timestamp: When the error occurred.
	Timestamp time.Time `json:"timestamp"`
	// ErrorID: A unique identifier for this specific error instance.
	// Useful for tracking and searching logs.
	ErrorID string `json:"error_id,omitempty"`
}

// NewAppError creates a new AppError instance.
// If flag_location is true, it captures the file, line, and function where it was called.
func NewAppError(code int, message string, reason string, flag_location bool) *AppError {
	err := &AppError{
		Code:        code,
		Message:     message,
		ReasonField: reason,
		Details:     make(map[string]any), // Initialize the map
		Timestamp:   time.Now().UTC(),
	}

	if flag_location {
		file, line, fn := getCallerInfo()
		err.File = file
		err.Line = line
		err.Function = fn
	}

	return err
}

// Helper functions for common HTTP errors, capturing location by default.
func NewBadRequestError(message string, reason string) *AppError {
	return NewAppError(http.StatusBadRequest, message, reason, true).WithReason(reason)
}

func NewInternalServerError(message string, reason string) *AppError {
	return NewAppError(http.StatusInternalServerError, message, reason, true).WithReason(reason)
}

func NewUnauthorizedError(message string) *AppError {
	return NewAppError(http.StatusUnauthorized, message, "Unauthorized access attempt", false)
}

func NewForbiddenError(message string) *AppError {
	return NewAppError(http.StatusForbidden, message, "Forbidden resource access", false)
}

func NewNotFoundError(message string, reason string) *AppError {
	return NewAppError(http.StatusNotFound, message, reason, false)
}

// StatusCode returns the HTTP status code associated with the error.
func (e *AppError) StatusCode() int {
	return e.Code
}

// WithMessage sets a new message for the error.
func (e *AppError) WithMessage(message string) *AppError {
	e.Message = message
	return e
}

// WithReason sets a new reason field for the error.
func (e *AppError) WithReason(reason string) *AppError {
	e.ReasonField = reason
	return e
}

// WithInner wraps an underlying error.
func (e *AppError) WithInner(inner error) *AppError {
	e.Inner = inner
	return e
}

// WithDetail adds a key-value pair to the details map.
func (e *AppError) WithDetail(key string, value any) *AppError {
	if e.Details == nil {
		e.Details = make(map[string]any)
	}
	e.Details[key] = value
	return e
}

// WithErrorID sets a unique identifier for the error instance.
func (e *AppError) WithErrorID(id string) *AppError {
	e.ErrorID = id
	return e
}

// Error implements the error interface.
// It provides a string representation of the error, primarily for logging.
// It includes the underlying error if present.
func (e *AppError) Error() string {
	baseMsg := fmt.Sprintf("Code: %d, Message: '%s', Reason: '%s'", e.Code, e.Message, e.ReasonField)
	if e.Inner != nil {
		baseMsg += fmt.Sprintf(", Inner: %v", e.Inner)
	}
	return baseMsg
}

// FullErrorString provides a more detailed string representation, including location info.
// Useful for detailed logging.
func (e *AppError) FullErrorString() string {
	baseMsg := e.Error()
	if e.File != "" && e.Line != 0 {
		baseMsg += fmt.Sprintf(" (at %s:%d", e.File, e.Line)
		if e.Function != "" {
			baseMsg += fmt.Sprintf(" in %s", e.Function)
		}
		baseMsg += ")"
	}
	if e.ErrorID != "" {
		baseMsg += fmt.Sprintf(" [ErrorID: %s]", e.ErrorID)
	}
	return baseMsg
}

// getCallerInfo retrieves the file, line, and function name of the caller.
// The 'skip' value is 3 because:
// 0: getCallerInfo
// 1: NewAppError (or helper like NewInternalServerError)
// 2: The function that called NewAppError/NewInternalServerError
// 3: The actual line in the code where the error was initiated
func getCallerInfo() (string, int, string) {
	pc, file, line, ok := runtime.Caller(3) // Adjusted skip
	if !ok {
		return "unknown", 0, "unknown"
	}
	fn := runtime.FuncForPC(pc)
	var funcName string
	if fn != nil {
		funcName = fn.Name()
		// Optionally, simplify the function name (remove package path)
		// funcName = filepath.Base(funcName)
	}
	return file, line, funcName
}
