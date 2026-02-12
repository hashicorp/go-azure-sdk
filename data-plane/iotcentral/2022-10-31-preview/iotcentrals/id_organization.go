package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &OrganizationId{}

// OrganizationId is a struct representing the Resource ID for a Organization
type OrganizationId struct {
	BaseURI        string
	OrganizationId string
}

// NewOrganizationID returns a new OrganizationId struct
func NewOrganizationID(baseURI string, organizationId string) OrganizationId {
	return OrganizationId{
		BaseURI:        strings.TrimSuffix(baseURI, "/"),
		OrganizationId: organizationId,
	}
}

// ParseOrganizationID parses 'input' into a OrganizationId
func ParseOrganizationID(input string) (*OrganizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OrganizationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OrganizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOrganizationIDInsensitively parses 'input' case-insensitively into a OrganizationId
// note: this method should only be used for API response data and not user input
func ParseOrganizationIDInsensitively(input string) (*OrganizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OrganizationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OrganizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OrganizationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.OrganizationId, ok = input.Parsed["organizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "organizationId", input)
	}

	return nil
}

// ValidateOrganizationID checks that 'input' can be parsed as a Organization ID
func ValidateOrganizationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOrganizationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Organization ID
func (id OrganizationId) ID() string {
	fmtString := "%s/organizations/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.OrganizationId)
}

// Path returns the formatted Organization ID without the BaseURI
func (id OrganizationId) Path() string {
	fmtString := "/organizations/%s"
	return fmt.Sprintf(fmtString, id.OrganizationId)
}

// PathElements returns the values of Organization ID Segments without the BaseURI
func (id OrganizationId) PathElements() []any {
	return []any{id.OrganizationId}
}

// Segments returns a slice of Resource ID Segments which comprise this Organization ID
func (id OrganizationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticOrganizations", "organizations", "organizations"),
		resourceids.UserSpecifiedSegment("organizationId", "organizationId"),
	}
}

// String returns a human-readable description of this Organization ID
func (id OrganizationId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Organization: %q", id.OrganizationId),
	}
	return fmt.Sprintf("Organization (%s)", strings.Join(components, "\n"))
}
