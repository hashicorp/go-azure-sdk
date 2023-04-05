package sapvirtualinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationType string

const (
	ConfigurationTypeCreateAndMount ConfigurationType = "CreateAndMount"
	ConfigurationTypeMount          ConfigurationType = "Mount"
	ConfigurationTypeSkip           ConfigurationType = "Skip"
)

func PossibleValuesForConfigurationType() []string {
	return []string{
		string(ConfigurationTypeCreateAndMount),
		string(ConfigurationTypeMount),
		string(ConfigurationTypeSkip),
	}
}

func (s *ConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForConfigurationType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ConfigurationType(decoded)
	return nil
}

type DiskSkuName string

const (
	DiskSkuNamePremiumLRS     DiskSkuName = "Premium_LRS"
	DiskSkuNamePremiumVTwoLRS DiskSkuName = "PremiumV2_LRS"
	DiskSkuNamePremiumZRS     DiskSkuName = "Premium_ZRS"
	DiskSkuNameStandardLRS    DiskSkuName = "Standard_LRS"
	DiskSkuNameStandardSSDLRS DiskSkuName = "StandardSSD_LRS"
	DiskSkuNameStandardSSDZRS DiskSkuName = "StandardSSD_ZRS"
	DiskSkuNameUltraSSDLRS    DiskSkuName = "UltraSSD_LRS"
)

func PossibleValuesForDiskSkuName() []string {
	return []string{
		string(DiskSkuNamePremiumLRS),
		string(DiskSkuNamePremiumVTwoLRS),
		string(DiskSkuNamePremiumZRS),
		string(DiskSkuNameStandardLRS),
		string(DiskSkuNameStandardSSDLRS),
		string(DiskSkuNameStandardSSDZRS),
		string(DiskSkuNameUltraSSDLRS),
	}
}

func (s *DiskSkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForDiskSkuName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = DiskSkuName(decoded)
	return nil
}

type NamingPatternType string

const (
	NamingPatternTypeFullResourceName NamingPatternType = "FullResourceName"
)

func PossibleValuesForNamingPatternType() []string {
	return []string{
		string(NamingPatternTypeFullResourceName),
	}
}

func (s *NamingPatternType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForNamingPatternType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = NamingPatternType(decoded)
	return nil
}

type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

func PossibleValuesForOSType() []string {
	return []string{
		string(OSTypeLinux),
		string(OSTypeWindows),
	}
}

func (s *OSType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForOSType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = OSType(decoded)
	return nil
}

type SAPConfigurationType string

const (
	SAPConfigurationTypeDeployment             SAPConfigurationType = "Deployment"
	SAPConfigurationTypeDeploymentWithOSConfig SAPConfigurationType = "DeploymentWithOSConfig"
	SAPConfigurationTypeDiscovery              SAPConfigurationType = "Discovery"
)

func PossibleValuesForSAPConfigurationType() []string {
	return []string{
		string(SAPConfigurationTypeDeployment),
		string(SAPConfigurationTypeDeploymentWithOSConfig),
		string(SAPConfigurationTypeDiscovery),
	}
}

func (s *SAPConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPConfigurationType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPConfigurationType(decoded)
	return nil
}

type SAPDatabaseType string

const (
	SAPDatabaseTypeDBTwo SAPDatabaseType = "DB2"
	SAPDatabaseTypeHANA  SAPDatabaseType = "HANA"
)

func PossibleValuesForSAPDatabaseType() []string {
	return []string{
		string(SAPDatabaseTypeDBTwo),
		string(SAPDatabaseTypeHANA),
	}
}

func (s *SAPDatabaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPDatabaseType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPDatabaseType(decoded)
	return nil
}

type SAPDeploymentType string

const (
	SAPDeploymentTypeSingleServer SAPDeploymentType = "SingleServer"
	SAPDeploymentTypeThreeTier    SAPDeploymentType = "ThreeTier"
)

func PossibleValuesForSAPDeploymentType() []string {
	return []string{
		string(SAPDeploymentTypeSingleServer),
		string(SAPDeploymentTypeThreeTier),
	}
}

func (s *SAPDeploymentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPDeploymentType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPDeploymentType(decoded)
	return nil
}

type SAPEnvironmentType string

