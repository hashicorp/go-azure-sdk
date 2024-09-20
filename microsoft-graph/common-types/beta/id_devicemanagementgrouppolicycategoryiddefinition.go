package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryIdDefinitionId{}

// DeviceManagementGroupPolicyCategoryIdDefinitionId is a struct representing the Resource ID for a Device Management Group Policy Category Id Definition
type DeviceManagementGroupPolicyCategoryIdDefinitionId struct {
	GroupPolicyCategoryId   string
	GroupPolicyDefinitionId string
}

// NewDeviceManagementGroupPolicyCategoryIdDefinitionID returns a new DeviceManagementGroupPolicyCategoryIdDefinitionId struct
func NewDeviceManagementGroupPolicyCategoryIdDefinitionID(groupPolicyCategoryId string, groupPolicyDefinitionId string) DeviceManagementGroupPolicyCategoryIdDefinitionId {
	return DeviceManagementGroupPolicyCategoryIdDefinitionId{
		GroupPolicyCategoryId:   groupPolicyCategoryId,
		GroupPolicyDefinitionId: groupPolicyDefinitionId,
	}
}

// ParseDeviceManagementGroupPolicyCategoryIdDefinitionID parses 'input' into a DeviceManagementGroupPolicyCategoryIdDefinitionId
func ParseDeviceManagementGroupPolicyCategoryIdDefinitionID(input string) (*DeviceManagementGroupPolicyCategoryIdDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryIdDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryIdDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyCategoryIdDefinitionIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyCategoryIdDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyCategoryIdDefinitionIDInsensitively(input string) (*DeviceManagementGroupPolicyCategoryIdDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryIdDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryIdDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyCategoryIdDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyCategoryId, ok = input.Parsed["groupPolicyCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyCategoryId", input)
	}

	if id.GroupPolicyDefinitionId, ok = input.Parsed["groupPolicyDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyCategoryIdDefinitionID checks that 'input' can be parsed as a Device Management Group Policy Category Id Definition ID
func ValidateDeviceManagementGroupPolicyCategoryIdDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyCategoryIdDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Category Id Definition ID
func (id DeviceManagementGroupPolicyCategoryIdDefinitionId) ID() string {
	fmtString := "/deviceManagement/groupPolicyCategories/%s/definitions/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyCategoryId, id.GroupPolicyDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Category Id Definition ID
func (id DeviceManagementGroupPolicyCategoryIdDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyCategories", "groupPolicyCategories", "groupPolicyCategories"),
		resourceids.UserSpecifiedSegment("groupPolicyCategoryId", "groupPolicyCategoryId"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("groupPolicyDefinitionId", "groupPolicyDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Category Id Definition ID
func (id DeviceManagementGroupPolicyCategoryIdDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Category: %q", id.GroupPolicyCategoryId),
		fmt.Sprintf("Group Policy Definition: %q", id.GroupPolicyDefinitionId),
	}
	return fmt.Sprintf("Device Management Group Policy Category Id Definition (%s)", strings.Join(components, "\n"))
}
