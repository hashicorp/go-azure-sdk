package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10GeneralConfiguration{}

type Windows10GeneralConfiguration struct {
	// Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a
	// Microsoft account.
	AccountsBlockAddingNonMicrosoftAccountEmail *bool `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`

	// Indicates whether or not to block the user from selecting an AntiTheft mode preference (Windows 10 Mobile only).
	AntiTheftModeBlocked *bool `json:"antiTheftModeBlocked,omitempty"`

	// State Management Setting.
	AppsAllowTrustedAppsSideloading *StateManagementSetting `json:"appsAllowTrustedAppsSideloading,omitempty"`

	// Indicates whether or not to disable the launch of all apps from Windows Store that came pre-installed or were
	// downloaded.
	AppsBlockWindowsStoreOriginatedApps *bool `json:"appsBlockWindowsStoreOriginatedApps,omitempty"`

	// Specify a list of allowed Bluetooth services and profiles in hex formatted strings.
	BluetoothAllowedServices *[]string `json:"bluetoothAllowedServices,omitempty"`

	// Whether or not to Block the user from using bluetooth advertising.
	BluetoothBlockAdvertising *bool `json:"bluetoothBlockAdvertising,omitempty"`

	// Whether or not to Block the user from using bluetooth discoverable mode.
	BluetoothBlockDiscoverableMode *bool `json:"bluetoothBlockDiscoverableMode,omitempty"`

	// Whether or not to block specific bundled Bluetooth peripherals to automatically pair with the host device.
	BluetoothBlockPrePairing *bool `json:"bluetoothBlockPrePairing,omitempty"`

	// Whether or not to Block the user from using bluetooth.
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`

	// Whether or not to Block the user from accessing the camera of the device.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`

	// Whether or not to Block the user from using data over cellular while roaming.
	CellularBlockDataWhenRoaming *bool `json:"cellularBlockDataWhenRoaming,omitempty"`

	// Whether or not to Block the user from using VPN over cellular.
	CellularBlockVpn *bool `json:"cellularBlockVpn,omitempty"`

	// Whether or not to Block the user from using VPN when roaming over cellular.
	CellularBlockVpnWhenRoaming *bool `json:"cellularBlockVpnWhenRoaming,omitempty"`

	// Whether or not to Block the user from doing manual root certificate installation.
	CertificatesBlockManualRootCertificateInstallation *bool `json:"certificatesBlockManualRootCertificateInstallation,omitempty"`

	// Whether or not to block Connected Devices Service which enables discovery and connection to other devices, remote
	// messaging, remote app sessions and other cross-device experiences.
	ConnectedDevicesServiceBlocked *bool `json:"connectedDevicesServiceBlocked,omitempty"`

	// Whether or not to Block the user from using copy paste.
	CopyPasteBlocked *bool `json:"copyPasteBlocked,omitempty"`

	// Whether or not to Block the user from using Cortana.
	CortanaBlocked *bool `json:"cortanaBlocked,omitempty"`

	// Whether or not to block end user access to Defender.
	DefenderBlockEndUserAccess *bool `json:"defenderBlockEndUserAccess,omitempty"`

	// Possible values of Cloud Block Level
	DefenderCloudBlockLevel *DefenderCloudBlockLevelType `json:"defenderCloudBlockLevel,omitempty"`

	// Number of days before deleting quarantined malware. Valid values 0 to 90
	DefenderDaysBeforeDeletingQuarantinedMalware nullable.Type[int64] `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`

	// Gets or sets Defender’s actions to take on detected Malware per threat level.
	DefenderDetectedMalwareActions *DefenderDetectedMalwareActions `json:"defenderDetectedMalwareActions,omitempty"`

	// File extensions to exclude from scans and real time protection.
	DefenderFileExtensionsToExclude *[]string `json:"defenderFileExtensionsToExclude,omitempty"`

	// Files and folder to exclude from scans and real time protection.
	DefenderFilesAndFoldersToExclude *[]string `json:"defenderFilesAndFoldersToExclude,omitempty"`

	// Possible values for monitoring file activity.
	DefenderMonitorFileActivity *DefenderMonitorFileActivity `json:"defenderMonitorFileActivity,omitempty"`

	// Processes to exclude from scans and real time protection.
	DefenderProcessesToExclude *[]string `json:"defenderProcessesToExclude,omitempty"`

	// Possible values for prompting user for samples submission.
	DefenderPromptForSampleSubmission *DefenderPromptForSampleSubmission `json:"defenderPromptForSampleSubmission,omitempty"`

	// Indicates whether or not to require behavior monitoring.
	DefenderRequireBehaviorMonitoring *bool `json:"defenderRequireBehaviorMonitoring,omitempty"`

	// Indicates whether or not to require cloud protection.
	DefenderRequireCloudProtection *bool `json:"defenderRequireCloudProtection,omitempty"`

	// Indicates whether or not to require network inspection system.
	DefenderRequireNetworkInspectionSystem *bool `json:"defenderRequireNetworkInspectionSystem,omitempty"`

	// Indicates whether or not to require real time monitoring.
	DefenderRequireRealTimeMonitoring *bool `json:"defenderRequireRealTimeMonitoring,omitempty"`

	// Indicates whether or not to scan archive files.
	DefenderScanArchiveFiles *bool `json:"defenderScanArchiveFiles,omitempty"`

	// Indicates whether or not to scan downloads.
	DefenderScanDownloads *bool `json:"defenderScanDownloads,omitempty"`

	// Indicates whether or not to scan incoming mail messages.
	DefenderScanIncomingMail *bool `json:"defenderScanIncomingMail,omitempty"`

	// Indicates whether or not to scan mapped network drives during full scan.
	DefenderScanMappedNetworkDrivesDuringFullScan *bool `json:"defenderScanMappedNetworkDrivesDuringFullScan,omitempty"`

	// Max CPU usage percentage during scan. Valid values 0 to 100
	DefenderScanMaxCpu nullable.Type[int64] `json:"defenderScanMaxCpu,omitempty"`

	// Indicates whether or not to scan files opened from a network folder.
	DefenderScanNetworkFiles *bool `json:"defenderScanNetworkFiles,omitempty"`

	// Indicates whether or not to scan removable drives during full scan.
	DefenderScanRemovableDrivesDuringFullScan *bool `json:"defenderScanRemovableDrivesDuringFullScan,omitempty"`

	// Indicates whether or not to scan scripts loaded in Internet Explorer browser.
	DefenderScanScriptsLoadedInInternetExplorer *bool `json:"defenderScanScriptsLoadedInInternetExplorer,omitempty"`

	// Possible values for system scan type.
	DefenderScanType *DefenderScanType `json:"defenderScanType,omitempty"`

	// The time to perform a daily quick scan.
	DefenderScheduledQuickScanTime nullable.Type[string] `json:"defenderScheduledQuickScanTime,omitempty"`

	// The defender time for the system scan.
	DefenderScheduledScanTime nullable.Type[string] `json:"defenderScheduledScanTime,omitempty"`

	// The signature update interval in hours. Specify 0 not to check. Valid values 0 to 24
	DefenderSignatureUpdateIntervalInHours nullable.Type[int64] `json:"defenderSignatureUpdateIntervalInHours,omitempty"`

	// Possible values for a weekly schedule.
	DefenderSystemScanSchedule *WeeklySchedule `json:"defenderSystemScanSchedule,omitempty"`

	// State Management Setting.
	DeveloperUnlockSetting *StateManagementSetting `json:"developerUnlockSetting,omitempty"`

	// Indicates whether or not to Block the user from resetting their phone.
	DeviceManagementBlockFactoryResetOnMobile *bool `json:"deviceManagementBlockFactoryResetOnMobile,omitempty"`

	// Indicates whether or not to Block the user from doing manual un-enrollment from device management.
	DeviceManagementBlockManualUnenroll *bool `json:"deviceManagementBlockManualUnenroll,omitempty"`

	// Allow the device to send diagnostic and usage telemetry data, such as Watson.
	DiagnosticsDataSubmissionMode *DiagnosticDataSubmissionMode `json:"diagnosticsDataSubmissionMode,omitempty"`

	// Allow users to change Start pages on Edge. Use the EdgeHomepageUrls to specify the Start pages that the user would
	// see by default when they open Edge.
	EdgeAllowStartPagesModification *bool `json:"edgeAllowStartPagesModification,omitempty"`

	// Indicates whether or not to prevent access to about flags on Edge browser.
	EdgeBlockAccessToAboutFlags *bool `json:"edgeBlockAccessToAboutFlags,omitempty"`

	// Block the address bar dropdown functionality in Microsoft Edge. Disable this settings to minimize network connections
	// from Microsoft Edge to Microsoft services.
	EdgeBlockAddressBarDropdown *bool `json:"edgeBlockAddressBarDropdown,omitempty"`

	// Indicates whether or not to block auto fill.
	EdgeBlockAutofill *bool `json:"edgeBlockAutofill,omitempty"`

	// Block Microsoft compatibility list in Microsoft Edge. This list from Microsoft helps Edge properly display sites with
	// known compatibility issues.
	EdgeBlockCompatibilityList *bool `json:"edgeBlockCompatibilityList,omitempty"`

	// Indicates whether or not to block developer tools in the Edge browser.
	EdgeBlockDeveloperTools *bool `json:"edgeBlockDeveloperTools,omitempty"`

	// Indicates whether or not to block extensions in the Edge browser.
	EdgeBlockExtensions *bool `json:"edgeBlockExtensions,omitempty"`

	// Indicates whether or not to block InPrivate browsing on corporate networks, in the Edge browser.
	EdgeBlockInPrivateBrowsing *bool `json:"edgeBlockInPrivateBrowsing,omitempty"`

	// Indicates whether or not to Block the user from using JavaScript.
	EdgeBlockJavaScript *bool `json:"edgeBlockJavaScript,omitempty"`

	// Block the collection of information by Microsoft for live tile creation when users pin a site to Start from Microsoft
	// Edge.
	EdgeBlockLiveTileDataCollection *bool `json:"edgeBlockLiveTileDataCollection,omitempty"`

	// Indicates whether or not to Block password manager.
	EdgeBlockPasswordManager *bool `json:"edgeBlockPasswordManager,omitempty"`

	// Indicates whether or not to block popups.
	EdgeBlockPopups *bool `json:"edgeBlockPopups,omitempty"`

	// Indicates whether or not to block the user from using the search suggestions in the address bar.
	EdgeBlockSearchSuggestions *bool `json:"edgeBlockSearchSuggestions,omitempty"`

	// Indicates whether or not to Block the user from sending the do not track header.
	EdgeBlockSendingDoNotTrackHeader *bool `json:"edgeBlockSendingDoNotTrackHeader,omitempty"`

	// Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. Note: the name of this
	// property is misleading; the property is obsolete, use EdgeSendIntranetTrafficToInternetExplorer instead.
	EdgeBlockSendingIntranetTrafficToInternetExplorer *bool `json:"edgeBlockSendingIntranetTrafficToInternetExplorer,omitempty"`

	// Indicates whether or not to Block the user from using the Edge browser.
	EdgeBlocked *bool `json:"edgeBlocked,omitempty"`

	// Clear browsing data on exiting Microsoft Edge.
	EdgeClearBrowsingDataOnExit *bool `json:"edgeClearBrowsingDataOnExit,omitempty"`

	// Possible values to specify which cookies are allowed in Microsoft Edge.
	EdgeCookiePolicy *EdgeCookiePolicy `json:"edgeCookiePolicy,omitempty"`

	// Block the Microsoft web page that opens on the first use of Microsoft Edge. This policy allows enterprises, like
	// those enrolled in zero emissions configurations, to block this page.
	EdgeDisableFirstRunPage *bool `json:"edgeDisableFirstRunPage,omitempty"`

	// Indicates the enterprise mode site list location. Could be a local file, local network or http location.
	EdgeEnterpriseModeSiteListLocation nullable.Type[string] `json:"edgeEnterpriseModeSiteListLocation,omitempty"`

	// The first run URL for when Edge browser is opened for the first time.
	EdgeFirstRunUrl nullable.Type[string] `json:"edgeFirstRunUrl,omitempty"`

	// The list of URLs for homepages shodwn on MDM-enrolled devices on Edge browser.
	EdgeHomepageUrls *[]string `json:"edgeHomepageUrls,omitempty"`

	// Indicates whether or not to Require the user to use the smart screen filter.
	EdgeRequireSmartScreen *bool `json:"edgeRequireSmartScreen,omitempty"`

	// Allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their
	// default search engine provided the AllowSearchEngineCustomization policy is not set.
	EdgeSearchEngine EdgeSearchEngineBase `json:"edgeSearchEngine"`

	// Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer.
	EdgeSendIntranetTrafficToInternetExplorer *bool `json:"edgeSendIntranetTrafficToInternetExplorer,omitempty"`

	// Enable favorites sync between Internet Explorer and Microsoft Edge. Additions, deletions, modifications and order
	// changes to favorites are shared between browsers.
	EdgeSyncFavoritesWithInternetExplorer *bool `json:"edgeSyncFavoritesWithInternetExplorer,omitempty"`

	// Endpoint for discovering cloud printers.
	EnterpriseCloudPrintDiscoveryEndPoint nullable.Type[string] `json:"enterpriseCloudPrintDiscoveryEndPoint,omitempty"`

	// Maximum number of printers that should be queried from a discovery endpoint. This is a mobile only setting. Valid
	// values 1 to 65535
	EnterpriseCloudPrintDiscoveryMaxLimit nullable.Type[int64] `json:"enterpriseCloudPrintDiscoveryMaxLimit,omitempty"`

	// OAuth resource URI for printer discovery service as configured in Azure portal.
	EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier nullable.Type[string] `json:"enterpriseCloudPrintMopriaDiscoveryResourceIdentifier,omitempty"`

	// Authentication endpoint for acquiring OAuth tokens.
	EnterpriseCloudPrintOAuthAuthority nullable.Type[string] `json:"enterpriseCloudPrintOAuthAuthority,omitempty"`

	// GUID of a client application authorized to retrieve OAuth tokens from the OAuth Authority.
	EnterpriseCloudPrintOAuthClientIdentifier nullable.Type[string] `json:"enterpriseCloudPrintOAuthClientIdentifier,omitempty"`

	// OAuth resource URI for print service as configured in the Azure portal.
	EnterpriseCloudPrintResourceIdentifier nullable.Type[string] `json:"enterpriseCloudPrintResourceIdentifier,omitempty"`

	// Indicates whether or not to enable device discovery UX.
	ExperienceBlockDeviceDiscovery *bool `json:"experienceBlockDeviceDiscovery,omitempty"`

	// Indicates whether or not to allow the error dialog from displaying if no SIM card is detected.
	ExperienceBlockErrorDialogWhenNoSIM *bool `json:"experienceBlockErrorDialogWhenNoSIM,omitempty"`

	// Indicates whether or not to enable task switching on the device.
	ExperienceBlockTaskSwitcher *bool `json:"experienceBlockTaskSwitcher,omitempty"`

	// Indicates whether or not to block DVR and broadcasting.
	GameDvrBlocked *bool `json:"gameDvrBlocked,omitempty"`

	// Indicates whether or not to Block the user from using internet sharing.
	InternetSharingBlocked *bool `json:"internetSharingBlocked,omitempty"`

	// Indicates whether or not to Block the user from location services.
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`

	// Specify whether to show a user-configurable setting to control the screen timeout while on the lock screen of Windows
	// 10 Mobile devices. If this policy is set to Allow, the value set by lockScreenTimeoutInSeconds is ignored.
	LockScreenAllowTimeoutConfiguration *bool `json:"lockScreenAllowTimeoutConfiguration,omitempty"`

	// Indicates whether or not to block action center notifications over lock screen.
	LockScreenBlockActionCenterNotifications *bool `json:"lockScreenBlockActionCenterNotifications,omitempty"`

	// Indicates whether or not the user can interact with Cortana using speech while the system is locked.
	LockScreenBlockCortana *bool `json:"lockScreenBlockCortana,omitempty"`

	// Indicates whether to allow toast notifications above the device lock screen.
	LockScreenBlockToastNotifications *bool `json:"lockScreenBlockToastNotifications,omitempty"`

	// Set the duration (in seconds) from the screen locking to the screen turning off for Windows 10 Mobile devices.
	// Supported values are 11-1800. Valid values 11 to 1800
	LockScreenTimeoutInSeconds nullable.Type[int64] `json:"lockScreenTimeoutInSeconds,omitempty"`

	// Disables the ability to quickly switch between users that are logged on simultaneously without logging off.
	LogonBlockFastUserSwitching *bool `json:"logonBlockFastUserSwitching,omitempty"`

	// Indicates whether or not to Block Microsoft account settings sync.
	MicrosoftAccountBlockSettingsSync *bool `json:"microsoftAccountBlockSettingsSync,omitempty"`

	// Indicates whether or not to Block a Microsoft account.
	MicrosoftAccountBlocked *bool `json:"microsoftAccountBlocked,omitempty"`

	// If set, proxy settings will be applied to all processes and accounts in the device. Otherwise, it will be applied to
	// the user account that’s enrolled into MDM.
	NetworkProxyApplySettingsDeviceWide *bool `json:"networkProxyApplySettingsDeviceWide,omitempty"`

	// Address to the proxy auto-config (PAC) script you want to use.
	NetworkProxyAutomaticConfigurationUrl nullable.Type[string] `json:"networkProxyAutomaticConfigurationUrl,omitempty"`

	// Disable automatic detection of settings. If enabled, the system will try to find the path to a proxy auto-config
	// (PAC) script.
	NetworkProxyDisableAutoDetect *bool `json:"networkProxyDisableAutoDetect,omitempty"`

	// Specifies manual proxy server settings.
	NetworkProxyServer *Windows10NetworkProxyServer `json:"networkProxyServer,omitempty"`

	// Indicates whether or not to Block the user from using near field communication.
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`

	// Gets or sets a value allowing IT admins to prevent apps and features from working with files on OneDrive.
	OneDriveDisableFileSync *bool `json:"oneDriveDisableFileSync,omitempty"`

	// Specify whether PINs or passwords such as '1111' or '1234' are allowed. For Windows 10 desktops, it also controls the
	// use of picture passwords.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// The password expiration in days. Valid values 0 to 730
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// The minimum password length. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// The minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// The number of previous passwords to prevent reuse of. Valid values 0 to 50
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Indicates whether or not to require a password upon resuming from an idle state.
	PasswordRequireWhenResumeFromIdleState *bool `json:"passwordRequireWhenResumeFromIdleState,omitempty"`

	// Indicates whether or not to require the user to have a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// The number of sign in failures before factory reset. Valid values 0 to 999
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// A http or https Url to a jpg, jpeg or png image that needs to be downloaded and used as the Desktop Image or a file
	// Url to a local image on the file system that needs to used as the Desktop Image.
	PersonalizationDesktopImageUrl nullable.Type[string] `json:"personalizationDesktopImageUrl,omitempty"`

	// A http or https Url to a jpg, jpeg or png image that neeeds to be downloaded and used as the Lock Screen Image or a
	// file Url to a local image on the file system that needs to be used as the Lock Screen Image.
	PersonalizationLockScreenImageUrl nullable.Type[string] `json:"personalizationLockScreenImageUrl,omitempty"`

	// State Management Setting.
	PrivacyAdvertisingId *StateManagementSetting `json:"privacyAdvertisingId,omitempty"`

	// Indicates whether or not to allow the automatic acceptance of the pairing and privacy user consent dialog when
	// launching apps.
	PrivacyAutoAcceptPairingAndConsentPrompts *bool `json:"privacyAutoAcceptPairingAndConsentPrompts,omitempty"`

	// Indicates whether or not to block the usage of cloud based speech services for Cortana, Dictation, or Store
	// applications.
	PrivacyBlockInputPersonalization *bool `json:"privacyBlockInputPersonalization,omitempty"`

	// Indicates whether or not to Block the user from reset protection mode.
	ResetProtectionModeBlocked *bool `json:"resetProtectionModeBlocked,omitempty"`

	// Specifies what level of safe search (filtering adult content) is required
	SafeSearchFilter *SafeSearchFilterType `json:"safeSearchFilter,omitempty"`

	// Indicates whether or not to Block the user from taking Screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Specifies if search can use diacritics.
	SearchBlockDiacritics *bool `json:"searchBlockDiacritics,omitempty"`

	// Specifies whether to use automatic language detection when indexing content and properties.
	SearchDisableAutoLanguageDetection *bool `json:"searchDisableAutoLanguageDetection,omitempty"`

	// Indicates whether or not to disable the search indexer backoff feature.
	SearchDisableIndexerBackoff *bool `json:"searchDisableIndexerBackoff,omitempty"`

	// Indicates whether or not to block indexing of WIP-protected items to prevent them from appearing in search results
	// for Cortana or Explorer.
	SearchDisableIndexingEncryptedItems *bool `json:"searchDisableIndexingEncryptedItems,omitempty"`

	// Indicates whether or not to allow users to add locations on removable drives to libraries and to be indexed.
	SearchDisableIndexingRemovableDrive *bool `json:"searchDisableIndexingRemovableDrive,omitempty"`

	// Specifies minimum amount of hard drive space on the same drive as the index location before indexing stops.
	SearchEnableAutomaticIndexSizeManangement *bool `json:"searchEnableAutomaticIndexSizeManangement,omitempty"`

	// Indicates whether or not to block remote queries of this computer’s index.
	SearchEnableRemoteQueries *bool `json:"searchEnableRemoteQueries,omitempty"`

	// Indicates whether or not to block access to Accounts in Settings app.
	SettingsBlockAccountsPage *bool `json:"settingsBlockAccountsPage,omitempty"`

	// Indicates whether or not to block the user from installing provisioning packages.
	SettingsBlockAddProvisioningPackage *bool `json:"settingsBlockAddProvisioningPackage,omitempty"`

	// Indicates whether or not to block access to Apps in Settings app.
	SettingsBlockAppsPage *bool `json:"settingsBlockAppsPage,omitempty"`

	// Indicates whether or not to block the user from changing the language settings.
	SettingsBlockChangeLanguage *bool `json:"settingsBlockChangeLanguage,omitempty"`

	// Indicates whether or not to block the user from changing power and sleep settings.
	SettingsBlockChangePowerSleep *bool `json:"settingsBlockChangePowerSleep,omitempty"`

	// Indicates whether or not to block the user from changing the region settings.
	SettingsBlockChangeRegion *bool `json:"settingsBlockChangeRegion,omitempty"`

	// Indicates whether or not to block the user from changing date and time settings.
	SettingsBlockChangeSystemTime *bool `json:"settingsBlockChangeSystemTime,omitempty"`

	// Indicates whether or not to block access to Devices in Settings app.
	SettingsBlockDevicesPage *bool `json:"settingsBlockDevicesPage,omitempty"`

	// Indicates whether or not to block access to Ease of Access in Settings app.
	SettingsBlockEaseOfAccessPage *bool `json:"settingsBlockEaseOfAccessPage,omitempty"`

	// Indicates whether or not to block the user from editing the device name.
	SettingsBlockEditDeviceName *bool `json:"settingsBlockEditDeviceName,omitempty"`

	// Indicates whether or not to block access to Gaming in Settings app.
	SettingsBlockGamingPage *bool `json:"settingsBlockGamingPage,omitempty"`

	// Indicates whether or not to block access to Network & Internet in Settings app.
	SettingsBlockNetworkInternetPage *bool `json:"settingsBlockNetworkInternetPage,omitempty"`

	// Indicates whether or not to block access to Personalization in Settings app.
	SettingsBlockPersonalizationPage *bool `json:"settingsBlockPersonalizationPage,omitempty"`

	// Indicates whether or not to block access to Privacy in Settings app.
	SettingsBlockPrivacyPage *bool `json:"settingsBlockPrivacyPage,omitempty"`

	// Indicates whether or not to block the runtime configuration agent from removing provisioning packages.
	SettingsBlockRemoveProvisioningPackage *bool `json:"settingsBlockRemoveProvisioningPackage,omitempty"`

	// Indicates whether or not to block access to Settings app.
	SettingsBlockSettingsApp *bool `json:"settingsBlockSettingsApp,omitempty"`

	// Indicates whether or not to block access to System in Settings app.
	SettingsBlockSystemPage *bool `json:"settingsBlockSystemPage,omitempty"`

	// Indicates whether or not to block access to Time & Language in Settings app.
	SettingsBlockTimeLanguagePage *bool `json:"settingsBlockTimeLanguagePage,omitempty"`

	// Indicates whether or not to block access to Update & Security in Settings app.
	SettingsBlockUpdateSecurityPage *bool `json:"settingsBlockUpdateSecurityPage,omitempty"`

	// Indicates whether or not to block multiple users of the same app to share data.
	SharedUserAppDataAllowed *bool `json:"sharedUserAppDataAllowed,omitempty"`

	// Indicates whether or not users can override SmartScreen Filter warnings about potentially malicious websites.
	SmartScreenBlockPromptOverride *bool `json:"smartScreenBlockPromptOverride,omitempty"`

	// Indicates whether or not users can override the SmartScreen Filter warnings about downloading unverified files
	SmartScreenBlockPromptOverrideForFiles *bool `json:"smartScreenBlockPromptOverrideForFiles,omitempty"`

	// This property will be deprecated in July 2019 and will be replaced by property SmartScreenAppInstallControl. Allows
	// IT Admins to control whether users are allowed to install apps from places other than the Store.
	SmartScreenEnableAppInstallControl *bool `json:"smartScreenEnableAppInstallControl,omitempty"`

	// Indicates whether or not to block the user from unpinning apps from taskbar.
	StartBlockUnpinningAppsFromTaskbar *bool `json:"startBlockUnpinningAppsFromTaskbar,omitempty"`

	// Type of start menu app list visibility.
	StartMenuAppListVisibility *WindowsStartMenuAppListVisibilityType `json:"startMenuAppListVisibility,omitempty"`

	// Enabling this policy hides the change account setting from appearing in the user tile in the start menu.
	StartMenuHideChangeAccountSettings *bool `json:"startMenuHideChangeAccountSettings,omitempty"`

	// Enabling this policy hides the most used apps from appearing on the start menu and disables the corresponding toggle
	// in the Settings app.
	StartMenuHideFrequentlyUsedApps *bool `json:"startMenuHideFrequentlyUsedApps,omitempty"`

	// Enabling this policy hides hibernate from appearing in the power button in the start menu.
	StartMenuHideHibernate *bool `json:"startMenuHideHibernate,omitempty"`

	// Enabling this policy hides lock from appearing in the user tile in the start menu.
	StartMenuHideLock *bool `json:"startMenuHideLock,omitempty"`

	// Enabling this policy hides the power button from appearing in the start menu.
	StartMenuHidePowerButton *bool `json:"startMenuHidePowerButton,omitempty"`

	// Enabling this policy hides recent jump lists from appearing on the start menu/taskbar and disables the corresponding
	// toggle in the Settings app.
	StartMenuHideRecentJumpLists *bool `json:"startMenuHideRecentJumpLists,omitempty"`

	// Enabling this policy hides recently added apps from appearing on the start menu and disables the corresponding toggle
	// in the Settings app.
	StartMenuHideRecentlyAddedApps *bool `json:"startMenuHideRecentlyAddedApps,omitempty"`

	// Enabling this policy hides 'Restart/Update and Restart' from appearing in the power button in the start menu.
	StartMenuHideRestartOptions *bool `json:"startMenuHideRestartOptions,omitempty"`

	// Enabling this policy hides shut down/update and shut down from appearing in the power button in the start menu.
	StartMenuHideShutDown *bool `json:"startMenuHideShutDown,omitempty"`

	// Enabling this policy hides sign out from appearing in the user tile in the start menu.
	StartMenuHideSignOut *bool `json:"startMenuHideSignOut,omitempty"`

	// Enabling this policy hides sleep from appearing in the power button in the start menu.
	StartMenuHideSleep *bool `json:"startMenuHideSleep,omitempty"`

	// Enabling this policy hides switch account from appearing in the user tile in the start menu.
	StartMenuHideSwitchAccount *bool `json:"startMenuHideSwitchAccount,omitempty"`

	// Enabling this policy hides the user tile from appearing in the start menu.
	StartMenuHideUserTile *bool `json:"startMenuHideUserTile,omitempty"`

	// This policy setting allows you to import Edge assets to be used with startMenuLayoutXml policy. Start layout can
	// contain secondary tile from Edge app which looks for Edge local asset file. Edge local asset would not exist and
	// cause Edge secondary tile to appear empty in this case. This policy only gets applied when startMenuLayoutXml policy
	// is modified. The value should be a UTF-8 Base64 encoded byte array.
	StartMenuLayoutEdgeAssetsXml nullable.Type[string] `json:"startMenuLayoutEdgeAssetsXml,omitempty"`

	// Allows admins to override the default Start menu layout and prevents the user from changing it. The layout is
	// modified by specifying an XML file based on a layout modification schema. XML needs to be in a UTF8 encoded byte
	// array format.
	StartMenuLayoutXml nullable.Type[string] `json:"startMenuLayoutXml,omitempty"`

	// Type of display modes for the start menu.
	StartMenuMode *WindowsStartMenuModeType `json:"startMenuMode,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderDocuments *VisibilitySetting `json:"startMenuPinnedFolderDocuments,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderDownloads *VisibilitySetting `json:"startMenuPinnedFolderDownloads,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderFileExplorer *VisibilitySetting `json:"startMenuPinnedFolderFileExplorer,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderHomeGroup *VisibilitySetting `json:"startMenuPinnedFolderHomeGroup,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderMusic *VisibilitySetting `json:"startMenuPinnedFolderMusic,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderNetwork *VisibilitySetting `json:"startMenuPinnedFolderNetwork,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderPersonalFolder *VisibilitySetting `json:"startMenuPinnedFolderPersonalFolder,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderPictures *VisibilitySetting `json:"startMenuPinnedFolderPictures,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderSettings *VisibilitySetting `json:"startMenuPinnedFolderSettings,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderVideos *VisibilitySetting `json:"startMenuPinnedFolderVideos,omitempty"`

	// Indicates whether or not to Block the user from using removable storage.
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`

	// Indicating whether or not to require encryption on a mobile device.
	StorageRequireMobileDeviceEncryption *bool `json:"storageRequireMobileDeviceEncryption,omitempty"`

	// Indicates whether application data is restricted to the system drive.
	StorageRestrictAppDataToSystemVolume *bool `json:"storageRestrictAppDataToSystemVolume,omitempty"`

	// Indicates whether the installation of applications is restricted to the system drive.
	StorageRestrictAppInstallToSystemVolume *bool `json:"storageRestrictAppInstallToSystemVolume,omitempty"`

	// Whether the device is required to connect to the network.
	TenantLockdownRequireNetworkDuringOutOfBoxExperience *bool `json:"tenantLockdownRequireNetworkDuringOutOfBoxExperience,omitempty"`

	// Indicates whether or not to Block the user from USB connection.
	UsbBlocked *bool `json:"usbBlocked,omitempty"`

	// Indicates whether or not to Block the user from voice recording.
	VoiceRecordingBlocked *bool `json:"voiceRecordingBlocked,omitempty"`

	// Indicates whether or not user's localhost IP address is displayed while making phone calls using the WebRTC
	WebRtcBlockLocalhostIPAddress *bool `json:"webRtcBlockLocalhostIpAddress,omitempty"`

	// Indicating whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
	WiFiBlockAutomaticConnectHotspots *bool `json:"wiFiBlockAutomaticConnectHotspots,omitempty"`

	// Indicates whether or not to Block the user from using Wi-Fi manual configuration.
	WiFiBlockManualConfiguration *bool `json:"wiFiBlockManualConfiguration,omitempty"`

	// Indicates whether or not to Block the user from using Wi-Fi.
	WiFiBlocked *bool `json:"wiFiBlocked,omitempty"`

	// Specify how often devices scan for Wi-Fi networks. Supported values are 1-500, where 100 = default, and 500 = low
	// frequency. Valid values 1 to 500
	WiFiScanInterval nullable.Type[int64] `json:"wiFiScanInterval,omitempty"`

	// Allows IT admins to block experiences that are typically for consumers only, such as Start suggestions, Membership
	// notifications, Post-OOBE app install and redirect tiles.
	WindowsSpotlightBlockConsumerSpecificFeatures *bool `json:"windowsSpotlightBlockConsumerSpecificFeatures,omitempty"`

	// Block suggestions from Microsoft that show after each OS clean install, upgrade or in an on-going basis to introduce
	// users to what is new or changed
	WindowsSpotlightBlockOnActionCenter *bool `json:"windowsSpotlightBlockOnActionCenter,omitempty"`

	// Block personalized content in Windows spotlight based on user’s device usage.
	WindowsSpotlightBlockTailoredExperiences *bool `json:"windowsSpotlightBlockTailoredExperiences,omitempty"`

	// Block third party content delivered via Windows Spotlight
	WindowsSpotlightBlockThirdPartyNotifications *bool `json:"windowsSpotlightBlockThirdPartyNotifications,omitempty"`

	// Block Windows Spotlight Windows welcome experience
	WindowsSpotlightBlockWelcomeExperience *bool `json:"windowsSpotlightBlockWelcomeExperience,omitempty"`

	// Allows IT admins to turn off the popup of Windows Tips.
	WindowsSpotlightBlockWindowsTips *bool `json:"windowsSpotlightBlockWindowsTips,omitempty"`

	// Allows IT admins to turn off all Windows Spotlight features
	WindowsSpotlightBlocked *bool `json:"windowsSpotlightBlocked,omitempty"`

	// Allows IT admind to set a predefined default search engine for MDM-Controlled devices
	WindowsSpotlightConfigureOnLockScreen *WindowsSpotlightEnablementSettings `json:"windowsSpotlightConfigureOnLockScreen,omitempty"`

	// Indicates whether or not to block automatic update of apps from Windows Store.
	WindowsStoreBlockAutoUpdate *bool `json:"windowsStoreBlockAutoUpdate,omitempty"`

	// Indicates whether or not to Block the user from using the Windows store.
	WindowsStoreBlocked *bool `json:"windowsStoreBlocked,omitempty"`

	// Indicates whether or not to enable Private Store Only.
	WindowsStoreEnablePrivateStoreOnly *bool `json:"windowsStoreEnablePrivateStoreOnly,omitempty"`

	// Indicates whether or not to allow other devices from discovering this PC for projection.
	WirelessDisplayBlockProjectionToThisDevice *bool `json:"wirelessDisplayBlockProjectionToThisDevice,omitempty"`

	// Indicates whether or not to allow user input from wireless display receiver.
	WirelessDisplayBlockUserInputFromReceiver *bool `json:"wirelessDisplayBlockUserInputFromReceiver,omitempty"`

	// Indicates whether or not to require a PIN for new devices to initiate pairing.
	WirelessDisplayRequirePinForPairing *bool `json:"wirelessDisplayRequirePinForPairing,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s Windows10GeneralConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:                 s.Assignments,
		CreatedDateTime:             s.CreatedDateTime,
		Description:                 s.Description,
		DeviceSettingStateSummaries: s.DeviceSettingStateSummaries,
		DeviceStatusOverview:        s.DeviceStatusOverview,
		DeviceStatuses:              s.DeviceStatuses,
		DisplayName:                 s.DisplayName,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		UserStatusOverview:          s.UserStatusOverview,
		UserStatuses:                s.UserStatuses,
		Version:                     s.Version,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s Windows10GeneralConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10GeneralConfiguration{}

func (s Windows10GeneralConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10GeneralConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10GeneralConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10GeneralConfiguration: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.windows10GeneralConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10GeneralConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Windows10GeneralConfiguration{}

func (s *Windows10GeneralConfiguration) UnmarshalJSON(bytes []byte) error {
	type alias Windows10GeneralConfiguration
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into Windows10GeneralConfiguration: %+v", err)
	}

	s.AccountsBlockAddingNonMicrosoftAccountEmail = decoded.AccountsBlockAddingNonMicrosoftAccountEmail
	s.AntiTheftModeBlocked = decoded.AntiTheftModeBlocked
	s.AppsAllowTrustedAppsSideloading = decoded.AppsAllowTrustedAppsSideloading
	s.AppsBlockWindowsStoreOriginatedApps = decoded.AppsBlockWindowsStoreOriginatedApps
	s.Assignments = decoded.Assignments
	s.BluetoothAllowedServices = decoded.BluetoothAllowedServices
	s.BluetoothBlockAdvertising = decoded.BluetoothBlockAdvertising
	s.BluetoothBlockDiscoverableMode = decoded.BluetoothBlockDiscoverableMode
	s.BluetoothBlockPrePairing = decoded.BluetoothBlockPrePairing
	s.BluetoothBlocked = decoded.BluetoothBlocked
	s.CameraBlocked = decoded.CameraBlocked
	s.CellularBlockDataWhenRoaming = decoded.CellularBlockDataWhenRoaming
	s.CellularBlockVpn = decoded.CellularBlockVpn
	s.CellularBlockVpnWhenRoaming = decoded.CellularBlockVpnWhenRoaming
	s.CertificatesBlockManualRootCertificateInstallation = decoded.CertificatesBlockManualRootCertificateInstallation
	s.ConnectedDevicesServiceBlocked = decoded.ConnectedDevicesServiceBlocked
	s.CopyPasteBlocked = decoded.CopyPasteBlocked
	s.CortanaBlocked = decoded.CortanaBlocked
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefenderBlockEndUserAccess = decoded.DefenderBlockEndUserAccess
	s.DefenderCloudBlockLevel = decoded.DefenderCloudBlockLevel
	s.DefenderDaysBeforeDeletingQuarantinedMalware = decoded.DefenderDaysBeforeDeletingQuarantinedMalware
	s.DefenderDetectedMalwareActions = decoded.DefenderDetectedMalwareActions
	s.DefenderFileExtensionsToExclude = decoded.DefenderFileExtensionsToExclude
	s.DefenderFilesAndFoldersToExclude = decoded.DefenderFilesAndFoldersToExclude
	s.DefenderMonitorFileActivity = decoded.DefenderMonitorFileActivity
	s.DefenderProcessesToExclude = decoded.DefenderProcessesToExclude
	s.DefenderPromptForSampleSubmission = decoded.DefenderPromptForSampleSubmission
	s.DefenderRequireBehaviorMonitoring = decoded.DefenderRequireBehaviorMonitoring
	s.DefenderRequireCloudProtection = decoded.DefenderRequireCloudProtection
	s.DefenderRequireNetworkInspectionSystem = decoded.DefenderRequireNetworkInspectionSystem
	s.DefenderRequireRealTimeMonitoring = decoded.DefenderRequireRealTimeMonitoring
	s.DefenderScanArchiveFiles = decoded.DefenderScanArchiveFiles
	s.DefenderScanDownloads = decoded.DefenderScanDownloads
	s.DefenderScanIncomingMail = decoded.DefenderScanIncomingMail
	s.DefenderScanMappedNetworkDrivesDuringFullScan = decoded.DefenderScanMappedNetworkDrivesDuringFullScan
	s.DefenderScanMaxCpu = decoded.DefenderScanMaxCpu
	s.DefenderScanNetworkFiles = decoded.DefenderScanNetworkFiles
	s.DefenderScanRemovableDrivesDuringFullScan = decoded.DefenderScanRemovableDrivesDuringFullScan
	s.DefenderScanScriptsLoadedInInternetExplorer = decoded.DefenderScanScriptsLoadedInInternetExplorer
	s.DefenderScanType = decoded.DefenderScanType
	s.DefenderScheduledQuickScanTime = decoded.DefenderScheduledQuickScanTime
	s.DefenderScheduledScanTime = decoded.DefenderScheduledScanTime
	s.DefenderSignatureUpdateIntervalInHours = decoded.DefenderSignatureUpdateIntervalInHours
	s.DefenderSystemScanSchedule = decoded.DefenderSystemScanSchedule
	s.Description = decoded.Description
	s.DeveloperUnlockSetting = decoded.DeveloperUnlockSetting
	s.DeviceManagementBlockFactoryResetOnMobile = decoded.DeviceManagementBlockFactoryResetOnMobile
	s.DeviceManagementBlockManualUnenroll = decoded.DeviceManagementBlockManualUnenroll
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DiagnosticsDataSubmissionMode = decoded.DiagnosticsDataSubmissionMode
	s.DisplayName = decoded.DisplayName
	s.EdgeAllowStartPagesModification = decoded.EdgeAllowStartPagesModification
	s.EdgeBlockAccessToAboutFlags = decoded.EdgeBlockAccessToAboutFlags
	s.EdgeBlockAddressBarDropdown = decoded.EdgeBlockAddressBarDropdown
	s.EdgeBlockAutofill = decoded.EdgeBlockAutofill
	s.EdgeBlockCompatibilityList = decoded.EdgeBlockCompatibilityList
	s.EdgeBlockDeveloperTools = decoded.EdgeBlockDeveloperTools
	s.EdgeBlockExtensions = decoded.EdgeBlockExtensions
	s.EdgeBlockInPrivateBrowsing = decoded.EdgeBlockInPrivateBrowsing
	s.EdgeBlockJavaScript = decoded.EdgeBlockJavaScript
	s.EdgeBlockLiveTileDataCollection = decoded.EdgeBlockLiveTileDataCollection
	s.EdgeBlockPasswordManager = decoded.EdgeBlockPasswordManager
	s.EdgeBlockPopups = decoded.EdgeBlockPopups
	s.EdgeBlockSearchSuggestions = decoded.EdgeBlockSearchSuggestions
	s.EdgeBlockSendingDoNotTrackHeader = decoded.EdgeBlockSendingDoNotTrackHeader
	s.EdgeBlockSendingIntranetTrafficToInternetExplorer = decoded.EdgeBlockSendingIntranetTrafficToInternetExplorer
	s.EdgeBlocked = decoded.EdgeBlocked
	s.EdgeClearBrowsingDataOnExit = decoded.EdgeClearBrowsingDataOnExit
	s.EdgeCookiePolicy = decoded.EdgeCookiePolicy
	s.EdgeDisableFirstRunPage = decoded.EdgeDisableFirstRunPage
	s.EdgeEnterpriseModeSiteListLocation = decoded.EdgeEnterpriseModeSiteListLocation
	s.EdgeFirstRunUrl = decoded.EdgeFirstRunUrl
	s.EdgeHomepageUrls = decoded.EdgeHomepageUrls
	s.EdgeRequireSmartScreen = decoded.EdgeRequireSmartScreen
	s.EdgeSendIntranetTrafficToInternetExplorer = decoded.EdgeSendIntranetTrafficToInternetExplorer
	s.EdgeSyncFavoritesWithInternetExplorer = decoded.EdgeSyncFavoritesWithInternetExplorer
	s.EnterpriseCloudPrintDiscoveryEndPoint = decoded.EnterpriseCloudPrintDiscoveryEndPoint
	s.EnterpriseCloudPrintDiscoveryMaxLimit = decoded.EnterpriseCloudPrintDiscoveryMaxLimit
	s.EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier = decoded.EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier
	s.EnterpriseCloudPrintOAuthAuthority = decoded.EnterpriseCloudPrintOAuthAuthority
	s.EnterpriseCloudPrintOAuthClientIdentifier = decoded.EnterpriseCloudPrintOAuthClientIdentifier
	s.EnterpriseCloudPrintResourceIdentifier = decoded.EnterpriseCloudPrintResourceIdentifier
	s.ExperienceBlockDeviceDiscovery = decoded.ExperienceBlockDeviceDiscovery
	s.ExperienceBlockErrorDialogWhenNoSIM = decoded.ExperienceBlockErrorDialogWhenNoSIM
	s.ExperienceBlockTaskSwitcher = decoded.ExperienceBlockTaskSwitcher
	s.GameDvrBlocked = decoded.GameDvrBlocked
	s.Id = decoded.Id
	s.InternetSharingBlocked = decoded.InternetSharingBlocked
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LocationServicesBlocked = decoded.LocationServicesBlocked
	s.LockScreenAllowTimeoutConfiguration = decoded.LockScreenAllowTimeoutConfiguration
	s.LockScreenBlockActionCenterNotifications = decoded.LockScreenBlockActionCenterNotifications
	s.LockScreenBlockCortana = decoded.LockScreenBlockCortana
	s.LockScreenBlockToastNotifications = decoded.LockScreenBlockToastNotifications
	s.LockScreenTimeoutInSeconds = decoded.LockScreenTimeoutInSeconds
	s.LogonBlockFastUserSwitching = decoded.LogonBlockFastUserSwitching
	s.MicrosoftAccountBlockSettingsSync = decoded.MicrosoftAccountBlockSettingsSync
	s.MicrosoftAccountBlocked = decoded.MicrosoftAccountBlocked
	s.NetworkProxyApplySettingsDeviceWide = decoded.NetworkProxyApplySettingsDeviceWide
	s.NetworkProxyAutomaticConfigurationUrl = decoded.NetworkProxyAutomaticConfigurationUrl
	s.NetworkProxyDisableAutoDetect = decoded.NetworkProxyDisableAutoDetect
	s.NetworkProxyServer = decoded.NetworkProxyServer
	s.NfcBlocked = decoded.NfcBlocked
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OneDriveDisableFileSync = decoded.OneDriveDisableFileSync
	s.PasswordBlockSimple = decoded.PasswordBlockSimple
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMinimumCharacterSetCount = decoded.PasswordMinimumCharacterSetCount
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinutesOfInactivityBeforeScreenTimeout = decoded.PasswordMinutesOfInactivityBeforeScreenTimeout
	s.PasswordPreviousPasswordBlockCount = decoded.PasswordPreviousPasswordBlockCount
	s.PasswordRequireWhenResumeFromIdleState = decoded.PasswordRequireWhenResumeFromIdleState
	s.PasswordRequired = decoded.PasswordRequired
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PasswordSignInFailureCountBeforeFactoryReset = decoded.PasswordSignInFailureCountBeforeFactoryReset
	s.PersonalizationDesktopImageUrl = decoded.PersonalizationDesktopImageUrl
	s.PersonalizationLockScreenImageUrl = decoded.PersonalizationLockScreenImageUrl
	s.PrivacyAdvertisingId = decoded.PrivacyAdvertisingId
	s.PrivacyAutoAcceptPairingAndConsentPrompts = decoded.PrivacyAutoAcceptPairingAndConsentPrompts
	s.PrivacyBlockInputPersonalization = decoded.PrivacyBlockInputPersonalization
	s.ResetProtectionModeBlocked = decoded.ResetProtectionModeBlocked
	s.SafeSearchFilter = decoded.SafeSearchFilter
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.SearchBlockDiacritics = decoded.SearchBlockDiacritics
	s.SearchDisableAutoLanguageDetection = decoded.SearchDisableAutoLanguageDetection
	s.SearchDisableIndexerBackoff = decoded.SearchDisableIndexerBackoff
	s.SearchDisableIndexingEncryptedItems = decoded.SearchDisableIndexingEncryptedItems
	s.SearchDisableIndexingRemovableDrive = decoded.SearchDisableIndexingRemovableDrive
	s.SearchEnableAutomaticIndexSizeManangement = decoded.SearchEnableAutomaticIndexSizeManangement
	s.SearchEnableRemoteQueries = decoded.SearchEnableRemoteQueries
	s.SettingsBlockAccountsPage = decoded.SettingsBlockAccountsPage
	s.SettingsBlockAddProvisioningPackage = decoded.SettingsBlockAddProvisioningPackage
	s.SettingsBlockAppsPage = decoded.SettingsBlockAppsPage
	s.SettingsBlockChangeLanguage = decoded.SettingsBlockChangeLanguage
	s.SettingsBlockChangePowerSleep = decoded.SettingsBlockChangePowerSleep
	s.SettingsBlockChangeRegion = decoded.SettingsBlockChangeRegion
	s.SettingsBlockChangeSystemTime = decoded.SettingsBlockChangeSystemTime
	s.SettingsBlockDevicesPage = decoded.SettingsBlockDevicesPage
	s.SettingsBlockEaseOfAccessPage = decoded.SettingsBlockEaseOfAccessPage
	s.SettingsBlockEditDeviceName = decoded.SettingsBlockEditDeviceName
	s.SettingsBlockGamingPage = decoded.SettingsBlockGamingPage
	s.SettingsBlockNetworkInternetPage = decoded.SettingsBlockNetworkInternetPage
	s.SettingsBlockPersonalizationPage = decoded.SettingsBlockPersonalizationPage
	s.SettingsBlockPrivacyPage = decoded.SettingsBlockPrivacyPage
	s.SettingsBlockRemoveProvisioningPackage = decoded.SettingsBlockRemoveProvisioningPackage
	s.SettingsBlockSettingsApp = decoded.SettingsBlockSettingsApp
	s.SettingsBlockSystemPage = decoded.SettingsBlockSystemPage
	s.SettingsBlockTimeLanguagePage = decoded.SettingsBlockTimeLanguagePage
	s.SettingsBlockUpdateSecurityPage = decoded.SettingsBlockUpdateSecurityPage
	s.SharedUserAppDataAllowed = decoded.SharedUserAppDataAllowed
	s.SmartScreenBlockPromptOverride = decoded.SmartScreenBlockPromptOverride
	s.SmartScreenBlockPromptOverrideForFiles = decoded.SmartScreenBlockPromptOverrideForFiles
	s.SmartScreenEnableAppInstallControl = decoded.SmartScreenEnableAppInstallControl
	s.StartBlockUnpinningAppsFromTaskbar = decoded.StartBlockUnpinningAppsFromTaskbar
	s.StartMenuAppListVisibility = decoded.StartMenuAppListVisibility
	s.StartMenuHideChangeAccountSettings = decoded.StartMenuHideChangeAccountSettings
	s.StartMenuHideFrequentlyUsedApps = decoded.StartMenuHideFrequentlyUsedApps
	s.StartMenuHideHibernate = decoded.StartMenuHideHibernate
	s.StartMenuHideLock = decoded.StartMenuHideLock
	s.StartMenuHidePowerButton = decoded.StartMenuHidePowerButton
	s.StartMenuHideRecentJumpLists = decoded.StartMenuHideRecentJumpLists
	s.StartMenuHideRecentlyAddedApps = decoded.StartMenuHideRecentlyAddedApps
	s.StartMenuHideRestartOptions = decoded.StartMenuHideRestartOptions
	s.StartMenuHideShutDown = decoded.StartMenuHideShutDown
	s.StartMenuHideSignOut = decoded.StartMenuHideSignOut
	s.StartMenuHideSleep = decoded.StartMenuHideSleep
	s.StartMenuHideSwitchAccount = decoded.StartMenuHideSwitchAccount
	s.StartMenuHideUserTile = decoded.StartMenuHideUserTile
	s.StartMenuLayoutEdgeAssetsXml = decoded.StartMenuLayoutEdgeAssetsXml
	s.StartMenuLayoutXml = decoded.StartMenuLayoutXml
	s.StartMenuMode = decoded.StartMenuMode
	s.StartMenuPinnedFolderDocuments = decoded.StartMenuPinnedFolderDocuments
	s.StartMenuPinnedFolderDownloads = decoded.StartMenuPinnedFolderDownloads
	s.StartMenuPinnedFolderFileExplorer = decoded.StartMenuPinnedFolderFileExplorer
	s.StartMenuPinnedFolderHomeGroup = decoded.StartMenuPinnedFolderHomeGroup
	s.StartMenuPinnedFolderMusic = decoded.StartMenuPinnedFolderMusic
	s.StartMenuPinnedFolderNetwork = decoded.StartMenuPinnedFolderNetwork
	s.StartMenuPinnedFolderPersonalFolder = decoded.StartMenuPinnedFolderPersonalFolder
	s.StartMenuPinnedFolderPictures = decoded.StartMenuPinnedFolderPictures
	s.StartMenuPinnedFolderSettings = decoded.StartMenuPinnedFolderSettings
	s.StartMenuPinnedFolderVideos = decoded.StartMenuPinnedFolderVideos
	s.StorageBlockRemovableStorage = decoded.StorageBlockRemovableStorage
	s.StorageRequireMobileDeviceEncryption = decoded.StorageRequireMobileDeviceEncryption
	s.StorageRestrictAppDataToSystemVolume = decoded.StorageRestrictAppDataToSystemVolume
	s.StorageRestrictAppInstallToSystemVolume = decoded.StorageRestrictAppInstallToSystemVolume
	s.TenantLockdownRequireNetworkDuringOutOfBoxExperience = decoded.TenantLockdownRequireNetworkDuringOutOfBoxExperience
	s.UsbBlocked = decoded.UsbBlocked
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version
	s.VoiceRecordingBlocked = decoded.VoiceRecordingBlocked
	s.WebRtcBlockLocalhostIPAddress = decoded.WebRtcBlockLocalhostIPAddress
	s.WiFiBlockAutomaticConnectHotspots = decoded.WiFiBlockAutomaticConnectHotspots
	s.WiFiBlockManualConfiguration = decoded.WiFiBlockManualConfiguration
	s.WiFiBlocked = decoded.WiFiBlocked
	s.WiFiScanInterval = decoded.WiFiScanInterval
	s.WindowsSpotlightBlockConsumerSpecificFeatures = decoded.WindowsSpotlightBlockConsumerSpecificFeatures
	s.WindowsSpotlightBlockOnActionCenter = decoded.WindowsSpotlightBlockOnActionCenter
	s.WindowsSpotlightBlockTailoredExperiences = decoded.WindowsSpotlightBlockTailoredExperiences
	s.WindowsSpotlightBlockThirdPartyNotifications = decoded.WindowsSpotlightBlockThirdPartyNotifications
	s.WindowsSpotlightBlockWelcomeExperience = decoded.WindowsSpotlightBlockWelcomeExperience
	s.WindowsSpotlightBlockWindowsTips = decoded.WindowsSpotlightBlockWindowsTips
	s.WindowsSpotlightBlocked = decoded.WindowsSpotlightBlocked
	s.WindowsSpotlightConfigureOnLockScreen = decoded.WindowsSpotlightConfigureOnLockScreen
	s.WindowsStoreBlockAutoUpdate = decoded.WindowsStoreBlockAutoUpdate
	s.WindowsStoreBlocked = decoded.WindowsStoreBlocked
	s.WindowsStoreEnablePrivateStoreOnly = decoded.WindowsStoreEnablePrivateStoreOnly
	s.WirelessDisplayBlockProjectionToThisDevice = decoded.WirelessDisplayBlockProjectionToThisDevice
	s.WirelessDisplayBlockUserInputFromReceiver = decoded.WirelessDisplayBlockUserInputFromReceiver
	s.WirelessDisplayRequirePinForPairing = decoded.WirelessDisplayRequirePinForPairing

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Windows10GeneralConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["edgeSearchEngine"]; ok {
		impl, err := UnmarshalEdgeSearchEngineBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EdgeSearchEngine' for 'Windows10GeneralConfiguration': %+v", err)
		}
		s.EdgeSearchEngine = impl
	}
	return nil
}
