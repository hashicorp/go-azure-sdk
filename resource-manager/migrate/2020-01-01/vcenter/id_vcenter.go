package vcenter

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VCenterId{}

// VCenterId is a struct representing the Resource ID for a V Center
type VCenterId struct {
	SubscriptionId    string
	ResourceGroupName string
	VmwareSiteName    string
	VCenterName       string
}

// NewVCenterID returns a new VCenterId struct
func NewVCenterID(subscriptionId string, resourceGroupName string, vmwareSiteName string, vCenterName string) VCenterId {
	return VCenterId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VmwareSiteName:    vmwareSiteName,
		VCenterName:       vCenterName,
	}
}

// ParseVCenterID parses 'input' into a VCenterId
func ParseVCenterID(input string) (*VCenterId, error) {
	parser := resourceids.NewParserFromResourceIdType(VCenterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VCenterId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVCenterIDInsensitively parses 'input' case-insensitively into a VCenterId
// note: this method should only be used for API response data and not user input
func ParseVCenterIDInsensitively(input string) (*VCenterId, error) {
	parser := resourceids.NewParserFromResourceIdType(VCenterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VCenterId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VCenterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VmwareSiteName, ok = input.Parsed["vmwareSiteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "vmwareSiteName", input)
	}

	if id.VCenterName, ok = input.Parsed["vCenterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "vCenterName", input)
	}

	return nil
}

// ValidateVCenterID checks that 'input' can be parsed as a V Center ID
func ValidateVCenterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVCenterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted V Center ID
func (id VCenterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OffAzure/vmwareSites/%s/vCenters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VmwareSiteName, id.VCenterName)
}

// Segments returns a slice of Resource ID Segments which comprise this V Center ID
func (id VCenterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOffAzure", "Microsoft.OffAzure", "Microsoft.OffAzure"),
		resourceids.StaticSegment("staticVmwareSites", "vmwareSites", "vmwareSites"),
		resourceids.UserSpecifiedSegment("vmwareSiteName", "vmwareSiteValue"),
		resourceids.StaticSegment("staticVCenters", "vCenters", "vCenters"),
		resourceids.UserSpecifiedSegment("vCenterName", "vCenterValue"),
	}
}

// String returns a human-readable description of this V Center ID
func (id VCenterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vmware Site Name: %q", id.VmwareSiteName),
		fmt.Sprintf("V Center Name: %q", id.VCenterName),
	}
	return fmt.Sprintf("V Center (%s)", strings.Join(components, "\n"))
}
