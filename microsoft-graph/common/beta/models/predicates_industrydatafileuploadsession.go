package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFileUploadSessionOperationPredicate struct {
	ContainerExpirationDateTime *string
	ContainerId                 *string
	ODataType                   *string
	SessionExpirationDateTime   *string
	SessionUrl                  *string
}

func (p IndustryDataFileUploadSessionOperationPredicate) Matches(input IndustryDataFileUploadSession) bool {

	if p.ContainerExpirationDateTime != nil && (input.ContainerExpirationDateTime == nil || *p.ContainerExpirationDateTime != *input.ContainerExpirationDateTime) {
		return false
	}

	if p.ContainerId != nil && (input.ContainerId == nil || *p.ContainerId != *input.ContainerId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SessionExpirationDateTime != nil && (input.SessionExpirationDateTime == nil || *p.SessionExpirationDateTime != *input.SessionExpirationDateTime) {
		return false
	}

	if p.SessionUrl != nil && (input.SessionUrl == nil || *p.SessionUrl != *input.SessionUrl) {
		return false
	}

	return true
}
