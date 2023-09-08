package userauthenticationfido2method

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationFido2MethodId{}

// UserAuthenticationFido2MethodId is a struct representing the Resource ID for a User Authentication Fido 2 Method
type UserAuthenticationFido2MethodId struct {
	UserId                      string
	Fido2AuthenticationMethodId string
}

// NewUserAuthenticationFido2MethodID returns a new UserAuthenticationFido2MethodId struct
func NewUserAuthenticationFido2MethodID(userId string, fido2AuthenticationMethodId string) UserAuthenticationFido2MethodId {
	return UserAuthenticationFido2MethodId{
		UserId:                      userId,
		Fido2AuthenticationMethodId: fido2AuthenticationMethodId,
	}
}

// ParseUserAuthenticationFido2MethodID parses 'input' into a UserAuthenticationFido2MethodId
func ParseUserAuthenticationFido2MethodID(input string) (*UserAuthenticationFido2MethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationFido2MethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationFido2MethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.Fido2AuthenticationMethodId, ok = parsed.Parsed["fido2AuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fido2AuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationFido2MethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationFido2MethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationFido2MethodIDInsensitively(input string) (*UserAuthenticationFido2MethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationFido2MethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationFido2MethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.Fido2AuthenticationMethodId, ok = parsed.Parsed["fido2AuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fido2AuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationFido2MethodID checks that 'input' can be parsed as a User Authentication Fido 2 Method ID
func ValidateUserAuthenticationFido2MethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationFido2MethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Fido 2 Method ID
func (id UserAuthenticationFido2MethodId) ID() string {
	fmtString := "/users/%s/authentication/fido2Methods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.Fido2AuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Fido 2 Method ID
func (id UserAuthenticationFido2MethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticFido2Methods", "fido2Methods", "fido2Methods"),
		resourceids.UserSpecifiedSegment("fido2AuthenticationMethodId", "fido2AuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Fido 2 Method ID
func (id UserAuthenticationFido2MethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Fido 2 Authentication Method: %q", id.Fido2AuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Fido 2 Method (%s)", strings.Join(components, "\n"))
}
