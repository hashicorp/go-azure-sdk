package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialType string

const (
	CredentialTypeAzureKeyVaultSThreeWithHMAC CredentialType = "AzureKeyVaultS3WithHMAC"
	CredentialTypeAzureKeyVaultSmb            CredentialType = "AzureKeyVaultSmb"
)

func PossibleValuesForCredentialType() []string {
	return []string{
		string(CredentialTypeAzureKeyVaultSThreeWithHMAC),
		string(CredentialTypeAzureKeyVaultSmb),
	}
}

func (s *CredentialType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCredentialType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCredentialType(input string) (*CredentialType, error) {
	vals := map[string]CredentialType{
		"azurekeyvaults3withhmac": CredentialTypeAzureKeyVaultSThreeWithHMAC,
		"azurekeyvaultsmb":        CredentialTypeAzureKeyVaultSmb,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialType(input)
	return &out, nil
}

type EndpointKind string

const (
	EndpointKindSource EndpointKind = "Source"
	EndpointKindTarget EndpointKind = "Target"
)

func PossibleValuesForEndpointKind() []string {
	return []string{
		string(EndpointKindSource),
		string(EndpointKindTarget),
	}
}

func (s *EndpointKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointKind(input string) (*EndpointKind, error) {
	vals := map[string]EndpointKind{
		"source": EndpointKindSource,
		"target": EndpointKindTarget,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointKind(input)
	return &out, nil
}

type EndpointType string

const (
	EndpointTypeAzureMultiCloudConnector  EndpointType = "AzureMultiCloudConnector"
	EndpointTypeAzureStorageBlobContainer EndpointType = "AzureStorageBlobContainer"
	EndpointTypeAzureStorageNfsFileShare  EndpointType = "AzureStorageNfsFileShare"
	EndpointTypeAzureStorageSmbFileShare  EndpointType = "AzureStorageSmbFileShare"
	EndpointTypeNfsMount                  EndpointType = "NfsMount"
	EndpointTypeSThreeWithHMAC            EndpointType = "S3WithHMAC"
	EndpointTypeSmbMount                  EndpointType = "SmbMount"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeAzureMultiCloudConnector),
		string(EndpointTypeAzureStorageBlobContainer),
		string(EndpointTypeAzureStorageNfsFileShare),
		string(EndpointTypeAzureStorageSmbFileShare),
		string(EndpointTypeNfsMount),
		string(EndpointTypeSThreeWithHMAC),
		string(EndpointTypeSmbMount),
	}
}

func (s *EndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointType(input string) (*EndpointType, error) {
	vals := map[string]EndpointType{
		"azuremulticloudconnector":  EndpointTypeAzureMultiCloudConnector,
		"azurestorageblobcontainer": EndpointTypeAzureStorageBlobContainer,
		"azurestoragenfsfileshare":  EndpointTypeAzureStorageNfsFileShare,
		"azurestoragesmbfileshare":  EndpointTypeAzureStorageSmbFileShare,
		"nfsmount":                  EndpointTypeNfsMount,
		"s3withhmac":                EndpointTypeSThreeWithHMAC,
		"smbmount":                  EndpointTypeSmbMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointType(input)
	return &out, nil
}

type NfsMountSourceType string

const (
	NfsMountSourceTypeFSXNegativeEFS NfsMountSourceType = "FSX-EFS"
	NfsMountSourceTypeNfsMount       NfsMountSourceType = "NfsMount"
)

func PossibleValuesForNfsMountSourceType() []string {
	return []string{
		string(NfsMountSourceTypeFSXNegativeEFS),
		string(NfsMountSourceTypeNfsMount),
	}
}

func (s *NfsMountSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNfsMountSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNfsMountSourceType(input string) (*NfsMountSourceType, error) {
	vals := map[string]NfsMountSourceType{
		"fsx-efs":  NfsMountSourceTypeFSXNegativeEFS,
		"nfsmount": NfsMountSourceTypeNfsMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NfsMountSourceType(input)
	return &out, nil
}

type NfsVersion string

const (
	NfsVersionNFSauto     NfsVersion = "NFSauto"
	NfsVersionNFSvFour    NfsVersion = "NFSv4"
	NfsVersionNFSvFourOne NfsVersion = "NFSv4_1"
	NfsVersionNFSvThree   NfsVersion = "NFSv3"
)

func PossibleValuesForNfsVersion() []string {
	return []string{
		string(NfsVersionNFSauto),
		string(NfsVersionNFSvFour),
		string(NfsVersionNFSvFourOne),
		string(NfsVersionNFSvThree),
	}
}

func (s *NfsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNfsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNfsVersion(input string) (*NfsVersion, error) {
	vals := map[string]NfsVersion{
		"nfsauto": NfsVersionNFSauto,
		"nfsv4":   NfsVersionNFSvFour,
		"nfsv4_1": NfsVersionNFSvFourOne,
		"nfsv3":   NfsVersionNFSvThree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NfsVersion(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type S3WithHmacSourceType string

const (
	S3WithHmacSourceTypeALIBABA S3WithHmacSourceType = "ALIBABA"
	S3WithHmacSourceTypeDELLEMC S3WithHmacSourceType = "DELL_EMC"
	S3WithHmacSourceTypeGCS     S3WithHmacSourceType = "GCS"
	S3WithHmacSourceTypeIBM     S3WithHmacSourceType = "IBM"
	S3WithHmacSourceTypeMINIO   S3WithHmacSourceType = "MINIO"
	S3WithHmacSourceTypeOTHER   S3WithHmacSourceType = "OTHER"
)

func PossibleValuesForS3WithHmacSourceType() []string {
	return []string{
		string(S3WithHmacSourceTypeALIBABA),
		string(S3WithHmacSourceTypeDELLEMC),
		string(S3WithHmacSourceTypeGCS),
		string(S3WithHmacSourceTypeIBM),
		string(S3WithHmacSourceTypeMINIO),
		string(S3WithHmacSourceTypeOTHER),
	}
}

func (s *S3WithHmacSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseS3WithHmacSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseS3WithHmacSourceType(input string) (*S3WithHmacSourceType, error) {
	vals := map[string]S3WithHmacSourceType{
		"alibaba":  S3WithHmacSourceTypeALIBABA,
		"dell_emc": S3WithHmacSourceTypeDELLEMC,
		"gcs":      S3WithHmacSourceTypeGCS,
		"ibm":      S3WithHmacSourceTypeIBM,
		"minio":    S3WithHmacSourceTypeMINIO,
		"other":    S3WithHmacSourceTypeOTHER,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := S3WithHmacSourceType(input)
	return &out, nil
}

type SmbMountSourceType string

const (
	SmbMountSourceTypeFSXNegativeSMB SmbMountSourceType = "FSX-SMB"
	SmbMountSourceTypeSmbMount       SmbMountSourceType = "SmbMount"
)

func PossibleValuesForSmbMountSourceType() []string {
	return []string{
		string(SmbMountSourceTypeFSXNegativeSMB),
		string(SmbMountSourceTypeSmbMount),
	}
}

func (s *SmbMountSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmbMountSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmbMountSourceType(input string) (*SmbMountSourceType, error) {
	vals := map[string]SmbMountSourceType{
		"fsx-smb":  SmbMountSourceTypeFSXNegativeSMB,
		"smbmount": SmbMountSourceTypeSmbMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmbMountSourceType(input)
	return &out, nil
}
