package virtualmachineschedules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineScheduleId{}

// VirtualMachineScheduleId is a struct representing the Resource ID for a Virtual Machine Schedule
type VirtualMachineScheduleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	LabName            string
	VirtualMachineName string
	Name               string
}

// NewVirtualMachineScheduleID returns a new VirtualMachineScheduleId struct
func NewVirtualMachineScheduleID(subscriptionId string, resourceGroupName string, labName string, virtualMachineName string, name string) VirtualMachineScheduleId {
	return VirtualMachineScheduleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		LabName:            labName,
		VirtualMachineName: virtualMachineName,
		Name:               name,
	}
}

// ParseVirtualMachineScheduleID parses 'input' into a VirtualMachineScheduleId
func ParseVirtualMachineScheduleID(input string) (*VirtualMachineScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScheduleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineName' was not found in the resource id %q", input)
	}

	if id.Name, ok = parsed.Parsed["name"]; !ok {
		return nil, fmt.Errorf("the segment 'name' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVirtualMachineScheduleIDInsensitively parses 'input' case-insensitively into a VirtualMachineScheduleId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScheduleIDInsensitively(input string) (*VirtualMachineScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScheduleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineName' was not found in the resource id %q", input)
	}

	if id.Name, ok = parsed.Parsed["name"]; !ok {
		return nil, fmt.Errorf("the segment 'name' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVirtualMachineScheduleID checks that 'input' can be parsed as a Virtual Machine Schedule ID
func ValidateVirtualMachineScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Schedule ID
func (id VirtualMachineScheduleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/virtualMachines/%s/schedules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.VirtualMachineName, id.Name)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Schedule ID
func (id VirtualMachineScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticSchedules", "schedules", "schedules"),
		resourceids.UserSpecifiedSegment("name", "nameValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Schedule ID
func (id VirtualMachineScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Name: %q", id.Name),
	}
	return fmt.Sprintf("Virtual Machine Schedule (%s)", strings.Join(components, "\n"))
}
