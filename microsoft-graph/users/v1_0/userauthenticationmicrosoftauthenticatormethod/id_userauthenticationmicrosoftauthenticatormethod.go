package userauthenticationmicrosoftauthenticatormethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationMicrosoftAuthenticatorMethodId{}

// UserAuthenticationMicrosoftAuthenticatorMethodId is a struct representing the Resource ID for a User Authentication Microsoft Authenticator Method
type UserAuthenticationMicrosoftAuthenticatorMethodId struct {
	UserId                                       string
	MicrosoftAuthenticatorAuthenticationMethodId string
}

// NewUserAuthenticationMicrosoftAuthenticatorMethodID returns a new UserAuthenticationMicrosoftAuthenticatorMethodId struct
func NewUserAuthenticationMicrosoftAuthenticatorMethodID(userId string, microsoftAuthenticatorAuthenticationMethodId string) UserAuthenticationMicrosoftAuthenticatorMethodId {
	return UserAuthenticationMicrosoftAuthenticatorMethodId{
		UserId: userId,
		MicrosoftAuthenticatorAuthenticationMethodId: microsoftAuthenticatorAuthenticationMethodId,
	}
}

// ParseUserAuthenticationMicrosoftAuthenticatorMethodID parses 'input' into a UserAuthenticationMicrosoftAuthenticatorMethodId
func ParseUserAuthenticationMicrosoftAuthenticatorMethodID(input string) (*UserAuthenticationMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationMicrosoftAuthenticatorMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MicrosoftAuthenticatorAuthenticationMethodId, ok = parsed.Parsed["microsoftAuthenticatorAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "microsoftAuthenticatorAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationMicrosoftAuthenticatorMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationMicrosoftAuthenticatorMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationMicrosoftAuthenticatorMethodIDInsensitively(input string) (*UserAuthenticationMicrosoftAuthenticatorMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationMicrosoftAuthenticatorMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationMicrosoftAuthenticatorMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MicrosoftAuthenticatorAuthenticationMethodId, ok = parsed.Parsed["microsoftAuthenticatorAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "microsoftAuthenticatorAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationMicrosoftAuthenticatorMethodID checks that 'input' can be parsed as a User Authentication Microsoft Authenticator Method ID
func ValidateUserAuthenticationMicrosoftAuthenticatorMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationMicrosoftAuthenticatorMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Microsoft Authenticator Method ID
func (id UserAuthenticationMicrosoftAuthenticatorMethodId) ID() string {
	fmtString := "/users/%s/authentication/microsoftAuthenticatorMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MicrosoftAuthenticatorAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Microsoft Authenticator Method ID
func (id UserAuthenticationMicrosoftAuthenticatorMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticMicrosoftAuthenticatorMethods", "microsoftAuthenticatorMethods", "microsoftAuthenticatorMethods"),
		resourceids.UserSpecifiedSegment("microsoftAuthenticatorAuthenticationMethodId", "microsoftAuthenticatorAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Microsoft Authenticator Method ID
func (id UserAuthenticationMicrosoftAuthenticatorMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Microsoft Authenticator Authentication Method: %q", id.MicrosoftAuthenticatorAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Microsoft Authenticator Method (%s)", strings.Join(components, "\n"))
}
