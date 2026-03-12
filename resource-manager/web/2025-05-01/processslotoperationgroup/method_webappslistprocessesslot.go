package processslotoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListProcessesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessInfo
}

type WebAppsListProcessesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessInfo
}

type WebAppsListProcessesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListProcessesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListProcessesSlot ...
func (c ProcessSlotOperationGroupClient) WebAppsListProcessesSlot(ctx context.Context, id SlotId) (result WebAppsListProcessesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListProcessesSlotCustomPager{},
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

// WebAppsListProcessesSlotComplete retrieves all the results into a single object
func (c ProcessSlotOperationGroupClient) WebAppsListProcessesSlotComplete(ctx context.Context, id SlotId) (WebAppsListProcessesSlotCompleteResult, error) {
	return c.WebAppsListProcessesSlotCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// WebAppsListProcessesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProcessSlotOperationGroupClient) WebAppsListProcessesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ProcessInfoOperationPredicate) (result WebAppsListProcessesSlotCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	resp, err := c.WebAppsListProcessesSlot(ctx, id)
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

	result = WebAppsListProcessesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
