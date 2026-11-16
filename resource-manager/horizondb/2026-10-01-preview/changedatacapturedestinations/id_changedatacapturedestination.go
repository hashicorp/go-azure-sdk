package changedatacapturedestinations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ChangeDataCaptureDestinationId{})
}

var _ resourceids.ResourceId = &ChangeDataCaptureDestinationId{}

// ChangeDataCaptureDestinationId is a struct representing the Resource ID for a Change Data Capture Destination
type ChangeDataCaptureDestinationId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	ClusterName                      string
	ChangeDataCaptureDestinationName string
}

// NewChangeDataCaptureDestinationID returns a new ChangeDataCaptureDestinationId struct
func NewChangeDataCaptureDestinationID(subscriptionId string, resourceGroupName string, clusterName string, changeDataCaptureDestinationName string) ChangeDataCaptureDestinationId {
	return ChangeDataCaptureDestinationId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		ClusterName:                      clusterName,
		ChangeDataCaptureDestinationName: changeDataCaptureDestinationName,
	}
}

// ParseChangeDataCaptureDestinationID parses 'input' into a ChangeDataCaptureDestinationId
func ParseChangeDataCaptureDestinationID(input string) (*ChangeDataCaptureDestinationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ChangeDataCaptureDestinationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ChangeDataCaptureDestinationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseChangeDataCaptureDestinationIDInsensitively parses 'input' case-insensitively into a ChangeDataCaptureDestinationId
// note: this method should only be used for API response data and not user input
func ParseChangeDataCaptureDestinationIDInsensitively(input string) (*ChangeDataCaptureDestinationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ChangeDataCaptureDestinationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ChangeDataCaptureDestinationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ChangeDataCaptureDestinationId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateChangeDataCaptureDestinationID checks that 'input' can be parsed as a Change Data Capture Destination ID
func ValidateChangeDataCaptureDestinationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseChangeDataCaptureDestinationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Change Data Capture Destination ID
func (id ChangeDataCaptureDestinationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HorizonDb/clusters/%s/changeDataCaptureDestinations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ChangeDataCaptureDestinationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Change Data Capture Destination ID
func (id ChangeDataCaptureDestinationId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Change Data Capture Destination ID
func (id ChangeDataCaptureDestinationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Change Data Capture Destination Name: %q", id.ChangeDataCaptureDestinationName),
	}
	return fmt.Sprintf("Change Data Capture Destination (%s)", strings.Join(components, "\n"))
}
