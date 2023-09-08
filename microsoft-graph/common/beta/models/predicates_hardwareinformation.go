package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareInformationOperationPredicate struct {
	BatteryChargeCycles                 *int64
	BatteryHealthPercentage             *int64
	BatterySerialNumber                 *string
	CellularTechnology                  *string
	DeviceFullQualifiedDomainName       *string
	DeviceLicensingLastErrorCode        *int64
	DeviceLicensingLastErrorDescription *string
	EsimIdentifier                      *string
	FreeStorageSpace                    *int64
	Imei                                *string
	IpAddressV4                         *string
	IsEncrypted                         *bool
	IsSharedDevice                      *bool
	IsSupervised                        *bool
	Manufacturer                        *string
	Meid                                *string
	Model                               *string
	ODataType                           *string
	OperatingSystemEdition              *string
	OperatingSystemLanguage             *string
	OperatingSystemProductType          *int64
	OsBuildNumber                       *string
	PhoneNumber                         *string
	ProductName                         *string
	ResidentUsersCount                  *int64
	SerialNumber                        *string
	SubnetAddress                       *string
	SubscriberCarrier                   *string
	SystemManagementBIOSVersion         *string
	TotalStorageSpace                   *int64
	TpmManufacturer                     *string
	TpmSpecificationVersion             *string
	TpmVersion                          *string
	WifiMac                             *string
}

func (p HardwareInformationOperationPredicate) Matches(input HardwareInformation) bool {

	if p.BatteryChargeCycles != nil && (input.BatteryChargeCycles == nil || *p.BatteryChargeCycles != *input.BatteryChargeCycles) {
		return false
	}

	if p.BatteryHealthPercentage != nil && (input.BatteryHealthPercentage == nil || *p.BatteryHealthPercentage != *input.BatteryHealthPercentage) {
		return false
	}

	if p.BatterySerialNumber != nil && (input.BatterySerialNumber == nil || *p.BatterySerialNumber != *input.BatterySerialNumber) {
		return false
	}

	if p.CellularTechnology != nil && (input.CellularTechnology == nil || *p.CellularTechnology != *input.CellularTechnology) {
		return false
	}

	if p.DeviceFullQualifiedDomainName != nil && (input.DeviceFullQualifiedDomainName == nil || *p.DeviceFullQualifiedDomainName != *input.DeviceFullQualifiedDomainName) {
		return false
	}

	if p.DeviceLicensingLastErrorCode != nil && (input.DeviceLicensingLastErrorCode == nil || *p.DeviceLicensingLastErrorCode != *input.DeviceLicensingLastErrorCode) {
		return false
	}

	if p.DeviceLicensingLastErrorDescription != nil && (input.DeviceLicensingLastErrorDescription == nil || *p.DeviceLicensingLastErrorDescription != *input.DeviceLicensingLastErrorDescription) {
		return false
	}

	if p.EsimIdentifier != nil && (input.EsimIdentifier == nil || *p.EsimIdentifier != *input.EsimIdentifier) {
		return false
	}

	if p.FreeStorageSpace != nil && (input.FreeStorageSpace == nil || *p.FreeStorageSpace != *input.FreeStorageSpace) {
		return false
	}

	if p.Imei != nil && (input.Imei == nil || *p.Imei != *input.Imei) {
		return false
	}

	if p.IpAddressV4 != nil && (input.IpAddressV4 == nil || *p.IpAddressV4 != *input.IpAddressV4) {
		return false
	}

	if p.IsEncrypted != nil && (input.IsEncrypted == nil || *p.IsEncrypted != *input.IsEncrypted) {
		return false
	}

	if p.IsSharedDevice != nil && (input.IsSharedDevice == nil || *p.IsSharedDevice != *input.IsSharedDevice) {
		return false
	}

	if p.IsSupervised != nil && (input.IsSupervised == nil || *p.IsSupervised != *input.IsSupervised) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Meid != nil && (input.Meid == nil || *p.Meid != *input.Meid) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperatingSystemEdition != nil && (input.OperatingSystemEdition == nil || *p.OperatingSystemEdition != *input.OperatingSystemEdition) {
		return false
	}

	if p.OperatingSystemLanguage != nil && (input.OperatingSystemLanguage == nil || *p.OperatingSystemLanguage != *input.OperatingSystemLanguage) {
		return false
	}

	if p.OperatingSystemProductType != nil && (input.OperatingSystemProductType == nil || *p.OperatingSystemProductType != *input.OperatingSystemProductType) {
		return false
	}

	if p.OsBuildNumber != nil && (input.OsBuildNumber == nil || *p.OsBuildNumber != *input.OsBuildNumber) {
		return false
	}

	if p.PhoneNumber != nil && (input.PhoneNumber == nil || *p.PhoneNumber != *input.PhoneNumber) {
		return false
	}

	if p.ProductName != nil && (input.ProductName == nil || *p.ProductName != *input.ProductName) {
		return false
	}

	if p.ResidentUsersCount != nil && (input.ResidentUsersCount == nil || *p.ResidentUsersCount != *input.ResidentUsersCount) {
		return false
	}

	if p.SerialNumber != nil && (input.SerialNumber == nil || *p.SerialNumber != *input.SerialNumber) {
		return false
	}

	if p.SubnetAddress != nil && (input.SubnetAddress == nil || *p.SubnetAddress != *input.SubnetAddress) {
		return false
	}

	if p.SubscriberCarrier != nil && (input.SubscriberCarrier == nil || *p.SubscriberCarrier != *input.SubscriberCarrier) {
		return false
	}

	if p.SystemManagementBIOSVersion != nil && (input.SystemManagementBIOSVersion == nil || *p.SystemManagementBIOSVersion != *input.SystemManagementBIOSVersion) {
		return false
	}

	if p.TotalStorageSpace != nil && (input.TotalStorageSpace == nil || *p.TotalStorageSpace != *input.TotalStorageSpace) {
		return false
	}

	if p.TpmManufacturer != nil && (input.TpmManufacturer == nil || *p.TpmManufacturer != *input.TpmManufacturer) {
		return false
	}

	if p.TpmSpecificationVersion != nil && (input.TpmSpecificationVersion == nil || *p.TpmSpecificationVersion != *input.TpmSpecificationVersion) {
		return false
	}

	if p.TpmVersion != nil && (input.TpmVersion == nil || *p.TpmVersion != *input.TpmVersion) {
		return false
	}

	if p.WifiMac != nil && (input.WifiMac == nil || *p.WifiMac != *input.WifiMac) {
		return false
	}

	return true
}
