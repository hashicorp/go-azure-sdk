package diagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &HostingEnvironmentDetectorId{}

// HostingEnvironmentDetectorId is a struct representing the Resource ID for a Hosting Environment Detector
type HostingEnvironmentDetectorId struct {
	SubscriptionId         string
	ResourceGroupName      string
	HostingEnvironmentName string
	DetectorName           string
}

// NewHostingEnvironmentDetectorID returns a new HostingEnvironmentDetectorId struct
func NewHostingEnvironmentDetectorID(subscriptionId string, resourceGroupName string, hostingEnvironmentName string, detectorName string) HostingEnvironmentDetectorId {
	return HostingEnvironmentDetectorId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		HostingEnvironmentName: hostingEnvironmentName,
		DetectorName:           detectorName,
	}
}

// ParseHostingEnvironmentDetectorID parses 'input' into a HostingEnvironmentDetectorId
func ParseHostingEnvironmentDetectorID(input string) (*HostingEnvironmentDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HostingEnvironmentDetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HostingEnvironmentDetectorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHostingEnvironmentDetectorIDInsensitively parses 'input' case-insensitively into a HostingEnvironmentDetectorId
// note: this method should only be used for API response data and not user input
func ParseHostingEnvironmentDetectorIDInsensitively(input string) (*HostingEnvironmentDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HostingEnvironmentDetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HostingEnvironmentDetectorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HostingEnvironmentDetectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HostingEnvironmentName, ok = input.Parsed["hostingEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hostingEnvironmentName", input)
	}

	if id.DetectorName, ok = input.Parsed["detectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectorName", input)
	}

	return nil
}

// ValidateHostingEnvironmentDetectorID checks that 'input' can be parsed as a Hosting Environment Detector ID
func ValidateHostingEnvironmentDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHostingEnvironmentDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hosting Environment Detector ID
func (id HostingEnvironmentDetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/hostingEnvironments/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HostingEnvironmentName, id.DetectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hosting Environment Detector ID
func (id HostingEnvironmentDetectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticHostingEnvironments", "hostingEnvironments", "hostingEnvironments"),
		resourceids.UserSpecifiedSegment("hostingEnvironmentName", "hostingEnvironmentValue"),
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorName", "detectorValue"),
	}
}

// String returns a human-readable description of this Hosting Environment Detector ID
func (id HostingEnvironmentDetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Hosting Environment Name: %q", id.HostingEnvironmentName),
		fmt.Sprintf("Detector Name: %q", id.DetectorName),
	}
	return fmt.Sprintf("Hosting Environment Detector (%s)", strings.Join(components, "\n"))
}
