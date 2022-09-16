package sqlpoolsreplicationlinks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationLinkId{}

// ReplicationLinkId is a struct representing the Resource ID for a Replication Link
type ReplicationLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SqlPoolName       string
	LinkId            string
}

// NewReplicationLinkID returns a new ReplicationLinkId struct
func NewReplicationLinkID(subscriptionId string, resourceGroupName string, workspaceName string, sqlPoolName string, linkId string) ReplicationLinkId {
	return ReplicationLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SqlPoolName:       sqlPoolName,
		LinkId:            linkId,
	}
}

// ParseReplicationLinkID parses 'input' into a ReplicationLinkId
func ParseReplicationLinkID(input string) (*ReplicationLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationLinkId{}

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

	if id.LinkId, ok = parsed.Parsed["linkId"]; !ok {
		return nil, fmt.Errorf("the segment 'linkId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationLinkIDInsensitively parses 'input' case-insensitively into a ReplicationLinkId
// note: this method should only be used for API response data and not user input
func ParseReplicationLinkIDInsensitively(input string) (*ReplicationLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationLinkId{}

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

	if id.LinkId, ok = parsed.Parsed["linkId"]; !ok {
		return nil, fmt.Errorf("the segment 'linkId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationLinkID checks that 'input' can be parsed as a Replication Link ID
func ValidateReplicationLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Link ID
func (id ReplicationLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/sqlPools/%s/replicationLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SqlPoolName, id.LinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Link ID
func (id ReplicationLinkId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticReplicationLinks", "replicationLinks", "replicationLinks"),
		resourceids.UserSpecifiedSegment("linkId", "linkIdValue"),
	}
}

// String returns a human-readable description of this Replication Link ID
func (id ReplicationLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
		fmt.Sprintf("Link: %q", id.LinkId),
	}
	return fmt.Sprintf("Replication Link (%s)", strings.Join(components, "\n"))
}
