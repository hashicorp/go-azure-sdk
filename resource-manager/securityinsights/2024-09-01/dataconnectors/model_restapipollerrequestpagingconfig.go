package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestApiPollerRequestPagingConfig struct {
	PageSize              *int64                         `json:"pageSize,omitempty"`
	PageSizeParameterName *string                        `json:"pageSizeParameterName,omitempty"`
	PagingType            RestApiPollerRequestPagingKind `json:"pagingType"`
}
