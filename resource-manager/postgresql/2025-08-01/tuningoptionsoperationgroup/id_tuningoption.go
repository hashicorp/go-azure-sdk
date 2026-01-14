package tuningoptionsoperationgroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TuningOptionId{})
}

var _ resourceids.ResourceId = &TuningOptionId{}

// TuningOptionId is a struct representing the Resource ID for a Tuning Option
type TuningOptionId struct {
	SubscriptionId     string
	ResourceGroupName  string
	FlexibleServerName string
	TuningOption       TuningOptionParameterEnum
}

// NewTuningOptionID returns a new TuningOptionId struct
func NewTuningOptionID(subscriptionId string, resourceGroupName string, flexibleServerName string, tuningOption TuningOptionParameterEnum) TuningOptionId {
	return TuningOptionId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		FlexibleServerName: flexibleServerName,
		TuningOption:       tuningOption,
	}
}

// ParseTuningOptionID parses 'input' into a TuningOptionId
func ParseTuningOptionID(input string) (*TuningOptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TuningOptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TuningOptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTuningOptionIDInsensitively parses 'input' case-insensitively into a TuningOptionId
// note: this method should only be used for API response data and not user input
func ParseTuningOptionIDInsensitively(input string) (*TuningOptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TuningOptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TuningOptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TuningOptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.FlexibleServerName, ok = input.Parsed["flexibleServerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "flexibleServerName", input)
	}

	if v, ok := input.Parsed["tuningOption"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "tuningOption", input)
		}

		tuningOption, err := parseTuningOptionParameterEnum(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.TuningOption = *tuningOption
	}

	return nil
}

// ValidateTuningOptionID checks that 'input' can be parsed as a Tuning Option ID
func ValidateTuningOptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTuningOptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Tuning Option ID
func (id TuningOptionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/flexibleServers/%s/tuningOptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FlexibleServerName, string(id.TuningOption))
}

// Segments returns a slice of Resource ID Segments which comprise this Tuning Option ID
func (id TuningOptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticFlexibleServers", "flexibleServers", "flexibleServers"),
		resourceids.UserSpecifiedSegment("flexibleServerName", "flexibleServerName"),
		resourceids.StaticSegment("staticTuningOptions", "tuningOptions", "tuningOptions"),
		resourceids.ConstantSegment("tuningOption", PossibleValuesForTuningOptionParameterEnum(), "index"),
	}
}

// String returns a human-readable description of this Tuning Option ID
func (id TuningOptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Flexible Server Name: %q", id.FlexibleServerName),
		fmt.Sprintf("Tuning Option: %q", string(id.TuningOption)),
	}
	return fmt.Sprintf("Tuning Option (%s)", strings.Join(components, "\n"))
}
