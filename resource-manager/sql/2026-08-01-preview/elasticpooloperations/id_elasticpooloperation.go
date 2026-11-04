package elasticpooloperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ElasticPoolOperationId{})
}

var _ resourceids.ResourceId = &ElasticPoolOperationId{}

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
	parser := resourceids.NewParserFromResourceIdType(&ElasticPoolOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ElasticPoolOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseElasticPoolOperationIDInsensitively parses 'input' case-insensitively into a ElasticPoolOperationId
// note: this method should only be used for API response data and not user input
func ParseElasticPoolOperationIDInsensitively(input string) (*ElasticPoolOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ElasticPoolOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ElasticPoolOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ElasticPoolOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServerName, ok = input.Parsed["serverName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverName", input)
	}

	if id.ElasticPoolName, ok = input.Parsed["elasticPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "elasticPoolName", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("serverName", "serverName"),
		resourceids.StaticSegment("staticElasticPools", "elasticPools", "elasticPools"),
		resourceids.UserSpecifiedSegment("elasticPoolName", "elasticPoolName"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
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
