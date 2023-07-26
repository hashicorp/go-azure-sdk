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

type SubAccountListVMHostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VMResources
}

type SubAccountListVMHostsCompleteResult struct {
	Items []VMResources
}

// SubAccountListVMHosts ...
func (c VMHostClient) SubAccountListVMHosts(ctx context.Context, id AccountId) (result SubAccountListVMHostsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listVMHosts", id.ID()),
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

// SubAccountListVMHostsComplete retrieves all the results into a single object
func (c VMHostClient) SubAccountListVMHostsComplete(ctx context.Context, id AccountId) (SubAccountListVMHostsCompleteResult, error) {
	return c.SubAccountListVMHostsCompleteMatchingPredicate(ctx, id, VMResourcesOperationPredicate{})
}

// SubAccountListVMHostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMHostClient) SubAccountListVMHostsCompleteMatchingPredicate(ctx context.Context, id AccountId, predicate VMResourcesOperationPredicate) (result SubAccountListVMHostsCompleteResult, err error) {
	items := make([]VMResources, 0)

	resp, err := c.SubAccountListVMHosts(ctx, id)
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

	result = SubAccountListVMHostsCompleteResult{
		Items: items,
	}
	return
}
