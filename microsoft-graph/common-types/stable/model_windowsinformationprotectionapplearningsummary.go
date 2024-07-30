package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionAppLearningSummary struct {
	ApplicationName *string                                                        `json:"applicationName,omitempty"`
	ApplicationType *WindowsInformationProtectionAppLearningSummaryApplicationType `json:"applicationType,omitempty"`
	DeviceCount     *int64                                                         `json:"deviceCount,omitempty"`
	Id              *string                                                        `json:"id,omitempty"`
	ODataType       *string                                                        `json:"@odata.type,omitempty"`
}
