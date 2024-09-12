package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ManagedAppProtection = DefaultManagedAppProtection{}

type DefaultManagedAppProtection struct {
	// Represents the level to which app data is encrypted for managed apps
	AppDataEncryptionType *ManagedAppDataEncryptionType `json:"appDataEncryptionType,omitempty"`

	// List of apps to which the policy is deployed.
	Apps *[]ManagedMobileApp `json:"apps,omitempty"`

	// A set of string key and string value pairs to be sent to the affected users, unalterned by this service
	CustomSettings *[]KeyValuePair `json:"customSettings,omitempty"`

	// Count of apps to which the current policy is deployed.
	DeployedAppCount *int64 `json:"deployedAppCount,omitempty"`

	// Navigation property to deployment summary of the configuration.
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`

	// When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
	DisableAppEncryptionIfDeviceEncryptionIsEnabled *bool `json:"disableAppEncryptionIfDeviceEncryptionIsEnabled,omitempty"`

	// Indicates whether managed-app data should be encrypted. (Android only)
	EncryptAppData *bool `json:"encryptAppData,omitempty"`

	// Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
	FaceIdBlocked *bool `json:"faceIdBlocked,omitempty"`

	// Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android
	// only)
	MinimumRequiredPatchVersion nullable.Type[string] `json:"minimumRequiredPatchVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
	MinimumRequiredSdkVersion nullable.Type[string] `json:"minimumRequiredSdkVersion,omitempty"`

	// Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android
	// only)
	MinimumWarningPatchVersion nullable.Type[string] `json:"minimumWarningPatchVersion,omitempty"`

	// Indicates whether screen capture is blocked. (Android only)
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Fields inherited from ManagedAppProtection

	// Data storage locations where a user may store managed data.
	AllowedDataStorageLocations *[]ManagedAppDataStorageLocation `json:"allowedDataStorageLocations,omitempty"`

	// Data can be transferred from/to these classes of apps
	AllowedInboundDataTransferSources *ManagedAppDataTransferLevel `json:"allowedInboundDataTransferSources,omitempty"`

	// Represents the level to which the device's clipboard may be shared between apps
	AllowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel `json:"allowedOutboundClipboardSharingLevel,omitempty"`

	// Data can be transferred from/to these classes of apps
	AllowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel `json:"allowedOutboundDataTransferDestinations,omitempty"`

	// Indicates whether contacts can be synced to the user's device.
	ContactSyncBlocked *bool `json:"contactSyncBlocked,omitempty"`

	// Indicates whether the backup of a managed app's data is blocked.
	DataBackupBlocked *bool `json:"dataBackupBlocked,omitempty"`

	// Indicates whether device compliance is required.
	DeviceComplianceRequired *bool `json:"deviceComplianceRequired,omitempty"`

	// Indicates whether use of the app pin is required if the device pin is set.
	DisableAppPinIfDevicePinIsSet *bool `json:"disableAppPinIfDevicePinIsSet,omitempty"`

	// Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
	FingerprintBlocked *bool `json:"fingerprintBlocked,omitempty"`

	// Type of managed browser
	ManagedBrowser *ManagedBrowserType `json:"managedBrowser,omitempty"`

	// Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by
	// CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
	ManagedBrowserToOpenLinksRequired *bool `json:"managedBrowserToOpenLinksRequired,omitempty"`

	// Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
	MaximumPinRetries *int64 `json:"maximumPinRetries,omitempty"`

	// Minimum pin length required for an app-level pin if PinRequired is set to True
	MinimumPinLength *int64 `json:"minimumPinLength,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredAppVersion nullable.Type[string] `json:"minimumRequiredAppVersion,omitempty"`

	// Versions less than the specified version will block the managed app from accessing company data.
	MinimumRequiredOsVersion nullable.Type[string] `json:"minimumRequiredOsVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app.
	MinimumWarningAppVersion nullable.Type[string] `json:"minimumWarningAppVersion,omitempty"`

	// Versions less than the specified version will result in warning message on the managed app from accessing company
	// data.
	MinimumWarningOsVersion nullable.Type[string] `json:"minimumWarningOsVersion,omitempty"`

	// Indicates whether organizational credentials are required for app use.
	OrganizationalCredentialsRequired *bool `json:"organizationalCredentialsRequired,omitempty"`

	// TimePeriod before the all-level pin must be reset if PinRequired is set to True.
	PeriodBeforePinReset *string `json:"periodBeforePinReset,omitempty"`

	// The period after which access is checked when the device is not connected to the internet.
	PeriodOfflineBeforeAccessCheck *string `json:"periodOfflineBeforeAccessCheck,omitempty"`

	// The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
	PeriodOfflineBeforeWipeIsEnforced *string `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`

	// The period after which access is checked when the device is connected to the internet.
	PeriodOnlineBeforeAccessCheck *string `json:"periodOnlineBeforeAccessCheck,omitempty"`

	// Character set which is to be used for a user's app PIN
	PinCharacterSet *ManagedAppPinCharacterSet `json:"pinCharacterSet,omitempty"`

	// Indicates whether an app-level pin is required.
	PinRequired *bool `json:"pinRequired,omitempty"`

	// Indicates whether printing is allowed from managed apps.
	PrintBlocked *bool `json:"printBlocked,omitempty"`

	// Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
	SaveAsBlocked *bool `json:"saveAsBlocked,omitempty"`

	// Indicates whether simplePin is blocked.
	SimplePinBlocked *bool `json:"simplePinBlocked,omitempty"`

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

func (s DefaultManagedAppProtection) ManagedAppProtection() BaseManagedAppProtectionImpl {
	return BaseManagedAppProtectionImpl{
		AllowedDataStorageLocations:             s.AllowedDataStorageLocations,
		AllowedInboundDataTransferSources:       s.AllowedInboundDataTransferSources,
		AllowedOutboundClipboardSharingLevel:    s.AllowedOutboundClipboardSharingLevel,
		AllowedOutboundDataTransferDestinations: s.AllowedOutboundDataTransferDestinations,
		ContactSyncBlocked:                      s.ContactSyncBlocked,
		DataBackupBlocked:                       s.DataBackupBlocked,
		DeviceComplianceRequired:                s.DeviceComplianceRequired,
		DisableAppPinIfDevicePinIsSet:           s.DisableAppPinIfDevicePinIsSet,
		FingerprintBlocked:                      s.FingerprintBlocked,
		ManagedBrowser:                          s.ManagedBrowser,
		ManagedBrowserToOpenLinksRequired:       s.ManagedBrowserToOpenLinksRequired,
		MaximumPinRetries:                       s.MaximumPinRetries,
		MinimumPinLength:                        s.MinimumPinLength,
		MinimumRequiredAppVersion:               s.MinimumRequiredAppVersion,
		MinimumRequiredOsVersion:                s.MinimumRequiredOsVersion,
		MinimumWarningAppVersion:                s.MinimumWarningAppVersion,
		MinimumWarningOsVersion:                 s.MinimumWarningOsVersion,
		OrganizationalCredentialsRequired:       s.OrganizationalCredentialsRequired,
		PeriodBeforePinReset:                    s.PeriodBeforePinReset,
		PeriodOfflineBeforeAccessCheck:          s.PeriodOfflineBeforeAccessCheck,
		PeriodOfflineBeforeWipeIsEnforced:       s.PeriodOfflineBeforeWipeIsEnforced,
		PeriodOnlineBeforeAccessCheck:           s.PeriodOnlineBeforeAccessCheck,
		PinCharacterSet:                         s.PinCharacterSet,
		PinRequired:                             s.PinRequired,
		PrintBlocked:                            s.PrintBlocked,
		SaveAsBlocked:                           s.SaveAsBlocked,
		SimplePinBlocked:                        s.SimplePinBlocked,
		CreatedDateTime:                         s.CreatedDateTime,
		Description:                             s.Description,
		DisplayName:                             s.DisplayName,
		LastModifiedDateTime:                    s.LastModifiedDateTime,
		Version:                                 s.Version,
		Id:                                      s.Id,
		ODataId:                                 s.ODataId,
		ODataType:                               s.ODataType,
	}
}

func (s DefaultManagedAppProtection) ManagedAppPolicy() BaseManagedAppPolicyImpl {
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

func (s DefaultManagedAppProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DefaultManagedAppProtection{}

func (s DefaultManagedAppProtection) MarshalJSON() ([]byte, error) {
	type wrapper DefaultManagedAppProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefaultManagedAppProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefaultManagedAppProtection: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.defaultManagedAppProtection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefaultManagedAppProtection: %+v", err)
	}

	return encoded, nil
}
