package recoverablesqlpools

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RecoverableSqlPoolId{}

// RecoverableSqlPoolId is a struct representing the Resource ID for a Recoverable Sql Pool
type RecoverableSqlPoolId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SqlPoolName       string
}

// NewRecoverableSqlPoolID returns a new RecoverableSqlPoolId struct
func NewRecoverableSqlPoolID(subscriptionId string, resourceGroupName string, workspaceName string, sqlPoolName string) RecoverableSqlPoolId {
	return RecoverableSqlPoolId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SqlPoolName:       sqlPoolName,
	}
}

// ParseRecoverableSqlPoolID parses 'input' into a RecoverableSqlPoolId
func ParseRecoverableSqlPoolID(input string) (*RecoverableSqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoverableSqlPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoverableSqlPoolId{}

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

	return &id, nil
}

// ParseRecoverableSqlPoolIDInsensitively parses 'input' case-insensitively into a RecoverableSqlPoolId
// note: this method should only be used for API response data and not user input
func ParseRecoverableSqlPoolIDInsensitively(input string) (*RecoverableSqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoverableSqlPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoverableSqlPoolId{}

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

	return &id, nil
}

// ValidateRecoverableSqlPoolID checks that 'input' can be parsed as a Recoverable Sql Pool ID
func ValidateRecoverableSqlPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecoverableSqlPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recoverable Sql Pool ID
func (id RecoverableSqlPoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/recoverableSqlPools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SqlPoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Recoverable Sql Pool ID
func (id RecoverableSqlPoolId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticRecoverableSqlPools", "recoverableSqlPools", "recoverableSqlPools"),
		resourceids.UserSpecifiedSegment("sqlPoolName", "sqlPoolValue"),
	}
}

// String returns a human-readable description of this Recoverable Sql Pool ID
func (id RecoverableSqlPoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
	}
	return fmt.Sprintf("Recoverable Sql Pool (%s)", strings.Join(components, "\n"))
}
