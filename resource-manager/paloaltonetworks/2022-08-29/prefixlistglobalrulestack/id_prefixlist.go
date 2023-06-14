package prefixlistglobalrulestack

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PrefixListId{}

// PrefixListId is a struct representing the Resource ID for a Prefix List
type PrefixListId struct {
	GlobalRuleStackName string
	PrefixListName      string
}

// NewPrefixListID returns a new PrefixListId struct
func NewPrefixListID(globalRuleStackName string, prefixListName string) PrefixListId {
	return PrefixListId{
		GlobalRuleStackName: globalRuleStackName,
		PrefixListName:      prefixListName,
	}
}

// ParsePrefixListID parses 'input' into a PrefixListId
func ParsePrefixListID(input string) (*PrefixListId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrefixListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrefixListId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	if id.PrefixListName, ok = parsed.Parsed["prefixListName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "prefixListName", *parsed)
	}

	return &id, nil
}

// ParsePrefixListIDInsensitively parses 'input' case-insensitively into a PrefixListId
// note: this method should only be used for API response data and not user input
func ParsePrefixListIDInsensitively(input string) (*PrefixListId, error) {
	parser := resourceids.NewParserFromResourceIdType(PrefixListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PrefixListId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	if id.PrefixListName, ok = parsed.Parsed["prefixListName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "prefixListName", *parsed)
	}

	return &id, nil
}

// ValidatePrefixListID checks that 'input' can be parsed as a Prefix List ID
func ValidatePrefixListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePrefixListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Prefix List ID
func (id PrefixListId) ID() string {
	fmtString := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/%s/prefixLists/%s"
	return fmt.Sprintf(fmtString, id.GlobalRuleStackName, id.PrefixListName)
}

// Segments returns a slice of Resource ID Segments which comprise this Prefix List ID
func (id PrefixListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticPaloAltoNetworksCloudNGFW", "PaloAltoNetworks.CloudNGFW", "PaloAltoNetworks.CloudNGFW"),
		resourceids.StaticSegment("staticGlobalRuleStacks", "globalRuleStacks", "globalRuleStacks"),
		resourceids.UserSpecifiedSegment("globalRuleStackName", "globalRuleStackValue"),
		resourceids.StaticSegment("staticPrefixLists", "prefixLists", "prefixLists"),
		resourceids.UserSpecifiedSegment("prefixListName", "prefixListValue"),
	}
}

// String returns a human-readable description of this Prefix List ID
func (id PrefixListId) String() string {
	components := []string{
		fmt.Sprintf("Global Rule Stack Name: %q", id.GlobalRuleStackName),
		fmt.Sprintf("Prefix List Name: %q", id.PrefixListName),
	}
	return fmt.Sprintf("Prefix List (%s)", strings.Join(components, "\n"))
}
