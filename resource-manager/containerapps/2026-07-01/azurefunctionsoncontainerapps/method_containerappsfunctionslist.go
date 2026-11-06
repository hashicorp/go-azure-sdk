package azurefunctionsoncontainerapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsFunctionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ContainerAppsFunction
}

type ContainerAppsFunctionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ContainerAppsFunction
}

type ContainerAppsFunctionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ContainerAppsFunctionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ContainerAppsFunctionsList ...
func (c AzureFunctionsOnContainerAppsClient) ContainerAppsFunctionsList(ctx context.Context, id ContainerAppId) (result ContainerAppsFunctionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ContainerAppsFunctionsListCustomPager{},
		Path:       fmt.Sprintf("%s/functions", id.ID()),
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
		Values *[]ContainerAppsFunction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ContainerAppsFunctionsListComplete retrieves all the results into a single object
func (c AzureFunctionsOnContainerAppsClient) ContainerAppsFunctionsListComplete(ctx context.Context, id ContainerAppId) (ContainerAppsFunctionsListCompleteResult, error) {
	return c.ContainerAppsFunctionsListCompleteMatchingPredicate(ctx, id, ContainerAppsFunctionOperationPredicate{})
}

// ContainerAppsFunctionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AzureFunctionsOnContainerAppsClient) ContainerAppsFunctionsListCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, predicate ContainerAppsFunctionOperationPredicate) (result ContainerAppsFunctionsListCompleteResult, err error) {
	items := make([]ContainerAppsFunction, 0)

	resp, err := c.ContainerAppsFunctionsList(ctx, id)
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

	result = ContainerAppsFunctionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
