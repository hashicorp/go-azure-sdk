package grouppolicycategorychildren

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementGroupPolicyCategoryIdChildrenId{}

// DeviceManagementGroupPolicyCategoryIdChildrenId is a struct representing the Resource ID for a Device Management Group Policy Category Id Children
type DeviceManagementGroupPolicyCategoryIdChildrenId struct {
	GroupPolicyCategoryId  string
	GroupPolicyCategoryId1 string
}

// NewDeviceManagementGroupPolicyCategoryIdChildrenID returns a new DeviceManagementGroupPolicyCategoryIdChildrenId struct
func NewDeviceManagementGroupPolicyCategoryIdChildrenID(groupPolicyCategoryId string, groupPolicyCategoryId1 string) DeviceManagementGroupPolicyCategoryIdChildrenId {
	return DeviceManagementGroupPolicyCategoryIdChildrenId{
		GroupPolicyCategoryId:  groupPolicyCategoryId,
		GroupPolicyCategoryId1: groupPolicyCategoryId1,
	}
}

// ParseDeviceManagementGroupPolicyCategoryIdChildrenID parses 'input' into a DeviceManagementGroupPolicyCategoryIdChildrenId
func ParseDeviceManagementGroupPolicyCategoryIdChildrenID(input string) (*DeviceManagementGroupPolicyCategoryIdChildrenId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryIdChildrenId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryIdChildrenId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementGroupPolicyCategoryIdChildrenIDInsensitively parses 'input' case-insensitively into a DeviceManagementGroupPolicyCategoryIdChildrenId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementGroupPolicyCategoryIdChildrenIDInsensitively(input string) (*DeviceManagementGroupPolicyCategoryIdChildrenId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementGroupPolicyCategoryIdChildrenId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementGroupPolicyCategoryIdChildrenId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementGroupPolicyCategoryIdChildrenId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupPolicyCategoryId, ok = input.Parsed["groupPolicyCategoryId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyCategoryId", input)
	}

	if id.GroupPolicyCategoryId1, ok = input.Parsed["groupPolicyCategoryId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupPolicyCategoryId1", input)
	}

	return nil
}

// ValidateDeviceManagementGroupPolicyCategoryIdChildrenID checks that 'input' can be parsed as a Device Management Group Policy Category Id Children ID
func ValidateDeviceManagementGroupPolicyCategoryIdChildrenID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementGroupPolicyCategoryIdChildrenID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Group Policy Category Id Children ID
func (id DeviceManagementGroupPolicyCategoryIdChildrenId) ID() string {
	fmtString := "/deviceManagement/groupPolicyCategories/%s/children/%s"
	return fmt.Sprintf(fmtString, id.GroupPolicyCategoryId, id.GroupPolicyCategoryId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Group Policy Category Id Children ID
func (id DeviceManagementGroupPolicyCategoryIdChildrenId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("groupPolicyCategories", "groupPolicyCategories", "groupPolicyCategories"),
		resourceids.UserSpecifiedSegment("groupPolicyCategoryId", "groupPolicyCategoryIdValue"),
		resourceids.StaticSegment("children", "children", "children"),
		resourceids.UserSpecifiedSegment("groupPolicyCategoryId1", "groupPolicyCategoryId1Value"),
	}
}

// String returns a human-readable description of this Device Management Group Policy Category Id Children ID
func (id DeviceManagementGroupPolicyCategoryIdChildrenId) String() string {
	components := []string{
		fmt.Sprintf("Group Policy Category: %q", id.GroupPolicyCategoryId),
		fmt.Sprintf("Group Policy Category Id 1: %q", id.GroupPolicyCategoryId1),
	}
	return fmt.Sprintf("Device Management Group Policy Category Id Children (%s)", strings.Join(components, "\n"))
}
