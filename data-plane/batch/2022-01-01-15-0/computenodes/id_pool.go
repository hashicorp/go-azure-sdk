package computenodes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PoolId{}

// PoolId is a struct representing the Resource ID for a Pool
type PoolId struct {
	BaseURI string
	PoolId  string
}

// NewPoolID returns a new PoolId struct
func NewPoolID(baseURI string, poolId string) PoolId {
	return PoolId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		PoolId:  poolId,
	}
}

// ParsePoolID parses 'input' into a PoolId
func ParsePoolID(input string) (*PoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePoolIDInsensitively parses 'input' case-insensitively into a PoolId
// note: this method should only be used for API response data and not user input
func ParsePoolIDInsensitively(input string) (*PoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PoolId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.PoolId, ok = input.Parsed["poolId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "poolId", input)
	}

	return nil
}

// ValidatePoolID checks that 'input' can be parsed as a Pool ID
func ValidatePoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pool ID
func (id PoolId) ID() string {
	fmtString := "%s/pools/%s"
	return fmt.Sprintf(fmtString, id.BaseURI, id.PoolId)
}

// Path returns the formatted Pool ID without the BaseURI
func (id PoolId) Path() string {
	fmtString := "/pools/%s"
	return fmt.Sprintf(fmtString, id.PoolId)
}

// PathElements returns the values of Pool ID Segments without the BaseURI
func (id PoolId) PathElements() []any {
	return []any{id.PoolId}
}

// Segments returns a slice of Resource ID Segments which comprise this Pool ID
func (id PoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint_url"),
		resourceids.StaticSegment("staticPools", "pools", "pools"),
		resourceids.UserSpecifiedSegment("poolId", "poolId"),
	}
}

// String returns a human-readable description of this Pool ID
func (id PoolId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Pool: %q", id.PoolId),
	}
	return fmt.Sprintf("Pool (%s)", strings.Join(components, "\n"))
}
