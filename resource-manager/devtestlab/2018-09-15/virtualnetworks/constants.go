package virtualnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransportProtocol string

const (
	TransportProtocolTcp TransportProtocol = "Tcp"
	TransportProtocolUdp TransportProtocol = "Udp"
)

func PossibleValuesForTransportProtocol() []string {
	return []string{
		string(TransportProtocolTcp),
		string(TransportProtocolUdp),
	}
}

type UsagePermissionType string

const (
	UsagePermissionTypeAllow   UsagePermissionType = "Allow"
	UsagePermissionTypeDefault UsagePermissionType = "Default"
	UsagePermissionTypeDeny    UsagePermissionType = "Deny"
)

func PossibleValuesForUsagePermissionType() []string {
	return []string{
		string(UsagePermissionTypeAllow),
		string(UsagePermissionTypeDefault),
		string(UsagePermissionTypeDeny),
	}
}
