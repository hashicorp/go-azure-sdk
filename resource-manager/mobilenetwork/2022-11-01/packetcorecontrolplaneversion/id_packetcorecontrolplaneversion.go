package packetcorecontrolplaneversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PacketCoreControlPlaneVersionId{}

// PacketCoreControlPlaneVersionId is a struct representing the Resource ID for a Packet Core Control Plane Version
type PacketCoreControlPlaneVersionId struct {
	VersionName string
}

// NewPacketCoreControlPlaneVersionID returns a new PacketCoreControlPlaneVersionId struct
func NewPacketCoreControlPlaneVersionID(versionName string) PacketCoreControlPlaneVersionId {
	return PacketCoreControlPlaneVersionId{
		VersionName: versionName,
	}
}

// ParsePacketCoreControlPlaneVersionID parses 'input' into a PacketCoreControlPlaneVersionId
func ParsePacketCoreControlPlaneVersionID(input string) (*PacketCoreControlPlaneVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(PacketCoreControlPlaneVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PacketCoreControlPlaneVersionId{}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, fmt.Errorf("the segment 'versionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePacketCoreControlPlaneVersionIDInsensitively parses 'input' case-insensitively into a PacketCoreControlPlaneVersionId
// note: this method should only be used for API response data and not user input
func ParsePacketCoreControlPlaneVersionIDInsensitively(input string) (*PacketCoreControlPlaneVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(PacketCoreControlPlaneVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PacketCoreControlPlaneVersionId{}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, fmt.Errorf("the segment 'versionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePacketCoreControlPlaneVersionID checks that 'input' can be parsed as a Packet Core Control Plane Version ID
func ValidatePacketCoreControlPlaneVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePacketCoreControlPlaneVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Packet Core Control Plane Version ID
func (id PacketCoreControlPlaneVersionId) ID() string {
	fmtString := "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/%s"
	return fmt.Sprintf(fmtString, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Packet Core Control Plane Version ID
func (id PacketCoreControlPlaneVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMobileNetwork", "Microsoft.MobileNetwork", "Microsoft.MobileNetwork"),
		resourceids.StaticSegment("staticPacketCoreControlPlaneVersions", "packetCoreControlPlaneVersions", "packetCoreControlPlaneVersions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Packet Core Control Plane Version ID
func (id PacketCoreControlPlaneVersionId) String() string {
	components := []string{
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Packet Core Control Plane Version (%s)", strings.Join(components, "\n"))
}
