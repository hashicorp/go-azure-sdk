package fqdnlistglobalrulestack

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FqdnListId{}

// FqdnListId is a struct representing the Resource ID for a Fqdn List
type FqdnListId struct {
	GlobalRuleStackName string
	FqdnListName        string
}

// NewFqdnListID returns a new FqdnListId struct
func NewFqdnListID(globalRuleStackName string, fqdnListName string) FqdnListId {
	return FqdnListId{
		GlobalRuleStackName: globalRuleStackName,
		FqdnListName:        fqdnListName,
	}
}

// ParseFqdnListID parses 'input' into a FqdnListId
func ParseFqdnListID(input string) (*FqdnListId, error) {
	parser := resourceids.NewParserFromResourceIdType(FqdnListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FqdnListId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	if id.FqdnListName, ok = parsed.Parsed["fqdnListName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fqdnListName", *parsed)
	}

	return &id, nil
}

// ParseFqdnListIDInsensitively parses 'input' case-insensitively into a FqdnListId
// note: this method should only be used for API response data and not user input
func ParseFqdnListIDInsensitively(input string) (*FqdnListId, error) {
	parser := resourceids.NewParserFromResourceIdType(FqdnListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FqdnListId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	if id.FqdnListName, ok = parsed.Parsed["fqdnListName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fqdnListName", *parsed)
	}

	return &id, nil
}

// ValidateFqdnListID checks that 'input' can be parsed as a Fqdn List ID
func ValidateFqdnListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFqdnListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Fqdn List ID
func (id FqdnListId) ID() string {
	fmtString := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/%s/fqdnLists/%s"
	return fmt.Sprintf(fmtString, id.GlobalRuleStackName, id.FqdnListName)
}

// Segments returns a slice of Resource ID Segments which comprise this Fqdn List ID
func (id FqdnListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticPaloAltoNetworksCloudNGFW", "PaloAltoNetworks.CloudNGFW", "PaloAltoNetworks.CloudNGFW"),
		resourceids.StaticSegment("staticGlobalRuleStacks", "globalRuleStacks", "globalRuleStacks"),
		resourceids.UserSpecifiedSegment("globalRuleStackName", "globalRuleStackValue"),
		resourceids.StaticSegment("staticFqdnLists", "fqdnLists", "fqdnLists"),
		resourceids.UserSpecifiedSegment("fqdnListName", "fqdnListValue"),
	}
}

// String returns a human-readable description of this Fqdn List ID
func (id FqdnListId) String() string {
	components := []string{
		fmt.Sprintf("Global Rule Stack Name: %q", id.GlobalRuleStackName),
		fmt.Sprintf("Fqdn List Name: %q", id.FqdnListName),
	}
	return fmt.Sprintf("Fqdn List (%s)", strings.Join(components, "\n"))
}
