package siteconfigslotresourceoperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoHealActionType string

const (
	AutoHealActionTypeCustomAction AutoHealActionType = "CustomAction"
	AutoHealActionTypeLogEvent     AutoHealActionType = "LogEvent"
	AutoHealActionTypeRecycle      AutoHealActionType = "Recycle"
)

func PossibleValuesForAutoHealActionType() []string {
	return []string{
		string(AutoHealActionTypeCustomAction),
		string(AutoHealActionTypeLogEvent),
		string(AutoHealActionTypeRecycle),
	}
}

func (s *AutoHealActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoHealActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoHealActionType(input string) (*AutoHealActionType, error) {
	vals := map[string]AutoHealActionType{
		"customaction": AutoHealActionTypeCustomAction,
		"logevent":     AutoHealActionTypeLogEvent,
		"recycle":      AutoHealActionTypeRecycle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoHealActionType(input)
	return &out, nil
}

type AzureStorageProtocol string

const (
	AzureStorageProtocolHTTP AzureStorageProtocol = "Http"
	AzureStorageProtocolNfs  AzureStorageProtocol = "Nfs"
	AzureStorageProtocolSmb  AzureStorageProtocol = "Smb"
)

func PossibleValuesForAzureStorageProtocol() []string {
	return []string{
		string(AzureStorageProtocolHTTP),
		string(AzureStorageProtocolNfs),
		string(AzureStorageProtocolSmb),
	}
}

func (s *AzureStorageProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureStorageProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureStorageProtocol(input string) (*AzureStorageProtocol, error) {
	vals := map[string]AzureStorageProtocol{
		"http": AzureStorageProtocolHTTP,
		"nfs":  AzureStorageProtocolNfs,
		"smb":  AzureStorageProtocolSmb,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureStorageProtocol(input)
	return &out, nil
}

type AzureStorageState string

const (
	AzureStorageStateInvalidCredentials AzureStorageState = "InvalidCredentials"
	AzureStorageStateInvalidShare       AzureStorageState = "InvalidShare"
	AzureStorageStateNotValidated       AzureStorageState = "NotValidated"
	AzureStorageStateOk                 AzureStorageState = "Ok"
)

func PossibleValuesForAzureStorageState() []string {
	return []string{
		string(AzureStorageStateInvalidCredentials),
		string(AzureStorageStateInvalidShare),
		string(AzureStorageStateNotValidated),
		string(AzureStorageStateOk),
	}
}

func (s *AzureStorageState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureStorageState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureStorageState(input string) (*AzureStorageState, error) {
	vals := map[string]AzureStorageState{
		"invalidcredentials": AzureStorageStateInvalidCredentials,
		"invalidshare":       AzureStorageStateInvalidShare,
		"notvalidated":       AzureStorageStateNotValidated,
		"ok":                 AzureStorageStateOk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureStorageState(input)
	return &out, nil
}

type AzureStorageType string

const (
	AzureStorageTypeAzureBlob  AzureStorageType = "AzureBlob"
	AzureStorageTypeAzureFiles AzureStorageType = "AzureFiles"
)

func PossibleValuesForAzureStorageType() []string {
	return []string{
		string(AzureStorageTypeAzureBlob),
		string(AzureStorageTypeAzureFiles),
	}
}

func (s *AzureStorageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAzureStorageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAzureStorageType(input string) (*AzureStorageType, error) {
	vals := map[string]AzureStorageType{
		"azureblob":  AzureStorageTypeAzureBlob,
		"azurefiles": AzureStorageTypeAzureFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureStorageType(input)
	return &out, nil
}

type ConnectionStringType string

const (
	ConnectionStringTypeApiHub          ConnectionStringType = "ApiHub"
	ConnectionStringTypeCustom          ConnectionStringType = "Custom"
	ConnectionStringTypeDocDb           ConnectionStringType = "DocDb"
	ConnectionStringTypeEventHub        ConnectionStringType = "EventHub"
	ConnectionStringTypeMySql           ConnectionStringType = "MySql"
	ConnectionStringTypeNotificationHub ConnectionStringType = "NotificationHub"
	ConnectionStringTypePostgreSQL      ConnectionStringType = "PostgreSQL"
	ConnectionStringTypeRedisCache      ConnectionStringType = "RedisCache"
	ConnectionStringTypeSQLAzure        ConnectionStringType = "SQLAzure"
	ConnectionStringTypeSQLServer       ConnectionStringType = "SQLServer"
	ConnectionStringTypeServiceBus      ConnectionStringType = "ServiceBus"
)

func PossibleValuesForConnectionStringType() []string {
	return []string{
		string(ConnectionStringTypeApiHub),
		string(ConnectionStringTypeCustom),
		string(ConnectionStringTypeDocDb),
		string(ConnectionStringTypeEventHub),
		string(ConnectionStringTypeMySql),
		string(ConnectionStringTypeNotificationHub),
		string(ConnectionStringTypePostgreSQL),
		string(ConnectionStringTypeRedisCache),
		string(ConnectionStringTypeSQLAzure),
		string(ConnectionStringTypeSQLServer),
		string(ConnectionStringTypeServiceBus),
	}
}

func (s *ConnectionStringType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionStringType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionStringType(input string) (*ConnectionStringType, error) {
	vals := map[string]ConnectionStringType{
		"apihub":          ConnectionStringTypeApiHub,
		"custom":          ConnectionStringTypeCustom,
		"docdb":           ConnectionStringTypeDocDb,
		"eventhub":        ConnectionStringTypeEventHub,
		"mysql":           ConnectionStringTypeMySql,
		"notificationhub": ConnectionStringTypeNotificationHub,
		"postgresql":      ConnectionStringTypePostgreSQL,
		"rediscache":      ConnectionStringTypeRedisCache,
		"sqlazure":        ConnectionStringTypeSQLAzure,
		"sqlserver":       ConnectionStringTypeSQLServer,
		"servicebus":      ConnectionStringTypeServiceBus,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionStringType(input)
	return &out, nil
}

type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleValuesForDefaultAction() []string {
	return []string{
		string(DefaultActionAllow),
		string(DefaultActionDeny),
	}
}

func (s *DefaultAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultAction(input string) (*DefaultAction, error) {
	vals := map[string]DefaultAction{
		"allow": DefaultActionAllow,
		"deny":  DefaultActionDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultAction(input)
	return &out, nil
}

type FtpsState string

const (
	FtpsStateAllAllowed FtpsState = "AllAllowed"
	FtpsStateDisabled   FtpsState = "Disabled"
	FtpsStateFtpsOnly   FtpsState = "FtpsOnly"
)

func PossibleValuesForFtpsState() []string {
	return []string{
		string(FtpsStateAllAllowed),
		string(FtpsStateDisabled),
		string(FtpsStateFtpsOnly),
	}
}

func (s *FtpsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFtpsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFtpsState(input string) (*FtpsState, error) {
	vals := map[string]FtpsState{
		"allallowed": FtpsStateAllAllowed,
		"disabled":   FtpsStateDisabled,
		"ftpsonly":   FtpsStateFtpsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FtpsState(input)
	return &out, nil
}

type IPFilterTag string

const (
	IPFilterTagDefault    IPFilterTag = "Default"
	IPFilterTagServiceTag IPFilterTag = "ServiceTag"
	IPFilterTagXffProxy   IPFilterTag = "XffProxy"
)

func PossibleValuesForIPFilterTag() []string {
	return []string{
		string(IPFilterTagDefault),
		string(IPFilterTagServiceTag),
		string(IPFilterTagXffProxy),
	}
}

func (s *IPFilterTag) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPFilterTag(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPFilterTag(input string) (*IPFilterTag, error) {
	vals := map[string]IPFilterTag{
		"default":    IPFilterTagDefault,
		"servicetag": IPFilterTagServiceTag,
		"xffproxy":   IPFilterTagXffProxy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPFilterTag(input)
	return &out, nil
}

type ManagedPipelineMode string

const (
	ManagedPipelineModeClassic    ManagedPipelineMode = "Classic"
	ManagedPipelineModeIntegrated ManagedPipelineMode = "Integrated"
)

func PossibleValuesForManagedPipelineMode() []string {
	return []string{
		string(ManagedPipelineModeClassic),
		string(ManagedPipelineModeIntegrated),
	}
}

func (s *ManagedPipelineMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedPipelineMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedPipelineMode(input string) (*ManagedPipelineMode, error) {
	vals := map[string]ManagedPipelineMode{
		"classic":    ManagedPipelineModeClassic,
		"integrated": ManagedPipelineModeIntegrated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedPipelineMode(input)
	return &out, nil
}

type ScmType string

const (
	ScmTypeBitbucketGit ScmType = "BitbucketGit"
	ScmTypeBitbucketHg  ScmType = "BitbucketHg"
	ScmTypeCodePlexGit  ScmType = "CodePlexGit"
	ScmTypeCodePlexHg   ScmType = "CodePlexHg"
	ScmTypeDropbox      ScmType = "Dropbox"
	ScmTypeExternalGit  ScmType = "ExternalGit"
	ScmTypeExternalHg   ScmType = "ExternalHg"
	ScmTypeGitHub       ScmType = "GitHub"
	ScmTypeLocalGit     ScmType = "LocalGit"
	ScmTypeNone         ScmType = "None"
	ScmTypeOneDrive     ScmType = "OneDrive"
	ScmTypeTfs          ScmType = "Tfs"
	ScmTypeVSO          ScmType = "VSO"
	ScmTypeVSTSRM       ScmType = "VSTSRM"
)

func PossibleValuesForScmType() []string {
	return []string{
		string(ScmTypeBitbucketGit),
		string(ScmTypeBitbucketHg),
		string(ScmTypeCodePlexGit),
		string(ScmTypeCodePlexHg),
		string(ScmTypeDropbox),
		string(ScmTypeExternalGit),
		string(ScmTypeExternalHg),
		string(ScmTypeGitHub),
		string(ScmTypeLocalGit),
		string(ScmTypeNone),
		string(ScmTypeOneDrive),
		string(ScmTypeTfs),
		string(ScmTypeVSO),
		string(ScmTypeVSTSRM),
	}
}

func (s *ScmType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScmType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScmType(input string) (*ScmType, error) {
	vals := map[string]ScmType{
		"bitbucketgit": ScmTypeBitbucketGit,
		"bitbuckethg":  ScmTypeBitbucketHg,
		"codeplexgit":  ScmTypeCodePlexGit,
		"codeplexhg":   ScmTypeCodePlexHg,
		"dropbox":      ScmTypeDropbox,
		"externalgit":  ScmTypeExternalGit,
		"externalhg":   ScmTypeExternalHg,
		"github":       ScmTypeGitHub,
		"localgit":     ScmTypeLocalGit,
		"none":         ScmTypeNone,
		"onedrive":     ScmTypeOneDrive,
		"tfs":          ScmTypeTfs,
		"vso":          ScmTypeVSO,
		"vstsrm":       ScmTypeVSTSRM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScmType(input)
	return &out, nil
}

type SiteLoadBalancing string

const (
	SiteLoadBalancingLeastRequests               SiteLoadBalancing = "LeastRequests"
	SiteLoadBalancingLeastRequestsWithTieBreaker SiteLoadBalancing = "LeastRequestsWithTieBreaker"
	SiteLoadBalancingLeastResponseTime           SiteLoadBalancing = "LeastResponseTime"
	SiteLoadBalancingPerSiteRoundRobin           SiteLoadBalancing = "PerSiteRoundRobin"
	SiteLoadBalancingRequestHash                 SiteLoadBalancing = "RequestHash"
	SiteLoadBalancingWeightedRoundRobin          SiteLoadBalancing = "WeightedRoundRobin"
	SiteLoadBalancingWeightedTotalTraffic        SiteLoadBalancing = "WeightedTotalTraffic"
)

func PossibleValuesForSiteLoadBalancing() []string {
	return []string{
		string(SiteLoadBalancingLeastRequests),
		string(SiteLoadBalancingLeastRequestsWithTieBreaker),
		string(SiteLoadBalancingLeastResponseTime),
		string(SiteLoadBalancingPerSiteRoundRobin),
		string(SiteLoadBalancingRequestHash),
		string(SiteLoadBalancingWeightedRoundRobin),
		string(SiteLoadBalancingWeightedTotalTraffic),
	}
}

func (s *SiteLoadBalancing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteLoadBalancing(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteLoadBalancing(input string) (*SiteLoadBalancing, error) {
	vals := map[string]SiteLoadBalancing{
		"leastrequests":               SiteLoadBalancingLeastRequests,
		"leastrequestswithtiebreaker": SiteLoadBalancingLeastRequestsWithTieBreaker,
		"leastresponsetime":           SiteLoadBalancingLeastResponseTime,
		"persiteroundrobin":           SiteLoadBalancingPerSiteRoundRobin,
		"requesthash":                 SiteLoadBalancingRequestHash,
		"weightedroundrobin":          SiteLoadBalancingWeightedRoundRobin,
		"weightedtotaltraffic":        SiteLoadBalancingWeightedTotalTraffic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteLoadBalancing(input)
	return &out, nil
}

type SupportedTlsVersions string

const (
	SupportedTlsVersionsOnePointOne   SupportedTlsVersions = "1.1"
	SupportedTlsVersionsOnePointThree SupportedTlsVersions = "1.3"
	SupportedTlsVersionsOnePointTwo   SupportedTlsVersions = "1.2"
	SupportedTlsVersionsOnePointZero  SupportedTlsVersions = "1.0"
)

func PossibleValuesForSupportedTlsVersions() []string {
	return []string{
		string(SupportedTlsVersionsOnePointOne),
		string(SupportedTlsVersionsOnePointThree),
		string(SupportedTlsVersionsOnePointTwo),
		string(SupportedTlsVersionsOnePointZero),
	}
}

func (s *SupportedTlsVersions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSupportedTlsVersions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSupportedTlsVersions(input string) (*SupportedTlsVersions, error) {
	vals := map[string]SupportedTlsVersions{
		"1.1": SupportedTlsVersionsOnePointOne,
		"1.3": SupportedTlsVersionsOnePointThree,
		"1.2": SupportedTlsVersionsOnePointTwo,
		"1.0": SupportedTlsVersionsOnePointZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportedTlsVersions(input)
	return &out, nil
}

type TlsCipherSuites string

const (
	TlsCipherSuitesTLSAESOneTwoEightGCMSHATwoFiveSix                  TlsCipherSuites = "TLS_AES_128_GCM_SHA256"
	TlsCipherSuitesTLSAESTwoFiveSixGCMSHAThreeEightFour               TlsCipherSuites = "TLS_AES_256_GCM_SHA384"
	TlsCipherSuitesTLSECDHEECDSAWITHAESOneTwoEightCBCSHATwoFiveSix    TlsCipherSuites = "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"
	TlsCipherSuitesTLSECDHEECDSAWITHAESOneTwoEightGCMSHATwoFiveSix    TlsCipherSuites = "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	TlsCipherSuitesTLSECDHEECDSAWITHAESTwoFiveSixGCMSHAThreeEightFour TlsCipherSuites = "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightCBCSHA                TlsCipherSuites = "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"
	TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightCBCSHATwoFiveSix      TlsCipherSuites = "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"
	TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightGCMSHATwoFiveSix      TlsCipherSuites = "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixCBCSHA                 TlsCipherSuites = "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA"
	TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixCBCSHAThreeEightFour   TlsCipherSuites = "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384"
	TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixGCMSHAThreeEightFour   TlsCipherSuites = "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	TlsCipherSuitesTLSRSAWITHAESOneTwoEightCBCSHA                     TlsCipherSuites = "TLS_RSA_WITH_AES_128_CBC_SHA"
	TlsCipherSuitesTLSRSAWITHAESOneTwoEightCBCSHATwoFiveSix           TlsCipherSuites = "TLS_RSA_WITH_AES_128_CBC_SHA256"
	TlsCipherSuitesTLSRSAWITHAESOneTwoEightGCMSHATwoFiveSix           TlsCipherSuites = "TLS_RSA_WITH_AES_128_GCM_SHA256"
	TlsCipherSuitesTLSRSAWITHAESTwoFiveSixCBCSHA                      TlsCipherSuites = "TLS_RSA_WITH_AES_256_CBC_SHA"
	TlsCipherSuitesTLSRSAWITHAESTwoFiveSixCBCSHATwoFiveSix            TlsCipherSuites = "TLS_RSA_WITH_AES_256_CBC_SHA256"
	TlsCipherSuitesTLSRSAWITHAESTwoFiveSixGCMSHAThreeEightFour        TlsCipherSuites = "TLS_RSA_WITH_AES_256_GCM_SHA384"
)

func PossibleValuesForTlsCipherSuites() []string {
	return []string{
		string(TlsCipherSuitesTLSAESOneTwoEightGCMSHATwoFiveSix),
		string(TlsCipherSuitesTLSAESTwoFiveSixGCMSHAThreeEightFour),
		string(TlsCipherSuitesTLSECDHEECDSAWITHAESOneTwoEightCBCSHATwoFiveSix),
		string(TlsCipherSuitesTLSECDHEECDSAWITHAESOneTwoEightGCMSHATwoFiveSix),
		string(TlsCipherSuitesTLSECDHEECDSAWITHAESTwoFiveSixGCMSHAThreeEightFour),
		string(TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightCBCSHA),
		string(TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightCBCSHATwoFiveSix),
		string(TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightGCMSHATwoFiveSix),
		string(TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixCBCSHA),
		string(TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixCBCSHAThreeEightFour),
		string(TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixGCMSHAThreeEightFour),
		string(TlsCipherSuitesTLSRSAWITHAESOneTwoEightCBCSHA),
		string(TlsCipherSuitesTLSRSAWITHAESOneTwoEightCBCSHATwoFiveSix),
		string(TlsCipherSuitesTLSRSAWITHAESOneTwoEightGCMSHATwoFiveSix),
		string(TlsCipherSuitesTLSRSAWITHAESTwoFiveSixCBCSHA),
		string(TlsCipherSuitesTLSRSAWITHAESTwoFiveSixCBCSHATwoFiveSix),
		string(TlsCipherSuitesTLSRSAWITHAESTwoFiveSixGCMSHAThreeEightFour),
	}
}

func (s *TlsCipherSuites) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTlsCipherSuites(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTlsCipherSuites(input string) (*TlsCipherSuites, error) {
	vals := map[string]TlsCipherSuites{
		"tls_aes_128_gcm_sha256":                  TlsCipherSuitesTLSAESOneTwoEightGCMSHATwoFiveSix,
		"tls_aes_256_gcm_sha384":                  TlsCipherSuitesTLSAESTwoFiveSixGCMSHAThreeEightFour,
		"tls_ecdhe_ecdsa_with_aes_128_cbc_sha256": TlsCipherSuitesTLSECDHEECDSAWITHAESOneTwoEightCBCSHATwoFiveSix,
		"tls_ecdhe_ecdsa_with_aes_128_gcm_sha256": TlsCipherSuitesTLSECDHEECDSAWITHAESOneTwoEightGCMSHATwoFiveSix,
		"tls_ecdhe_ecdsa_with_aes_256_gcm_sha384": TlsCipherSuitesTLSECDHEECDSAWITHAESTwoFiveSixGCMSHAThreeEightFour,
		"tls_ecdhe_rsa_with_aes_128_cbc_sha":      TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightCBCSHA,
		"tls_ecdhe_rsa_with_aes_128_cbc_sha256":   TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightCBCSHATwoFiveSix,
		"tls_ecdhe_rsa_with_aes_128_gcm_sha256":   TlsCipherSuitesTLSECDHERSAWITHAESOneTwoEightGCMSHATwoFiveSix,
		"tls_ecdhe_rsa_with_aes_256_cbc_sha":      TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixCBCSHA,
		"tls_ecdhe_rsa_with_aes_256_cbc_sha384":   TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixCBCSHAThreeEightFour,
		"tls_ecdhe_rsa_with_aes_256_gcm_sha384":   TlsCipherSuitesTLSECDHERSAWITHAESTwoFiveSixGCMSHAThreeEightFour,
		"tls_rsa_with_aes_128_cbc_sha":            TlsCipherSuitesTLSRSAWITHAESOneTwoEightCBCSHA,
		"tls_rsa_with_aes_128_cbc_sha256":         TlsCipherSuitesTLSRSAWITHAESOneTwoEightCBCSHATwoFiveSix,
		"tls_rsa_with_aes_128_gcm_sha256":         TlsCipherSuitesTLSRSAWITHAESOneTwoEightGCMSHATwoFiveSix,
		"tls_rsa_with_aes_256_cbc_sha":            TlsCipherSuitesTLSRSAWITHAESTwoFiveSixCBCSHA,
		"tls_rsa_with_aes_256_cbc_sha256":         TlsCipherSuitesTLSRSAWITHAESTwoFiveSixCBCSHATwoFiveSix,
		"tls_rsa_with_aes_256_gcm_sha384":         TlsCipherSuitesTLSRSAWITHAESTwoFiveSixGCMSHAThreeEightFour,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TlsCipherSuites(input)
	return &out, nil
}
