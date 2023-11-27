package roles

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleId{}

// RoleId is a struct representing the Resource ID for a Role
type RoleId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DataBoxEdgeDeviceName string
	RoleName              string
}

// NewRoleID returns a new RoleId struct
func NewRoleID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, roleName string) RoleId {
	return RoleId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DataBoxEdgeDeviceName: dataBoxEdgeDeviceName,
		RoleName:              roleName,
	}
}

// ParseRoleID parses 'input' into a RoleId
func ParseRoleID(input string) (*RoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleIDInsensitively parses 'input' case-insensitively into a RoleId
// note: this method should only be used for API response data and not user input
func ParseRoleIDInsensitively(input string) (*RoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DataBoxEdgeDeviceName, ok = input.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataBoxEdgeDeviceName", input)
	}

	if id.RoleName, ok = input.Parsed["roleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleName", input)
	}

	return nil
}

// ValidateRoleID checks that 'input' can be parsed as a Role ID
func ValidateRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role ID
func (id RoleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/roles/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.RoleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Role ID
func (id RoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticRoles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("roleName", "roleValue"),
	}
}

// String returns a human-readable description of this Role ID
func (id RoleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Role Name: %q", id.RoleName),
	}
	return fmt.Sprintf("Role (%s)", strings.Join(components, "\n"))
}
