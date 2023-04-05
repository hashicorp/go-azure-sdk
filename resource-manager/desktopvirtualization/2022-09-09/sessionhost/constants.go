package sessionhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthCheckName string

const (
	HealthCheckNameAppAttachHealthCheck     HealthCheckName = "AppAttachHealthCheck"
	HealthCheckNameDomainJoinedCheck        HealthCheckName = "DomainJoinedCheck"
	HealthCheckNameDomainReachable          HealthCheckName = "DomainReachable"
	HealthCheckNameDomainTrustCheck         HealthCheckName = "DomainTrustCheck"
	HealthCheckNameFSLogixHealthCheck       HealthCheckName = "FSLogixHealthCheck"
	HealthCheckNameMetaDataServiceCheck     HealthCheckName = "MetaDataServiceCheck"
	HealthCheckNameMonitoringAgentCheck     HealthCheckName = "MonitoringAgentCheck"
	HealthCheckNameSupportedEncryptionCheck HealthCheckName = "SupportedEncryptionCheck"
	HealthCheckNameSxSStackListenerCheck    HealthCheckName = "SxSStackListenerCheck"
	HealthCheckNameUrlsAccessibleCheck      HealthCheckName = "UrlsAccessibleCheck"
	HealthCheckNameWebRTCRedirectorCheck    HealthCheckName = "WebRTCRedirectorCheck"
)

func PossibleValuesForHealthCheckName() []string {
	return []string{
		string(HealthCheckNameAppAttachHealthCheck),
		string(HealthCheckNameDomainJoinedCheck),
		string(HealthCheckNameDomainReachable),
		string(HealthCheckNameDomainTrustCheck),
		string(HealthCheckNameFSLogixHealthCheck),
		string(HealthCheckNameMetaDataServiceCheck),
		string(HealthCheckNameMonitoringAgentCheck),
		string(HealthCheckNameSupportedEncryptionCheck),
		string(HealthCheckNameSxSStackListenerCheck),
		string(HealthCheckNameUrlsAccessibleCheck),
		string(HealthCheckNameWebRTCRedirectorCheck),
	}
}

type HealthCheckResult string

const (
	HealthCheckResultHealthCheckFailed    HealthCheckResult = "HealthCheckFailed"
	HealthCheckResultHealthCheckSucceeded HealthCheckResult = "HealthCheckSucceeded"
	HealthCheckResultSessionHostShutdown  HealthCheckResult = "SessionHostShutdown"
	HealthCheckResultUnknown              HealthCheckResult = "Unknown"
)

func PossibleValuesForHealthCheckResult() []string {
	return []string{
		string(HealthCheckResultHealthCheckFailed),
		string(HealthCheckResultHealthCheckSucceeded),
		string(HealthCheckResultSessionHostShutdown),
		string(HealthCheckResultUnknown),
	}
}

type Status string

const (
	StatusAvailable                   Status = "Available"
	StatusDisconnected                Status = "Disconnected"
	StatusDomainTrustRelationshipLost Status = "DomainTrustRelationshipLost"
	StatusFSLogixNotHealthy           Status = "FSLogixNotHealthy"
	StatusNeedsAssistance             Status = "NeedsAssistance"
	StatusNoHeartbeat                 Status = "NoHeartbeat"
	StatusNotJoinedToDomain           Status = "NotJoinedToDomain"
	StatusShutdown                    Status = "Shutdown"
	StatusSxSStackListenerNotReady    Status = "SxSStackListenerNotReady"
	StatusUnavailable                 Status = "Unavailable"
	StatusUpgradeFailed               Status = "UpgradeFailed"
	StatusUpgrading                   Status = "Upgrading"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAvailable),
		string(StatusDisconnected),
		string(StatusDomainTrustRelationshipLost),
		string(StatusFSLogixNotHealthy),
		string(StatusNeedsAssistance),
		string(StatusNoHeartbeat),
		string(StatusNotJoinedToDomain),
		string(StatusShutdown),
		string(StatusSxSStackListenerNotReady),
		string(StatusUnavailable),
		string(StatusUpgradeFailed),
		string(StatusUpgrading),
	}
}

type UpdateState string

const (
	UpdateStateFailed    UpdateState = "Failed"
	UpdateStateInitial   UpdateState = "Initial"
	UpdateStatePending   UpdateState = "Pending"
	UpdateStateStarted   UpdateState = "Started"
	UpdateStateSucceeded UpdateState = "Succeeded"
)

func PossibleValuesForUpdateState() []string {
	return []string{
		string(UpdateStateFailed),
		string(UpdateStateInitial),
		string(UpdateStatePending),
		string(UpdateStateStarted),
		string(UpdateStateSucceeded),
	}
}
