package sqlpoolsusages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SqlPoolId{}

// SqlPoolId is a struct representing the Resource ID for a Sql Pool
type SqlPoolId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SqlPoolName       string
}

// NewSqlPoolID returns a new SqlPoolId struct
func NewSqlPoolID(subscriptionId string, resourceGroupName string, workspaceName string, sqlPoolName string) SqlPoolId {
	return SqlPoolId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SqlPoolName:       sqlPoolName,
	}
}

// ParseSqlPoolID parses 'input' into a SqlPoolId
func ParseSqlPoolID(input string) (*SqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(SqlPoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SqlPoolId{}

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

// ParseSqlPoolIDInsensitively parses 'input' case-insensitively into a SqlPoolId
// note: this method should only be used for API response data and not user input
func ParseSqlPoolIDInsensitively(input string) (*SqlPoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(SqlPoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SqlPoolId{}

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

// ValidateSqlPoolID checks that 'input' can be parsed as a Sql Pool ID
func ValidateSqlPoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlPoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Pool ID
func (id SqlPoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/sqlPools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SqlPoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Pool ID
func (id SqlPoolId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Sql Pool ID
func (id SqlPoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
	}
	return fmt.Sprintf("Sql Pool (%s)", strings.Join(components, "\n"))
}
