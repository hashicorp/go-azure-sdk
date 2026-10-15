package managednamespaces

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ManagedClusterManagedNamespaceId{})
}

var _ resourceids.ResourceId = &ManagedClusterManagedNamespaceId{}

// ManagedClusterManagedNamespaceId is a struct representing the Resource ID for a Managed Cluster Managed Namespace
type ManagedClusterManagedNamespaceId struct {
	SubscriptionId       string
	ResourceGroupName    string
	ManagedClusterName   string
	ManagedNamespaceName string
}

// NewManagedClusterManagedNamespaceID returns a new ManagedClusterManagedNamespaceId struct
func NewManagedClusterManagedNamespaceID(subscriptionId string, resourceGroupName string, managedClusterName string, managedNamespaceName string) ManagedClusterManagedNamespaceId {
	return ManagedClusterManagedNamespaceId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		ManagedClusterName:   managedClusterName,
		ManagedNamespaceName: managedNamespaceName,
	}
}

// ParseManagedClusterManagedNamespaceID parses 'input' into a ManagedClusterManagedNamespaceId
func ParseManagedClusterManagedNamespaceID(input string) (*ManagedClusterManagedNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedClusterManagedNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedClusterManagedNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedClusterManagedNamespaceIDInsensitively parses 'input' case-insensitively into a ManagedClusterManagedNamespaceId
// note: this method should only be used for API response data and not user input
func ParseManagedClusterManagedNamespaceIDInsensitively(input string) (*ManagedClusterManagedNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedClusterManagedNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedClusterManagedNamespaceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedClusterManagedNamespaceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedClusterName, ok = input.Parsed["managedClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedClusterName", input)
	}

	if id.ManagedNamespaceName, ok = input.Parsed["managedNamespaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedNamespaceName", input)
	}

	return nil
}

// ValidateManagedClusterManagedNamespaceID checks that 'input' can be parsed as a Managed Cluster Managed Namespace ID
func ValidateManagedClusterManagedNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedClusterManagedNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Cluster Managed Namespace ID
func (id ManagedClusterManagedNamespaceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerService/managedClusters/%s/managedNamespaces/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedClusterName, id.ManagedNamespaceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Cluster Managed Namespace ID
func (id ManagedClusterManagedNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftContainerService", "Microsoft.ContainerService", "Microsoft.ContainerService"),
		resourceids.StaticSegment("staticManagedClusters", "managedClusters", "managedClusters"),
		resourceids.UserSpecifiedSegment("managedClusterName", "managedClusterName"),
		resourceids.StaticSegment("staticManagedNamespaces", "managedNamespaces", "managedNamespaces"),
		resourceids.UserSpecifiedSegment("managedNamespaceName", "managedNamespaceName"),
	}
}

// String returns a human-readable description of this Managed Cluster Managed Namespace ID
func (id ManagedClusterManagedNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Cluster Name: %q", id.ManagedClusterName),
		fmt.Sprintf("Managed Namespace Name: %q", id.ManagedNamespaceName),
	}
	return fmt.Sprintf("Managed Cluster Managed Namespace (%s)", strings.Join(components, "\n"))
}
