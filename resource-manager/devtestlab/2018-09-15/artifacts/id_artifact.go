package artifacts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ArtifactId{})
}

var _ resourceids.ResourceId = &ArtifactId{}

// ArtifactId is a struct representing the Resource ID for a Artifact
type ArtifactId struct {
	SubscriptionId     string
	ResourceGroupName  string
	LabName            string
	ArtifactSourceName string
	ArtifactName       string
}

// NewArtifactID returns a new ArtifactId struct
func NewArtifactID(subscriptionId string, resourceGroupName string, labName string, artifactSourceName string, artifactName string) ArtifactId {
	return ArtifactId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		LabName:            labName,
		ArtifactSourceName: artifactSourceName,
		ArtifactName:       artifactName,
	}
}

// ParseArtifactID parses 'input' into a ArtifactId
func ParseArtifactID(input string) (*ArtifactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArtifactId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArtifactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseArtifactIDInsensitively parses 'input' case-insensitively into a ArtifactId
// note: this method should only be used for API response data and not user input
func ParseArtifactIDInsensitively(input string) (*ArtifactId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArtifactId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArtifactId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ArtifactId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ArtifactName, ok = input.Parsed["artifactName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "artifactName", input)
	}

	return nil
}

// ValidateArtifactID checks that 'input' can be parsed as a Artifact ID
func ValidateArtifactID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseArtifactID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Artifact ID
func (id ArtifactId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/artifactSources/%s/artifacts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.ArtifactSourceName, id.ArtifactName)
}

// Segments returns a slice of Resource ID Segments which comprise this Artifact ID
func (id ArtifactId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticArtifacts", "artifacts", "artifacts"),
		resourceids.UserSpecifiedSegment("artifactName", "artifactValue"),
	}
}

// String returns a human-readable description of this Artifact ID
func (id ArtifactId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Artifact Source Name: %q", id.ArtifactSourceName),
		fmt.Sprintf("Artifact Name: %q", id.ArtifactName),
	}
	return fmt.Sprintf("Artifact (%s)", strings.Join(components, "\n"))
}
