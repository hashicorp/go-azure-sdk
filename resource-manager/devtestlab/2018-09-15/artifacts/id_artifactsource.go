// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifacts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ArtifactSourceId{}

// ArtifactSourceId is a struct representing the Resource ID for a Artifact Source
type ArtifactSourceId struct {
	SubscriptionId     string
	ResourceGroupName  string
	LabName            string
	ArtifactSourceName string
}

// NewArtifactSourceID returns a new ArtifactSourceId struct
func NewArtifactSourceID(subscriptionId string, resourceGroupName string, labName string, artifactSourceName string) ArtifactSourceId {
	return ArtifactSourceId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		LabName:            labName,
		ArtifactSourceName: artifactSourceName,
	}
}

// ParseArtifactSourceID parses 'input' into a ArtifactSourceId
func ParseArtifactSourceID(input string) (*ArtifactSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ArtifactSourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ArtifactSourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.ArtifactSourceName, ok = parsed.Parsed["artifactSourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'artifactSourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseArtifactSourceIDInsensitively parses 'input' case-insensitively into a ArtifactSourceId
// note: this method should only be used for API response data and not user input
func ParseArtifactSourceIDInsensitively(input string) (*ArtifactSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ArtifactSourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ArtifactSourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.ArtifactSourceName, ok = parsed.Parsed["artifactSourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'artifactSourceName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateArtifactSourceID checks that 'input' can be parsed as a Artifact Source ID
func ValidateArtifactSourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseArtifactSourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Artifact Source ID
func (id ArtifactSourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/artifactSources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.ArtifactSourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Artifact Source ID
func (id ArtifactSourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticArtifactSources", "artifactSources", "artifactSources"),
		resourceids.UserSpecifiedSegment("artifactSourceName", "artifactSourceValue"),
	}
}

// String returns a human-readable description of this Artifact Source ID
func (id ArtifactSourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Artifact Source Name: %q", id.ArtifactSourceName),
	}
	return fmt.Sprintf("Artifact Source (%s)", strings.Join(components, "\n"))
}
