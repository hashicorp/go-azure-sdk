package virtualendpointshareduseserviceplan

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementVirtualEndpointSharedUseServicePlanId{}

// DeviceManagementVirtualEndpointSharedUseServicePlanId is a struct representing the Resource ID for a Device Management Virtual Endpoint Shared Use Service Plan
type DeviceManagementVirtualEndpointSharedUseServicePlanId struct {
	CloudPCSharedUseServicePlanId string
}

// NewDeviceManagementVirtualEndpointSharedUseServicePlanID returns a new DeviceManagementVirtualEndpointSharedUseServicePlanId struct
func NewDeviceManagementVirtualEndpointSharedUseServicePlanID(cloudPCSharedUseServicePlanId string) DeviceManagementVirtualEndpointSharedUseServicePlanId {
	return DeviceManagementVirtualEndpointSharedUseServicePlanId{
		CloudPCSharedUseServicePlanId: cloudPCSharedUseServicePlanId,
	}
}

// ParseDeviceManagementVirtualEndpointSharedUseServicePlanID parses 'input' into a DeviceManagementVirtualEndpointSharedUseServicePlanId
func ParseDeviceManagementVirtualEndpointSharedUseServicePlanID(input string) (*DeviceManagementVirtualEndpointSharedUseServicePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointSharedUseServicePlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointSharedUseServicePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementVirtualEndpointSharedUseServicePlanIDInsensitively parses 'input' case-insensitively into a DeviceManagementVirtualEndpointSharedUseServicePlanId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementVirtualEndpointSharedUseServicePlanIDInsensitively(input string) (*DeviceManagementVirtualEndpointSharedUseServicePlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementVirtualEndpointSharedUseServicePlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementVirtualEndpointSharedUseServicePlanId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementVirtualEndpointSharedUseServicePlanId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCSharedUseServicePlanId, ok = input.Parsed["cloudPCSharedUseServicePlanId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCSharedUseServicePlanId", input)
	}

	return nil
}

// ValidateDeviceManagementVirtualEndpointSharedUseServicePlanID checks that 'input' can be parsed as a Device Management Virtual Endpoint Shared Use Service Plan ID
func ValidateDeviceManagementVirtualEndpointSharedUseServicePlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementVirtualEndpointSharedUseServicePlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Virtual Endpoint Shared Use Service Plan ID
func (id DeviceManagementVirtualEndpointSharedUseServicePlanId) ID() string {
	fmtString := "/deviceManagement/virtualEndpoint/sharedUseServicePlans/%s"
	return fmt.Sprintf(fmtString, id.CloudPCSharedUseServicePlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Virtual Endpoint Shared Use Service Plan ID
func (id DeviceManagementVirtualEndpointSharedUseServicePlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("virtualEndpoint", "virtualEndpoint", "virtualEndpoint"),
		resourceids.StaticSegment("sharedUseServicePlans", "sharedUseServicePlans", "sharedUseServicePlans"),
		resourceids.UserSpecifiedSegment("cloudPCSharedUseServicePlanId", "cloudPCSharedUseServicePlanIdValue"),
	}
}

// String returns a human-readable description of this Device Management Virtual Endpoint Shared Use Service Plan ID
func (id DeviceManagementVirtualEndpointSharedUseServicePlanId) String() string {
	components := []string{
		fmt.Sprintf("Cloud P C Shared Use Service Plan: %q", id.CloudPCSharedUseServicePlanId),
	}
	return fmt.Sprintf("Device Management Virtual Endpoint Shared Use Service Plan (%s)", strings.Join(components, "\n"))
}
