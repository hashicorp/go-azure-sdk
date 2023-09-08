package memanagedappregistration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedAppRegistrationId{}

// MeManagedAppRegistrationId is a struct representing the Resource ID for a Me Managed App Registration
type MeManagedAppRegistrationId struct {
	ManagedAppRegistrationId string
}

// NewMeManagedAppRegistrationID returns a new MeManagedAppRegistrationId struct
func NewMeManagedAppRegistrationID(managedAppRegistrationId string) MeManagedAppRegistrationId {
	return MeManagedAppRegistrationId{
		ManagedAppRegistrationId: managedAppRegistrationId,
	}
}

// ParseMeManagedAppRegistrationID parses 'input' into a MeManagedAppRegistrationId
func ParseMeManagedAppRegistrationID(input string) (*MeManagedAppRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedAppRegistrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedAppRegistrationId{}

	if id.ManagedAppRegistrationId, ok = parsed.Parsed["managedAppRegistrationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedAppRegistrationId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedAppRegistrationIDInsensitively parses 'input' case-insensitively into a MeManagedAppRegistrationId
// note: this method should only be used for API response data and not user input
func ParseMeManagedAppRegistrationIDInsensitively(input string) (*MeManagedAppRegistrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedAppRegistrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedAppRegistrationId{}

	if id.ManagedAppRegistrationId, ok = parsed.Parsed["managedAppRegistrationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedAppRegistrationId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedAppRegistrationID checks that 'input' can be parsed as a Me Managed App Registration ID
func ValidateMeManagedAppRegistrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedAppRegistrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed App Registration ID
func (id MeManagedAppRegistrationId) ID() string {
	fmtString := "/me/managedAppRegistrations/%s"
	return fmt.Sprintf(fmtString, id.ManagedAppRegistrationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed App Registration ID
func (id MeManagedAppRegistrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedAppRegistrations", "managedAppRegistrations", "managedAppRegistrations"),
		resourceids.UserSpecifiedSegment("managedAppRegistrationId", "managedAppRegistrationIdValue"),
	}
}

// String returns a human-readable description of this Me Managed App Registration ID
func (id MeManagedAppRegistrationId) String() string {
	components := []string{
		fmt.Sprintf("Managed App Registration: %q", id.ManagedAppRegistrationId),
	}
	return fmt.Sprintf("Me Managed App Registration (%s)", strings.Join(components, "\n"))
}
