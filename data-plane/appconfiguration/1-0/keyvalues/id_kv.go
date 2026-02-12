package keyvalues

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &KvId{}

// KvId is a struct representing the Resource ID for a Kv
type KvId struct {
	BaseURI string
	KvName  string
}

// NewKvID returns a new KvId struct
func NewKvID(baseURI string, kvName string) KvId {
	return KvId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		KvName:  kvName,
	}
}

// ParseKvID parses 'input' into a KvId
func ParseKvID(input string) (*KvId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KvId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KvId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseKvIDInsensitively parses 'input' case-insensitively into a KvId
// note: this method should only be used for API response data and not user input
func ParseKvIDInsensitively(input string) (*KvId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KvId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KvId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *KvId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.KvName, ok = input.Parsed["kvName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "kvName", input)
	}

	return nil
}

// ValidateKvID checks that 'input' can be parsed as a Kv ID
func ValidateKvID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseKvID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Kv ID
func (id KvId) ID() string {
	fmtString := "%s/kv/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.KvName)
}

// Path returns the formatted Kv ID without the BaseURI
func (id KvId) Path() string {
	fmtString := "/kv/%s"
	return fmt.Sprintf(fmtString, id.KvName)
}

// PathElements returns the values of Kv ID Segments without the BaseURI
func (id KvId) PathElements() []any {
	return []any{id.KvName}
}

// Segments returns a slice of Resource ID Segments which comprise this Kv ID
func (id KvId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticKv", "kv", "kv"),
		resourceids.UserSpecifiedSegment("kvName", "kvName"),
	}
}

// String returns a human-readable description of this Kv ID
func (id KvId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Kv Name: %q", id.KvName),
	}
	return fmt.Sprintf("Kv (%s)", strings.Join(components, "\n"))
}
