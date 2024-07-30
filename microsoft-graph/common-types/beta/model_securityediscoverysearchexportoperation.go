package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperation struct {
	Action             *SecurityEdiscoverySearchExportOperationAction            `json:"action,omitempty"`
	AdditionalOptions  *SecurityEdiscoverySearchExportOperationAdditionalOptions `json:"additionalOptions,omitempty"`
	CompletedDateTime  *string                                                   `json:"completedDateTime,omitempty"`
	CreatedBy          *IdentitySet                                              `json:"createdBy,omitempty"`
	CreatedDateTime    *string                                                   `json:"createdDateTime,omitempty"`
	Description        *string                                                   `json:"description,omitempty"`
	DisplayName        *string                                                   `json:"displayName,omitempty"`
	ExportCriteria     *SecurityEdiscoverySearchExportOperationExportCriteria    `json:"exportCriteria,omitempty"`
	ExportFileMetadata *[]SecurityExportFileMetadata                             `json:"exportFileMetadata,omitempty"`
	ExportFormat       *SecurityEdiscoverySearchExportOperationExportFormat      `json:"exportFormat,omitempty"`
	ExportLocation     *SecurityEdiscoverySearchExportOperationExportLocation    `json:"exportLocation,omitempty"`
	ExportSingleItems  *bool                                                     `json:"exportSingleItems,omitempty"`
	Id                 *string                                                   `json:"id,omitempty"`
	ODataType          *string                                                   `json:"@odata.type,omitempty"`
	PercentProgress    *int64                                                    `json:"percentProgress,omitempty"`
	ResultInfo         *ResultInfo                                               `json:"resultInfo,omitempty"`
	Search             *SecurityEdiscoverySearch                                 `json:"search,omitempty"`
	Status             *SecurityEdiscoverySearchExportOperationStatus            `json:"status,omitempty"`
}
