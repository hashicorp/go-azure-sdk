package runbook

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunbookProvisioningState string

const (
	RunbookProvisioningStateSucceeded RunbookProvisioningState = "Succeeded"
)

func PossibleValuesForRunbookProvisioningState() []string {
	return []string{
		string(RunbookProvisioningStateSucceeded),
	}
}

func (s *RunbookProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRunbookProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = RunbookProvisioningState(decoded)
	return nil
}

type RunbookState string

const (
	RunbookStateEdit      RunbookState = "Edit"
	RunbookStateNew       RunbookState = "New"
	RunbookStatePublished RunbookState = "Published"
)

func PossibleValuesForRunbookState() []string {
	return []string{
		string(RunbookStateEdit),
		string(RunbookStateNew),
		string(RunbookStatePublished),
	}
}

func (s *RunbookState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRunbookState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = RunbookState(decoded)
	return nil
}

type RunbookTypeEnum string

const (
	RunbookTypeEnumGraph                   RunbookTypeEnum = "Graph"
	RunbookTypeEnumGraphPowerShell         RunbookTypeEnum = "GraphPowerShell"
	RunbookTypeEnumGraphPowerShellWorkflow RunbookTypeEnum = "GraphPowerShellWorkflow"
	RunbookTypeEnumPowerShell              RunbookTypeEnum = "PowerShell"
	RunbookTypeEnumPowerShellWorkflow      RunbookTypeEnum = "PowerShellWorkflow"
	RunbookTypeEnumPythonThree             RunbookTypeEnum = "Python3"
	RunbookTypeEnumPythonTwo               RunbookTypeEnum = "Python2"
	RunbookTypeEnumScript                  RunbookTypeEnum = "Script"
)

func PossibleValuesForRunbookTypeEnum() []string {
	return []string{
		string(RunbookTypeEnumGraph),
		string(RunbookTypeEnumGraphPowerShell),
		string(RunbookTypeEnumGraphPowerShellWorkflow),
		string(RunbookTypeEnumPowerShell),
		string(RunbookTypeEnumPowerShellWorkflow),
		string(RunbookTypeEnumPythonThree),
		string(RunbookTypeEnumPythonTwo),
		string(RunbookTypeEnumScript),
	}
}

func (s *RunbookTypeEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForRunbookTypeEnum() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = RunbookTypeEnum(decoded)
	return nil
}
