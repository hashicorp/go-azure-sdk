package userauthenticationemailmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationEmailMethodId{}

// UserAuthenticationEmailMethodId is a struct representing the Resource ID for a User Authentication Email Method
type UserAuthenticationEmailMethodId struct {
	UserId                      string
	EmailAuthenticationMethodId string
}

// NewUserAuthenticationEmailMethodID returns a new UserAuthenticationEmailMethodId struct
func NewUserAuthenticationEmailMethodID(userId string, emailAuthenticationMethodId string) UserAuthenticationEmailMethodId {
	return UserAuthenticationEmailMethodId{
		UserId:                      userId,
		EmailAuthenticationMethodId: emailAuthenticationMethodId,
	}
}

// ParseUserAuthenticationEmailMethodID parses 'input' into a UserAuthenticationEmailMethodId
func ParseUserAuthenticationEmailMethodID(input string) (*UserAuthenticationEmailMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationEmailMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationEmailMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EmailAuthenticationMethodId, ok = parsed.Parsed["emailAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "emailAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationEmailMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationEmailMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationEmailMethodIDInsensitively(input string) (*UserAuthenticationEmailMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationEmailMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationEmailMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EmailAuthenticationMethodId, ok = parsed.Parsed["emailAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "emailAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationEmailMethodID checks that 'input' can be parsed as a User Authentication Email Method ID
func ValidateUserAuthenticationEmailMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationEmailMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Email Method ID
func (id UserAuthenticationEmailMethodId) ID() string {
	fmtString := "/users/%s/authentication/emailMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EmailAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Email Method ID
func (id UserAuthenticationEmailMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticEmailMethods", "emailMethods", "emailMethods"),
		resourceids.UserSpecifiedSegment("emailAuthenticationMethodId", "emailAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Email Method ID
func (id UserAuthenticationEmailMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Email Authentication Method: %q", id.EmailAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Email Method (%s)", strings.Join(components, "\n"))
}
