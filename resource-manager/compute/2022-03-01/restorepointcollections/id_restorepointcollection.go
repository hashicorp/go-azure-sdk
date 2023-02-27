// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package restorepointcollections

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = RestorePointCollectionId{}

// RestorePointCollectionId is a struct representing the Resource ID for a Restore Point Collection
type RestorePointCollectionId struct {
	SubscriptionId             string
	ResourceGroupName          string
	RestorePointCollectionName string
}

// NewRestorePointCollectionID returns a new RestorePointCollectionId struct
func NewRestorePointCollectionID(subscriptionId string, resourceGroupName string, restorePointCollectionName string) RestorePointCollectionId {
	return RestorePointCollectionId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		RestorePointCollectionName: restorePointCollectionName,
	}
}

// ParseRestorePointCollectionID parses 'input' into a RestorePointCollectionId
func ParseRestorePointCollectionID(input string) (*RestorePointCollectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorePointCollectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorePointCollectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RestorePointCollectionName, ok = parsed.Parsed["restorePointCollectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointCollectionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRestorePointCollectionIDInsensitively parses 'input' case-insensitively into a RestorePointCollectionId
// note: this method should only be used for API response data and not user input
func ParseRestorePointCollectionIDInsensitively(input string) (*RestorePointCollectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorePointCollectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorePointCollectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RestorePointCollectionName, ok = parsed.Parsed["restorePointCollectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'restorePointCollectionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRestorePointCollectionID checks that 'input' can be parsed as a Restore Point Collection ID
func ValidateRestorePointCollectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRestorePointCollectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Restore Point Collection ID
func (id RestorePointCollectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/restorePointCollections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RestorePointCollectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Restore Point Collection ID
func (id RestorePointCollectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticRestorePointCollections", "restorePointCollections", "restorePointCollections"),
		resourceids.UserSpecifiedSegment("restorePointCollectionName", "restorePointCollectionValue"),
	}
}

// String returns a human-readable description of this Restore Point Collection ID
func (id RestorePointCollectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Restore Point Collection Name: %q", id.RestorePointCollectionName),
	}
	return fmt.Sprintf("Restore Point Collection (%s)", strings.Join(components, "\n"))
}
