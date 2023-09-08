package directorycustomsecurityattributedefinitionallowedvalue

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryCustomSecurityAttributeDefinitionAllowedValueId{}

// DirectoryCustomSecurityAttributeDefinitionAllowedValueId is a struct representing the Resource ID for a Directory Custom Security Attribute Definition Allowed Value
type DirectoryCustomSecurityAttributeDefinitionAllowedValueId struct {
	CustomSecurityAttributeDefinitionId string
	AllowedValueId                      string
}

// NewDirectoryCustomSecurityAttributeDefinitionAllowedValueID returns a new DirectoryCustomSecurityAttributeDefinitionAllowedValueId struct
func NewDirectoryCustomSecurityAttributeDefinitionAllowedValueID(customSecurityAttributeDefinitionId string, allowedValueId string) DirectoryCustomSecurityAttributeDefinitionAllowedValueId {
	return DirectoryCustomSecurityAttributeDefinitionAllowedValueId{
		CustomSecurityAttributeDefinitionId: customSecurityAttributeDefinitionId,
		AllowedValueId:                      allowedValueId,
	}
}

// ParseDirectoryCustomSecurityAttributeDefinitionAllowedValueID parses 'input' into a DirectoryCustomSecurityAttributeDefinitionAllowedValueId
func ParseDirectoryCustomSecurityAttributeDefinitionAllowedValueID(input string) (*DirectoryCustomSecurityAttributeDefinitionAllowedValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryCustomSecurityAttributeDefinitionAllowedValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryCustomSecurityAttributeDefinitionAllowedValueId{}

	if id.CustomSecurityAttributeDefinitionId, ok = parsed.Parsed["customSecurityAttributeDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "customSecurityAttributeDefinitionId", *parsed)
	}

	if id.AllowedValueId, ok = parsed.Parsed["allowedValueId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "allowedValueId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryCustomSecurityAttributeDefinitionAllowedValueIDInsensitively parses 'input' case-insensitively into a DirectoryCustomSecurityAttributeDefinitionAllowedValueId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCustomSecurityAttributeDefinitionAllowedValueIDInsensitively(input string) (*DirectoryCustomSecurityAttributeDefinitionAllowedValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryCustomSecurityAttributeDefinitionAllowedValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryCustomSecurityAttributeDefinitionAllowedValueId{}

	if id.CustomSecurityAttributeDefinitionId, ok = parsed.Parsed["customSecurityAttributeDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "customSecurityAttributeDefinitionId", *parsed)
	}

	if id.AllowedValueId, ok = parsed.Parsed["allowedValueId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "allowedValueId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryCustomSecurityAttributeDefinitionAllowedValueID checks that 'input' can be parsed as a Directory Custom Security Attribute Definition Allowed Value ID
func ValidateDirectoryCustomSecurityAttributeDefinitionAllowedValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCustomSecurityAttributeDefinitionAllowedValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Custom Security Attribute Definition Allowed Value ID
func (id DirectoryCustomSecurityAttributeDefinitionAllowedValueId) ID() string {
	fmtString := "/directory/customSecurityAttributeDefinitions/%s/allowedValues/%s"
	return fmt.Sprintf(fmtString, id.CustomSecurityAttributeDefinitionId, id.AllowedValueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Custom Security Attribute Definition Allowed Value ID
func (id DirectoryCustomSecurityAttributeDefinitionAllowedValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticCustomSecurityAttributeDefinitions", "customSecurityAttributeDefinitions", "customSecurityAttributeDefinitions"),
		resourceids.UserSpecifiedSegment("customSecurityAttributeDefinitionId", "customSecurityAttributeDefinitionIdValue"),
		resourceids.StaticSegment("staticAllowedValues", "allowedValues", "allowedValues"),
		resourceids.UserSpecifiedSegment("allowedValueId", "allowedValueIdValue"),
	}
}

// String returns a human-readable description of this Directory Custom Security Attribute Definition Allowed Value ID
func (id DirectoryCustomSecurityAttributeDefinitionAllowedValueId) String() string {
	components := []string{
		fmt.Sprintf("Custom Security Attribute Definition: %q", id.CustomSecurityAttributeDefinitionId),
		fmt.Sprintf("Allowed Value: %q", id.AllowedValueId),
	}
	return fmt.Sprintf("Directory Custom Security Attribute Definition Allowed Value (%s)", strings.Join(components, "\n"))
}
