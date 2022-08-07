// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hollow/v1/forest.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on PushLeafRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PushLeafRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushLeafRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PushLeafRequestMultiError, or nil if none found.
func (m *PushLeafRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PushLeafRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Userid

	// no validation rules for Message

	if len(errors) > 0 {
		return PushLeafRequestMultiError(errors)
	}

	return nil
}

// PushLeafRequestMultiError is an error wrapping multiple validation errors
// returned by PushLeafRequest.ValidateAll() if the designated constraints
// aren't met.
type PushLeafRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushLeafRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushLeafRequestMultiError) AllErrors() []error { return m }

// PushLeafRequestValidationError is the validation error returned by
// PushLeafRequest.Validate if the designated constraints aren't met.
type PushLeafRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushLeafRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushLeafRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushLeafRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushLeafRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushLeafRequestValidationError) ErrorName() string { return "PushLeafRequestValidationError" }

// Error satisfies the builtin error interface
func (e PushLeafRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushLeafRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushLeafRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushLeafRequestValidationError{}

// Validate checks the field values on GetLeafsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetLeafsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLeafsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLeafsRequestMultiError, or nil if none found.
func (m *GetLeafsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLeafsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Pagesize

	if len(errors) > 0 {
		return GetLeafsRequestMultiError(errors)
	}

	return nil
}

// GetLeafsRequestMultiError is an error wrapping multiple validation errors
// returned by GetLeafsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetLeafsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLeafsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLeafsRequestMultiError) AllErrors() []error { return m }

// GetLeafsRequestValidationError is the validation error returned by
// GetLeafsRequest.Validate if the designated constraints aren't met.
type GetLeafsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLeafsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLeafsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLeafsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLeafsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLeafsRequestValidationError) ErrorName() string { return "GetLeafsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetLeafsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLeafsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLeafsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLeafsRequestValidationError{}

// Validate checks the field values on Leaf with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Leaf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Leaf with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LeafMultiError, or nil if none found.
func (m *Leaf) ValidateAll() error {
	return m.validate(true)
}

func (m *Leaf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Owner

	// no validation rules for Message

	// no validation rules for CreateAt

	// no validation rules for Status

	if len(errors) > 0 {
		return LeafMultiError(errors)
	}

	return nil
}

// LeafMultiError is an error wrapping multiple validation errors returned by
// Leaf.ValidateAll() if the designated constraints aren't met.
type LeafMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LeafMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LeafMultiError) AllErrors() []error { return m }

// LeafValidationError is the validation error returned by Leaf.Validate if the
// designated constraints aren't met.
type LeafValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LeafValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LeafValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LeafValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LeafValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LeafValidationError) ErrorName() string { return "LeafValidationError" }

// Error satisfies the builtin error interface
func (e LeafValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLeaf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LeafValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LeafValidationError{}

// Validate checks the field values on MultipleTodoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MultipleTodoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MultipleTodoReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MultipleTodoReplyMultiError, or nil if none found.
func (m *MultipleTodoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *MultipleTodoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MultipleTodoReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MultipleTodoReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultipleTodoReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return MultipleTodoReplyMultiError(errors)
	}

	return nil
}

// MultipleTodoReplyMultiError is an error wrapping multiple validation errors
// returned by MultipleTodoReply.ValidateAll() if the designated constraints
// aren't met.
type MultipleTodoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MultipleTodoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MultipleTodoReplyMultiError) AllErrors() []error { return m }

// MultipleTodoReplyValidationError is the validation error returned by
// MultipleTodoReply.Validate if the designated constraints aren't met.
type MultipleTodoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultipleTodoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultipleTodoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultipleTodoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultipleTodoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultipleTodoReplyValidationError) ErrorName() string {
	return "MultipleTodoReplyValidationError"
}

// Error satisfies the builtin error interface
func (e MultipleTodoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultipleTodoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultipleTodoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultipleTodoReplyValidationError{}

// Validate checks the field values on PushLeafReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PushLeafReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushLeafReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PushLeafReplyMultiError, or
// nil if none found.
func (m *PushLeafReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PushLeafReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Msg

	if len(errors) > 0 {
		return PushLeafReplyMultiError(errors)
	}

	return nil
}

// PushLeafReplyMultiError is an error wrapping multiple validation errors
// returned by PushLeafReply.ValidateAll() if the designated constraints
// aren't met.
type PushLeafReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushLeafReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushLeafReplyMultiError) AllErrors() []error { return m }

// PushLeafReplyValidationError is the validation error returned by
// PushLeafReply.Validate if the designated constraints aren't met.
type PushLeafReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushLeafReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushLeafReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushLeafReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushLeafReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushLeafReplyValidationError) ErrorName() string { return "PushLeafReplyValidationError" }

// Error satisfies the builtin error interface
func (e PushLeafReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushLeafReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushLeafReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushLeafReplyValidationError{}

// Validate checks the field values on GetLeafsReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetLeafsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLeafsReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetLeafsReplyMultiError, or
// nil if none found.
func (m *GetLeafsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLeafsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Msg

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetLeafsReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetLeafsReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetLeafsReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetLeafsReplyMultiError(errors)
	}

	return nil
}

// GetLeafsReplyMultiError is an error wrapping multiple validation errors
// returned by GetLeafsReply.ValidateAll() if the designated constraints
// aren't met.
type GetLeafsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLeafsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLeafsReplyMultiError) AllErrors() []error { return m }

// GetLeafsReplyValidationError is the validation error returned by
// GetLeafsReply.Validate if the designated constraints aren't met.
type GetLeafsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLeafsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLeafsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLeafsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLeafsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLeafsReplyValidationError) ErrorName() string { return "GetLeafsReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetLeafsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLeafsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLeafsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLeafsReplyValidationError{}
