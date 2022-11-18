package nodetype

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Access string

const (
	AccessAllow Access = "allow"
	AccessDeny  Access = "deny"
)

func PossibleValuesForAccess() []string {
	return []string{
		string(AccessAllow),
		string(AccessDeny),
	}
}

func parseAccess(input string) (*Access, error) {
	vals := map[string]Access{
		"allow": AccessAllow,
		"deny":  AccessDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Access(input)
	return &out, nil
}

type Direction string

const (
	DirectionInbound  Direction = "inbound"
	DirectionOutbound Direction = "outbound"
)

func PossibleValuesForDirection() []string {
	return []string{
		string(DirectionInbound),
		string(DirectionOutbound),
	}
}

func parseDirection(input string) (*Direction, error) {
	vals := map[string]Direction{
		"inbound":  DirectionInbound,
		"outbound": DirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Direction(input)
	return &out, nil
}

type DiskType string

const (
	DiskTypePremiumLRS     DiskType = "Premium_LRS"
	DiskTypeStandardLRS    DiskType = "Standard_LRS"
	DiskTypeStandardSSDLRS DiskType = "StandardSSD_LRS"
)

func PossibleValuesForDiskType() []string {
	return []string{
		string(DiskTypePremiumLRS),
		string(DiskTypeStandardLRS),
		string(DiskTypeStandardSSDLRS),
	}
}

func parseDiskType(input string) (*DiskType, error) {
	vals := map[string]DiskType{
		"premium_lrs":     DiskTypePremiumLRS,
		"standard_lrs":    DiskTypeStandardLRS,
		"standardssd_lrs": DiskTypeStandardSSDLRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskType(input)
	return &out, nil
}

type IPAddressType string

const (
	IPAddressTypeIPvFour IPAddressType = "IPv4"
	IPAddressTypeIPvSix  IPAddressType = "IPv6"
)

func PossibleValuesForIPAddressType() []string {
	return []string{
		string(IPAddressTypeIPvFour),
		string(IPAddressTypeIPvSix),
	}
}

func parseIPAddressType(input string) (*IPAddressType, error) {
	vals := map[string]IPAddressType{
		"ipv4": IPAddressTypeIPvFour,
		"ipv6": IPAddressTypeIPvSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPAddressType(input)
	return &out, nil
}

type ManagedResourceProvisioningState string

const (
	ManagedResourceProvisioningStateCanceled  ManagedResourceProvisioningState = "Canceled"
	ManagedResourceProvisioningStateCreated   ManagedResourceProvisioningState = "Created"
	ManagedResourceProvisioningStateCreating  ManagedResourceProvisioningState = "Creating"
	ManagedResourceProvisioningStateDeleted   ManagedResourceProvisioningState = "Deleted"
	ManagedResourceProvisioningStateDeleting  ManagedResourceProvisioningState = "Deleting"
	ManagedResourceProvisioningStateFailed    ManagedResourceProvisioningState = "Failed"
	ManagedResourceProvisioningStateNone      ManagedResourceProvisioningState = "None"
	ManagedResourceProvisioningStateOther     ManagedResourceProvisioningState = "Other"
	ManagedResourceProvisioningStateSucceeded ManagedResourceProvisioningState = "Succeeded"
	ManagedResourceProvisioningStateUpdating  ManagedResourceProvisioningState = "Updating"
)

func PossibleValuesForManagedResourceProvisioningState() []string {
	return []string{
		string(ManagedResourceProvisioningStateCanceled),
		string(ManagedResourceProvisioningStateCreated),
		string(ManagedResourceProvisioningStateCreating),
		string(ManagedResourceProvisioningStateDeleted),
		string(ManagedResourceProvisioningStateDeleting),
		string(ManagedResourceProvisioningStateFailed),
		string(ManagedResourceProvisioningStateNone),
		string(ManagedResourceProvisioningStateOther),
		string(ManagedResourceProvisioningStateSucceeded),
		string(ManagedResourceProvisioningStateUpdating),
	}
}

func parseManagedResourceProvisioningState(input string) (*ManagedResourceProvisioningState, error) {
	vals := map[string]ManagedResourceProvisioningState{
		"canceled":  ManagedResourceProvisioningStateCanceled,
		"created":   ManagedResourceProvisioningStateCreated,
		"creating":  ManagedResourceProvisioningStateCreating,
		"deleted":   ManagedResourceProvisioningStateDeleted,
		"deleting":  ManagedResourceProvisioningStateDeleting,
		"failed":    ManagedResourceProvisioningStateFailed,
		"none":      ManagedResourceProvisioningStateNone,
		"other":     ManagedResourceProvisioningStateOther,
		"succeeded": ManagedResourceProvisioningStateSucceeded,
		"updating":  ManagedResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedResourceProvisioningState(input)
	return &out, nil
}

type NodeTypeSkuScaleType string

const (
	NodeTypeSkuScaleTypeAutomatic NodeTypeSkuScaleType = "Automatic"
	NodeTypeSkuScaleTypeManual    NodeTypeSkuScaleType = "Manual"
	NodeTypeSkuScaleTypeNone      NodeTypeSkuScaleType = "None"
)

func PossibleValuesForNodeTypeSkuScaleType() []string {
	return []string{
		string(NodeTypeSkuScaleTypeAutomatic),
		string(NodeTypeSkuScaleTypeManual),
		string(NodeTypeSkuScaleTypeNone),
	}
}

func parseNodeTypeSkuScaleType(input string) (*NodeTypeSkuScaleType, error) {
	vals := map[string]NodeTypeSkuScaleType{
		"automatic": NodeTypeSkuScaleTypeAutomatic,
		"manual":    NodeTypeSkuScaleTypeManual,
		"none":      NodeTypeSkuScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeTypeSkuScaleType(input)
	return &out, nil
}

type NsgProtocol string

const (
	NsgProtocolAh    NsgProtocol = "ah"
	NsgProtocolEsp   NsgProtocol = "esp"
	NsgProtocolHTTP  NsgProtocol = "http"
	NsgProtocolHTTPS NsgProtocol = "https"
	NsgProtocolIcmp  NsgProtocol = "icmp"
	NsgProtocolTcp   NsgProtocol = "tcp"
	NsgProtocolUdp   NsgProtocol = "udp"
)

func PossibleValuesForNsgProtocol() []string {
	return []string{
		string(NsgProtocolAh),
		string(NsgProtocolEsp),
		string(NsgProtocolHTTP),
		string(NsgProtocolHTTPS),
		string(NsgProtocolIcmp),
		string(NsgProtocolTcp),
		string(NsgProtocolUdp),
	}
}

func parseNsgProtocol(input string) (*NsgProtocol, error) {
	vals := map[string]NsgProtocol{
		"ah":    NsgProtocolAh,
		"esp":   NsgProtocolEsp,
		"http":  NsgProtocolHTTP,
		"https": NsgProtocolHTTPS,
		"icmp":  NsgProtocolIcmp,
		"tcp":   NsgProtocolTcp,
		"udp":   NsgProtocolUdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NsgProtocol(input)
	return &out, nil
}
