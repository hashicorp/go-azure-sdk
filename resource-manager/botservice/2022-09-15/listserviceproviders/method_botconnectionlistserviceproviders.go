package listserviceproviders

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

type BotConnectionListServiceProvidersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServiceProvider
}

type BotConnectionListServiceProvidersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServiceProvider
}

type BotConnectionListServiceProvidersCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BotConnectionListServiceProvidersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BotConnectionListServiceProviders ...
func (c ListServiceProvidersClient) BotConnectionListServiceProviders(ctx context.Context, id commonids.SubscriptionId) (result BotConnectionListServiceProvidersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &BotConnectionListServiceProvidersCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.BotService/listAuthServiceProviders", id.ID()),
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
		Values *[]ServiceProvider `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BotConnectionListServiceProvidersComplete retrieves all the results into a single object
func (c ListServiceProvidersClient) BotConnectionListServiceProvidersComplete(ctx context.Context, id commonids.SubscriptionId) (BotConnectionListServiceProvidersCompleteResult, error) {
	return c.BotConnectionListServiceProvidersCompleteMatchingPredicate(ctx, id, ServiceProviderOperationPredicate{})
}

// BotConnectionListServiceProvidersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ListServiceProvidersClient) BotConnectionListServiceProvidersCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ServiceProviderOperationPredicate) (result BotConnectionListServiceProvidersCompleteResult, err error) {
	items := make([]ServiceProvider, 0)

	resp, err := c.BotConnectionListServiceProviders(ctx, id)
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

	result = BotConnectionListServiceProvidersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
