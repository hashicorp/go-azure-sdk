package instanceprocessslotoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListInstanceProcessesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessInfo
}

type WebAppsListInstanceProcessesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessInfo
}

type WebAppsListInstanceProcessesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListInstanceProcessesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListInstanceProcessesSlot ...
func (c InstanceProcessSlotOperationGroupClient) WebAppsListInstanceProcessesSlot(ctx context.Context, id SlotInstanceId) (result WebAppsListInstanceProcessesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListInstanceProcessesSlotCustomPager{},
		Path:       fmt.Sprintf("%s/processes", id.ID()),
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
		Values *[]ProcessInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListInstanceProcessesSlotComplete retrieves all the results into a single object
func (c InstanceProcessSlotOperationGroupClient) WebAppsListInstanceProcessesSlotComplete(ctx context.Context, id SlotInstanceId) (WebAppsListInstanceProcessesSlotCompleteResult, error) {
	return c.WebAppsListInstanceProcessesSlotCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// WebAppsListInstanceProcessesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InstanceProcessSlotOperationGroupClient) WebAppsListInstanceProcessesSlotCompleteMatchingPredicate(ctx context.Context, id SlotInstanceId, predicate ProcessInfoOperationPredicate) (result WebAppsListInstanceProcessesSlotCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	resp, err := c.WebAppsListInstanceProcessesSlot(ctx, id)
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

	result = WebAppsListInstanceProcessesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
