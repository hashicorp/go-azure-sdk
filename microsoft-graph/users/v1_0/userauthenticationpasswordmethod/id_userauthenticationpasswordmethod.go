package userauthenticationpasswordmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationPasswordMethodId{}

// UserAuthenticationPasswordMethodId is a struct representing the Resource ID for a User Authentication Password Method
type UserAuthenticationPasswordMethodId struct {
	UserId                         string
	PasswordAuthenticationMethodId string
}

// NewUserAuthenticationPasswordMethodID returns a new UserAuthenticationPasswordMethodId struct
func NewUserAuthenticationPasswordMethodID(userId string, passwordAuthenticationMethodId string) UserAuthenticationPasswordMethodId {
	return UserAuthenticationPasswordMethodId{
		UserId:                         userId,
		PasswordAuthenticationMethodId: passwordAuthenticationMethodId,
	}
}

// ParseUserAuthenticationPasswordMethodID parses 'input' into a UserAuthenticationPasswordMethodId
func ParseUserAuthenticationPasswordMethodID(input string) (*UserAuthenticationPasswordMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationPasswordMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationPasswordMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PasswordAuthenticationMethodId, ok = parsed.Parsed["passwordAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationPasswordMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationPasswordMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationPasswordMethodIDInsensitively(input string) (*UserAuthenticationPasswordMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationPasswordMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationPasswordMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PasswordAuthenticationMethodId, ok = parsed.Parsed["passwordAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "passwordAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationPasswordMethodID checks that 'input' can be parsed as a User Authentication Password Method ID
func ValidateUserAuthenticationPasswordMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationPasswordMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Password Method ID
func (id UserAuthenticationPasswordMethodId) ID() string {
	fmtString := "/users/%s/authentication/passwordMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PasswordAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Password Method ID
func (id UserAuthenticationPasswordMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticPasswordMethods", "passwordMethods", "passwordMethods"),
		resourceids.UserSpecifiedSegment("passwordAuthenticationMethodId", "passwordAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Password Method ID
func (id UserAuthenticationPasswordMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Password Authentication Method: %q", id.PasswordAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Password Method (%s)", strings.Join(components, "\n"))
}
