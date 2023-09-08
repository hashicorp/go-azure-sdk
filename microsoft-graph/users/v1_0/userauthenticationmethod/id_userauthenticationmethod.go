package userauthenticationmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationMethodId{}

// UserAuthenticationMethodId is a struct representing the Resource ID for a User Authentication Method
type UserAuthenticationMethodId struct {
	UserId                 string
	AuthenticationMethodId string
}

// NewUserAuthenticationMethodID returns a new UserAuthenticationMethodId struct
func NewUserAuthenticationMethodID(userId string, authenticationMethodId string) UserAuthenticationMethodId {
	return UserAuthenticationMethodId{
		UserId:                 userId,
		AuthenticationMethodId: authenticationMethodId,
	}
}

// ParseUserAuthenticationMethodID parses 'input' into a UserAuthenticationMethodId
func ParseUserAuthenticationMethodID(input string) (*UserAuthenticationMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AuthenticationMethodId, ok = parsed.Parsed["authenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationMethodIDInsensitively(input string) (*UserAuthenticationMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AuthenticationMethodId, ok = parsed.Parsed["authenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationMethodID checks that 'input' can be parsed as a User Authentication Method ID
func ValidateUserAuthenticationMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Method ID
func (id UserAuthenticationMethodId) ID() string {
	fmtString := "/users/%s/authentication/methods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Method ID
func (id UserAuthenticationMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticMethods", "methods", "methods"),
		resourceids.UserSpecifiedSegment("authenticationMethodId", "authenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Method ID
func (id UserAuthenticationMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Authentication Method: %q", id.AuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Method (%s)", strings.Join(components, "\n"))
}
