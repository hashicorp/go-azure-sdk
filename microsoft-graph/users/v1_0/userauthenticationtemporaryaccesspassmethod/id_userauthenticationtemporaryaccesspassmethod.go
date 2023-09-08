package userauthenticationtemporaryaccesspassmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationTemporaryAccessPassMethodId{}

// UserAuthenticationTemporaryAccessPassMethodId is a struct representing the Resource ID for a User Authentication Temporary Access Pass Method
type UserAuthenticationTemporaryAccessPassMethodId struct {
	UserId                                    string
	TemporaryAccessPassAuthenticationMethodId string
}

// NewUserAuthenticationTemporaryAccessPassMethodID returns a new UserAuthenticationTemporaryAccessPassMethodId struct
func NewUserAuthenticationTemporaryAccessPassMethodID(userId string, temporaryAccessPassAuthenticationMethodId string) UserAuthenticationTemporaryAccessPassMethodId {
	return UserAuthenticationTemporaryAccessPassMethodId{
		UserId: userId,
		TemporaryAccessPassAuthenticationMethodId: temporaryAccessPassAuthenticationMethodId,
	}
}

// ParseUserAuthenticationTemporaryAccessPassMethodID parses 'input' into a UserAuthenticationTemporaryAccessPassMethodId
func ParseUserAuthenticationTemporaryAccessPassMethodID(input string) (*UserAuthenticationTemporaryAccessPassMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationTemporaryAccessPassMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationTemporaryAccessPassMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TemporaryAccessPassAuthenticationMethodId, ok = parsed.Parsed["temporaryAccessPassAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "temporaryAccessPassAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationTemporaryAccessPassMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationTemporaryAccessPassMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationTemporaryAccessPassMethodIDInsensitively(input string) (*UserAuthenticationTemporaryAccessPassMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationTemporaryAccessPassMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationTemporaryAccessPassMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TemporaryAccessPassAuthenticationMethodId, ok = parsed.Parsed["temporaryAccessPassAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "temporaryAccessPassAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationTemporaryAccessPassMethodID checks that 'input' can be parsed as a User Authentication Temporary Access Pass Method ID
func ValidateUserAuthenticationTemporaryAccessPassMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationTemporaryAccessPassMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Temporary Access Pass Method ID
func (id UserAuthenticationTemporaryAccessPassMethodId) ID() string {
	fmtString := "/users/%s/authentication/temporaryAccessPassMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TemporaryAccessPassAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Temporary Access Pass Method ID
func (id UserAuthenticationTemporaryAccessPassMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticTemporaryAccessPassMethods", "temporaryAccessPassMethods", "temporaryAccessPassMethods"),
		resourceids.UserSpecifiedSegment("temporaryAccessPassAuthenticationMethodId", "temporaryAccessPassAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Temporary Access Pass Method ID
func (id UserAuthenticationTemporaryAccessPassMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Temporary Access Pass Authentication Method: %q", id.TemporaryAccessPassAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Temporary Access Pass Method (%s)", strings.Join(components, "\n"))
}
