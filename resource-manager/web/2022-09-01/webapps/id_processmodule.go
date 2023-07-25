package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProcessModuleId{}

// ProcessModuleId is a struct representing the Resource ID for a Process Module
type ProcessModuleId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	ProcessId         string
	ModuleName        string
}

// NewProcessModuleID returns a new ProcessModuleId struct
func NewProcessModuleID(subscriptionId string, resourceGroupName string, siteName string, slotName string, processId string, moduleName string) ProcessModuleId {
	return ProcessModuleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		ProcessId:         processId,
		ModuleName:        moduleName,
	}
}

// ParseProcessModuleID parses 'input' into a ProcessModuleId
func ParseProcessModuleID(input string) (*ProcessModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProcessModuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProcessModuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "moduleName", *parsed)
	}

	return &id, nil
}

// ParseProcessModuleIDInsensitively parses 'input' case-insensitively into a ProcessModuleId
// note: this method should only be used for API response data and not user input
func ParseProcessModuleIDInsensitively(input string) (*ProcessModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProcessModuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProcessModuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "moduleName", *parsed)
	}

	return &id, nil
}

// ValidateProcessModuleID checks that 'input' can be parsed as a Process Module ID
func ValidateProcessModuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProcessModuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Process Module ID
func (id ProcessModuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/processes/%s/modules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.ProcessId, id.ModuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Process Module ID
func (id ProcessModuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotValue"),
		resourceids.StaticSegment("staticProcesses", "processes", "processes"),
		resourceids.UserSpecifiedSegment("processId", "processIdValue"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleValue"),
	}
}

// String returns a human-readable description of this Process Module ID
func (id ProcessModuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Process: %q", id.ProcessId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
	}
	return fmt.Sprintf("Process Module (%s)", strings.Join(components, "\n"))
}
