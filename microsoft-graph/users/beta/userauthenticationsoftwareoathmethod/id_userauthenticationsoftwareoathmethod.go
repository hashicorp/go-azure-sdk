package userauthenticationsoftwareoathmethod

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAuthenticationSoftwareOathMethodId{}

// UserAuthenticationSoftwareOathMethodId is a struct representing the Resource ID for a User Authentication Software Oath Method
type UserAuthenticationSoftwareOathMethodId struct {
	UserId                             string
	SoftwareOathAuthenticationMethodId string
}

// NewUserAuthenticationSoftwareOathMethodID returns a new UserAuthenticationSoftwareOathMethodId struct
func NewUserAuthenticationSoftwareOathMethodID(userId string, softwareOathAuthenticationMethodId string) UserAuthenticationSoftwareOathMethodId {
	return UserAuthenticationSoftwareOathMethodId{
		UserId:                             userId,
		SoftwareOathAuthenticationMethodId: softwareOathAuthenticationMethodId,
	}
}

// ParseUserAuthenticationSoftwareOathMethodID parses 'input' into a UserAuthenticationSoftwareOathMethodId
func ParseUserAuthenticationSoftwareOathMethodID(input string) (*UserAuthenticationSoftwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationSoftwareOathMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationSoftwareOathMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SoftwareOathAuthenticationMethodId, ok = parsed.Parsed["softwareOathAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "softwareOathAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ParseUserAuthenticationSoftwareOathMethodIDInsensitively parses 'input' case-insensitively into a UserAuthenticationSoftwareOathMethodId
// note: this method should only be used for API response data and not user input
func ParseUserAuthenticationSoftwareOathMethodIDInsensitively(input string) (*UserAuthenticationSoftwareOathMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAuthenticationSoftwareOathMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAuthenticationSoftwareOathMethodId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.SoftwareOathAuthenticationMethodId, ok = parsed.Parsed["softwareOathAuthenticationMethodId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "softwareOathAuthenticationMethodId", *parsed)
	}

	return &id, nil
}

// ValidateUserAuthenticationSoftwareOathMethodID checks that 'input' can be parsed as a User Authentication Software Oath Method ID
func ValidateUserAuthenticationSoftwareOathMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAuthenticationSoftwareOathMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Authentication Software Oath Method ID
func (id UserAuthenticationSoftwareOathMethodId) ID() string {
	fmtString := "/users/%s/authentication/softwareOathMethods/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.SoftwareOathAuthenticationMethodId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Authentication Software Oath Method ID
func (id UserAuthenticationSoftwareOathMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAuthentication", "authentication", "authentication"),
		resourceids.StaticSegment("staticSoftwareOathMethods", "softwareOathMethods", "softwareOathMethods"),
		resourceids.UserSpecifiedSegment("softwareOathAuthenticationMethodId", "softwareOathAuthenticationMethodIdValue"),
	}
}

// String returns a human-readable description of this User Authentication Software Oath Method ID
func (id UserAuthenticationSoftwareOathMethodId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Software Oath Authentication Method: %q", id.SoftwareOathAuthenticationMethodId),
	}
	return fmt.Sprintf("User Authentication Software Oath Method (%s)", strings.Join(components, "\n"))
}
