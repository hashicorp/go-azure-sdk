package adaptivenetworkhardenings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Rule struct {
	DestinationPort *int64               `json:"destinationPort,omitempty"`
	Direction       *Direction           `json:"direction,omitempty"`
	IPAddresses     *[]string            `json:"ipAddresses,omitempty"`
	Name            *string              `json:"name,omitempty"`
	Protocols       *[]TransportProtocol `json:"protocols,omitempty"`
}
