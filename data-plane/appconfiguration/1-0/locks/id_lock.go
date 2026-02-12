package locks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LockId{}

// LockId is a struct representing the Resource ID for a Lock
type LockId struct {
	BaseURI  string
	LockName string
}

// NewLockID returns a new LockId struct
func NewLockID(baseURI string, lockName string) LockId {
	return LockId{
		BaseURI:  strings.TrimSuffix(baseURI, "/"),
		LockName: lockName,
	}
}

// ParseLockID parses 'input' into a LockId
func ParseLockID(input string) (*LockId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LockId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LockId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLockIDInsensitively parses 'input' case-insensitively into a LockId
// note: this method should only be used for API response data and not user input
func ParseLockIDInsensitively(input string) (*LockId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LockId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LockId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LockId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.LockName, ok = input.Parsed["lockName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "lockName", input)
	}

	return nil
}

// ValidateLockID checks that 'input' can be parsed as a Lock ID
func ValidateLockID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLockID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Lock ID
func (id LockId) ID() string {
	fmtString := "%s/locks/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.LockName)
}

// Path returns the formatted Lock ID without the BaseURI
func (id LockId) Path() string {
	fmtString := "/locks/%s"
	return fmt.Sprintf(fmtString, id.LockName)
}

// PathElements returns the values of Lock ID Segments without the BaseURI
func (id LockId) PathElements() []any {
	return []any{id.LockName}
}

// Segments returns a slice of Resource ID Segments which comprise this Lock ID
func (id LockId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticLocks", "locks", "locks"),
		resourceids.UserSpecifiedSegment("lockName", "lockName"),
	}
}

// String returns a human-readable description of this Lock ID
func (id LockId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Lock Name: %q", id.LockName),
	}
	return fmt.Sprintf("Lock (%s)", strings.Join(components, "\n"))
}
