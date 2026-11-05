package containerappsauthconfigs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByContainerAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AuthConfig
}

type ListByContainerAppCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AuthConfig
}

type ListByContainerAppCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByContainerAppCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByContainerApp ...
func (c ContainerAppsAuthConfigsClient) ListByContainerApp(ctx context.Context, id ContainerAppId) (result ListByContainerAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByContainerAppCustomPager{},
		Path:       fmt.Sprintf("%s/authConfigs", id.ID()),
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
		Values *[]AuthConfig `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByContainerAppComplete retrieves all the results into a single object
func (c ContainerAppsAuthConfigsClient) ListByContainerAppComplete(ctx context.Context, id ContainerAppId) (ListByContainerAppCompleteResult, error) {
	return c.ListByContainerAppCompleteMatchingPredicate(ctx, id, AuthConfigOperationPredicate{})
}

// ListByContainerAppCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainerAppsAuthConfigsClient) ListByContainerAppCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, predicate AuthConfigOperationPredicate) (result ListByContainerAppCompleteResult, err error) {
	items := make([]AuthConfig, 0)

	resp, err := c.ListByContainerApp(ctx, id)
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

	result = ListByContainerAppCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
