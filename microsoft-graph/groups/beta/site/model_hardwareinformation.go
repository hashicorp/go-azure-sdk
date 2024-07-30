package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformation struct {
	BatteryChargeCycles                                            *int64                                                                             `json:"batteryChargeCycles,omitempty"`
	BatteryHealthPercentage                                        *int64                                                                             `json:"batteryHealthPercentage,omitempty"`
	BatteryLevelPercentage                                         *float64                                                                           `json:"batteryLevelPercentage,omitempty"`
	BatterySerialNumber                                            *string                                                                            `json:"batterySerialNumber,omitempty"`
	CellularTechnology                                             *string                                                                            `json:"cellularTechnology,omitempty"`
	DeviceFullQualifiedDomainName                                  *string                                                                            `json:"deviceFullQualifiedDomainName,omitempty"`
	DeviceGuardLocalSystemAuthorityCredentialGuardState            *HardwareInformationDeviceGuardLocalSystemAuthorityCredentialGuardState            `json:"deviceGuardLocalSystemAuthorityCredentialGuardState,omitempty"`
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState *HardwareInformationDeviceGuardVirtualizationBasedSecurityHardwareRequirementState `json:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState,omitempty"`
	DeviceGuardVirtualizationBasedSecurityState                    *HardwareInformationDeviceGuardVirtualizationBasedSecurityState                    `json:"deviceGuardVirtualizationBasedSecurityState,omitempty"`
	DeviceLicensingLastErrorCode                                   *int64                                                                             `json:"deviceLicensingLastErrorCode,omitempty"`
	DeviceLicensingLastErrorDescription                            *string                                                                            `json:"deviceLicensingLastErrorDescription,omitempty"`
	DeviceLicensingStatus                                          *HardwareInformationDeviceLicensingStatus                                          `json:"deviceLicensingStatus,omitempty"`
	EsimIdentifier                                                 *string                                                                            `json:"esimIdentifier,omitempty"`
	FreeStorageSpace                                               *int64                                                                             `json:"freeStorageSpace,omitempty"`
	Imei                                                           *string                                                                            `json:"imei,omitempty"`
	IpAddressV4                                                    *string                                                                            `json:"ipAddressV4,omitempty"`
	IsEncrypted                                                    *bool                                                                              `json:"isEncrypted,omitempty"`
	IsSharedDevice                                                 *bool                                                                              `json:"isSharedDevice,omitempty"`
	IsSupervised                                                   *bool                                                                              `json:"isSupervised,omitempty"`
	Manufacturer                                                   *string                                                                            `json:"manufacturer,omitempty"`
	Meid                                                           *string                                                                            `json:"meid,omitempty"`
	Model                                                          *string                                                                            `json:"model,omitempty"`
	ODataType                                                      *string                                                                            `json:"@odata.type,omitempty"`
	OperatingSystemEdition                                         *string                                                                            `json:"operatingSystemEdition,omitempty"`
	OperatingSystemLanguage                                        *string                                                                            `json:"operatingSystemLanguage,omitempty"`
	OperatingSystemProductType                                     *int64                                                                             `json:"operatingSystemProductType,omitempty"`
	OsBuildNumber                                                  *string                                                                            `json:"osBuildNumber,omitempty"`
	PhoneNumber                                                    *string                                                                            `json:"phoneNumber,omitempty"`
	ProductName                                                    *string                                                                            `json:"productName,omitempty"`
	ResidentUsersCount                                             *int64                                                                             `json:"residentUsersCount,omitempty"`
	SerialNumber                                                   *string                                                                            `json:"serialNumber,omitempty"`
	SharedDeviceCachedUsers                                        *[]SharedAppleDeviceUser                                                           `json:"sharedDeviceCachedUsers,omitempty"`
	SubnetAddress                                                  *string                                                                            `json:"subnetAddress,omitempty"`
	SubscriberCarrier                                              *string                                                                            `json:"subscriberCarrier,omitempty"`
	SystemManagementBIOSVersion                                    *string                                                                            `json:"systemManagementBIOSVersion,omitempty"`
	TotalStorageSpace                                              *int64                                                                             `json:"totalStorageSpace,omitempty"`
	TpmManufacturer                                                *string                                                                            `json:"tpmManufacturer,omitempty"`
	TpmSpecificationVersion                                        *string                                                                            `json:"tpmSpecificationVersion,omitempty"`
	TpmVersion                                                     *string                                                                            `json:"tpmVersion,omitempty"`
	WifiMac                                                        *string                                                                            `json:"wifiMac,omitempty"`
	WiredIPv4Addresses                                             *[]string                                                                          `json:"wiredIPv4Addresses,omitempty"`
}
