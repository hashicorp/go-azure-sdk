package instructions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &InstructionId{}

// InstructionId is a struct representing the Resource ID for a Instruction
type InstructionId struct {
	BillingAccountName string
	BillingProfileName string
	InstructionName    string
}

// NewInstructionID returns a new InstructionId struct
func NewInstructionID(billingAccountName string, billingProfileName string, instructionName string) InstructionId {
	return InstructionId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		InstructionName:    instructionName,
	}
}

// ParseInstructionID parses 'input' into a InstructionId
func ParseInstructionID(input string) (*InstructionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InstructionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InstructionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInstructionIDInsensitively parses 'input' case-insensitively into a InstructionId
// note: this method should only be used for API response data and not user input
func ParseInstructionIDInsensitively(input string) (*InstructionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InstructionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InstructionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InstructionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.InstructionName, ok = input.Parsed["instructionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "instructionName", input)
	}

	return nil
}

// ValidateInstructionID checks that 'input' can be parsed as a Instruction ID
func ValidateInstructionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInstructionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Instruction ID
func (id InstructionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/instructions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InstructionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Instruction ID
func (id InstructionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticInstructions", "instructions", "instructions"),
		resourceids.UserSpecifiedSegment("instructionName", "instructionValue"),
	}
}

// String returns a human-readable description of this Instruction ID
func (id InstructionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Instruction Name: %q", id.InstructionName),
	}
	return fmt.Sprintf("Instruction (%s)", strings.Join(components, "\n"))
}
