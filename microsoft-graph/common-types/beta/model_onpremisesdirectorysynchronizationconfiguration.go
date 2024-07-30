package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesDirectorySynchronizationConfiguration struct {
	AccidentalDeletionPrevention             *OnPremisesAccidentalDeletionPrevention `json:"accidentalDeletionPrevention,omitempty"`
	AnchorAttribute                          *string                                 `json:"anchorAttribute,omitempty"`
	ApplicationId                            *string                                 `json:"applicationId,omitempty"`
	CurrentExportData                        *OnPremisesCurrentExportData            `json:"currentExportData,omitempty"`
	CustomerRequestedSynchronizationInterval *string                                 `json:"customerRequestedSynchronizationInterval,omitempty"`
	ODataType                                *string                                 `json:"@odata.type,omitempty"`
	SynchronizationClientVersion             *string                                 `json:"synchronizationClientVersion,omitempty"`
	SynchronizationInterval                  *string                                 `json:"synchronizationInterval,omitempty"`
	WritebackConfiguration                   *OnPremisesWritebackConfiguration       `json:"writebackConfiguration,omitempty"`
}
