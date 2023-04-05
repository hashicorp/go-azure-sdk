package roleeligibilityschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MemberType string

const (
	MemberTypeDirect    MemberType = "Direct"
	MemberTypeGroup     MemberType = "Group"
	MemberTypeInherited MemberType = "Inherited"
)

func PossibleValuesForMemberType() []string {
	return []string{
		string(MemberTypeDirect),
		string(MemberTypeGroup),
		string(MemberTypeInherited),
	}
}

type PrincipalType string

const (
	PrincipalTypeDevice           PrincipalType = "Device"
	PrincipalTypeForeignGroup     PrincipalType = "ForeignGroup"
	PrincipalTypeGroup            PrincipalType = "Group"
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	PrincipalTypeUser             PrincipalType = "User"
)

func PossibleValuesForPrincipalType() []string {
	return []string{
		string(PrincipalTypeDevice),
		string(PrincipalTypeForeignGroup),
		string(PrincipalTypeGroup),
		string(PrincipalTypeServicePrincipal),
		string(PrincipalTypeUser),
	}
}

type Status string

const (
	StatusAccepted                    Status = "Accepted"
	StatusAdminApproved               Status = "AdminApproved"
	StatusAdminDenied                 Status = "AdminDenied"
	StatusCanceled                    Status = "Canceled"
	StatusDenied                      Status = "Denied"
	StatusFailed                      Status = "Failed"
	StatusFailedAsResourceIsLocked    Status = "FailedAsResourceIsLocked"
	StatusGranted                     Status = "Granted"
	StatusInvalid                     Status = "Invalid"
	StatusPendingAdminDecision        Status = "PendingAdminDecision"
	StatusPendingApproval             Status = "PendingApproval"
	StatusPendingApprovalProvisioning Status = "PendingApprovalProvisioning"
	StatusPendingEvaluation           Status = "PendingEvaluation"
	StatusPendingExternalProvisioning Status = "PendingExternalProvisioning"
	StatusPendingProvisioning         Status = "PendingProvisioning"
	StatusPendingRevocation           Status = "PendingRevocation"
	StatusPendingScheduleCreation     Status = "PendingScheduleCreation"
	StatusProvisioned                 Status = "Provisioned"
	StatusProvisioningStarted         Status = "ProvisioningStarted"
	StatusRevoked                     Status = "Revoked"
	StatusScheduleCreated             Status = "ScheduleCreated"
	StatusTimedOut                    Status = "TimedOut"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAccepted),
		string(StatusAdminApproved),
		string(StatusAdminDenied),
		string(StatusCanceled),
		string(StatusDenied),
		string(StatusFailed),
		string(StatusFailedAsResourceIsLocked),
		string(StatusGranted),
		string(StatusInvalid),
		string(StatusPendingAdminDecision),
		string(StatusPendingApproval),
		string(StatusPendingApprovalProvisioning),
		string(StatusPendingEvaluation),
		string(StatusPendingExternalProvisioning),
		string(StatusPendingProvisioning),
		string(StatusPendingRevocation),
		string(StatusPendingScheduleCreation),
		string(StatusProvisioned),
		string(StatusProvisioningStarted),
		string(StatusRevoked),
		string(StatusScheduleCreated),
		string(StatusTimedOut),
	}
}
