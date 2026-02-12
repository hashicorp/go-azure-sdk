package datasets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetGetDatasetsByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DatasetResource
}

type DatasetGetDatasetsByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DatasetResource
}

type DatasetGetDatasetsByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DatasetGetDatasetsByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DatasetGetDatasetsByWorkspace ...
func (c DatasetsClient) DatasetGetDatasetsByWorkspace(ctx context.Context) (result DatasetGetDatasetsByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DatasetGetDatasetsByWorkspaceCustomPager{},
		Path:       "/datasets",
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
		Values *[]DatasetResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DatasetGetDatasetsByWorkspaceComplete retrieves all the results into a single object
func (c DatasetsClient) DatasetGetDatasetsByWorkspaceComplete(ctx context.Context) (DatasetGetDatasetsByWorkspaceCompleteResult, error) {
	return c.DatasetGetDatasetsByWorkspaceCompleteMatchingPredicate(ctx, DatasetResourceOperationPredicate{})
}

// DatasetGetDatasetsByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatasetsClient) DatasetGetDatasetsByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate DatasetResourceOperationPredicate) (result DatasetGetDatasetsByWorkspaceCompleteResult, err error) {
	items := make([]DatasetResource, 0)

	resp, err := c.DatasetGetDatasetsByWorkspace(ctx)
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

	result = DatasetGetDatasetsByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
