package elasticpooloperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ElasticPoolOperationId{}

// ElasticPoolOperationId is a struct representing the Resource ID for a Elastic Pool Operation
type ElasticPoolOperationId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	ElasticPoolName   string
	OperationId       string
}

// NewElasticPoolOperationID returns a new ElasticPoolOperationId struct
func NewElasticPoolOperationID(subscriptionId string, resourceGroupName string, serverName string, elasticPoolName string, operationId string) ElasticPoolOperationId {
	return ElasticPoolOperationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		ElasticPoolName:   elasticPoolName,
		OperationId:       operationId,
	}
}

// ParseElasticPoolOperationID parses 'input' into a ElasticPoolOperationId
func ParseElasticPoolOperationID(input string) (*ElasticPoolOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ElasticPoolOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ElasticPoolOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.ElasticPoolName, ok = parsed.Parsed["elasticPoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "elasticPoolName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseElasticPoolOperationIDInsensitively parses 'input' case-insensitively into a ElasticPoolOperationId
// note: this method should only be used for API response data and not user input
func ParseElasticPoolOperationIDInsensitively(input string) (*ElasticPoolOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ElasticPoolOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ElasticPoolOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.ElasticPoolName, ok = parsed.Parsed["elasticPoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "elasticPoolName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateElasticPoolOperationID checks that 'input' can be parsed as a Elastic Pool Operation ID
func ValidateElasticPoolOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseElasticPoolOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Elastic Pool Operation ID
func (id ElasticPoolOperationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/elasticPools/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.ElasticPoolName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Elastic Pool Operation ID
func (id ElasticPoolOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticElasticPools", "elasticPools", "elasticPools"),
		resourceids.UserSpecifiedSegment("elasticPoolName", "elasticPoolValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Elastic Pool Operation ID
func (id ElasticPoolOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Elastic Pool Name: %q", id.ElasticPoolName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Elastic Pool Operation (%s)", strings.Join(components, "\n"))
}
