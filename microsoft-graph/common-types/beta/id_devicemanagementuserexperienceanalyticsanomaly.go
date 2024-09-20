package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementUserExperienceAnalyticsAnomalyId{}

// DeviceManagementUserExperienceAnalyticsAnomalyId is a struct representing the Resource ID for a Device Management User Experience Analytics Anomaly
type DeviceManagementUserExperienceAnalyticsAnomalyId struct {
	UserExperienceAnalyticsAnomalyId string
}

// NewDeviceManagementUserExperienceAnalyticsAnomalyID returns a new DeviceManagementUserExperienceAnalyticsAnomalyId struct
func NewDeviceManagementUserExperienceAnalyticsAnomalyID(userExperienceAnalyticsAnomalyId string) DeviceManagementUserExperienceAnalyticsAnomalyId {
	return DeviceManagementUserExperienceAnalyticsAnomalyId{
		UserExperienceAnalyticsAnomalyId: userExperienceAnalyticsAnomalyId,
	}
}

// ParseDeviceManagementUserExperienceAnalyticsAnomalyID parses 'input' into a DeviceManagementUserExperienceAnalyticsAnomalyId
func ParseDeviceManagementUserExperienceAnalyticsAnomalyID(input string) (*DeviceManagementUserExperienceAnalyticsAnomalyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAnomalyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAnomalyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementUserExperienceAnalyticsAnomalyIDInsensitively parses 'input' case-insensitively into a DeviceManagementUserExperienceAnalyticsAnomalyId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementUserExperienceAnalyticsAnomalyIDInsensitively(input string) (*DeviceManagementUserExperienceAnalyticsAnomalyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementUserExperienceAnalyticsAnomalyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementUserExperienceAnalyticsAnomalyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementUserExperienceAnalyticsAnomalyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserExperienceAnalyticsAnomalyId, ok = input.Parsed["userExperienceAnalyticsAnomalyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userExperienceAnalyticsAnomalyId", input)
	}

	return nil
}

// ValidateDeviceManagementUserExperienceAnalyticsAnomalyID checks that 'input' can be parsed as a Device Management User Experience Analytics Anomaly ID
func ValidateDeviceManagementUserExperienceAnalyticsAnomalyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementUserExperienceAnalyticsAnomalyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management User Experience Analytics Anomaly ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyId) ID() string {
	fmtString := "/deviceManagement/userExperienceAnalyticsAnomaly/%s"
	return fmt.Sprintf(fmtString, id.UserExperienceAnalyticsAnomalyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management User Experience Analytics Anomaly ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("userExperienceAnalyticsAnomaly", "userExperienceAnalyticsAnomaly", "userExperienceAnalyticsAnomaly"),
		resourceids.UserSpecifiedSegment("userExperienceAnalyticsAnomalyId", "userExperienceAnalyticsAnomalyId"),
	}
}

// String returns a human-readable description of this Device Management User Experience Analytics Anomaly ID
func (id DeviceManagementUserExperienceAnalyticsAnomalyId) String() string {
	components := []string{
		fmt.Sprintf("User Experience Analytics Anomaly: %q", id.UserExperienceAnalyticsAnomalyId),
	}
	return fmt.Sprintf("Device Management User Experience Analytics Anomaly (%s)", strings.Join(components, "\n"))
}
