package containerappspatches

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PatchId{})
}

var _ resourceids.ResourceId = &PatchId{}

// PatchId is a struct representing the Resource ID for a Patch
type PatchId struct {
	SubscriptionId    string
	ResourceGroupName string
	ContainerAppName  string
	PatchName         string
}

// NewPatchID returns a new PatchId struct
func NewPatchID(subscriptionId string, resourceGroupName string, containerAppName string, patchName string) PatchId {
	return PatchId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ContainerAppName:  containerAppName,
		PatchName:         patchName,
	}
}

// ParsePatchID parses 'input' into a PatchId
func ParsePatchID(input string) (*PatchId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PatchId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PatchId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePatchIDInsensitively parses 'input' case-insensitively into a PatchId
// note: this method should only be used for API response data and not user input
func ParsePatchIDInsensitively(input string) (*PatchId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PatchId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PatchId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PatchId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ContainerAppName, ok = input.Parsed["containerAppName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "containerAppName", input)
	}

	if id.PatchName, ok = input.Parsed["patchName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "patchName", input)
	}

	return nil
}

// ValidatePatchID checks that 'input' can be parsed as a Patch ID
func ValidatePatchID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePatchID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Patch ID
func (id PatchId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/containerApps/%s/patches/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ContainerAppName, id.PatchName)
}

// Segments returns a slice of Resource ID Segments which comprise this Patch ID
func (id PatchId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticContainerApps", "containerApps", "containerApps"),
		resourceids.UserSpecifiedSegment("containerAppName", "containerAppValue"),
		resourceids.StaticSegment("staticPatches", "patches", "patches"),
		resourceids.UserSpecifiedSegment("patchName", "patchValue"),
	}
}

// String returns a human-readable description of this Patch ID
func (id PatchId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Container App Name: %q", id.ContainerAppName),
		fmt.Sprintf("Patch Name: %q", id.PatchName),
	}
	return fmt.Sprintf("Patch (%s)", strings.Join(components, "\n"))
}
