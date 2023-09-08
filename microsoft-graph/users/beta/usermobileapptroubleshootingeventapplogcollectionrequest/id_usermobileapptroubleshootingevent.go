package usermobileapptroubleshootingeventapplogcollectionrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMobileAppTroubleshootingEventId{}

// UserMobileAppTroubleshootingEventId is a struct representing the Resource ID for a User Mobile App Troubleshooting Event
type UserMobileAppTroubleshootingEventId struct {
	UserId                          string
	MobileAppTroubleshootingEventId string
}

// NewUserMobileAppTroubleshootingEventID returns a new UserMobileAppTroubleshootingEventId struct
func NewUserMobileAppTroubleshootingEventID(userId string, mobileAppTroubleshootingEventId string) UserMobileAppTroubleshootingEventId {
	return UserMobileAppTroubleshootingEventId{
		UserId:                          userId,
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
	}
}

// ParseUserMobileAppTroubleshootingEventID parses 'input' into a UserMobileAppTroubleshootingEventId
func ParseUserMobileAppTroubleshootingEventID(input string) (*UserMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMobileAppTroubleshootingEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MobileAppTroubleshootingEventId, ok = parsed.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", *parsed)
	}

	return &id, nil
}

// ParseUserMobileAppTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a UserMobileAppTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseUserMobileAppTroubleshootingEventIDInsensitively(input string) (*UserMobileAppTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMobileAppTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMobileAppTroubleshootingEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MobileAppTroubleshootingEventId, ok = parsed.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", *parsed)
	}

	return &id, nil
}

// ValidateUserMobileAppTroubleshootingEventID checks that 'input' can be parsed as a User Mobile App Troubleshooting Event ID
func ValidateUserMobileAppTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMobileAppTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mobile App Troubleshooting Event ID
func (id UserMobileAppTroubleshootingEventId) ID() string {
	fmtString := "/users/%s/mobileAppTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MobileAppTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mobile App Troubleshooting Event ID
func (id UserMobileAppTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventIdValue"),
	}
}

// String returns a human-readable description of this User Mobile App Troubleshooting Event ID
func (id UserMobileAppTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
	}
	return fmt.Sprintf("User Mobile App Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
