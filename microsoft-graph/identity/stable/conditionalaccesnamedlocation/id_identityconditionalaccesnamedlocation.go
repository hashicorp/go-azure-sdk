package conditionalaccesnamedlocation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccesNamedLocationId{}

// IdentityConditionalAccesNamedLocationId is a struct representing the Resource ID for a Identity Conditional Acces Named Location
type IdentityConditionalAccesNamedLocationId struct {
	NamedLocationId string
}

// NewIdentityConditionalAccesNamedLocationID returns a new IdentityConditionalAccesNamedLocationId struct
func NewIdentityConditionalAccesNamedLocationID(namedLocationId string) IdentityConditionalAccesNamedLocationId {
	return IdentityConditionalAccesNamedLocationId{
		NamedLocationId: namedLocationId,
	}
}

// ParseIdentityConditionalAccesNamedLocationID parses 'input' into a IdentityConditionalAccesNamedLocationId
func ParseIdentityConditionalAccesNamedLocationID(input string) (*IdentityConditionalAccesNamedLocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesNamedLocationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesNamedLocationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccesNamedLocationIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesNamedLocationId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesNamedLocationIDInsensitively(input string) (*IdentityConditionalAccesNamedLocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesNamedLocationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesNamedLocationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccesNamedLocationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.NamedLocationId, ok = input.Parsed["namedLocationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "namedLocationId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccesNamedLocationID checks that 'input' can be parsed as a Identity Conditional Acces Named Location ID
func ValidateIdentityConditionalAccesNamedLocationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesNamedLocationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Named Location ID
func (id IdentityConditionalAccesNamedLocationId) ID() string {
	fmtString := "/identity/conditionalAccess/namedLocations/%s"
	return fmt.Sprintf(fmtString, id.NamedLocationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Named Location ID
func (id IdentityConditionalAccesNamedLocationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("namedLocations", "namedLocations", "namedLocations"),
		resourceids.UserSpecifiedSegment("namedLocationId", "namedLocationIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Named Location ID
func (id IdentityConditionalAccesNamedLocationId) String() string {
	components := []string{
		fmt.Sprintf("Named Location: %q", id.NamedLocationId),
	}
	return fmt.Sprintf("Identity Conditional Acces Named Location (%s)", strings.Join(components, "\n"))
}
