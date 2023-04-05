package runbook

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
