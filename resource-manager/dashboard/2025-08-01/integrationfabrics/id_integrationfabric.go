package integrationfabrics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&IntegrationFabricId{})
}

var _ resourceids.ResourceId = &IntegrationFabricId{}

// IntegrationFabricId is a struct representing the Resource ID for a Integration Fabric
type IntegrationFabricId struct {
	SubscriptionId        string
	ResourceGroupName     string
	GrafanaName           string
	IntegrationFabricName string
}

// NewIntegrationFabricID returns a new IntegrationFabricId struct
func NewIntegrationFabricID(subscriptionId string, resourceGroupName string, grafanaName string, integrationFabricName string) IntegrationFabricId {
	return IntegrationFabricId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		GrafanaName:           grafanaName,
		IntegrationFabricName: integrationFabricName,
	}
}

// ParseIntegrationFabricID parses 'input' into a IntegrationFabricId
func ParseIntegrationFabricID(input string) (*IntegrationFabricId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IntegrationFabricId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IntegrationFabricId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIntegrationFabricIDInsensitively parses 'input' case-insensitively into a IntegrationFabricId
// note: this method should only be used for API response data and not user input
func ParseIntegrationFabricIDInsensitively(input string) (*IntegrationFabricId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IntegrationFabricId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IntegrationFabricId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IntegrationFabricId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.GrafanaName, ok = input.Parsed["grafanaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "grafanaName", input)
	}

	if id.IntegrationFabricName, ok = input.Parsed["integrationFabricName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "integrationFabricName", input)
	}

	return nil
}

// ValidateIntegrationFabricID checks that 'input' can be parsed as a Integration Fabric ID
func ValidateIntegrationFabricID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIntegrationFabricID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Integration Fabric ID
func (id IntegrationFabricId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Dashboard/grafana/%s/integrationFabrics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.GrafanaName, id.IntegrationFabricName)
}

// Segments returns a slice of Resource ID Segments which comprise this Integration Fabric ID
func (id IntegrationFabricId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDashboard", "Microsoft.Dashboard", "Microsoft.Dashboard"),
		resourceids.StaticSegment("staticGrafana", "grafana", "grafana"),
		resourceids.UserSpecifiedSegment("grafanaName", "grafanaName"),
		resourceids.StaticSegment("staticIntegrationFabrics", "integrationFabrics", "integrationFabrics"),
		resourceids.UserSpecifiedSegment("integrationFabricName", "integrationFabricName"),
	}
}

// String returns a human-readable description of this Integration Fabric ID
func (id IntegrationFabricId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Grafana Name: %q", id.GrafanaName),
		fmt.Sprintf("Integration Fabric Name: %q", id.IntegrationFabricName),
	}
	return fmt.Sprintf("Integration Fabric (%s)", strings.Join(components, "\n"))
}
