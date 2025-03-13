package standbyvirtualmachines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByStandbyVirtualMachinePoolResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StandbyVirtualMachineResource
}

type ListByStandbyVirtualMachinePoolResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StandbyVirtualMachineResource
}

type ListByStandbyVirtualMachinePoolResourceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByStandbyVirtualMachinePoolResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByStandbyVirtualMachinePoolResource ...
func (c StandbyVirtualMachinesClient) ListByStandbyVirtualMachinePoolResource(ctx context.Context, id StandbyVirtualMachinePoolId) (result ListByStandbyVirtualMachinePoolResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByStandbyVirtualMachinePoolResourceCustomPager{},
		Path:       fmt.Sprintf("%s/standbyVirtualMachines", id.ID()),
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
		Values *[]StandbyVirtualMachineResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByStandbyVirtualMachinePoolResourceComplete retrieves all the results into a single object
func (c StandbyVirtualMachinesClient) ListByStandbyVirtualMachinePoolResourceComplete(ctx context.Context, id StandbyVirtualMachinePoolId) (ListByStandbyVirtualMachinePoolResourceCompleteResult, error) {
	return c.ListByStandbyVirtualMachinePoolResourceCompleteMatchingPredicate(ctx, id, StandbyVirtualMachineResourceOperationPredicate{})
}

// ListByStandbyVirtualMachinePoolResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StandbyVirtualMachinesClient) ListByStandbyVirtualMachinePoolResourceCompleteMatchingPredicate(ctx context.Context, id StandbyVirtualMachinePoolId, predicate StandbyVirtualMachineResourceOperationPredicate) (result ListByStandbyVirtualMachinePoolResourceCompleteResult, err error) {
	items := make([]StandbyVirtualMachineResource, 0)

	resp, err := c.ListByStandbyVirtualMachinePoolResource(ctx, id)
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

	result = ListByStandbyVirtualMachinePoolResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
