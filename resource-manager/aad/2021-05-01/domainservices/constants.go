package domainservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalAccess string

const (
	ExternalAccessDisabled ExternalAccess = "Disabled"
	ExternalAccessEnabled  ExternalAccess = "Enabled"
)

func PossibleValuesForExternalAccess() []string {
	return []string{
		string(ExternalAccessDisabled),
		string(ExternalAccessEnabled),
	}
}

type FilteredSync string

const (
	FilteredSyncDisabled FilteredSync = "Disabled"
	FilteredSyncEnabled  FilteredSync = "Enabled"
)

func PossibleValuesForFilteredSync() []string {
	return []string{
		string(FilteredSyncDisabled),
		string(FilteredSyncEnabled),
	}
}

type KerberosArmoring string

const (
	KerberosArmoringDisabled KerberosArmoring = "Disabled"
	KerberosArmoringEnabled  KerberosArmoring = "Enabled"
)

func PossibleValuesForKerberosArmoring() []string {
	return []string{
		string(KerberosArmoringDisabled),
		string(KerberosArmoringEnabled),
	}
}

type KerberosRc4Encryption string

const (
	KerberosRc4EncryptionDisabled KerberosRc4Encryption = "Disabled"
	KerberosRc4EncryptionEnabled  KerberosRc4Encryption = "Enabled"
)

func PossibleValuesForKerberosRc4Encryption() []string {
	return []string{
		string(KerberosRc4EncryptionDisabled),
		string(KerberosRc4EncryptionEnabled),
	}
}

type Ldaps string

const (
	LdapsDisabled Ldaps = "Disabled"
	LdapsEnabled  Ldaps = "Enabled"
)

func PossibleValuesForLdaps() []string {
	return []string{
		string(LdapsDisabled),
		string(LdapsEnabled),
	}
}

type NotifyDcAdmins string

const (
	NotifyDcAdminsDisabled NotifyDcAdmins = "Disabled"
	NotifyDcAdminsEnabled  NotifyDcAdmins = "Enabled"
)

func PossibleValuesForNotifyDcAdmins() []string {
	return []string{
		string(NotifyDcAdminsDisabled),
		string(NotifyDcAdminsEnabled),
	}
}

type NotifyGlobalAdmins string

const (
	NotifyGlobalAdminsDisabled NotifyGlobalAdmins = "Disabled"
	NotifyGlobalAdminsEnabled  NotifyGlobalAdmins = "Enabled"
)

func PossibleValuesForNotifyGlobalAdmins() []string {
	return []string{
		string(NotifyGlobalAdminsDisabled),
		string(NotifyGlobalAdminsEnabled),
	}
}

type NtlmV1 string

const (
	NtlmV1Disabled NtlmV1 = "Disabled"
	NtlmV1Enabled  NtlmV1 = "Enabled"
)

func PossibleValuesForNtlmV1() []string {
	return []string{
		string(NtlmV1Disabled),
		string(NtlmV1Enabled),
	}
}

type Status string

const (
	StatusFailure Status = "Failure"
	StatusNone    Status = "None"
	StatusOK      Status = "OK"
	StatusRunning Status = "Running"
	StatusSkipped Status = "Skipped"
	StatusWarning Status = "Warning"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusFailure),
		string(StatusNone),
		string(StatusOK),
		string(StatusRunning),
		string(StatusSkipped),
		string(StatusWarning),
	}
}

type SyncKerberosPasswords string

const (
	SyncKerberosPasswordsDisabled SyncKerberosPasswords = "Disabled"
	SyncKerberosPasswordsEnabled  SyncKerberosPasswords = "Enabled"
)

func PossibleValuesForSyncKerberosPasswords() []string {
	return []string{
		string(SyncKerberosPasswordsDisabled),
		string(SyncKerberosPasswordsEnabled),
	}
}

type SyncNtlmPasswords string

const (
	SyncNtlmPasswordsDisabled SyncNtlmPasswords = "Disabled"
	SyncNtlmPasswordsEnabled  SyncNtlmPasswords = "Enabled"
)

func PossibleValuesForSyncNtlmPasswords() []string {
	return []string{
		string(SyncNtlmPasswordsDisabled),
		string(SyncNtlmPasswordsEnabled),
	}
}

type SyncOnPremPasswords string

const (
	SyncOnPremPasswordsDisabled SyncOnPremPasswords = "Disabled"
	SyncOnPremPasswordsEnabled  SyncOnPremPasswords = "Enabled"
)

func PossibleValuesForSyncOnPremPasswords() []string {
	return []string{
		string(SyncOnPremPasswordsDisabled),
		string(SyncOnPremPasswordsEnabled),
	}
}

type TlsV1 string

const (
	TlsV1Disabled TlsV1 = "Disabled"
	TlsV1Enabled  TlsV1 = "Enabled"
)

func PossibleValuesForTlsV1() []string {
	return []string{
		string(TlsV1Disabled),
		string(TlsV1Enabled),
	}
}
