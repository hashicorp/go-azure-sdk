package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryExportOperation struct {
	Action             *SecurityEdiscoveryExportOperationAction          `json:"action,omitempty"`
	CompletedDateTime  *string                                           `json:"completedDateTime,omitempty"`
	CreatedBy          *IdentitySet                                      `json:"createdBy,omitempty"`
	CreatedDateTime    *string                                           `json:"createdDateTime,omitempty"`
	Description        *string                                           `json:"description,omitempty"`
	ExportFileMetadata *[]SecurityExportFileMetadata                     `json:"exportFileMetadata,omitempty"`
	ExportOptions      *SecurityEdiscoveryExportOperationExportOptions   `json:"exportOptions,omitempty"`
	ExportStructure    *SecurityEdiscoveryExportOperationExportStructure `json:"exportStructure,omitempty"`
	Id                 *string                                           `json:"id,omitempty"`
	ODataType          *string                                           `json:"@odata.type,omitempty"`
	OutputName         *string                                           `json:"outputName,omitempty"`
	PercentProgress    *int64                                            `json:"percentProgress,omitempty"`
	ResultInfo         *ResultInfo                                       `json:"resultInfo,omitempty"`
	ReviewSet          *SecurityEdiscoveryReviewSet                      `json:"reviewSet,omitempty"`
	ReviewSetQuery     *SecurityEdiscoveryReviewSetQuery                 `json:"reviewSetQuery,omitempty"`
	Status             *SecurityEdiscoveryExportOperationStatus          `json:"status,omitempty"`
}
