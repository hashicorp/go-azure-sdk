package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsolesListByVirtualMachineOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Console
}

type ConsolesListByVirtualMachineCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Console
}

type ConsolesListByVirtualMachineCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConsolesListByVirtualMachineCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConsolesListByVirtualMachine ...
func (c NetworkcloudsClient) ConsolesListByVirtualMachine(ctx context.Context, id VirtualMachineId) (result ConsolesListByVirtualMachineOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ConsolesListByVirtualMachineCustomPager{},
		Path:       fmt.Sprintf("%s/consoles", id.ID()),
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
		Values *[]Console `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConsolesListByVirtualMachineComplete retrieves all the results into a single object
func (c NetworkcloudsClient) ConsolesListByVirtualMachineComplete(ctx context.Context, id VirtualMachineId) (ConsolesListByVirtualMachineCompleteResult, error) {
	return c.ConsolesListByVirtualMachineCompleteMatchingPredicate(ctx, id, ConsoleOperationPredicate{})
}

// ConsolesListByVirtualMachineCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) ConsolesListByVirtualMachineCompleteMatchingPredicate(ctx context.Context, id VirtualMachineId, predicate ConsoleOperationPredicate) (result ConsolesListByVirtualMachineCompleteResult, err error) {
	items := make([]Console, 0)

	resp, err := c.ConsolesListByVirtualMachine(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ConsolesListByVirtualMachineCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
