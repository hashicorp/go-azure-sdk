package integrationruntimenodes

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeUpdateResult string

const (
	IntegrationRuntimeUpdateResultFail    IntegrationRuntimeUpdateResult = "Fail"
	IntegrationRuntimeUpdateResultNone    IntegrationRuntimeUpdateResult = "None"
	IntegrationRuntimeUpdateResultSucceed IntegrationRuntimeUpdateResult = "Succeed"
)

func PossibleValuesForIntegrationRuntimeUpdateResult() []string {
	return []string{
		string(IntegrationRuntimeUpdateResultFail),
		string(IntegrationRuntimeUpdateResultNone),
		string(IntegrationRuntimeUpdateResultSucceed),
	}
}

func parseIntegrationRuntimeUpdateResult(input string) (*IntegrationRuntimeUpdateResult, error) {
	vals := map[string]IntegrationRuntimeUpdateResult{
		"fail":    IntegrationRuntimeUpdateResultFail,
		"none":    IntegrationRuntimeUpdateResultNone,
		"succeed": IntegrationRuntimeUpdateResultSucceed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeUpdateResult(input)
	return &out, nil
}

type SelfHostedIntegrationRuntimeNodeStatus string

const (
	SelfHostedIntegrationRuntimeNodeStatusInitializeFailed SelfHostedIntegrationRuntimeNodeStatus = "InitializeFailed"
	SelfHostedIntegrationRuntimeNodeStatusInitializing     SelfHostedIntegrationRuntimeNodeStatus = "Initializing"
	SelfHostedIntegrationRuntimeNodeStatusLimited          SelfHostedIntegrationRuntimeNodeStatus = "Limited"
	SelfHostedIntegrationRuntimeNodeStatusNeedRegistration SelfHostedIntegrationRuntimeNodeStatus = "NeedRegistration"
	SelfHostedIntegrationRuntimeNodeStatusOffline          SelfHostedIntegrationRuntimeNodeStatus = "Offline"
	SelfHostedIntegrationRuntimeNodeStatusOnline           SelfHostedIntegrationRuntimeNodeStatus = "Online"
	SelfHostedIntegrationRuntimeNodeStatusUpgrading        SelfHostedIntegrationRuntimeNodeStatus = "Upgrading"
)

func PossibleValuesForSelfHostedIntegrationRuntimeNodeStatus() []string {
	return []string{
		string(SelfHostedIntegrationRuntimeNodeStatusInitializeFailed),
		string(SelfHostedIntegrationRuntimeNodeStatusInitializing),
		string(SelfHostedIntegrationRuntimeNodeStatusLimited),
		string(SelfHostedIntegrationRuntimeNodeStatusNeedRegistration),
		string(SelfHostedIntegrationRuntimeNodeStatusOffline),
		string(SelfHostedIntegrationRuntimeNodeStatusOnline),
		string(SelfHostedIntegrationRuntimeNodeStatusUpgrading),
	}
}

func parseSelfHostedIntegrationRuntimeNodeStatus(input string) (*SelfHostedIntegrationRuntimeNodeStatus, error) {
	vals := map[string]SelfHostedIntegrationRuntimeNodeStatus{
		"initializefailed": SelfHostedIntegrationRuntimeNodeStatusInitializeFailed,
		"initializing":     SelfHostedIntegrationRuntimeNodeStatusInitializing,
		"limited":          SelfHostedIntegrationRuntimeNodeStatusLimited,
		"needregistration": SelfHostedIntegrationRuntimeNodeStatusNeedRegistration,
		"offline":          SelfHostedIntegrationRuntimeNodeStatusOffline,
		"online":           SelfHostedIntegrationRuntimeNodeStatusOnline,
		"upgrading":        SelfHostedIntegrationRuntimeNodeStatusUpgrading,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SelfHostedIntegrationRuntimeNodeStatus(input)
	return &out, nil
}
