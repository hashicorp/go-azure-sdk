package userauthenticationphonemethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationPhoneMethodId{}

// UserAuthenticationPhoneMethodId is a struct representing the Resource ID for a User Authentication Phone Method
type UserAuthenticationPhoneMethodId struct {
	UserId                      string
	PhoneAuthenticationMethodId string
}

// NewUserAuthenticationPhoneMethodID returns a new UserAuthenticationPhoneMethodId struct
func NewUserAuthenticationPhoneMethodID(userId string, phoneAuthenticationMethodId string) UserAuthenticationPhoneMethodId {
	return UserAuthenticationPhoneMethodId{
		UserId:                      userId,
		PhoneAuthenticationMethodId: phoneAuthenticationMethodId,
	}
}

// ParseUserAuthenticationPhoneMethodID parses 'input' into a UserAuthenticationPhoneMethodId
func ParseUserAuthenticationPhoneMethodID(input string) (*UserAuthenticationPhoneMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationPhoneMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationPhoneMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PhoneAuthenticationMethodId, ok = parsed.Parsed["phoneAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "phoneAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationPhoneMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationPhoneMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationPhoneMethodIDInsensitively(input string) (*UserAuthenticationPhoneMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationPhoneMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationPhoneMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.PhoneAuthenticationMethodId, ok = parsed.Parsed["phoneAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "phoneAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationPhoneMethodID checks that 'input' can be parsed as a User Authentication Phone Method ID
func ValidateUserAuthenticationPhoneMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationPhoneMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Phone Method ID
func (id UserAuthenticationPhoneMethodId) ID() string {
	fmtString := "/users/%s/authentication/phoneMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.PhoneAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Phone Method ID
func (id UserAuthenticationPhoneMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticPhoneMethods", "phoneMethods", "phoneMethods"),
		resourceids.UserSpecifiedSegment("phoneAuthenticationMethodId", "phoneAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Phone Method ID
func (id UserAuthenticationPhoneMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Phone Authentication Method: %q", id.PhoneAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Phone Method (%s)", strings.Join(components, "\n"))
}
