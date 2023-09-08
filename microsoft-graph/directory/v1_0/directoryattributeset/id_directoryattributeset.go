package directoryattributeset

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryAttributeSetId{}

// DirectoryAttributeSetId is a struct representing the Resource ID for a Directory Attribute Set
type DirectoryAttributeSetId struct {
	AttributeSetId string
}

// NewDirectoryAttributeSetID returns a new DirectoryAttributeSetId struct
func NewDirectoryAttributeSetID(attributeSetId string) DirectoryAttributeSetId {
	return DirectoryAttributeSetId{
		AttributeSetId: attributeSetId,
	}
}

// ParseDirectoryAttributeSetID parses 'input' into a DirectoryAttributeSetId
func ParseDirectoryAttributeSetID(input string) (*DirectoryAttributeSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAttributeSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAttributeSetId{}

	if id.AttributeSetId, ok = parsed.Parsed["attributeSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attributeSetId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryAttributeSetIDInsensitively parses 'input' case-insensitively into a DirectoryAttributeSetId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAttributeSetIDInsensitively(input string) (*DirectoryAttributeSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAttributeSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAttributeSetId{}

	if id.AttributeSetId, ok = parsed.Parsed["attributeSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attributeSetId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryAttributeSetID checks that 'input' can be parsed as a Directory Attribute Set ID
func ValidateDirectoryAttributeSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAttributeSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Attribute Set ID
func (id DirectoryAttributeSetId) ID() string {
	fmtString := "/directory/attributeSets/%s"
	return fmt.Sprintf(fmtString, id.AttributeSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Attribute Set ID
func (id DirectoryAttributeSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticAttributeSets", "attributeSets", "attributeSets"),
		resourceids.UserSpecifiedSegment("attributeSetId", "attributeSetIdValue"),
	}
}

// String returns a human-readable description of this Directory Attribute Set ID
func (id DirectoryAttributeSetId) String() string {
	components := []string{
		fmt.Sprintf("Attribute Set: %q", id.AttributeSetId),
	}
	return fmt.Sprintf("Directory Attribute Set (%s)", strings.Join(components, "\n"))
}
