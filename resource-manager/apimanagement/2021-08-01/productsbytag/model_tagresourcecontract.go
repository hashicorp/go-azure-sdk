package productsbytag

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagResourceContract struct {
	Api       *ApiTagResourceContractProperties       `json:"api"`
	Operation *OperationTagResourceContractProperties `json:"operation"`
	Product   *ProductTagResourceContractProperties   `json:"product"`
	Tag       TagTagResourceContractProperties        `json:"tag"`
}
