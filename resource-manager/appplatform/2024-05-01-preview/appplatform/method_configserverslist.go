package appplatform

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

type ConfigServersListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ConfigServerResource
}

type ConfigServersListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ConfigServerResource
}

type ConfigServersListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConfigServersListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConfigServersList ...
func (c AppPlatformClient) ConfigServersList(ctx context.Context, id commonids.SpringCloudServiceId) (result ConfigServersListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ConfigServersListCustomPager{},
		Path:       fmt.Sprintf("%s/configServers", id.ID()),
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
		Values *[]ConfigServerResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConfigServersListComplete retrieves all the results into a single object
func (c AppPlatformClient) ConfigServersListComplete(ctx context.Context, id commonids.SpringCloudServiceId) (ConfigServersListCompleteResult, error) {
	return c.ConfigServersListCompleteMatchingPredicate(ctx, id, ConfigServerResourceOperationPredicate{})
}

// ConfigServersListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppPlatformClient) ConfigServersListCompleteMatchingPredicate(ctx context.Context, id commonids.SpringCloudServiceId, predicate ConfigServerResourceOperationPredicate) (result ConfigServersListCompleteResult, err error) {
	items := make([]ConfigServerResource, 0)

	resp, err := c.ConfigServersList(ctx, id)
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

	result = ConfigServersListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
