package memobileapptroubleshootingeventapplogcollectionrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMobileAppTroubleshootingEventAppLogCollectionRequestId{}

// MeMobileAppTroubleshootingEventAppLogCollectionRequestId is a struct representing the Resource ID for a Me Mobile App Troubleshooting Event App Log Collection Request
type MeMobileAppTroubleshootingEventAppLogCollectionRequestId struct {
	MobileAppTroubleshootingEventId string
	AppLogCollectionRequestId       string
}

// NewMeMobileAppTroubleshootingEventAppLogCollectionRequestID returns a new MeMobileAppTroubleshootingEventAppLogCollectionRequestId struct
func NewMeMobileAppTroubleshootingEventAppLogCollectionRequestID(mobileAppTroubleshootingEventId string, appLogCollectionRequestId string) MeMobileAppTroubleshootingEventAppLogCollectionRequestId {
	return MeMobileAppTroubleshootingEventAppLogCollectionRequestId{
		MobileAppTroubleshootingEventId: mobileAppTroubleshootingEventId,
		AppLogCollectionRequestId:       appLogCollectionRequestId,
	}
}

// ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestID parses 'input' into a MeMobileAppTroubleshootingEventAppLogCollectionRequestId
func ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestID(input string) (*MeMobileAppTroubleshootingEventAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMobileAppTroubleshootingEventAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMobileAppTroubleshootingEventAppLogCollectionRequestId{}

	if id.MobileAppTroubleshootingEventId, ok = parsed.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", *parsed)
	}

	if id.AppLogCollectionRequestId, ok = parsed.Parsed["appLogCollectionRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", *parsed)
	}

	return &id, nil
}

// ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a MeMobileAppTroubleshootingEventAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestIDInsensitively(input string) (*MeMobileAppTroubleshootingEventAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMobileAppTroubleshootingEventAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMobileAppTroubleshootingEventAppLogCollectionRequestId{}

	if id.MobileAppTroubleshootingEventId, ok = parsed.Parsed["mobileAppTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobileAppTroubleshootingEventId", *parsed)
	}

	if id.AppLogCollectionRequestId, ok = parsed.Parsed["appLogCollectionRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appLogCollectionRequestId", *parsed)
	}

	return &id, nil
}

// ValidateMeMobileAppTroubleshootingEventAppLogCollectionRequestID checks that 'input' can be parsed as a Me Mobile App Troubleshooting Event App Log Collection Request ID
func ValidateMeMobileAppTroubleshootingEventAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMobileAppTroubleshootingEventAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Mobile App Troubleshooting Event App Log Collection Request ID
func (id MeMobileAppTroubleshootingEventAppLogCollectionRequestId) ID() string {
	fmtString := "/me/mobileAppTroubleshootingEvents/%s/appLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.MobileAppTroubleshootingEventId, id.AppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Mobile App Troubleshooting Event App Log Collection Request ID
func (id MeMobileAppTroubleshootingEventAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents", "mobileAppTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("mobileAppTroubleshootingEventId", "mobileAppTroubleshootingEventIdValue"),
		resourceids.StaticSegment("staticAppLogCollectionRequests", "appLogCollectionRequests", "appLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("appLogCollectionRequestId", "appLogCollectionRequestIdValue"),
	}
}

// String returns a human-readable description of this Me Mobile App Troubleshooting Event App Log Collection Request ID
func (id MeMobileAppTroubleshootingEventAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Mobile App Troubleshooting Event: %q", id.MobileAppTroubleshootingEventId),
		fmt.Sprintf("App Log Collection Request: %q", id.AppLogCollectionRequestId),
	}
	return fmt.Sprintf("Me Mobile App Troubleshooting Event App Log Collection Request (%s)", strings.Join(components, "\n"))
}
