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
	recaser.RegisterResourceId(&ArmTemplateId{})
}

var _ resourceids.ResourceId = &ArmTemplateId{}

// ArmTemplateId is a struct representing the Resource ID for a Arm Template
type ArmTemplateId struct {
	SubscriptionId     string
	ResourceGroupName  string
	LabName            string
	ArtifactSourceName string
	ArmTemplateName    string
}

// NewArmTemplateID returns a new ArmTemplateId struct
func NewArmTemplateID(subscriptionId string, resourceGroupName string, labName string, artifactSourceName string, armTemplateName string) ArmTemplateId {
	return ArmTemplateId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		LabName:            labName,
		ArtifactSourceName: artifactSourceName,
		ArmTemplateName:    armTemplateName,
	}
}

// ParseArmTemplateID parses 'input' into a ArmTemplateId
func ParseArmTemplateID(input string) (*ArmTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArmTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArmTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseArmTemplateIDInsensitively parses 'input' case-insensitively into a ArmTemplateId
// note: this method should only be used for API response data and not user input
func ParseArmTemplateIDInsensitively(input string) (*ArmTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ArmTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ArmTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ArmTemplateId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ArmTemplateName, ok = input.Parsed["armTemplateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "armTemplateName", input)
	}

	return nil
}

// ValidateArmTemplateID checks that 'input' can be parsed as a Arm Template ID
func ValidateArmTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseArmTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Arm Template ID
func (id ArmTemplateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/artifactSources/%s/armTemplates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.ArtifactSourceName, id.ArmTemplateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Arm Template ID
func (id ArmTemplateId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticArmTemplates", "armTemplates", "armTemplates"),
		resourceids.UserSpecifiedSegment("armTemplateName", "armTemplateValue"),
	}
}

// String returns a human-readable description of this Arm Template ID
func (id ArmTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Artifact Source Name: %q", id.ArtifactSourceName),
		fmt.Sprintf("Arm Template Name: %q", id.ArmTemplateName),
	}
	return fmt.Sprintf("Arm Template (%s)", strings.Join(components, "\n"))
}
