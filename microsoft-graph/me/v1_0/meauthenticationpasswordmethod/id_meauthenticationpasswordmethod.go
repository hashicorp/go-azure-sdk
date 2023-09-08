package meauthenticationpasswordmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAuthenticationPasswordMethodId{}

// MeAuthenticationPasswordMethodId is a struct representing the Resource ID for a Me Authentication Password Method
type MeAuthenticationPasswordMethodId struct {
	PasswordAuthenticationMethodId string
}

// NewMeAuthenticationPasswordMethodID returns a new MeAuthenticationPasswordMethodId struct
func NewMeAuthenticationPasswordMethodID(passwordAuthenticationMethodId string) MeAuthenticationPasswordMethodId {
	return MeAuthenticationPasswordMethodId{
		PasswordAuthenticationMethodId: passwordAuthenticationMethodId,
	}
}

// ParseMeAuthenticationPasswordMethodID parses 'input' into a MeAuthenticationPasswordMethodId
func ParseMeAuthenticationPasswordMethodID(input string) (*MeAuthenticationPasswordMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationPasswordMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationPasswordMethodId{}

	if id.PasswordAuthenticationMethodId, ok = parsed.Parsed["passwordAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseMeAuthenticationPasswordMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationPasswordMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationPasswordMethodIDInsensitively(input string) (*MeAuthenticationPasswordMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationPasswordMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationPasswordMethodId{}

	if id.PasswordAuthenticationMethodId, ok = parsed.Parsed["passwordAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateMeAuthenticationPasswordMethodID checks that 'input' can be parsed as a Me Authentication Password Method ID
func ValidateMeAuthenticationPasswordMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationPasswordMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Password Method ID
func (id MeAuthenticationPasswordMethodId) ID() string {
	fmtString := "/me/authentication/passwordMethods/%s"
	return fmt.Sprintf(fmtString, id.PasswordAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Password Method ID
func (id MeAuthenticationPasswordMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticPasswordMethods", "passwordMethods", "passwordMethods"),
		resourceids.UserSpecifiedSegment("passwordAuthenticationMethodId", "passwordAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this Me Authentication Password Method ID
func (id MeAuthenticationPasswordMethodId) String() string {
	components := []string{
		fmt.Sprintf("Password Authentication Method: %q", id.PasswordAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Password Method (%s)", strings.Join(components, "\n"))
}
