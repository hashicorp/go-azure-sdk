package guestagents

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVMOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GuestAgent
}

type ListByVMCompleteResult struct {
	Items []GuestAgent
}

// ListByVM ...
func (c GuestAgentsClient) ListByVM(ctx context.Context, id VirtualMachineId) (result ListByVMOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/guestAgents", id.ID()),
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
		Values *[]GuestAgent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByVMComplete retrieves all the results into a single object
func (c GuestAgentsClient) ListByVMComplete(ctx context.Context, id VirtualMachineId) (ListByVMCompleteResult, error) {
	return c.ListByVMCompleteMatchingPredicate(ctx, id, GuestAgentOperationPredicate{})
}

// ListByVMCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GuestAgentsClient) ListByVMCompleteMatchingPredicate(ctx context.Context, id VirtualMachineId, predicate GuestAgentOperationPredicate) (result ListByVMCompleteResult, err error) {
	items := make([]GuestAgent, 0)

	resp, err := c.ListByVM(ctx, id)
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

	result = ListByVMCompleteResult{
		Items: items,
	}
	return
}
