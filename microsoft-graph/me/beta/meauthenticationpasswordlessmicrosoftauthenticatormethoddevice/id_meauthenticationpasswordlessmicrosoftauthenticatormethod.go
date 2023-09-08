package meauthenticationpasswordlessmicrosoftauthenticatormethoddevice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId is a struct representing the Resource ID for a Me Authentication Passwordless Microsoft Authenticator Method
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId struct {
	PasswordlessMicrosoftAuthenticatorAuthenticationMethodId string
}

// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID returns a new MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId struct
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(passwordlessMicrosoftAuthenticatorAuthenticationMethodId string) MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId {
	return MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
		PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: passwordlessMicrosoftAuthenticatorAuthenticationMethodId,
	}
}

// ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID parses 'input' into a MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
func ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(input string) (*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, ok = parsed.Parsed["passwordlessMicrosoftAuthenticatorAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordlessMicrosoftAuthenticatorAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively parses 'input' case-insensitively into a MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
// note: this method should only be used for API response data and not user input
func ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(input string) (*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, ok = parsed.Parsed["passwordlessMicrosoftAuthenticatorAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordlessMicrosoftAuthenticatorAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID checks that 'input' can be parsed as a Me Authentication Passwordless Microsoft Authenticator Method ID
func ValidateMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Authentication Passwordless Microsoft Authenticator Method ID
func (id MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) ID() string {
	fmtString := "/me/authentication/passwordlessMicrosoftAuthenticatorMethods/%s"
	return fmt.Sprintf(fmtString, id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Authentication Passwordless Microsoft Authenticator Method ID
func (id MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticPasswordlessMicrosoftAuthenticatorMethods", "passwordlessMicrosoftAuthenticatorMethods", "passwordlessMicrosoftAuthenticatorMethods"),
		resourceids.UserSpecifiedSegment("passwordlessMicrosoftAuthenticatorAuthenticationMethodId", "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this Me Authentication Passwordless Microsoft Authenticator Method ID
func (id MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) String() string {
	components := []string{
		fmt.Sprintf("Passwordless Microsoft Authenticator Authentication Method: %q", id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId),
	}
	return fmt.Sprintf("Me Authentication Passwordless Microsoft Authenticator Method (%s)", strings.Join(components, "\n"))
}
