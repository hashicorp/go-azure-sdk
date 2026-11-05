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

type ContainerAppsRevisionFunctionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ContainerAppsFunction
}

type ContainerAppsRevisionFunctionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ContainerAppsFunction
}

type ContainerAppsRevisionFunctionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ContainerAppsRevisionFunctionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ContainerAppsRevisionFunctionsList ...
func (c AzureFunctionsOnContainerAppsClient) ContainerAppsRevisionFunctionsList(ctx context.Context, id RevisionId) (result ContainerAppsRevisionFunctionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ContainerAppsRevisionFunctionsListCustomPager{},
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

// ContainerAppsRevisionFunctionsListComplete retrieves all the results into a single object
func (c AzureFunctionsOnContainerAppsClient) ContainerAppsRevisionFunctionsListComplete(ctx context.Context, id RevisionId) (ContainerAppsRevisionFunctionsListCompleteResult, error) {
	return c.ContainerAppsRevisionFunctionsListCompleteMatchingPredicate(ctx, id, ContainerAppsFunctionOperationPredicate{})
}

// ContainerAppsRevisionFunctionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AzureFunctionsOnContainerAppsClient) ContainerAppsRevisionFunctionsListCompleteMatchingPredicate(ctx context.Context, id RevisionId, predicate ContainerAppsFunctionOperationPredicate) (result ContainerAppsRevisionFunctionsListCompleteResult, err error) {
	items := make([]ContainerAppsFunction, 0)

	resp, err := c.ContainerAppsRevisionFunctionsList(ctx, id)
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

	result = ContainerAppsRevisionFunctionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
