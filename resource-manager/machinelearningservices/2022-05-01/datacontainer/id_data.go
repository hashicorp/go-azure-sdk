package datacontainer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DataId{}

// DataId is a struct representing the Resource ID for a Data
type DataId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	Name              string
}

// NewDataID returns a new DataId struct
func NewDataID(subscriptionId string, resourceGroupName string, workspaceName string, name string) DataId {
	return DataId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		Name:              name,
	}
}

// ParseDataID parses 'input' into a DataId
func ParseDataID(input string) (*DataId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.Name, ok = parsed.Parsed["name"]; !ok {
		return nil, fmt.Errorf("the segment 'name' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDataIDInsensitively parses 'input' case-insensitively into a DataId
// note: this method should only be used for API response data and not user input
func ParseDataIDInsensitively(input string) (*DataId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.Name, ok = parsed.Parsed["name"]; !ok {
		return nil, fmt.Errorf("the segment 'name' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDataID checks that 'input' can be parsed as a Data ID
func ValidateDataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data ID
func (id DataId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/data/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.Name)
}

// Segments returns a slice of Resource ID Segments which comprise this Data ID
func (id DataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticData", "data", "data"),
		resourceids.UserSpecifiedSegment("name", "nameValue"),
	}
}

// String returns a human-readable description of this Data ID
func (id DataId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Name: %q", id.Name),
	}
	return fmt.Sprintf("Data (%s)", strings.Join(components, "\n"))
}
