package userauthenticationpasswordlessmicrosoftauthenticatormethoddevice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

// UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId is a struct representing the Resource ID for a User Authentication Passwordless Microsoft Authenticator Method
type UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId struct {
	UserId                                                   string
	PasswordlessMicrosoftAuthenticatorAuthenticationMethodId string
}

// NewUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID returns a new UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId struct
func NewUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(userId string, passwordlessMicrosoftAuthenticatorAuthenticationMethodId string) UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId {
	return UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{
		UserId: userId,
		PasswordlessMicrosoftAuthenticatorAuthenticationMethodId: passwordlessMicrosoftAuthenticatorAuthenticationMethodId,
	}
}

// ParseUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID parses 'input' into a UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
func ParseUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(input string) (*UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, ok = parsed.Parsed["passwordlessMicrosoftAuthenticatorAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordlessMicrosoftAuthenticatorAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodIDInsensitively(input string) (*UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId, ok = parsed.Parsed["passwordlessMicrosoftAuthenticatorAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordlessMicrosoftAuthenticatorAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID checks that 'input' can be parsed as a User Authentication Passwordless Microsoft Authenticator Method ID
func ValidateUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Passwordless Microsoft Authenticator Method ID
func (id UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) ID() string {
	fmtString := "/users/%s/authentication/passwordlessMicrosoftAuthenticatorMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Passwordless Microsoft Authenticator Method ID
func (id UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticPasswordlessMicrosoftAuthenticatorMethods", "passwordlessMicrosoftAuthenticatorMethods", "passwordlessMicrosoftAuthenticatorMethods"),
		resourceids.UserSpecifiedSegment("passwordlessMicrosoftAuthenticatorAuthenticationMethodId", "passwordlessMicrosoftAuthenticatorAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Passwordless Microsoft Authenticator Method ID
func (id UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Passwordless Microsoft Authenticator Authentication Method: %q", id.PasswordlessMicrosoftAuthenticatorAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Passwordless Microsoft Authenticator Method (%s)", strings.Join(components, "\n"))
}
