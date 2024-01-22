package products

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductsListResult struct {
	NextLink *string    `json:"nextLink,omitempty"`
	Value    *[]Product `json:"value,omitempty"`
}
