package integrationruntimenodes

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
