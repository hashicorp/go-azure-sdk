package roleeligibilityschedulerequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type RequestType string

const (
	RequestTypeAdminAssign    RequestType = "AdminAssign"
	RequestTypeAdminExtend    RequestType = "AdminExtend"
	RequestTypeAdminRemove    RequestType = "AdminRemove"
	RequestTypeAdminRenew     RequestType = "AdminRenew"
	RequestTypeAdminUpdate    RequestType = "AdminUpdate"
	RequestTypeSelfActivate   RequestType = "SelfActivate"
	RequestTypeSelfDeactivate RequestType = "SelfDeactivate"
	RequestTypeSelfExtend     RequestType = "SelfExtend"
	RequestTypeSelfRenew      RequestType = "SelfRenew"
)

func PossibleValuesForRequestType() []string {
	return []string{
		string(RequestTypeAdminAssign),
		string(RequestTypeAdminExtend),
		string(RequestTypeAdminRemove),
		string(RequestTypeAdminRenew),
		string(RequestTypeAdminUpdate),
		string(RequestTypeSelfActivate),
		string(RequestTypeSelfDeactivate),
		string(RequestTypeSelfExtend),
		string(RequestTypeSelfRenew),
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

type Type string

const (
	TypeAfterDateTime Type = "AfterDateTime"
	TypeAfterDuration Type = "AfterDuration"
	TypeNoExpiration  Type = "NoExpiration"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeAfterDateTime),
		string(TypeAfterDuration),
		string(TypeNoExpiration),
	}
}
