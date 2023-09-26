// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/matching/common_inputs/ssl/v3/ssl_inputs.proto

package sslv3

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

// Validate checks the field values on UriSanInput with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UriSanInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UriSanInput with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UriSanInputMultiError, or
// nil if none found.
func (m *UriSanInput) ValidateAll() error {
	return m.validate(true)
}

func (m *UriSanInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UriSanInputMultiError(errors)
	}
	return nil
}

// UriSanInputMultiError is an error wrapping multiple validation errors
// returned by UriSanInput.ValidateAll() if the designated constraints aren't met.
type UriSanInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UriSanInputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UriSanInputMultiError) AllErrors() []error { return m }

// UriSanInputValidationError is the validation error returned by
// UriSanInput.Validate if the designated constraints aren't met.
type UriSanInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UriSanInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UriSanInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UriSanInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UriSanInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UriSanInputValidationError) ErrorName() string { return "UriSanInputValidationError" }

// Error satisfies the builtin error interface
func (e UriSanInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUriSanInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UriSanInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UriSanInputValidationError{}

// Validate checks the field values on DnsSanInput with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DnsSanInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsSanInput with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DnsSanInputMultiError, or
// nil if none found.
func (m *DnsSanInput) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsSanInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DnsSanInputMultiError(errors)
	}
	return nil
}

// DnsSanInputMultiError is an error wrapping multiple validation errors
// returned by DnsSanInput.ValidateAll() if the designated constraints aren't met.
type DnsSanInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsSanInputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsSanInputMultiError) AllErrors() []error { return m }

// DnsSanInputValidationError is the validation error returned by
// DnsSanInput.Validate if the designated constraints aren't met.
type DnsSanInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsSanInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsSanInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsSanInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsSanInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsSanInputValidationError) ErrorName() string { return "DnsSanInputValidationError" }

// Error satisfies the builtin error interface
func (e DnsSanInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsSanInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsSanInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsSanInputValidationError{}

// Validate checks the field values on SubjectInput with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubjectInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectInput with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubjectInputMultiError, or
// nil if none found.
func (m *SubjectInput) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SubjectInputMultiError(errors)
	}
	return nil
}

// SubjectInputMultiError is an error wrapping multiple validation errors
// returned by SubjectInput.ValidateAll() if the designated constraints aren't met.
type SubjectInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectInputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectInputMultiError) AllErrors() []error { return m }

// SubjectInputValidationError is the validation error returned by
// SubjectInput.Validate if the designated constraints aren't met.
type SubjectInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectInputValidationError) ErrorName() string { return "SubjectInputValidationError" }

// Error satisfies the builtin error interface
func (e SubjectInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectInputValidationError{}
