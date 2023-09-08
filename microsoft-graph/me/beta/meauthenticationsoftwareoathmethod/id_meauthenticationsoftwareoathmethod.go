package meauthenticationsoftwareoathmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAuthenticationSoftwareOathMethodId{}

// MeAuthenticationSoftwareOathMethodId is a struct representing the Resource ID for a Me Authentication Software Oath Method
type MeAuthenticationSoftwareOathMethodId struct {
	SoftwareOathAuthenticationMethodId string
}

// NewMeAuthenticationSoftwareOathMethodID returns a new MeAuthenticationSoftwareOathMethodId struct
func NewMeAuthenticationSoftwareOathMethodID(softwareOathAuthenticationMethodId string) MeAuthenticationSoftwareOathMethodId {
	return MeAuthenticationSoftwareOathMethodId{
		SoftwareOathAuthenticationMethodId: softwareOathAuthenticationMethodId,
	}
}

// ParseMeAuthenticationSoftwareOathMethodID parses 'input' into a MeAuthenticationSoftwareOathMethodId
func ParseMeAuthenticationSoftwareOathMethodID(input string) (*MeAuthenticationSoftwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationSoftwareOathMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationSoftwareOathMethodId{}

	if id.SoftwareOathAuthenticationMethodId, ok = parsed.Parsed["softwareOathAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "softwareOathAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseMeAuthenticationSoftwareOathMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationSoftwareOathMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationSoftwareOathMethodIDInsensitively(input string) (*MeAuthenticationSoftwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationSoftwareOathMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationSoftwareOathMethodId{}

	if id.SoftwareOathAuthenticationMethodId, ok = parsed.Parsed["softwareOathAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "softwareOathAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateMeAuthenticationSoftwareOathMethodID checks that 'input' can be parsed as a Me Authentication Software Oath Method ID
func ValidateMeAuthenticationSoftwareOathMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationSoftwareOathMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Software Oath Method ID
func (id MeAuthenticationSoftwareOathMethodId) ID() string {
	fmtString := "/me/authentication/softwareOathMethods/%s"
	return fmt.Sprintf(fmtString, id.SoftwareOathAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Software Oath Method ID
func (id MeAuthenticationSoftwareOathMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticSoftwareOathMethods", "softwareOathMethods", "softwareOathMethods"),
		resourceids.UserSpecifiedSegment("softwareOathAuthenticationMethodId", "softwareOathAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this Me Authentication Software Oath Method ID
func (id MeAuthenticationSoftwareOathMethodId) String() string {
	components := []string{
		fmt.Sprintf("Software Oath Authentication Method: %q", id.SoftwareOathAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Software Oath Method (%s)", strings.Join(components, "\n"))
}
