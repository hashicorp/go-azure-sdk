package datastore

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DataStoreId{}

// DataStoreId is a struct representing the Resource ID for a Data Store
type DataStoreId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	Name              string
}

// NewDataStoreID returns a new DataStoreId struct
func NewDataStoreID(subscriptionId string, resourceGroupName string, workspaceName string, name string) DataStoreId {
	return DataStoreId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		Name:              name,
	}
}

// ParseDataStoreID parses 'input' into a DataStoreId
func ParseDataStoreID(input string) (*DataStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataStoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataStoreId{}

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

// ParseDataStoreIDInsensitively parses 'input' case-insensitively into a DataStoreId
// note: this method should only be used for API response data and not user input
func ParseDataStoreIDInsensitively(input string) (*DataStoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataStoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataStoreId{}

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

// ValidateDataStoreID checks that 'input' can be parsed as a Data Store ID
func ValidateDataStoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataStoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Data Store ID
func (id DataStoreId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/dataStores/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.Name)
}

// Segments returns a slice of Resource ID Segments which comprise this Data Store ID
func (id DataStoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticDataStores", "dataStores", "dataStores"),
		resourceids.UserSpecifiedSegment("name", "nameValue"),
	}
}

// String returns a human-readable description of this Data Store ID
func (id DataStoreId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Name: %q", id.Name),
	}
	return fmt.Sprintf("Data Store (%s)", strings.Join(components, "\n"))
}
