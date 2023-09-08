package meauthenticationemailmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAuthenticationEmailMethodId{}

// MeAuthenticationEmailMethodId is a struct representing the Resource ID for a Me Authentication Email Method
type MeAuthenticationEmailMethodId struct {
	EmailAuthenticationMethodId string
}

// NewMeAuthenticationEmailMethodID returns a new MeAuthenticationEmailMethodId struct
func NewMeAuthenticationEmailMethodID(emailAuthenticationMethodId string) MeAuthenticationEmailMethodId {
	return MeAuthenticationEmailMethodId{
		EmailAuthenticationMethodId: emailAuthenticationMethodId,
	}
}

// ParseMeAuthenticationEmailMethodID parses 'input' into a MeAuthenticationEmailMethodId
func ParseMeAuthenticationEmailMethodID(input string) (*MeAuthenticationEmailMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationEmailMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationEmailMethodId{}

	if id.EmailAuthenticationMethodId, ok = parsed.Parsed["emailAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "emailAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseMeAuthenticationEmailMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationEmailMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationEmailMethodIDInsensitively(input string) (*MeAuthenticationEmailMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationEmailMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationEmailMethodId{}

	if id.EmailAuthenticationMethodId, ok = parsed.Parsed["emailAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "emailAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateMeAuthenticationEmailMethodID checks that 'input' can be parsed as a Me Authentication Email Method ID
func ValidateMeAuthenticationEmailMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationEmailMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Email Method ID
func (id MeAuthenticationEmailMethodId) ID() string {
	fmtString := "/me/authentication/emailMethods/%s"
	return fmt.Sprintf(fmtString, id.EmailAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Email Method ID
func (id MeAuthenticationEmailMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticEmailMethods", "emailMethods", "emailMethods"),
		resourceids.UserSpecifiedSegment("emailAuthenticationMethodId", "emailAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this Me Authentication Email Method ID
func (id MeAuthenticationEmailMethodId) String() string {
	components := []string{
		fmt.Sprintf("Email Authentication Method: %q", id.EmailAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Email Method (%s)", strings.Join(components, "\n"))
}
