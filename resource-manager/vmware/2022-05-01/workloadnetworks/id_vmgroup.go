package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VMGroupId{}

// VMGroupId is a struct representing the Resource ID for a V M Group
type VMGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	VmGroupId         string
}

// NewVMGroupID returns a new VMGroupId struct
func NewVMGroupID(subscriptionId string, resourceGroupName string, privateCloudName string, vmGroupId string) VMGroupId {
	return VMGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		VmGroupId:         vmGroupId,
	}
}

// ParseVMGroupID parses 'input' into a VMGroupId
func ParseVMGroupID(input string) (*VMGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(VMGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VMGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.VmGroupId, ok = parsed.Parsed["vmGroupId"]; !ok {
		return nil, fmt.Errorf("the segment 'vmGroupId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVMGroupIDInsensitively parses 'input' case-insensitively into a VMGroupId
// note: this method should only be used for API response data and not user input
func ParseVMGroupIDInsensitively(input string) (*VMGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(VMGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VMGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.VmGroupId, ok = parsed.Parsed["vmGroupId"]; !ok {
		return nil, fmt.Errorf("the segment 'vmGroupId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVMGroupID checks that 'input' can be parsed as a V M Group ID
func ValidateVMGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVMGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted V M Group ID
func (id VMGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/vmGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.VmGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this V M Group ID
func (id VMGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticWorkloadNetworks", "workloadNetworks", "workloadNetworks"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticVmGroups", "vmGroups", "vmGroups"),
		resourceids.UserSpecifiedSegment("vmGroupId", "vmGroupIdValue"),
	}
}

// String returns a human-readable description of this V M Group ID
func (id VMGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Vm Group: %q", id.VmGroupId),
	}
	return fmt.Sprintf("V M Group (%s)", strings.Join(components, "\n"))
}
