package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFile struct {
	Content              *string                        `json:"content,omitempty"`
	DateTime             *string                        `json:"dateTime,omitempty"`
	Extension            *string                        `json:"extension,omitempty"`
	ExtractedTextContent *string                        `json:"extractedTextContent,omitempty"`
	Id                   *string                        `json:"id,omitempty"`
	MediaType            *string                        `json:"mediaType,omitempty"`
	Name                 *string                        `json:"name,omitempty"`
	ODataType            *string                        `json:"@odata.type,omitempty"`
	OtherProperties      *SecurityStringValueDictionary `json:"otherProperties,omitempty"`
	ProcessingStatus     *SecurityFileProcessingStatus  `json:"processingStatus,omitempty"`
	SenderOrAuthors      *[]string                      `json:"senderOrAuthors,omitempty"`
	Size                 *int64                         `json:"size,omitempty"`
	SourceType           *SecurityFileSourceType        `json:"sourceType,omitempty"`
	SubjectTitle         *string                        `json:"subjectTitle,omitempty"`
}
