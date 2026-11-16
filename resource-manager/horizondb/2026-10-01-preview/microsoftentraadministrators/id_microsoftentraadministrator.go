package microsoftentraadministrators

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MicrosoftEntraAdministratorId{})
}

var _ resourceids.ResourceId = &MicrosoftEntraAdministratorId{}

// MicrosoftEntraAdministratorId is a struct representing the Resource ID for a Microsoft Entra Administrator
type MicrosoftEntraAdministratorId struct {
	SubscriptionId    string
	ResourceGroupName string
	ClusterName       string
	ObjectId          string
}

// NewMicrosoftEntraAdministratorID returns a new MicrosoftEntraAdministratorId struct
func NewMicrosoftEntraAdministratorID(subscriptionId string, resourceGroupName string, clusterName string, objectId string) MicrosoftEntraAdministratorId {
	return MicrosoftEntraAdministratorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ClusterName:       clusterName,
		ObjectId:          objectId,
	}
}

// ParseMicrosoftEntraAdministratorID parses 'input' into a MicrosoftEntraAdministratorId
func ParseMicrosoftEntraAdministratorID(input string) (*MicrosoftEntraAdministratorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MicrosoftEntraAdministratorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MicrosoftEntraAdministratorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMicrosoftEntraAdministratorIDInsensitively parses 'input' case-insensitively into a MicrosoftEntraAdministratorId
// note: this method should only be used for API response data and not user input
func ParseMicrosoftEntraAdministratorIDInsensitively(input string) (*MicrosoftEntraAdministratorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MicrosoftEntraAdministratorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MicrosoftEntraAdministratorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MicrosoftEntraAdministratorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.ObjectId, ok = input.Parsed["objectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "objectId", input)
	}

	return nil
}

// ValidateMicrosoftEntraAdministratorID checks that 'input' can be parsed as a Microsoft Entra Administrator ID
func ValidateMicrosoftEntraAdministratorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMicrosoftEntraAdministratorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Microsoft Entra Administrator ID
func (id MicrosoftEntraAdministratorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/clusters/%s/microsoftEntraAdministrators/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Microsoft Entra Administrator ID
func (id MicrosoftEntraAdministratorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHorizonDb", "Microsoft.HorizonDb", "Microsoft.HorizonDb"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticMicrosoftEntraAdministrators", "microsoftEntraAdministrators", "microsoftEntraAdministrators"),
		resourceids.UserSpecifiedSegment("objectId", "objectId"),
	}
}

// String returns a human-readable description of this Microsoft Entra Administrator ID
func (id MicrosoftEntraAdministratorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Object: %q", id.ObjectId),
	}
	return fmt.Sprintf("Microsoft Entra Administrator (%s)", strings.Join(components, "\n"))
}
