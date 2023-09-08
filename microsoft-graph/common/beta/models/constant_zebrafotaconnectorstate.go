package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaConnectorState string

const (
	ZebraFotaConnectorStateconnected    ZebraFotaConnectorState = "Connected"
	ZebraFotaConnectorStatedisconnected ZebraFotaConnectorState = "Disconnected"
	ZebraFotaConnectorStatenone         ZebraFotaConnectorState = "None"
)

func PossibleValuesForZebraFotaConnectorState() []string {
	return []string{
		string(ZebraFotaConnectorStateconnected),
		string(ZebraFotaConnectorStatedisconnected),
		string(ZebraFotaConnectorStatenone),
	}
}

func parseZebraFotaConnectorState(input string) (*ZebraFotaConnectorState, error) {
	vals := map[string]ZebraFotaConnectorState{
		"connected":    ZebraFotaConnectorStateconnected,
		"disconnected": ZebraFotaConnectorStatedisconnected,
		"none":         ZebraFotaConnectorStatenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaConnectorState(input)
	return &out, nil
}
