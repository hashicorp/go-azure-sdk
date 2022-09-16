package sqlpoolsworkloadgroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = WorkloadGroupId{}

// WorkloadGroupId is a struct representing the Resource ID for a Workload Group
type WorkloadGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SqlPoolName       string
	WorkloadGroupName string
}

// NewWorkloadGroupID returns a new WorkloadGroupId struct
func NewWorkloadGroupID(subscriptionId string, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string) WorkloadGroupId {
	return WorkloadGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SqlPoolName:       sqlPoolName,
		WorkloadGroupName: workloadGroupName,
	}
}

// ParseWorkloadGroupID parses 'input' into a WorkloadGroupId
func ParseWorkloadGroupID(input string) (*WorkloadGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkloadGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkloadGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.SqlPoolName, ok = parsed.Parsed["sqlPoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'sqlPoolName' was not found in the resource id %q", input)
	}

	if id.WorkloadGroupName, ok = parsed.Parsed["workloadGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'workloadGroupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseWorkloadGroupIDInsensitively parses 'input' case-insensitively into a WorkloadGroupId
// note: this method should only be used for API response data and not user input
func ParseWorkloadGroupIDInsensitively(input string) (*WorkloadGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkloadGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkloadGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.SqlPoolName, ok = parsed.Parsed["sqlPoolName"]; !ok {
		return nil, fmt.Errorf("the segment 'sqlPoolName' was not found in the resource id %q", input)
	}

	if id.WorkloadGroupName, ok = parsed.Parsed["workloadGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'workloadGroupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateWorkloadGroupID checks that 'input' can be parsed as a Workload Group ID
func ValidateWorkloadGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkloadGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workload Group ID
func (id WorkloadGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/sqlPools/%s/workloadGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SqlPoolName, id.WorkloadGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workload Group ID
func (id WorkloadGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticSqlPools", "sqlPools", "sqlPools"),
		resourceids.UserSpecifiedSegment("sqlPoolName", "sqlPoolValue"),
		resourceids.StaticSegment("staticWorkloadGroups", "workloadGroups", "workloadGroups"),
		resourceids.UserSpecifiedSegment("workloadGroupName", "workloadGroupValue"),
	}
}

// String returns a human-readable description of this Workload Group ID
func (id WorkloadGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
		fmt.Sprintf("Workload Group Name: %q", id.WorkloadGroupName),
	}
	return fmt.Sprintf("Workload Group (%s)", strings.Join(components, "\n"))
}
