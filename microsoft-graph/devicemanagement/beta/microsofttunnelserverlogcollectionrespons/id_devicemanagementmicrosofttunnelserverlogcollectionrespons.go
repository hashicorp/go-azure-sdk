package microsofttunnelserverlogcollectionrespons

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{}

// DeviceManagementMicrosoftTunnelServerLogCollectionResponsId is a struct representing the Resource ID for a Device Management Microsoft Tunnel Server Log Collection Respons
type DeviceManagementMicrosoftTunnelServerLogCollectionResponsId struct {
	MicrosoftTunnelServerLogCollectionResponseId string
}

// NewDeviceManagementMicrosoftTunnelServerLogCollectionResponsID returns a new DeviceManagementMicrosoftTunnelServerLogCollectionResponsId struct
func NewDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(microsoftTunnelServerLogCollectionResponseId string) DeviceManagementMicrosoftTunnelServerLogCollectionResponsId {
	return DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{
		MicrosoftTunnelServerLogCollectionResponseId: microsoftTunnelServerLogCollectionResponseId,
	}
}

// ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsID parses 'input' into a DeviceManagementMicrosoftTunnelServerLogCollectionResponsId
func ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(input string) (*DeviceManagementMicrosoftTunnelServerLogCollectionResponsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsIDInsensitively parses 'input' case-insensitively into a DeviceManagementMicrosoftTunnelServerLogCollectionResponsId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsIDInsensitively(input string) (*DeviceManagementMicrosoftTunnelServerLogCollectionResponsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelServerLogCollectionResponsId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMicrosoftTunnelServerLogCollectionResponsId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftTunnelServerLogCollectionResponseId, ok = input.Parsed["microsoftTunnelServerLogCollectionResponseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelServerLogCollectionResponseId", input)
	}

	return nil
}

// ValidateDeviceManagementMicrosoftTunnelServerLogCollectionResponsID checks that 'input' can be parsed as a Device Management Microsoft Tunnel Server Log Collection Respons ID
func ValidateDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMicrosoftTunnelServerLogCollectionResponsID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Microsoft Tunnel Server Log Collection Respons ID
func (id DeviceManagementMicrosoftTunnelServerLogCollectionResponsId) ID() string {
	fmtString := "/deviceManagement/microsoftTunnelServerLogCollectionResponses/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftTunnelServerLogCollectionResponseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Microsoft Tunnel Server Log Collection Respons ID
func (id DeviceManagementMicrosoftTunnelServerLogCollectionResponsId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("microsoftTunnelServerLogCollectionResponses", "microsoftTunnelServerLogCollectionResponses", "microsoftTunnelServerLogCollectionResponses"),
		resourceids.UserSpecifiedSegment("microsoftTunnelServerLogCollectionResponseId", "microsoftTunnelServerLogCollectionResponseIdValue"),
	}
}

// String returns a human-readable description of this Device Management Microsoft Tunnel Server Log Collection Respons ID
func (id DeviceManagementMicrosoftTunnelServerLogCollectionResponsId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Tunnel Server Log Collection Response: %q", id.MicrosoftTunnelServerLogCollectionResponseId),
	}
	return fmt.Sprintf("Device Management Microsoft Tunnel Server Log Collection Respons (%s)", strings.Join(components, "\n"))
}
