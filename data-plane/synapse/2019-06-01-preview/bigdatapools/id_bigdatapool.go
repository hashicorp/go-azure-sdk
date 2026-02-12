package bigdatapools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BigDataPoolId{}

// BigDataPoolId is a struct representing the Resource ID for a Big Data Pool
type BigDataPoolId struct {
	BaseURI         string
	BigDataPoolName string
}

// NewBigDataPoolID returns a new BigDataPoolId struct
func NewBigDataPoolID(baseURI string, bigDataPoolName string) BigDataPoolId {
	return BigDataPoolId{
		BaseURI:         strings.TrimSuffix(baseURI, "/"),
		BigDataPoolName: bigDataPoolName,
	}
}

// ParseBigDataPoolID parses 'input' into a BigDataPoolId
func ParseBigDataPoolID(input string) (*BigDataPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BigDataPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BigDataPoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBigDataPoolIDInsensitively parses 'input' case-insensitively into a BigDataPoolId
// note: this method should only be used for API response data and not user input
func ParseBigDataPoolIDInsensitively(input string) (*BigDataPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BigDataPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BigDataPoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BigDataPoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.BigDataPoolName, ok = input.Parsed["bigDataPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bigDataPoolName", input)
	}

	return nil
}

// ValidateBigDataPoolID checks that 'input' can be parsed as a Big Data Pool ID
func ValidateBigDataPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBigDataPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Big Data Pool ID
func (id BigDataPoolId) ID() string {
	fmtString := "%s/bigDataPools/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.BigDataPoolName)
}

// Path returns the formatted Big Data Pool ID without the BaseURI
func (id BigDataPoolId) Path() string {
	fmtString := "/bigDataPools/%s"
	return fmt.Sprintf(fmtString, id.BigDataPoolName)
}

// PathElements returns the values of Big Data Pool ID Segments without the BaseURI
func (id BigDataPoolId) PathElements() []any {
	return []any{id.BigDataPoolName}
}

// Segments returns a slice of Resource ID Segments which comprise this Big Data Pool ID
func (id BigDataPoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticBigDataPools", "bigDataPools", "bigDataPools"),
		resourceids.UserSpecifiedSegment("bigDataPoolName", "bigDataPoolName"),
	}
}

// String returns a human-readable description of this Big Data Pool ID
func (id BigDataPoolId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Big Data Pool Name: %q", id.BigDataPoolName),
	}
	return fmt.Sprintf("Big Data Pool (%s)", strings.Join(components, "\n"))
}
