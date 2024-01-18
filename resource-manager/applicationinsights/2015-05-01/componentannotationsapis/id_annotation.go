package componentannotationsapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AnnotationId{}

// AnnotationId is a struct representing the Resource ID for a Annotation
type AnnotationId struct {
	SubscriptionId    string
	ResourceGroupName string
	ComponentName     string
	AnnotationId      string
}

// NewAnnotationID returns a new AnnotationId struct
func NewAnnotationID(subscriptionId string, resourceGroupName string, componentName string, annotationId string) AnnotationId {
	return AnnotationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ComponentName:     componentName,
		AnnotationId:      annotationId,
	}
}

// ParseAnnotationID parses 'input' into a AnnotationId
func ParseAnnotationID(input string) (*AnnotationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AnnotationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AnnotationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAnnotationIDInsensitively parses 'input' case-insensitively into a AnnotationId
// note: this method should only be used for API response data and not user input
func ParseAnnotationIDInsensitively(input string) (*AnnotationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AnnotationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AnnotationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AnnotationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	if id.AnnotationId, ok = input.Parsed["annotationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "annotationId", input)
	}

	return nil
}

// ValidateAnnotationID checks that 'input' can be parsed as a Annotation ID
func ValidateAnnotationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAnnotationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Annotation ID
func (id AnnotationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/components/%s/annotations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ComponentName, id.AnnotationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Annotation ID
func (id AnnotationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentValue"),
		resourceids.StaticSegment("staticAnnotations", "annotations", "annotations"),
		resourceids.UserSpecifiedSegment("annotationId", "annotationIdValue"),
	}
}

// String returns a human-readable description of this Annotation ID
func (id AnnotationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Annotation: %q", id.AnnotationId),
	}
	return fmt.Sprintf("Annotation (%s)", strings.Join(components, "\n"))
}
