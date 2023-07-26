package vmhost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubAccountListVMHostUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VMResources
}

type SubAccountListVMHostUpdateCompleteResult struct {
	Items []VMResources
}

// SubAccountListVMHostUpdate ...
func (c VMHostClient) SubAccountListVMHostUpdate(ctx context.Context, id AccountId, input VMHostUpdateRequest) (result SubAccountListVMHostUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/vmHostUpdate", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]VMResources `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SubAccountListVMHostUpdateComplete retrieves all the results into a single object
func (c VMHostClient) SubAccountListVMHostUpdateComplete(ctx context.Context, id AccountId, input VMHostUpdateRequest) (SubAccountListVMHostUpdateCompleteResult, error) {
	return c.SubAccountListVMHostUpdateCompleteMatchingPredicate(ctx, id, input, VMResourcesOperationPredicate{})
}

// SubAccountListVMHostUpdateCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMHostClient) SubAccountListVMHostUpdateCompleteMatchingPredicate(ctx context.Context, id AccountId, input VMHostUpdateRequest, predicate VMResourcesOperationPredicate) (result SubAccountListVMHostUpdateCompleteResult, err error) {
	items := make([]VMResources, 0)

	resp, err := c.SubAccountListVMHostUpdate(ctx, id, input)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = SubAccountListVMHostUpdateCompleteResult{
		Items: items,
	}
	return
}
