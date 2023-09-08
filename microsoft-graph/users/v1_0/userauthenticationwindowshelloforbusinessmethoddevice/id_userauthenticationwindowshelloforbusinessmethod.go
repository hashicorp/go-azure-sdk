package userauthenticationwindowshelloforbusinessmethoddevice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationWindowsHelloForBusinessMethodId{}

// UserAuthenticationWindowsHelloForBusinessMethodId is a struct representing the Resource ID for a User Authentication Windows Hello For Business Method
type UserAuthenticationWindowsHelloForBusinessMethodId struct {
	UserId                                        string
	WindowsHelloForBusinessAuthenticationMethodId string
}

// NewUserAuthenticationWindowsHelloForBusinessMethodID returns a new UserAuthenticationWindowsHelloForBusinessMethodId struct
func NewUserAuthenticationWindowsHelloForBusinessMethodID(userId string, windowsHelloForBusinessAuthenticationMethodId string) UserAuthenticationWindowsHelloForBusinessMethodId {
	return UserAuthenticationWindowsHelloForBusinessMethodId{
		UserId: userId,
		WindowsHelloForBusinessAuthenticationMethodId: windowsHelloForBusinessAuthenticationMethodId,
	}
}

// ParseUserAuthenticationWindowsHelloForBusinessMethodID parses 'input' into a UserAuthenticationWindowsHelloForBusinessMethodId
func ParseUserAuthenticationWindowsHelloForBusinessMethodID(input string) (*UserAuthenticationWindowsHelloForBusinessMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationWindowsHelloForBusinessMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationWindowsHelloForBusinessMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WindowsHelloForBusinessAuthenticationMethodId, ok = parsed.Parsed["windowsHelloForBusinessAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "windowsHelloForBusinessAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationWindowsHelloForBusinessMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationWindowsHelloForBusinessMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationWindowsHelloForBusinessMethodIDInsensitively(input string) (*UserAuthenticationWindowsHelloForBusinessMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationWindowsHelloForBusinessMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationWindowsHelloForBusinessMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.WindowsHelloForBusinessAuthenticationMethodId, ok = parsed.Parsed["windowsHelloForBusinessAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "windowsHelloForBusinessAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationWindowsHelloForBusinessMethodID checks that 'input' can be parsed as a User Authentication Windows Hello For Business Method ID
func ValidateUserAuthenticationWindowsHelloForBusinessMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationWindowsHelloForBusinessMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Windows Hello For Business Method ID
func (id UserAuthenticationWindowsHelloForBusinessMethodId) ID() string {
	fmtString := "/users/%s/authentication/windowsHelloForBusinessMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.WindowsHelloForBusinessAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Windows Hello For Business Method ID
func (id UserAuthenticationWindowsHelloForBusinessMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticWindowsHelloForBusinessMethods", "windowsHelloForBusinessMethods", "windowsHelloForBusinessMethods"),
		resourceids.UserSpecifiedSegment("windowsHelloForBusinessAuthenticationMethodId", "windowsHelloForBusinessAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Windows Hello For Business Method ID
func (id UserAuthenticationWindowsHelloForBusinessMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Windows Hello For Business Authentication Method: %q", id.WindowsHelloForBusinessAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Windows Hello For Business Method (%s)", strings.Join(components, "\n"))
}
