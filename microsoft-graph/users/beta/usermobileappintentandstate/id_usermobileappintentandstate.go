package usermobileappintentandstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMobileAppIntentAndStateId{}

// UserMobileAppIntentAndStateId is a struct representing the Resource ID for a User Mobile App Intent And State
type UserMobileAppIntentAndStateId struct {
	UserId                    string
	MobileAppIntentAndStateId string
}

// NewUserMobileAppIntentAndStateID returns a new UserMobileAppIntentAndStateId struct
func NewUserMobileAppIntentAndStateID(userId string, mobileAppIntentAndStateId string) UserMobileAppIntentAndStateId {
	return UserMobileAppIntentAndStateId{
		UserId:                    userId,
		MobileAppIntentAndStateId: mobileAppIntentAndStateId,
	}
}

// ParseUserMobileAppIntentAndStateID parses 'input' into a UserMobileAppIntentAndStateId
func ParseUserMobileAppIntentAndStateID(input string) (*UserMobileAppIntentAndStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMobileAppIntentAndStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMobileAppIntentAndStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MobileAppIntentAndStateId, ok = parsed.Parsed["mobileAppIntentAndStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppIntentAndStateId", *parsed)
	}

	return &id, nil
}

// ParseUserMobileAppIntentAndStateIDInsensitively parses 'input' case-insensitively into a UserMobileAppIntentAndStateId
// note: this method should only be used for API response data and not user input
func ParseUserMobileAppIntentAndStateIDInsensitively(input string) (*UserMobileAppIntentAndStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMobileAppIntentAndStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMobileAppIntentAndStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MobileAppIntentAndStateId, ok = parsed.Parsed["mobileAppIntentAndStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppIntentAndStateId", *parsed)
	}

	return &id, nil
}

// ValidateUserMobileAppIntentAndStateID checks that 'input' can be parsed as a User Mobile App Intent And State ID
func ValidateUserMobileAppIntentAndStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMobileAppIntentAndStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mobile App Intent And State ID
func (id UserMobileAppIntentAndStateId) ID() string {
	fmtString := "/users/%s/mobileAppIntentAndStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MobileAppIntentAndStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mobile App Intent And State ID
func (id UserMobileAppIntentAndStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMobileAppIntentAndStates", "mobileAppIntentAndStates", "mobileAppIntentAndStates"),
		resourceids.UserSpecifiedSegment("mobileAppIntentAndStateId", "mobileAppIntentAndStateIdValue"),
	}
}

// String returns a human-readable description of this User Mobile App Intent And State ID
func (id UserMobileAppIntentAndStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mobile App Intent And State: %q", id.MobileAppIntentAndStateId),
	}
	return fmt.Sprintf("User Mobile App Intent And State (%s)", strings.Join(components, "\n"))
}
