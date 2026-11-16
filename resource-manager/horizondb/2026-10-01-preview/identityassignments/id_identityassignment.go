package identityassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&IdentityAssignmentId{})
}

var _ resourceids.ResourceId = &IdentityAssignmentId{}

// IdentityAssignmentId is a struct representing the Resource ID for a Identity Assignment
type IdentityAssignmentId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	ClusterName                      string
	ChangeDataCaptureDestinationName string
	IdentityAssignmentName           string
}

// NewIdentityAssignmentID returns a new IdentityAssignmentId struct
func NewIdentityAssignmentID(subscriptionId string, resourceGroupName string, clusterName string, changeDataCaptureDestinationName string, identityAssignmentName string) IdentityAssignmentId {
	return IdentityAssignmentId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		ClusterName:                      clusterName,
		ChangeDataCaptureDestinationName: changeDataCaptureDestinationName,
		IdentityAssignmentName:           identityAssignmentName,
	}
}

// ParseIdentityAssignmentID parses 'input' into a IdentityAssignmentId
func ParseIdentityAssignmentID(input string) (*IdentityAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityAssignmentIDInsensitively(input string) (*IdentityAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityAssignmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ChangeDataCaptureDestinationName, ok = input.Parsed["changeDataCaptureDestinationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "changeDataCaptureDestinationName", input)
	}

	if id.IdentityAssignmentName, ok = input.Parsed["identityAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityAssignmentName", input)
	}

	return nil
}

// ValidateIdentityAssignmentID checks that 'input' can be parsed as a Identity Assignment ID
func ValidateIdentityAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Assignment ID
func (id IdentityAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/clusters/%s/changeDataCaptureDestinations/%s/identityAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ChangeDataCaptureDestinationName, id.IdentityAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Assignment ID
func (id IdentityAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHorizonDb", "Microsoft.HorizonDb", "Microsoft.HorizonDb"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticChangeDataCaptureDestinations", "changeDataCaptureDestinations", "changeDataCaptureDestinations"),
		resourceids.UserSpecifiedSegment("changeDataCaptureDestinationName", "changeDataCaptureDestinationName"),
		resourceids.StaticSegment("staticIdentityAssignments", "identityAssignments", "identityAssignments"),
		resourceids.UserSpecifiedSegment("identityAssignmentName", "identityAssignmentName"),
	}
}

// String returns a human-readable description of this Identity Assignment ID
func (id IdentityAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Change Data Capture Destination Name: %q", id.ChangeDataCaptureDestinationName),
		fmt.Sprintf("Identity Assignment Name: %q", id.IdentityAssignmentName),
	}
	return fmt.Sprintf("Identity Assignment (%s)", strings.Join(components, "\n"))
}
