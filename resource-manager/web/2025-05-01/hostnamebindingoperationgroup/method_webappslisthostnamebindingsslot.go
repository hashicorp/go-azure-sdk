package hostnamebindingoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListHostNameBindingsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HostNameBinding
}

type WebAppsListHostNameBindingsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HostNameBinding
}

type WebAppsListHostNameBindingsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListHostNameBindingsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListHostNameBindingsSlot ...
func (c HostNameBindingOperationGroupClient) WebAppsListHostNameBindingsSlot(ctx context.Context, id SlotId) (result WebAppsListHostNameBindingsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListHostNameBindingsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/hostNameBindings", id.ID()),
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
		Values *[]HostNameBinding `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListHostNameBindingsSlotComplete retrieves all the results into a single object
func (c HostNameBindingOperationGroupClient) WebAppsListHostNameBindingsSlotComplete(ctx context.Context, id SlotId) (WebAppsListHostNameBindingsSlotCompleteResult, error) {
	return c.WebAppsListHostNameBindingsSlotCompleteMatchingPredicate(ctx, id, HostNameBindingOperationPredicate{})
}

// WebAppsListHostNameBindingsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HostNameBindingOperationGroupClient) WebAppsListHostNameBindingsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate HostNameBindingOperationPredicate) (result WebAppsListHostNameBindingsSlotCompleteResult, err error) {
	items := make([]HostNameBinding, 0)

	resp, err := c.WebAppsListHostNameBindingsSlot(ctx, id)
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

	result = WebAppsListHostNameBindingsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
