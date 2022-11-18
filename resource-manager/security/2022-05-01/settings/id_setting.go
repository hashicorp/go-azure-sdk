package settings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SettingId{}

// SettingId is a struct representing the Resource ID for a Setting
type SettingId struct {
	SubscriptionId string
	SettingName    SettingName
}

// NewSettingID returns a new SettingId struct
func NewSettingID(subscriptionId string, settingName SettingName) SettingId {
	return SettingId{
		SubscriptionId: subscriptionId,
		SettingName:    settingName,
	}
}

// ParseSettingID parses 'input' into a SettingId
func ParseSettingID(input string) (*SettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(SettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["settingName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'settingName' was not found in the resource id %q", input)
		}

		settingName, err := parseSettingName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.SettingName = *settingName
	}

	return &id, nil
}

// ParseSettingIDInsensitively parses 'input' case-insensitively into a SettingId
// note: this method should only be used for API response data and not user input
func ParseSettingIDInsensitively(input string) (*SettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(SettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if v, ok := parsed.Parsed["settingName"]; true {
		if !ok {
			return nil, fmt.Errorf("the segment 'settingName' was not found in the resource id %q", input)
		}

		settingName, err := parseSettingName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.SettingName = *settingName
	}

	return &id, nil
}

// ValidateSettingID checks that 'input' can be parsed as a Setting ID
func ValidateSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Setting ID
func (id SettingId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/settings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, string(id.SettingName))
}

// Segments returns a slice of Resource ID Segments which comprise this Setting ID
func (id SettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSettings", "settings", "settings"),
		resourceids.ConstantSegment("settingName", PossibleValuesForSettingName(), "MCAS"),
	}
}

// String returns a human-readable description of this Setting ID
func (id SettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Setting Name: %q", string(id.SettingName)),
	}
	return fmt.Sprintf("Setting (%s)", strings.Join(components, "\n"))
}
