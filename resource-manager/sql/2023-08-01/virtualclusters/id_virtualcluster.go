package virtualclusters

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VirtualClusterId{})
}

var _ resourceids.ResourceId = &VirtualClusterId{}

// VirtualClusterId is a struct representing the Resource ID for a Virtual Cluster
type VirtualClusterId struct {
	SubscriptionId     string
	ResourceGroupName  string
	VirtualClusterName string
}

// NewVirtualClusterID returns a new VirtualClusterId struct
func NewVirtualClusterID(subscriptionId string, resourceGroupName string, virtualClusterName string) VirtualClusterId {
	return VirtualClusterId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		VirtualClusterName: virtualClusterName,
	}
}

// ParseVirtualClusterID parses 'input' into a VirtualClusterId
func ParseVirtualClusterID(input string) (*VirtualClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualClusterId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualClusterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVirtualClusterIDInsensitively parses 'input' case-insensitively into a VirtualClusterId
// note: this method should only be used for API response data and not user input
func ParseVirtualClusterIDInsensitively(input string) (*VirtualClusterId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualClusterId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualClusterId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VirtualClusterId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VirtualClusterName, ok = input.Parsed["virtualClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualClusterName", input)
	}

	return nil
}

// ValidateVirtualClusterID checks that 'input' can be parsed as a Virtual Cluster ID
func ValidateVirtualClusterID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualClusterID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Cluster ID
func (id VirtualClusterId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/virtualClusters/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualClusterName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Cluster ID
func (id VirtualClusterId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticVirtualClusters", "virtualClusters", "virtualClusters"),
		resourceids.UserSpecifiedSegment("virtualClusterName", "virtualClusterName"),
	}
}

// String returns a human-readable description of this Virtual Cluster ID
func (id VirtualClusterId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Cluster Name: %q", id.VirtualClusterName),
	}
	return fmt.Sprintf("Virtual Cluster (%s)", strings.Join(components, "\n"))
}
