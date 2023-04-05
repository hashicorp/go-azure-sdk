package apiproduct

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductState string

const (
	ProductStateNotPublished ProductState = "notPublished"
	ProductStatePublished    ProductState = "published"
)

func PossibleValuesForProductState() []string {
	return []string{
		string(ProductStateNotPublished),
		string(ProductStatePublished),
	}
}