const (
	SAPEnvironmentTypeNonProd SAPEnvironmentType = "NonProd"
	SAPEnvironmentTypeProd    SAPEnvironmentType = "Prod"
)

func PossibleValuesForSAPEnvironmentType() []string {
	return []string{
		string(SAPEnvironmentTypeNonProd),
		string(SAPEnvironmentTypeProd),
	}
}

func (s *SAPEnvironmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPEnvironmentType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPEnvironmentType(decoded)
	return nil
}

type SAPHealthState string

const (
	SAPHealthStateDegraded  SAPHealthState = "Degraded"
	SAPHealthStateHealthy   SAPHealthState = "Healthy"
	SAPHealthStateUnhealthy SAPHealthState = "Unhealthy"
	SAPHealthStateUnknown   SAPHealthState = "Unknown"
)

func PossibleValuesForSAPHealthState() []string {
	return []string{
		string(SAPHealthStateDegraded),
		string(SAPHealthStateHealthy),
		string(SAPHealthStateUnhealthy),
		string(SAPHealthStateUnknown),
	}
}

func (s *SAPHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPHealthState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPHealthState(decoded)
	return nil
}

type SAPHighAvailabilityType string

const (
	SAPHighAvailabilityTypeAvailabilitySet  SAPHighAvailabilityType = "AvailabilitySet"
	SAPHighAvailabilityTypeAvailabilityZone SAPHighAvailabilityType = "AvailabilityZone"
)

func PossibleValuesForSAPHighAvailabilityType() []string {
	return []string{
		string(SAPHighAvailabilityTypeAvailabilitySet),
		string(SAPHighAvailabilityTypeAvailabilityZone),
	}
}

func (s *SAPHighAvailabilityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPHighAvailabilityType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPHighAvailabilityType(decoded)
	return nil
}

type SAPProductType string

const (
	SAPProductTypeECC       SAPProductType = "ECC"
	SAPProductTypeOther     SAPProductType = "Other"
	SAPProductTypeSFourHANA SAPProductType = "S4HANA"
)

func PossibleValuesForSAPProductType() []string {
	return []string{
		string(SAPProductTypeECC),
		string(SAPProductTypeOther),
		string(SAPProductTypeSFourHANA),
	}
}

func (s *SAPProductType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPProductType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPProductType(decoded)
	return nil
}

type SAPSoftwareInstallationType string

const (
	SAPSoftwareInstallationTypeExternal                  SAPSoftwareInstallationType = "External"
	SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig SAPSoftwareInstallationType = "SAPInstallWithoutOSConfig"
	SAPSoftwareInstallationTypeServiceInitiated          SAPSoftwareInstallationType = "ServiceInitiated"
)

func PossibleValuesForSAPSoftwareInstallationType() []string {
	return []string{
		string(SAPSoftwareInstallationTypeExternal),
		string(SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig),
		string(SAPSoftwareInstallationTypeServiceInitiated),
	}
}

func (s *SAPSoftwareInstallationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPSoftwareInstallationType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPSoftwareInstallationType(decoded)
	return nil
}

type SAPVirtualInstanceState string

const (
	SAPVirtualInstanceStateDiscoveryFailed                    SAPVirtualInstanceState = "DiscoveryFailed"
	SAPVirtualInstanceStateDiscoveryInProgress                SAPVirtualInstanceState = "DiscoveryInProgress"
	SAPVirtualInstanceStateDiscoveryPending                   SAPVirtualInstanceState = "DiscoveryPending"
	SAPVirtualInstanceStateInfrastructureDeploymentFailed     SAPVirtualInstanceState = "InfrastructureDeploymentFailed"
	SAPVirtualInstanceStateInfrastructureDeploymentInProgress SAPVirtualInstanceState = "InfrastructureDeploymentInProgress"
	SAPVirtualInstanceStateInfrastructureDeploymentPending    SAPVirtualInstanceState = "InfrastructureDeploymentPending"
	SAPVirtualInstanceStateRegistrationComplete               SAPVirtualInstanceState = "RegistrationComplete"
	SAPVirtualInstanceStateSoftwareDetectionFailed            SAPVirtualInstanceState = "SoftwareDetectionFailed"
	SAPVirtualInstanceStateSoftwareDetectionInProgress        SAPVirtualInstanceState = "SoftwareDetectionInProgress"
	SAPVirtualInstanceStateSoftwareInstallationFailed         SAPVirtualInstanceState = "SoftwareInstallationFailed"
	SAPVirtualInstanceStateSoftwareInstallationInProgress     SAPVirtualInstanceState = "SoftwareInstallationInProgress"
	SAPVirtualInstanceStateSoftwareInstallationPending        SAPVirtualInstanceState = "SoftwareInstallationPending"
)

