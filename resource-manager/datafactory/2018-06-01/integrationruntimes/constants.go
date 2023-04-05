package integrationruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeAuthKeyName string

const (
	IntegrationRuntimeAuthKeyNameAuthKeyOne IntegrationRuntimeAuthKeyName = "authKey1"
	IntegrationRuntimeAuthKeyNameAuthKeyTwo IntegrationRuntimeAuthKeyName = "authKey2"
)

func PossibleValuesForIntegrationRuntimeAuthKeyName() []string {
	return []string{
		string(IntegrationRuntimeAuthKeyNameAuthKeyOne),
		string(IntegrationRuntimeAuthKeyNameAuthKeyTwo),
	}
}

type IntegrationRuntimeAutoUpdate string

const (
	IntegrationRuntimeAutoUpdateOff IntegrationRuntimeAutoUpdate = "Off"
	IntegrationRuntimeAutoUpdateOn  IntegrationRuntimeAutoUpdate = "On"
)

func PossibleValuesForIntegrationRuntimeAutoUpdate() []string {
	return []string{
		string(IntegrationRuntimeAutoUpdateOff),
		string(IntegrationRuntimeAutoUpdateOn),
	}
}

type IntegrationRuntimeState string

const (
	IntegrationRuntimeStateAccessDenied     IntegrationRuntimeState = "AccessDenied"
	IntegrationRuntimeStateInitial          IntegrationRuntimeState = "Initial"
	IntegrationRuntimeStateLimited          IntegrationRuntimeState = "Limited"
	IntegrationRuntimeStateNeedRegistration IntegrationRuntimeState = "NeedRegistration"
	IntegrationRuntimeStateOffline          IntegrationRuntimeState = "Offline"
	IntegrationRuntimeStateOnline           IntegrationRuntimeState = "Online"
	IntegrationRuntimeStateStarted          IntegrationRuntimeState = "Started"
	IntegrationRuntimeStateStarting         IntegrationRuntimeState = "Starting"
	IntegrationRuntimeStateStopped          IntegrationRuntimeState = "Stopped"
	IntegrationRuntimeStateStopping         IntegrationRuntimeState = "Stopping"
)

func PossibleValuesForIntegrationRuntimeState() []string {
	return []string{
		string(IntegrationRuntimeStateAccessDenied),
		string(IntegrationRuntimeStateInitial),
		string(IntegrationRuntimeStateLimited),
		string(IntegrationRuntimeStateNeedRegistration),
		string(IntegrationRuntimeStateOffline),
		string(IntegrationRuntimeStateOnline),
		string(IntegrationRuntimeStateStarted),
		string(IntegrationRuntimeStateStarting),
		string(IntegrationRuntimeStateStopped),
		string(IntegrationRuntimeStateStopping),
	}
}

type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    IntegrationRuntimeType = "Managed"
	IntegrationRuntimeTypeSelfHosted IntegrationRuntimeType = "SelfHosted"
)

func PossibleValuesForIntegrationRuntimeType() []string {
	return []string{
		string(IntegrationRuntimeTypeManaged),
		string(IntegrationRuntimeTypeSelfHosted),
	}
}
