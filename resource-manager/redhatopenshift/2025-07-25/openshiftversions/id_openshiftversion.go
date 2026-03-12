package openshiftversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&OpenShiftVersionId{})
}

var _ resourceids.ResourceId = &OpenShiftVersionId{}

// OpenShiftVersionId is a struct representing the Resource ID for a Open Shift Version
type OpenShiftVersionId struct {
	SubscriptionId       string
	LocationName         string
	OpenShiftVersionName string
}

// NewOpenShiftVersionID returns a new OpenShiftVersionId struct
func NewOpenShiftVersionID(subscriptionId string, locationName string, openShiftVersionName string) OpenShiftVersionId {
	return OpenShiftVersionId{
		SubscriptionId:       subscriptionId,
		LocationName:         locationName,
		OpenShiftVersionName: openShiftVersionName,
	}
}

// ParseOpenShiftVersionID parses 'input' into a OpenShiftVersionId
func ParseOpenShiftVersionID(input string) (*OpenShiftVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OpenShiftVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OpenShiftVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOpenShiftVersionIDInsensitively parses 'input' case-insensitively into a OpenShiftVersionId
// note: this method should only be used for API response data and not user input
func ParseOpenShiftVersionIDInsensitively(input string) (*OpenShiftVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OpenShiftVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OpenShiftVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OpenShiftVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.OpenShiftVersionName, ok = input.Parsed["openShiftVersionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "openShiftVersionName", input)
	}

	return nil
}

// ValidateOpenShiftVersionID checks that 'input' can be parsed as a Open Shift Version ID
func ValidateOpenShiftVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOpenShiftVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Open Shift Version ID
func (id OpenShiftVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.RedHatOpenShift/locations/%s/openShiftVersions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.OpenShiftVersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Open Shift Version ID
func (id OpenShiftVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRedHatOpenShift", "Microsoft.RedHatOpenShift", "Microsoft.RedHatOpenShift"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticOpenShiftVersions", "openShiftVersions", "openShiftVersions"),
		resourceids.UserSpecifiedSegment("openShiftVersionName", "openShiftVersionName"),
	}
}

// String returns a human-readable description of this Open Shift Version ID
func (id OpenShiftVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Open Shift Version Name: %q", id.OpenShiftVersionName),
	}
	return fmt.Sprintf("Open Shift Version (%s)", strings.Join(components, "\n"))
}
