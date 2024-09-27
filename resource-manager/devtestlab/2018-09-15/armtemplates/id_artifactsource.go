package armtemplates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ArtifactSourceId{})
}

var _ resourceids.ResourceId = &ArtifactSourceId{}

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
	parser := resourceids.NewParserFromResourceIdType(&ArtifactSourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArtifactSourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseArtifactSourceIDInsensitively parses 'input' case-insensitively into a ArtifactSourceId
// note: this method should only be used for API response data and not user input
func ParseArtifactSourceIDInsensitively(input string) (*ArtifactSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArtifactSourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArtifactSourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ArtifactSourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LabName, ok = input.Parsed["labName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labName", input)
	}

	if id.ArtifactSourceName, ok = input.Parsed["artifactSourceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "artifactSourceName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("labName", "labName"),
		resourceids.StaticSegment("staticArtifactSources", "artifactSources", "artifactSources"),
		resourceids.UserSpecifiedSegment("artifactSourceName", "name"),
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
