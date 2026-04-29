package caches

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheLifeCycleState string

const (
	CacheLifeCycleStateClusterPeeringOfferSent CacheLifeCycleState = "ClusterPeeringOfferSent"
	CacheLifeCycleStateCreating                CacheLifeCycleState = "Creating"
	CacheLifeCycleStateFailed                  CacheLifeCycleState = "Failed"
	CacheLifeCycleStateSucceeded               CacheLifeCycleState = "Succeeded"
	CacheLifeCycleStateVserverPeeringOfferSent CacheLifeCycleState = "VserverPeeringOfferSent"
)

func PossibleValuesForCacheLifeCycleState() []string {
	return []string{
		string(CacheLifeCycleStateClusterPeeringOfferSent),
		string(CacheLifeCycleStateCreating),
		string(CacheLifeCycleStateFailed),
		string(CacheLifeCycleStateSucceeded),
		string(CacheLifeCycleStateVserverPeeringOfferSent),
	}
}

func (s *CacheLifeCycleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCacheLifeCycleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCacheLifeCycleState(input string) (*CacheLifeCycleState, error) {
	vals := map[string]CacheLifeCycleState{
		"clusterpeeringoffersent": CacheLifeCycleStateClusterPeeringOfferSent,
		"creating":                CacheLifeCycleStateCreating,
		"failed":                  CacheLifeCycleStateFailed,
		"succeeded":               CacheLifeCycleStateSucceeded,
		"vserverpeeringoffersent": CacheLifeCycleStateVserverPeeringOfferSent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CacheLifeCycleState(input)
	return &out, nil
}

type CacheProvisioningState string

const (
	CacheProvisioningStateCanceled  CacheProvisioningState = "Canceled"
	CacheProvisioningStateCreating  CacheProvisioningState = "Creating"
	CacheProvisioningStateDeleting  CacheProvisioningState = "Deleting"
	CacheProvisioningStateFailed    CacheProvisioningState = "Failed"
	CacheProvisioningStateSucceeded CacheProvisioningState = "Succeeded"
	CacheProvisioningStateUpdating  CacheProvisioningState = "Updating"
)

func PossibleValuesForCacheProvisioningState() []string {
	return []string{
		string(CacheProvisioningStateCanceled),
		string(CacheProvisioningStateCreating),
		string(CacheProvisioningStateDeleting),
		string(CacheProvisioningStateFailed),
		string(CacheProvisioningStateSucceeded),
		string(CacheProvisioningStateUpdating),
	}
}

func (s *CacheProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCacheProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCacheProvisioningState(input string) (*CacheProvisioningState, error) {
	vals := map[string]CacheProvisioningState{
		"canceled":  CacheProvisioningStateCanceled,
		"creating":  CacheProvisioningStateCreating,
		"deleting":  CacheProvisioningStateDeleting,
		"failed":    CacheProvisioningStateFailed,
		"succeeded": CacheProvisioningStateSucceeded,
		"updating":  CacheProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CacheProvisioningState(input)
	return &out, nil
}

type ChownMode string

const (
	ChownModeRestricted   ChownMode = "Restricted"
	ChownModeUnrestricted ChownMode = "Unrestricted"
)

func PossibleValuesForChownMode() []string {
	return []string{
		string(ChownModeRestricted),
		string(ChownModeUnrestricted),
	}
}

func (s *ChownMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChownMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChownMode(input string) (*ChownMode, error) {
	vals := map[string]ChownMode{
		"restricted":   ChownModeRestricted,
		"unrestricted": ChownModeUnrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChownMode(input)
	return &out, nil
}

type CifsChangeNotifyState string

const (
	CifsChangeNotifyStateDisabled CifsChangeNotifyState = "Disabled"
	CifsChangeNotifyStateEnabled  CifsChangeNotifyState = "Enabled"
)

func PossibleValuesForCifsChangeNotifyState() []string {
	return []string{
		string(CifsChangeNotifyStateDisabled),
		string(CifsChangeNotifyStateEnabled),
	}
}

func (s *CifsChangeNotifyState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCifsChangeNotifyState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCifsChangeNotifyState(input string) (*CifsChangeNotifyState, error) {
	vals := map[string]CifsChangeNotifyState{
		"disabled": CifsChangeNotifyStateDisabled,
		"enabled":  CifsChangeNotifyStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CifsChangeNotifyState(input)
	return &out, nil
}

type EnableWriteBackState string

const (
	EnableWriteBackStateDisabled EnableWriteBackState = "Disabled"
	EnableWriteBackStateEnabled  EnableWriteBackState = "Enabled"
)

func PossibleValuesForEnableWriteBackState() []string {
	return []string{
		string(EnableWriteBackStateDisabled),
		string(EnableWriteBackStateEnabled),
	}
}

func (s *EnableWriteBackState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnableWriteBackState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnableWriteBackState(input string) (*EnableWriteBackState, error) {
	vals := map[string]EnableWriteBackState{
		"disabled": EnableWriteBackStateDisabled,
		"enabled":  EnableWriteBackStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnableWriteBackState(input)
	return &out, nil
}

type EncryptionKeySource string

const (
	EncryptionKeySourceMicrosoftPointKeyVault EncryptionKeySource = "Microsoft.KeyVault"
	EncryptionKeySourceMicrosoftPointNetApp   EncryptionKeySource = "Microsoft.NetApp"
)

func PossibleValuesForEncryptionKeySource() []string {
	return []string{
		string(EncryptionKeySourceMicrosoftPointKeyVault),
		string(EncryptionKeySourceMicrosoftPointNetApp),
	}
}

func (s *EncryptionKeySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionKeySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionKeySource(input string) (*EncryptionKeySource, error) {
	vals := map[string]EncryptionKeySource{
		"microsoft.keyvault": EncryptionKeySourceMicrosoftPointKeyVault,
		"microsoft.netapp":   EncryptionKeySourceMicrosoftPointNetApp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionKeySource(input)
	return &out, nil
}

type EncryptionState string

const (
	EncryptionStateDisabled EncryptionState = "Disabled"
	EncryptionStateEnabled  EncryptionState = "Enabled"
)

func PossibleValuesForEncryptionState() []string {
	return []string{
		string(EncryptionStateDisabled),
		string(EncryptionStateEnabled),
	}
}

func (s *EncryptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptionState(input string) (*EncryptionState, error) {
	vals := map[string]EncryptionState{
		"disabled": EncryptionStateDisabled,
		"enabled":  EncryptionStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionState(input)
	return &out, nil
}

type GlobalFileLockingState string

const (
	GlobalFileLockingStateDisabled GlobalFileLockingState = "Disabled"
	GlobalFileLockingStateEnabled  GlobalFileLockingState = "Enabled"
)

func PossibleValuesForGlobalFileLockingState() []string {
	return []string{
		string(GlobalFileLockingStateDisabled),
		string(GlobalFileLockingStateEnabled),
	}
}

func (s *GlobalFileLockingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGlobalFileLockingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGlobalFileLockingState(input string) (*GlobalFileLockingState, error) {
	vals := map[string]GlobalFileLockingState{
		"disabled": GlobalFileLockingStateDisabled,
		"enabled":  GlobalFileLockingStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GlobalFileLockingState(input)
	return &out, nil
}

type KerberosState string

const (
	KerberosStateDisabled KerberosState = "Disabled"
	KerberosStateEnabled  KerberosState = "Enabled"
)

func PossibleValuesForKerberosState() []string {
	return []string{
		string(KerberosStateDisabled),
		string(KerberosStateEnabled),
	}
}

func (s *KerberosState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKerberosState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKerberosState(input string) (*KerberosState, error) {
	vals := map[string]KerberosState{
		"disabled": KerberosStateDisabled,
		"enabled":  KerberosStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KerberosState(input)
	return &out, nil
}

type LdapServerType string

const (
	LdapServerTypeActiveDirectory LdapServerType = "ActiveDirectory"
	LdapServerTypeOpenLDAP        LdapServerType = "OpenLDAP"
)

func PossibleValuesForLdapServerType() []string {
	return []string{
		string(LdapServerTypeActiveDirectory),
		string(LdapServerTypeOpenLDAP),
	}
}

func (s *LdapServerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLdapServerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLdapServerType(input string) (*LdapServerType, error) {
	vals := map[string]LdapServerType{
		"activedirectory": LdapServerTypeActiveDirectory,
		"openldap":        LdapServerTypeOpenLDAP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LdapServerType(input)
	return &out, nil
}

type LdapState string

const (
	LdapStateDisabled LdapState = "Disabled"
	LdapStateEnabled  LdapState = "Enabled"
)

func PossibleValuesForLdapState() []string {
	return []string{
		string(LdapStateDisabled),
		string(LdapStateEnabled),
	}
}

func (s *LdapState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLdapState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLdapState(input string) (*LdapState, error) {
	vals := map[string]LdapState{
		"disabled": LdapStateDisabled,
		"enabled":  LdapStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LdapState(input)
	return &out, nil
}

type ProtocolTypes string

const (
	ProtocolTypesNFSvFour  ProtocolTypes = "NFSv4"
	ProtocolTypesNFSvThree ProtocolTypes = "NFSv3"
	ProtocolTypesSMB       ProtocolTypes = "SMB"
)

func PossibleValuesForProtocolTypes() []string {
	return []string{
		string(ProtocolTypesNFSvFour),
		string(ProtocolTypesNFSvThree),
		string(ProtocolTypesSMB),
	}
}

func (s *ProtocolTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtocolTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtocolTypes(input string) (*ProtocolTypes, error) {
	vals := map[string]ProtocolTypes{
		"nfsv4": ProtocolTypesNFSvFour,
		"nfsv3": ProtocolTypesNFSvThree,
		"smb":   ProtocolTypesSMB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtocolTypes(input)
	return &out, nil
}

type SmbAccessBasedEnumeration string

const (
	SmbAccessBasedEnumerationDisabled SmbAccessBasedEnumeration = "Disabled"
	SmbAccessBasedEnumerationEnabled  SmbAccessBasedEnumeration = "Enabled"
)

func PossibleValuesForSmbAccessBasedEnumeration() []string {
	return []string{
		string(SmbAccessBasedEnumerationDisabled),
		string(SmbAccessBasedEnumerationEnabled),
	}
}

func (s *SmbAccessBasedEnumeration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmbAccessBasedEnumeration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmbAccessBasedEnumeration(input string) (*SmbAccessBasedEnumeration, error) {
	vals := map[string]SmbAccessBasedEnumeration{
		"disabled": SmbAccessBasedEnumerationDisabled,
		"enabled":  SmbAccessBasedEnumerationEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmbAccessBasedEnumeration(input)
	return &out, nil
}

type SmbEncryptionState string

const (
	SmbEncryptionStateDisabled SmbEncryptionState = "Disabled"
	SmbEncryptionStateEnabled  SmbEncryptionState = "Enabled"
)

func PossibleValuesForSmbEncryptionState() []string {
	return []string{
		string(SmbEncryptionStateDisabled),
		string(SmbEncryptionStateEnabled),
	}
}

func (s *SmbEncryptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmbEncryptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmbEncryptionState(input string) (*SmbEncryptionState, error) {
	vals := map[string]SmbEncryptionState{
		"disabled": SmbEncryptionStateDisabled,
		"enabled":  SmbEncryptionStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmbEncryptionState(input)
	return &out, nil
}

type SmbNonBrowsable string

const (
	SmbNonBrowsableDisabled SmbNonBrowsable = "Disabled"
	SmbNonBrowsableEnabled  SmbNonBrowsable = "Enabled"
)

func PossibleValuesForSmbNonBrowsable() []string {
	return []string{
		string(SmbNonBrowsableDisabled),
		string(SmbNonBrowsableEnabled),
	}
}

func (s *SmbNonBrowsable) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSmbNonBrowsable(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSmbNonBrowsable(input string) (*SmbNonBrowsable, error) {
	vals := map[string]SmbNonBrowsable{
		"disabled": SmbNonBrowsableDisabled,
		"enabled":  SmbNonBrowsableEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SmbNonBrowsable(input)
	return &out, nil
}

type VolumeLanguage string

const (
	VolumeLanguageAr                                                    VolumeLanguage = "ar"
	VolumeLanguageArPointutfNegativeEight                               VolumeLanguage = "ar.utf-8"
	VolumeLanguageC                                                     VolumeLanguage = "c"
	VolumeLanguageCPointutfNegativeEight                                VolumeLanguage = "c.utf-8"
	VolumeLanguageCs                                                    VolumeLanguage = "cs"
	VolumeLanguageCsPointutfNegativeEight                               VolumeLanguage = "cs.utf-8"
	VolumeLanguageDa                                                    VolumeLanguage = "da"
	VolumeLanguageDaPointutfNegativeEight                               VolumeLanguage = "da.utf-8"
	VolumeLanguageDe                                                    VolumeLanguage = "de"
	VolumeLanguageDePointutfNegativeEight                               VolumeLanguage = "de.utf-8"
	VolumeLanguageEn                                                    VolumeLanguage = "en"
	VolumeLanguageEnNegativeus                                          VolumeLanguage = "en-us"
	VolumeLanguageEnNegativeusPointutfNegativeEight                     VolumeLanguage = "en-us.utf-8"
	VolumeLanguageEnPointutfNegativeEight                               VolumeLanguage = "en.utf-8"
	VolumeLanguageEs                                                    VolumeLanguage = "es"
	VolumeLanguageEsPointutfNegativeEight                               VolumeLanguage = "es.utf-8"
	VolumeLanguageFi                                                    VolumeLanguage = "fi"
	VolumeLanguageFiPointutfNegativeEight                               VolumeLanguage = "fi.utf-8"
	VolumeLanguageFr                                                    VolumeLanguage = "fr"
	VolumeLanguageFrPointutfNegativeEight                               VolumeLanguage = "fr.utf-8"
	VolumeLanguageHe                                                    VolumeLanguage = "he"
	VolumeLanguageHePointutfNegativeEight                               VolumeLanguage = "he.utf-8"
	VolumeLanguageHr                                                    VolumeLanguage = "hr"
	VolumeLanguageHrPointutfNegativeEight                               VolumeLanguage = "hr.utf-8"
	VolumeLanguageHu                                                    VolumeLanguage = "hu"
	VolumeLanguageHuPointutfNegativeEight                               VolumeLanguage = "hu.utf-8"
	VolumeLanguageIt                                                    VolumeLanguage = "it"
	VolumeLanguageItPointutfNegativeEight                               VolumeLanguage = "it.utf-8"
	VolumeLanguageJa                                                    VolumeLanguage = "ja"
	VolumeLanguageJaNegativejpPointNineThreeTwo                         VolumeLanguage = "ja-jp.932"
	VolumeLanguageJaNegativejpPointNineThreeTwoPointutfNegativeEight    VolumeLanguage = "ja-jp.932.utf-8"
	VolumeLanguageJaNegativejpPointpck                                  VolumeLanguage = "ja-jp.pck"
	VolumeLanguageJaNegativejpPointpckNegativevTwo                      VolumeLanguage = "ja-jp.pck-v2"
	VolumeLanguageJaNegativejpPointpckNegativevTwoPointutfNegativeEight VolumeLanguage = "ja-jp.pck-v2.utf-8"
	VolumeLanguageJaNegativejpPointpckPointutfNegativeEight             VolumeLanguage = "ja-jp.pck.utf-8"
	VolumeLanguageJaNegativevOne                                        VolumeLanguage = "ja-v1"
	VolumeLanguageJaNegativevOnePointutfNegativeEight                   VolumeLanguage = "ja-v1.utf-8"
	VolumeLanguageJaPointutfNegativeEight                               VolumeLanguage = "ja.utf-8"
	VolumeLanguageKo                                                    VolumeLanguage = "ko"
	VolumeLanguageKoPointutfNegativeEight                               VolumeLanguage = "ko.utf-8"
	VolumeLanguageNl                                                    VolumeLanguage = "nl"
	VolumeLanguageNlPointutfNegativeEight                               VolumeLanguage = "nl.utf-8"
	VolumeLanguageNo                                                    VolumeLanguage = "no"
	VolumeLanguageNoPointutfNegativeEight                               VolumeLanguage = "no.utf-8"
	VolumeLanguagePl                                                    VolumeLanguage = "pl"
	VolumeLanguagePlPointutfNegativeEight                               VolumeLanguage = "pl.utf-8"
	VolumeLanguagePt                                                    VolumeLanguage = "pt"
	VolumeLanguagePtPointutfNegativeEight                               VolumeLanguage = "pt.utf-8"
	VolumeLanguageRo                                                    VolumeLanguage = "ro"
	VolumeLanguageRoPointutfNegativeEight                               VolumeLanguage = "ro.utf-8"
	VolumeLanguageRu                                                    VolumeLanguage = "ru"
	VolumeLanguageRuPointutfNegativeEight                               VolumeLanguage = "ru.utf-8"
	VolumeLanguageSk                                                    VolumeLanguage = "sk"
	VolumeLanguageSkPointutfNegativeEight                               VolumeLanguage = "sk.utf-8"
	VolumeLanguageSl                                                    VolumeLanguage = "sl"
	VolumeLanguageSlPointutfNegativeEight                               VolumeLanguage = "sl.utf-8"
	VolumeLanguageSv                                                    VolumeLanguage = "sv"
	VolumeLanguageSvPointutfNegativeEight                               VolumeLanguage = "sv.utf-8"
	VolumeLanguageTr                                                    VolumeLanguage = "tr"
	VolumeLanguageTrPointutfNegativeEight                               VolumeLanguage = "tr.utf-8"
	VolumeLanguageUtfEightmbFour                                        VolumeLanguage = "utf8mb4"
	VolumeLanguageZh                                                    VolumeLanguage = "zh"
	VolumeLanguageZhNegativetw                                          VolumeLanguage = "zh-tw"
	VolumeLanguageZhNegativetwPointbigFive                              VolumeLanguage = "zh-tw.big5"
	VolumeLanguageZhNegativetwPointbigFivePointutfNegativeEight         VolumeLanguage = "zh-tw.big5.utf-8"
	VolumeLanguageZhNegativetwPointutfNegativeEight                     VolumeLanguage = "zh-tw.utf-8"
	VolumeLanguageZhPointgbk                                            VolumeLanguage = "zh.gbk"
	VolumeLanguageZhPointgbkPointutfNegativeEight                       VolumeLanguage = "zh.gbk.utf-8"
	VolumeLanguageZhPointutfNegativeEight                               VolumeLanguage = "zh.utf-8"
)

func PossibleValuesForVolumeLanguage() []string {
	return []string{
		string(VolumeLanguageAr),
		string(VolumeLanguageArPointutfNegativeEight),
		string(VolumeLanguageC),
		string(VolumeLanguageCPointutfNegativeEight),
		string(VolumeLanguageCs),
		string(VolumeLanguageCsPointutfNegativeEight),
		string(VolumeLanguageDa),
		string(VolumeLanguageDaPointutfNegativeEight),
		string(VolumeLanguageDe),
		string(VolumeLanguageDePointutfNegativeEight),
		string(VolumeLanguageEn),
		string(VolumeLanguageEnNegativeus),
		string(VolumeLanguageEnNegativeusPointutfNegativeEight),
		string(VolumeLanguageEnPointutfNegativeEight),
		string(VolumeLanguageEs),
		string(VolumeLanguageEsPointutfNegativeEight),
		string(VolumeLanguageFi),
		string(VolumeLanguageFiPointutfNegativeEight),
		string(VolumeLanguageFr),
		string(VolumeLanguageFrPointutfNegativeEight),
		string(VolumeLanguageHe),
		string(VolumeLanguageHePointutfNegativeEight),
		string(VolumeLanguageHr),
		string(VolumeLanguageHrPointutfNegativeEight),
		string(VolumeLanguageHu),
		string(VolumeLanguageHuPointutfNegativeEight),
		string(VolumeLanguageIt),
		string(VolumeLanguageItPointutfNegativeEight),
		string(VolumeLanguageJa),
		string(VolumeLanguageJaNegativejpPointNineThreeTwo),
		string(VolumeLanguageJaNegativejpPointNineThreeTwoPointutfNegativeEight),
		string(VolumeLanguageJaNegativejpPointpck),
		string(VolumeLanguageJaNegativejpPointpckNegativevTwo),
		string(VolumeLanguageJaNegativejpPointpckNegativevTwoPointutfNegativeEight),
		string(VolumeLanguageJaNegativejpPointpckPointutfNegativeEight),
		string(VolumeLanguageJaNegativevOne),
		string(VolumeLanguageJaNegativevOnePointutfNegativeEight),
		string(VolumeLanguageJaPointutfNegativeEight),
		string(VolumeLanguageKo),
		string(VolumeLanguageKoPointutfNegativeEight),
		string(VolumeLanguageNl),
		string(VolumeLanguageNlPointutfNegativeEight),
		string(VolumeLanguageNo),
		string(VolumeLanguageNoPointutfNegativeEight),
		string(VolumeLanguagePl),
		string(VolumeLanguagePlPointutfNegativeEight),
		string(VolumeLanguagePt),
		string(VolumeLanguagePtPointutfNegativeEight),
		string(VolumeLanguageRo),
		string(VolumeLanguageRoPointutfNegativeEight),
		string(VolumeLanguageRu),
		string(VolumeLanguageRuPointutfNegativeEight),
		string(VolumeLanguageSk),
		string(VolumeLanguageSkPointutfNegativeEight),
		string(VolumeLanguageSl),
		string(VolumeLanguageSlPointutfNegativeEight),
		string(VolumeLanguageSv),
		string(VolumeLanguageSvPointutfNegativeEight),
		string(VolumeLanguageTr),
		string(VolumeLanguageTrPointutfNegativeEight),
		string(VolumeLanguageUtfEightmbFour),
		string(VolumeLanguageZh),
		string(VolumeLanguageZhNegativetw),
		string(VolumeLanguageZhNegativetwPointbigFive),
		string(VolumeLanguageZhNegativetwPointbigFivePointutfNegativeEight),
		string(VolumeLanguageZhNegativetwPointutfNegativeEight),
		string(VolumeLanguageZhPointgbk),
		string(VolumeLanguageZhPointgbkPointutfNegativeEight),
		string(VolumeLanguageZhPointutfNegativeEight),
	}
}

func (s *VolumeLanguage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVolumeLanguage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVolumeLanguage(input string) (*VolumeLanguage, error) {
	vals := map[string]VolumeLanguage{
		"ar":                 VolumeLanguageAr,
		"ar.utf-8":           VolumeLanguageArPointutfNegativeEight,
		"c":                  VolumeLanguageC,
		"c.utf-8":            VolumeLanguageCPointutfNegativeEight,
		"cs":                 VolumeLanguageCs,
		"cs.utf-8":           VolumeLanguageCsPointutfNegativeEight,
		"da":                 VolumeLanguageDa,
		"da.utf-8":           VolumeLanguageDaPointutfNegativeEight,
		"de":                 VolumeLanguageDe,
		"de.utf-8":           VolumeLanguageDePointutfNegativeEight,
		"en":                 VolumeLanguageEn,
		"en-us":              VolumeLanguageEnNegativeus,
		"en-us.utf-8":        VolumeLanguageEnNegativeusPointutfNegativeEight,
		"en.utf-8":           VolumeLanguageEnPointutfNegativeEight,
		"es":                 VolumeLanguageEs,
		"es.utf-8":           VolumeLanguageEsPointutfNegativeEight,
		"fi":                 VolumeLanguageFi,
		"fi.utf-8":           VolumeLanguageFiPointutfNegativeEight,
		"fr":                 VolumeLanguageFr,
		"fr.utf-8":           VolumeLanguageFrPointutfNegativeEight,
		"he":                 VolumeLanguageHe,
		"he.utf-8":           VolumeLanguageHePointutfNegativeEight,
		"hr":                 VolumeLanguageHr,
		"hr.utf-8":           VolumeLanguageHrPointutfNegativeEight,
		"hu":                 VolumeLanguageHu,
		"hu.utf-8":           VolumeLanguageHuPointutfNegativeEight,
		"it":                 VolumeLanguageIt,
		"it.utf-8":           VolumeLanguageItPointutfNegativeEight,
		"ja":                 VolumeLanguageJa,
		"ja-jp.932":          VolumeLanguageJaNegativejpPointNineThreeTwo,
		"ja-jp.932.utf-8":    VolumeLanguageJaNegativejpPointNineThreeTwoPointutfNegativeEight,
		"ja-jp.pck":          VolumeLanguageJaNegativejpPointpck,
		"ja-jp.pck-v2":       VolumeLanguageJaNegativejpPointpckNegativevTwo,
		"ja-jp.pck-v2.utf-8": VolumeLanguageJaNegativejpPointpckNegativevTwoPointutfNegativeEight,
		"ja-jp.pck.utf-8":    VolumeLanguageJaNegativejpPointpckPointutfNegativeEight,
		"ja-v1":              VolumeLanguageJaNegativevOne,
		"ja-v1.utf-8":        VolumeLanguageJaNegativevOnePointutfNegativeEight,
		"ja.utf-8":           VolumeLanguageJaPointutfNegativeEight,
		"ko":                 VolumeLanguageKo,
		"ko.utf-8":           VolumeLanguageKoPointutfNegativeEight,
		"nl":                 VolumeLanguageNl,
		"nl.utf-8":           VolumeLanguageNlPointutfNegativeEight,
		"no":                 VolumeLanguageNo,
		"no.utf-8":           VolumeLanguageNoPointutfNegativeEight,
		"pl":                 VolumeLanguagePl,
		"pl.utf-8":           VolumeLanguagePlPointutfNegativeEight,
		"pt":                 VolumeLanguagePt,
		"pt.utf-8":           VolumeLanguagePtPointutfNegativeEight,
		"ro":                 VolumeLanguageRo,
		"ro.utf-8":           VolumeLanguageRoPointutfNegativeEight,
		"ru":                 VolumeLanguageRu,
		"ru.utf-8":           VolumeLanguageRuPointutfNegativeEight,
		"sk":                 VolumeLanguageSk,
		"sk.utf-8":           VolumeLanguageSkPointutfNegativeEight,
		"sl":                 VolumeLanguageSl,
		"sl.utf-8":           VolumeLanguageSlPointutfNegativeEight,
		"sv":                 VolumeLanguageSv,
		"sv.utf-8":           VolumeLanguageSvPointutfNegativeEight,
		"tr":                 VolumeLanguageTr,
		"tr.utf-8":           VolumeLanguageTrPointutfNegativeEight,
		"utf8mb4":            VolumeLanguageUtfEightmbFour,
		"zh":                 VolumeLanguageZh,
		"zh-tw":              VolumeLanguageZhNegativetw,
		"zh-tw.big5":         VolumeLanguageZhNegativetwPointbigFive,
		"zh-tw.big5.utf-8":   VolumeLanguageZhNegativetwPointbigFivePointutfNegativeEight,
		"zh-tw.utf-8":        VolumeLanguageZhNegativetwPointutfNegativeEight,
		"zh.gbk":             VolumeLanguageZhPointgbk,
		"zh.gbk.utf-8":       VolumeLanguageZhPointgbkPointutfNegativeEight,
		"zh.utf-8":           VolumeLanguageZhPointutfNegativeEight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VolumeLanguage(input)
	return &out, nil
}
