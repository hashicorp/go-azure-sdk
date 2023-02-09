package actions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = WorkflowId{}

// WorkflowId is a struct representing the Resource ID for a Workflow
type WorkflowId struct {
	SubscriptionId         string
	ResourceGroupName      string
	StorageSyncServiceName string
	WorkflowId             string
}

// NewWorkflowID returns a new WorkflowId struct
func NewWorkflowID(subscriptionId string, resourceGroupName string, storageSyncServiceName string, workflowId string) WorkflowId {
	return WorkflowId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		StorageSyncServiceName: storageSyncServiceName,
		WorkflowId:             workflowId,
	}
}

// ParseWorkflowID parses 'input' into a WorkflowId
func ParseWorkflowID(input string) (*WorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkflowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkflowId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageSyncServiceName' was not found in the resource id %q", input)
	}

	if id.WorkflowId, ok = parsed.Parsed["workflowId"]; !ok {
		return nil, fmt.Errorf("the segment 'workflowId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseWorkflowIDInsensitively parses 'input' case-insensitively into a WorkflowId
// note: this method should only be used for API response data and not user input
func ParseWorkflowIDInsensitively(input string) (*WorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkflowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkflowId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageSyncServiceName' was not found in the resource id %q", input)
	}

	if id.WorkflowId, ok = parsed.Parsed["workflowId"]; !ok {
		return nil, fmt.Errorf("the segment 'workflowId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateWorkflowID checks that 'input' can be parsed as a Workflow ID
func ValidateWorkflowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkflowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workflow ID
func (id WorkflowId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StorageSync/storageSyncServices/%s/workflows/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StorageSyncServiceName, id.WorkflowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Workflow ID
func (id WorkflowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorageSync", "Microsoft.StorageSync", "Microsoft.StorageSync"),
		resourceids.StaticSegment("staticStorageSyncServices", "storageSyncServices", "storageSyncServices"),
		resourceids.UserSpecifiedSegment("storageSyncServiceName", "storageSyncServiceValue"),
		resourceids.StaticSegment("staticWorkflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowId", "workflowIdValue"),
	}
}

// String returns a human-readable description of this Workflow ID
func (id WorkflowId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Storage Sync Service Name: %q", id.StorageSyncServiceName),
		fmt.Sprintf("Workflow: %q", id.WorkflowId),
	}
	return fmt.Sprintf("Workflow (%s)", strings.Join(components, "\n"))
}
