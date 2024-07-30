package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItem struct {
	Accessibility                *MacOSPrivacyAccessControlItemAccessibility                `json:"accessibility,omitempty"`
	AddressBook                  *MacOSPrivacyAccessControlItemAddressBook                  `json:"addressBook,omitempty"`
	AppleEventsAllowedReceivers  *[]MacOSAppleEventReceiver                                 `json:"appleEventsAllowedReceivers,omitempty"`
	BlockCamera                  *bool                                                      `json:"blockCamera,omitempty"`
	BlockListenEvent             *bool                                                      `json:"blockListenEvent,omitempty"`
	BlockMicrophone              *bool                                                      `json:"blockMicrophone,omitempty"`
	BlockScreenCapture           *bool                                                      `json:"blockScreenCapture,omitempty"`
	Calendar                     *MacOSPrivacyAccessControlItemCalendar                     `json:"calendar,omitempty"`
	CodeRequirement              *string                                                    `json:"codeRequirement,omitempty"`
	DisplayName                  *string                                                    `json:"displayName,omitempty"`
	FileProviderPresence         *MacOSPrivacyAccessControlItemFileProviderPresence         `json:"fileProviderPresence,omitempty"`
	Identifier                   *string                                                    `json:"identifier,omitempty"`
	IdentifierType               *MacOSPrivacyAccessControlItemIdentifierType               `json:"identifierType,omitempty"`
	MediaLibrary                 *MacOSPrivacyAccessControlItemMediaLibrary                 `json:"mediaLibrary,omitempty"`
	ODataType                    *string                                                    `json:"@odata.type,omitempty"`
	Photos                       *MacOSPrivacyAccessControlItemPhotos                       `json:"photos,omitempty"`
	PostEvent                    *MacOSPrivacyAccessControlItemPostEvent                    `json:"postEvent,omitempty"`
	Reminders                    *MacOSPrivacyAccessControlItemReminders                    `json:"reminders,omitempty"`
	SpeechRecognition            *MacOSPrivacyAccessControlItemSpeechRecognition            `json:"speechRecognition,omitempty"`
	StaticCodeValidation         *bool                                                      `json:"staticCodeValidation,omitempty"`
	SystemPolicyAllFiles         *MacOSPrivacyAccessControlItemSystemPolicyAllFiles         `json:"systemPolicyAllFiles,omitempty"`
	SystemPolicyDesktopFolder    *MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder    `json:"systemPolicyDesktopFolder,omitempty"`
	SystemPolicyDocumentsFolder  *MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder  `json:"systemPolicyDocumentsFolder,omitempty"`
	SystemPolicyDownloadsFolder  *MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder  `json:"systemPolicyDownloadsFolder,omitempty"`
	SystemPolicyNetworkVolumes   *MacOSPrivacyAccessControlItemSystemPolicyNetworkVolumes   `json:"systemPolicyNetworkVolumes,omitempty"`
	SystemPolicyRemovableVolumes *MacOSPrivacyAccessControlItemSystemPolicyRemovableVolumes `json:"systemPolicyRemovableVolumes,omitempty"`
	SystemPolicySystemAdminFiles *MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles `json:"systemPolicySystemAdminFiles,omitempty"`
}
