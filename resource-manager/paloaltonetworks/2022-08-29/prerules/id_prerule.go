package prerules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PreRuleId{}

// PreRuleId is a struct representing the Resource ID for a Pre Rule
type PreRuleId struct {
	GlobalRuleStackName string
	PreRuleName         string
}

// NewPreRuleID returns a new PreRuleId struct
func NewPreRuleID(globalRuleStackName string, preRuleName string) PreRuleId {
	return PreRuleId{
		GlobalRuleStackName: globalRuleStackName,
		PreRuleName:         preRuleName,
	}
}

// ParsePreRuleID parses 'input' into a PreRuleId
func ParsePreRuleID(input string) (*PreRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PreRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PreRuleId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	if id.PreRuleName, ok = parsed.Parsed["preRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "preRuleName", *parsed)
	}

	return &id, nil
}

// ParsePreRuleIDInsensitively parses 'input' case-insensitively into a PreRuleId
// note: this method should only be used for API response data and not user input
func ParsePreRuleIDInsensitively(input string) (*PreRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PreRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PreRuleId{}

	if id.GlobalRuleStackName, ok = parsed.Parsed["globalRuleStackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "globalRuleStackName", *parsed)
	}

	if id.PreRuleName, ok = parsed.Parsed["preRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "preRuleName", *parsed)
	}

	return &id, nil
}

// ValidatePreRuleID checks that 'input' can be parsed as a Pre Rule ID
func ValidatePreRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePreRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pre Rule ID
func (id PreRuleId) ID() string {
	fmtString := "/providers/PaloAltoNetworks.CloudNGFW/globalRuleStacks/%s/preRules/%s"
	return fmt.Sprintf(fmtString, id.GlobalRuleStackName, id.PreRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Pre Rule ID
func (id PreRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticPaloAltoNetworksCloudNGFW", "PaloAltoNetworks.CloudNGFW", "PaloAltoNetworks.CloudNGFW"),
		resourceids.StaticSegment("staticGlobalRuleStacks", "globalRuleStacks", "globalRuleStacks"),
		resourceids.UserSpecifiedSegment("globalRuleStackName", "globalRuleStackValue"),
		resourceids.StaticSegment("staticPreRules", "preRules", "preRules"),
		resourceids.UserSpecifiedSegment("preRuleName", "preRuleValue"),
	}
}

// String returns a human-readable description of this Pre Rule ID
func (id PreRuleId) String() string {
	components := []string{
		fmt.Sprintf("Global Rule Stack Name: %q", id.GlobalRuleStackName),
		fmt.Sprintf("Pre Rule Name: %q", id.PreRuleName),
	}
	return fmt.Sprintf("Pre Rule (%s)", strings.Join(components, "\n"))
}
