package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsInformationProtection = WindowsInformationProtectionPolicy{}

type WindowsInformationProtectionPolicy struct {
	// Offline interval before app data is wiped (days)
	DaysWithoutContactBeforeUnenroll *int64 `json:"daysWithoutContactBeforeUnenroll,omitempty"`

	// Enrollment url for the MDM
	MdmEnrollmentUrl nullable.Type[string] `json:"mdmEnrollmentUrl,omitempty"`

	// Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to
	// become PIN or password locked. Range is an integer X where 0 <= X <= 999.
	MinutesOfInactivityBeforeDeviceLock *int64 `json:"minutesOfInactivityBeforeDeviceLock,omitempty"`

	// Integer value that specifies the number of past PINs that can be associated to a user account that can't be reused.
	// The largest number you can configure for this policy setting is 50. The lowest number you can configure for this
	// policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added
	// in Windows 10, version 1511. Default is 0.
	NumberOfPastPinsRemembered *int64 `json:"numberOfPastPinsRemembered,omitempty"`

	// The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe
	// functionality. Range is an integer X where 4 <= X <= 16 for desktop and 0 <= X <= 999 for mobile devices.
	PasswordMaximumAttemptCount *int64 `json:"passwordMaximumAttemptCount,omitempty"`

	// Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to
	// change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure
	// for this policy setting is 0. If this policy is set to 0, then the user's PIN will never expire. This node was added
	// in Windows 10, version 1511. Default is 0.
	PinExpirationDays *int64 `json:"pinExpirationDays,omitempty"`

	// Pin Character Requirements
	PinLowercaseLetters *WindowsInformationProtectionPinCharacterRequirements `json:"pinLowercaseLetters,omitempty"`

	// Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number
	// you can configure for this policy setting is 4. The largest number you can configure must be less than the number
	// configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest.
	PinMinimumLength *int64 `json:"pinMinimumLength,omitempty"`

	// Pin Character Requirements
	PinSpecialCharacters *WindowsInformationProtectionPinCharacterRequirements `json:"pinSpecialCharacters,omitempty"`

	// Pin Character Requirements
	PinUppercaseLetters *WindowsInformationProtectionPinCharacterRequirements `json:"pinUppercaseLetters,omitempty"`

	// New property in RS2, pending documentation
	RevokeOnMdmHandoffDisabled *bool `json:"revokeOnMdmHandoffDisabled,omitempty"`

	// Boolean value that sets Windows Hello for Business as a method for signing into Windows.
	WindowsHelloForBusinessBlocked *bool `json:"windowsHelloForBusinessBlocked,omitempty"`

	// Fields inherited from WindowsInformationProtection

	// Navigation property to list of security groups targeted for policy.
	Assignments *[]TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`

	// Specifies whether to allow Azure RMS encryption for WIP
	AzureRightsManagementServicesAllowed *bool `json:"azureRightsManagementServicesAllowed,omitempty"`

	// Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data
	// recovery agent(DRA) certificate for encrypting file system(EFS)
	DataRecoveryCertificate *WindowsInformationProtectionDataRecoveryCertificate `json:"dataRecoveryCertificate,omitempty"`

	// Possible values for WIP Protection enforcement levels
	EnforcementLevel *WindowsInformationProtectionEnforcementLevel `json:"enforcementLevel,omitempty"`

	// Primary enterprise domain
	EnterpriseDomain nullable.Type[string] `json:"enterpriseDomain,omitempty"`

	// Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those
	// computers will be considered part of the enterprise and protected. These locations will be considered a safe
	// destination for enterprise data to be shared to
	EnterpriseIPRanges *[]WindowsInformationProtectionIPRangeCollection `json:"enterpriseIPRanges,omitempty"`

	// Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other
	// subnets. Default is false
	EnterpriseIPRangesAreAuthoritative *bool `json:"enterpriseIPRangesAreAuthoritative,omitempty"`

	// This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167,
	// 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on
	// the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring
	// the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies
	EnterpriseInternalProxyServers *[]WindowsInformationProtectionResourceCollection `json:"enterpriseInternalProxyServers,omitempty"`

	// This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is
	// sent to a device will be considered enterprise data and protected These locations will be considered a safe
	// destination for enterprise data to be shared to
	EnterpriseNetworkDomainNames *[]WindowsInformationProtectionResourceCollection `json:"enterpriseNetworkDomainNames,omitempty"`

	// List of enterprise domains to be protected
	EnterpriseProtectedDomainNames *[]WindowsInformationProtectionResourceCollection `json:"enterpriseProtectedDomainNames,omitempty"`

	// Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these
	// resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource
	// will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this
	// purpose must also be configured using the EnterpriseInternalProxyServers policy
	EnterpriseProxiedDomains *[]WindowsInformationProtectionProxiedDomainCollection `json:"enterpriseProxiedDomains,omitempty"`

	// This is a list of proxy servers. Any server not on this list is considered non-enterprise
	EnterpriseProxyServers *[]WindowsInformationProtectionResourceCollection `json:"enterpriseProxyServers,omitempty"`

	// Boolean value that tells the client to accept the configured list of proxies and not try to detect other work
	// proxies. Default is false
	EnterpriseProxyServersAreAuthoritative *bool `json:"enterpriseProxyServersAreAuthoritative,omitempty"`

	// Another way to input exempt apps through xml files
	ExemptAppLockerFiles *[]WindowsInformationProtectionAppLockerFile `json:"exemptAppLockerFiles,omitempty"`

	// Exempt applications can also access enterprise data, but the data handled by those applications are not protected.
	// This is because some critical enterprise applications may have compatibility problems with encrypted data.
	ExemptApps *[]WindowsInformationProtectionApp `json:"exemptApps,omitempty"`

	// Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in
	// the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in
	// the title bar of a WIP-protected app
	IconsVisible *bool `json:"iconsVisible,omitempty"`

	// This switch is for the Windows Search Indexer, to allow or disallow indexing of items
	IndexingEncryptedStoresOrItemsBlocked *bool `json:"indexingEncryptedStoresOrItemsBlocked,omitempty"`

	// Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned *bool `json:"isAssigned,omitempty"`

	// List of domain names that can used for work or personal resource
	NeutralDomainResources *[]WindowsInformationProtectionResourceCollection `json:"neutralDomainResources,omitempty"`

	// Another way to input protected apps through xml files
	ProtectedAppLockerFiles *[]WindowsInformationProtectionAppLockerFile `json:"protectedAppLockerFiles,omitempty"`

	// Protected applications can access enterprise data and the data handled by those applications are protected with
	// encryption
	ProtectedApps *[]WindowsInformationProtectionApp `json:"protectedApps,omitempty"`

	// Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured
	ProtectionUnderLockConfigRequired *bool `json:"protectionUnderLockConfigRequired,omitempty"`

	// This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1
	// (Don't revoke keys), the keys will not be revoked and the user will continue to have access to protected files after
	// unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently.
	RevokeOnUnenrollDisabled *bool `json:"revokeOnUnenrollDisabled,omitempty"`

	// TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who
	// has access to RMS-protected file and how long they have access
	RightsManagementServicesTemplateId nullable.Type[string] `json:"rightsManagementServicesTemplateId,omitempty"`

	// Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share
	// within the corporate boundary
	SmbAutoEncryptedFileExtensions *[]WindowsInformationProtectionResourceCollection `json:"smbAutoEncryptedFileExtensions,omitempty"`

	// Fields inherited from ManagedAppPolicy

	// The date and time the policy was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The policy's description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy display name.
	DisplayName *string `json:"displayName,omitempty"`

	// Last time the policy was modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s WindowsInformationProtectionPolicy) WindowsInformationProtection() BaseWindowsInformationProtectionImpl {
	return BaseWindowsInformationProtectionImpl{
		Assignments:                            s.Assignments,
		AzureRightsManagementServicesAllowed:   s.AzureRightsManagementServicesAllowed,
		DataRecoveryCertificate:                s.DataRecoveryCertificate,
		EnforcementLevel:                       s.EnforcementLevel,
		EnterpriseDomain:                       s.EnterpriseDomain,
		EnterpriseIPRanges:                     s.EnterpriseIPRanges,
		EnterpriseIPRangesAreAuthoritative:     s.EnterpriseIPRangesAreAuthoritative,
		EnterpriseInternalProxyServers:         s.EnterpriseInternalProxyServers,
		EnterpriseNetworkDomainNames:           s.EnterpriseNetworkDomainNames,
		EnterpriseProtectedDomainNames:         s.EnterpriseProtectedDomainNames,
		EnterpriseProxiedDomains:               s.EnterpriseProxiedDomains,
		EnterpriseProxyServers:                 s.EnterpriseProxyServers,
		EnterpriseProxyServersAreAuthoritative: s.EnterpriseProxyServersAreAuthoritative,
		ExemptAppLockerFiles:                   s.ExemptAppLockerFiles,
		ExemptApps:                             s.ExemptApps,
		IconsVisible:                           s.IconsVisible,
		IndexingEncryptedStoresOrItemsBlocked:  s.IndexingEncryptedStoresOrItemsBlocked,
		IsAssigned:                             s.IsAssigned,
		NeutralDomainResources:                 s.NeutralDomainResources,
		ProtectedAppLockerFiles:                s.ProtectedAppLockerFiles,
		ProtectedApps:                          s.ProtectedApps,
		ProtectionUnderLockConfigRequired:      s.ProtectionUnderLockConfigRequired,
		RevokeOnUnenrollDisabled:               s.RevokeOnUnenrollDisabled,
		RightsManagementServicesTemplateId:     s.RightsManagementServicesTemplateId,
		SmbAutoEncryptedFileExtensions:         s.SmbAutoEncryptedFileExtensions,
		CreatedDateTime:                        s.CreatedDateTime,
		Description:                            s.Description,
		DisplayName:                            s.DisplayName,
		LastModifiedDateTime:                   s.LastModifiedDateTime,
		Version:                                s.Version,
		Id:                                     s.Id,
		ODataId:                                s.ODataId,
		ODataType:                              s.ODataType,
	}
}

func (s WindowsInformationProtectionPolicy) ManagedAppPolicy() BaseManagedAppPolicyImpl {
	return BaseManagedAppPolicyImpl{
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Version:              s.Version,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s WindowsInformationProtectionPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsInformationProtectionPolicy{}

func (s WindowsInformationProtectionPolicy) MarshalJSON() ([]byte, error) {
	type wrapper WindowsInformationProtectionPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsInformationProtectionPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionPolicy: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windowsInformationProtectionPolicy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsInformationProtectionPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsInformationProtectionPolicy{}

func (s *WindowsInformationProtectionPolicy) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		DaysWithoutContactBeforeUnenroll       *int64                                                 `json:"daysWithoutContactBeforeUnenroll,omitempty"`
		MdmEnrollmentUrl                       nullable.Type[string]                                  `json:"mdmEnrollmentUrl,omitempty"`
		MinutesOfInactivityBeforeDeviceLock    *int64                                                 `json:"minutesOfInactivityBeforeDeviceLock,omitempty"`
		NumberOfPastPinsRemembered             *int64                                                 `json:"numberOfPastPinsRemembered,omitempty"`
		PasswordMaximumAttemptCount            *int64                                                 `json:"passwordMaximumAttemptCount,omitempty"`
		PinExpirationDays                      *int64                                                 `json:"pinExpirationDays,omitempty"`
		PinLowercaseLetters                    *WindowsInformationProtectionPinCharacterRequirements  `json:"pinLowercaseLetters,omitempty"`
		PinMinimumLength                       *int64                                                 `json:"pinMinimumLength,omitempty"`
		PinSpecialCharacters                   *WindowsInformationProtectionPinCharacterRequirements  `json:"pinSpecialCharacters,omitempty"`
		PinUppercaseLetters                    *WindowsInformationProtectionPinCharacterRequirements  `json:"pinUppercaseLetters,omitempty"`
		RevokeOnMdmHandoffDisabled             *bool                                                  `json:"revokeOnMdmHandoffDisabled,omitempty"`
		WindowsHelloForBusinessBlocked         *bool                                                  `json:"windowsHelloForBusinessBlocked,omitempty"`
		Assignments                            *[]TargetedManagedAppPolicyAssignment                  `json:"assignments,omitempty"`
		AzureRightsManagementServicesAllowed   *bool                                                  `json:"azureRightsManagementServicesAllowed,omitempty"`
		DataRecoveryCertificate                *WindowsInformationProtectionDataRecoveryCertificate   `json:"dataRecoveryCertificate,omitempty"`
		EnforcementLevel                       *WindowsInformationProtectionEnforcementLevel          `json:"enforcementLevel,omitempty"`
		EnterpriseDomain                       nullable.Type[string]                                  `json:"enterpriseDomain,omitempty"`
		EnterpriseIPRanges                     *[]WindowsInformationProtectionIPRangeCollection       `json:"enterpriseIPRanges,omitempty"`
		EnterpriseIPRangesAreAuthoritative     *bool                                                  `json:"enterpriseIPRangesAreAuthoritative,omitempty"`
		EnterpriseInternalProxyServers         *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseInternalProxyServers,omitempty"`
		EnterpriseNetworkDomainNames           *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseNetworkDomainNames,omitempty"`
		EnterpriseProtectedDomainNames         *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseProtectedDomainNames,omitempty"`
		EnterpriseProxiedDomains               *[]WindowsInformationProtectionProxiedDomainCollection `json:"enterpriseProxiedDomains,omitempty"`
		EnterpriseProxyServers                 *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseProxyServers,omitempty"`
		EnterpriseProxyServersAreAuthoritative *bool                                                  `json:"enterpriseProxyServersAreAuthoritative,omitempty"`
		ExemptAppLockerFiles                   *[]WindowsInformationProtectionAppLockerFile           `json:"exemptAppLockerFiles,omitempty"`
		ExemptApps                             *[]WindowsInformationProtectionApp                     `json:"exemptApps,omitempty"`
		IconsVisible                           *bool                                                  `json:"iconsVisible,omitempty"`
		IndexingEncryptedStoresOrItemsBlocked  *bool                                                  `json:"indexingEncryptedStoresOrItemsBlocked,omitempty"`
		IsAssigned                             *bool                                                  `json:"isAssigned,omitempty"`
		NeutralDomainResources                 *[]WindowsInformationProtectionResourceCollection      `json:"neutralDomainResources,omitempty"`
		ProtectedAppLockerFiles                *[]WindowsInformationProtectionAppLockerFile           `json:"protectedAppLockerFiles,omitempty"`
		ProtectedApps                          *[]WindowsInformationProtectionApp                     `json:"protectedApps,omitempty"`
		ProtectionUnderLockConfigRequired      *bool                                                  `json:"protectionUnderLockConfigRequired,omitempty"`
		RevokeOnUnenrollDisabled               *bool                                                  `json:"revokeOnUnenrollDisabled,omitempty"`
		RightsManagementServicesTemplateId     nullable.Type[string]                                  `json:"rightsManagementServicesTemplateId,omitempty"`
		SmbAutoEncryptedFileExtensions         *[]WindowsInformationProtectionResourceCollection      `json:"smbAutoEncryptedFileExtensions,omitempty"`
		CreatedDateTime                        *string                                                `json:"createdDateTime,omitempty"`
		Description                            nullable.Type[string]                                  `json:"description,omitempty"`
		DisplayName                            *string                                                `json:"displayName,omitempty"`
		LastModifiedDateTime                   *string                                                `json:"lastModifiedDateTime,omitempty"`
		Version                                nullable.Type[string]                                  `json:"version,omitempty"`
		Id                                     *string                                                `json:"id,omitempty"`
		ODataId                                *string                                                `json:"@odata.id,omitempty"`
		ODataType                              *string                                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DaysWithoutContactBeforeUnenroll = decoded.DaysWithoutContactBeforeUnenroll
	s.MdmEnrollmentUrl = decoded.MdmEnrollmentUrl
	s.MinutesOfInactivityBeforeDeviceLock = decoded.MinutesOfInactivityBeforeDeviceLock
	s.NumberOfPastPinsRemembered = decoded.NumberOfPastPinsRemembered
	s.PasswordMaximumAttemptCount = decoded.PasswordMaximumAttemptCount
	s.PinExpirationDays = decoded.PinExpirationDays
	s.PinLowercaseLetters = decoded.PinLowercaseLetters
	s.PinMinimumLength = decoded.PinMinimumLength
	s.PinSpecialCharacters = decoded.PinSpecialCharacters
	s.PinUppercaseLetters = decoded.PinUppercaseLetters
	s.RevokeOnMdmHandoffDisabled = decoded.RevokeOnMdmHandoffDisabled
	s.WindowsHelloForBusinessBlocked = decoded.WindowsHelloForBusinessBlocked
	s.Assignments = decoded.Assignments
	s.AzureRightsManagementServicesAllowed = decoded.AzureRightsManagementServicesAllowed
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DataRecoveryCertificate = decoded.DataRecoveryCertificate
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EnforcementLevel = decoded.EnforcementLevel
	s.EnterpriseDomain = decoded.EnterpriseDomain
	s.EnterpriseIPRanges = decoded.EnterpriseIPRanges
	s.EnterpriseIPRangesAreAuthoritative = decoded.EnterpriseIPRangesAreAuthoritative
	s.EnterpriseInternalProxyServers = decoded.EnterpriseInternalProxyServers
	s.EnterpriseNetworkDomainNames = decoded.EnterpriseNetworkDomainNames
	s.EnterpriseProtectedDomainNames = decoded.EnterpriseProtectedDomainNames
	s.EnterpriseProxiedDomains = decoded.EnterpriseProxiedDomains
	s.EnterpriseProxyServers = decoded.EnterpriseProxyServers
	s.EnterpriseProxyServersAreAuthoritative = decoded.EnterpriseProxyServersAreAuthoritative
	s.ExemptAppLockerFiles = decoded.ExemptAppLockerFiles
	s.IconsVisible = decoded.IconsVisible
	s.Id = decoded.Id
	s.IndexingEncryptedStoresOrItemsBlocked = decoded.IndexingEncryptedStoresOrItemsBlocked
	s.IsAssigned = decoded.IsAssigned
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.NeutralDomainResources = decoded.NeutralDomainResources
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ProtectedAppLockerFiles = decoded.ProtectedAppLockerFiles
	s.ProtectionUnderLockConfigRequired = decoded.ProtectionUnderLockConfigRequired
	s.RevokeOnUnenrollDisabled = decoded.RevokeOnUnenrollDisabled
	s.RightsManagementServicesTemplateId = decoded.RightsManagementServicesTemplateId
	s.SmbAutoEncryptedFileExtensions = decoded.SmbAutoEncryptedFileExtensions
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsInformationProtectionPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["exemptApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ExemptApps into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsInformationProtectionApp, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsInformationProtectionAppImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ExemptApps' for 'WindowsInformationProtectionPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ExemptApps = &output
	}

	if v, ok := temp["protectedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ProtectedApps into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsInformationProtectionApp, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsInformationProtectionAppImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ProtectedApps' for 'WindowsInformationProtectionPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ProtectedApps = &output
	}

	return nil
}
