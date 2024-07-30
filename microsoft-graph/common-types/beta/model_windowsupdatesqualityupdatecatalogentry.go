package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCatalogEntry struct {
	CatalogName                 *string                                                             `json:"catalogName,omitempty"`
	CveSeverityInformation      *WindowsUpdatesQualityUpdateCveSeverityInformation                  `json:"cveSeverityInformation,omitempty"`
	DeployableUntilDateTime     *string                                                             `json:"deployableUntilDateTime,omitempty"`
	DisplayName                 *string                                                             `json:"displayName,omitempty"`
	Id                          *string                                                             `json:"id,omitempty"`
	IsExpeditable               *bool                                                               `json:"isExpeditable,omitempty"`
	ODataType                   *string                                                             `json:"@odata.type,omitempty"`
	ProductRevisions            *[]WindowsUpdatesProductRevision                                    `json:"productRevisions,omitempty"`
	QualityUpdateCadence        *WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence        `json:"qualityUpdateCadence,omitempty"`
	QualityUpdateClassification *WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification `json:"qualityUpdateClassification,omitempty"`
	ReleaseDateTime             *string                                                             `json:"releaseDateTime,omitempty"`
	ShortName                   *string                                                             `json:"shortName,omitempty"`
}
