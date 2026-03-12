package hostnamebindings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListHostNameBindingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HostNameBinding
}

type WebAppsListHostNameBindingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HostNameBinding
}

type WebAppsListHostNameBindingsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListHostNameBindingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListHostNameBindings ...
func (c HostNameBindingsClient) WebAppsListHostNameBindings(ctx context.Context, id commonids.AppServiceId) (result WebAppsListHostNameBindingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListHostNameBindingsCustomPager{},
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

// WebAppsListHostNameBindingsComplete retrieves all the results into a single object
func (c HostNameBindingsClient) WebAppsListHostNameBindingsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListHostNameBindingsCompleteResult, error) {
	return c.WebAppsListHostNameBindingsCompleteMatchingPredicate(ctx, id, HostNameBindingOperationPredicate{})
}

// WebAppsListHostNameBindingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HostNameBindingsClient) WebAppsListHostNameBindingsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate HostNameBindingOperationPredicate) (result WebAppsListHostNameBindingsCompleteResult, err error) {
	items := make([]HostNameBinding, 0)

	resp, err := c.WebAppsListHostNameBindings(ctx, id)
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

	result = WebAppsListHostNameBindingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
