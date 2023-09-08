package usermobileapptroubleshootingeventapplogcollectionrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMobileAppTroubleshootingEventAppLogCollectionRequestId{}

// UserMobileAppTroubleshootingEventAppLogCollectionRequestId is a struct representing the Resource ID for a User Mobile App Troubleshooting Event App Log Collection Request
type UserMobileAppTroubleshootingEventAppLogCollectionRequestId struct {
	UserId                          string
	MobileAppTroubleshootingEventId string
	AppLogCollectionRequestId       string
}

// NewUserMobileAppTroubleshootingEventAppLogCollectionRequestID returns a new UserMobileAppTroubleshootingEventAppLogCollectionRequestId struct
func NewUserMobileAppTroubleshootingEventAppLogCollectionRequestID(userId string, mobileAppTroubleshootingEventId string, appLogCollectionRequestId string) UserMobileAppTroubleshootingEventAppLogCollectionRequestId {
	return UserMobileAppTroubleshootingEventAppLogCollectionRequestId{
		UserId:                          userId,
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
		AppLogCollectionRequestId:       appLogCollectionRequestId,
	}
}

// ParseUserMobileAppTroubleshootingEventAppLogCollectionRequestID parses 'input' into a UserMobileAppTroubleshootingEventAppLogCollectionRequestId
func ParseUserMobileAppTroubleshootingEventAppLogCollectionRequestID(input string) (*UserMobileAppTroubleshootingEventAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMobileAppTroubleshootingEventAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMobileAppTroubleshootingEventAppLogCollectionRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MobileAppTroubleshootingEventId, ok = parsed.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", *parsed)
	}

	if id.AppLogCollectionRequestId, ok = parsed.Parsed["appLogCollectionRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserMobileAppTroubleshootingEventAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a UserMobileAppTroubleshootingEventAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseUserMobileAppTroubleshootingEventAppLogCollectionRequestIDInsensitively(input string) (*UserMobileAppTroubleshootingEventAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMobileAppTroubleshootingEventAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMobileAppTroubleshootingEventAppLogCollectionRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MobileAppTroubleshootingEventId, ok = parsed.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", *parsed)
	}

	if id.AppLogCollectionRequestId, ok = parsed.Parsed["appLogCollectionRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserMobileAppTroubleshootingEventAppLogCollectionRequestID checks that 'input' can be parsed as a User Mobile App Troubleshooting Event App Log Collection Request ID
func ValidateUserMobileAppTroubleshootingEventAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMobileAppTroubleshootingEventAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Mobile App Troubleshooting Event App Log Collection Request ID
func (id UserMobileAppTroubleshootingEventAppLogCollectionRequestId) ID() string {
	fmtString := "/users/%s/mobileAppTroubleshootingEvents/%s/appLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MobileAppTroubleshootingEventId, id.AppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Mobile App Troubleshooting Event App Log Collection Request ID
func (id UserMobileAppTroubleshootingEventAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventIdValue"),
		resourceids.StaticSegment("staticAppLogCollectionRequests", "appLogCollectionRequests", "appLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("appLogCollectionRequestId", "appLogCollectionRequestIdValue"),
	}
}

// String returns a human-readable description of this User Mobile App Troubleshooting Event App Log Collection Request ID
func (id UserMobileAppTroubleshootingEventAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
		fmt.Sprintf("App Log Collection Request: %q", id.AppLogCollectionRequestId),
	}
	return fmt.Sprintf("User Mobile App Troubleshooting Event App Log Collection Request (%s)", strings.Join(components, "\n"))
}