func PossibleValuesForSAPVirtualInstanceState() []string {
	return []string{
		string(SAPVirtualInstanceStateDiscoveryFailed),
		string(SAPVirtualInstanceStateDiscoveryInProgress),
		string(SAPVirtualInstanceStateDiscoveryPending),
		string(SAPVirtualInstanceStateInfrastructureDeploymentFailed),
		string(SAPVirtualInstanceStateInfrastructureDeploymentInProgress),
		string(SAPVirtualInstanceStateInfrastructureDeploymentPending),
		string(SAPVirtualInstanceStateRegistrationComplete),
		string(SAPVirtualInstanceStateSoftwareDetectionFailed),
		string(SAPVirtualInstanceStateSoftwareDetectionInProgress),
		string(SAPVirtualInstanceStateSoftwareInstallationFailed),
		string(SAPVirtualInstanceStateSoftwareInstallationInProgress),
		string(SAPVirtualInstanceStateSoftwareInstallationPending),
	}
}

func (s *SAPVirtualInstanceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPVirtualInstanceState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPVirtualInstanceState(decoded)
	return nil
}

type SAPVirtualInstanceStatus string

const (
	SAPVirtualInstanceStatusOffline          SAPVirtualInstanceStatus = "Offline"
	SAPVirtualInstanceStatusPartiallyRunning SAPVirtualInstanceStatus = "PartiallyRunning"
	SAPVirtualInstanceStatusRunning          SAPVirtualInstanceStatus = "Running"
	SAPVirtualInstanceStatusSoftShutdown     SAPVirtualInstanceStatus = "SoftShutdown"
	SAPVirtualInstanceStatusStarting         SAPVirtualInstanceStatus = "Starting"
	SAPVirtualInstanceStatusStopping         SAPVirtualInstanceStatus = "Stopping"
	SAPVirtualInstanceStatusUnavailable      SAPVirtualInstanceStatus = "Unavailable"
)

func PossibleValuesForSAPVirtualInstanceStatus() []string {
	return []string{
		string(SAPVirtualInstanceStatusOffline),
		string(SAPVirtualInstanceStatusPartiallyRunning),
		string(SAPVirtualInstanceStatusRunning),
		string(SAPVirtualInstanceStatusSoftShutdown),
		string(SAPVirtualInstanceStatusStarting),
		string(SAPVirtualInstanceStatusStopping),
		string(SAPVirtualInstanceStatusUnavailable),
	}
}

func (s *SAPVirtualInstanceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPVirtualInstanceStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPVirtualInstanceStatus(decoded)
	return nil
}

type SapVirtualInstanceProvisioningState string

const (
	SapVirtualInstanceProvisioningStateCreating  SapVirtualInstanceProvisioningState = "Creating"
	SapVirtualInstanceProvisioningStateDeleting  SapVirtualInstanceProvisioningState = "Deleting"
	SapVirtualInstanceProvisioningStateFailed    SapVirtualInstanceProvisioningState = "Failed"
	SapVirtualInstanceProvisioningStateSucceeded SapVirtualInstanceProvisioningState = "Succeeded"
	SapVirtualInstanceProvisioningStateUpdating  SapVirtualInstanceProvisioningState = "Updating"
)

func PossibleValuesForSapVirtualInstanceProvisioningState() []string {
	return []string{
		string(SapVirtualInstanceProvisioningStateCreating),
		string(SapVirtualInstanceProvisioningStateDeleting),
		string(SapVirtualInstanceProvisioningStateFailed),
		string(SapVirtualInstanceProvisioningStateSucceeded),
		string(SapVirtualInstanceProvisioningStateUpdating),
	}
}

func (s *SapVirtualInstanceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSapVirtualInstanceProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SapVirtualInstanceProvisioningState(decoded)
	return nil
}
