package pipelines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineGetPipelinesByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PipelineResource
}

type PipelineGetPipelinesByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PipelineResource
}

type PipelineGetPipelinesByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PipelineGetPipelinesByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PipelineGetPipelinesByWorkspace ...
func (c PipelinesClient) PipelineGetPipelinesByWorkspace(ctx context.Context) (result PipelineGetPipelinesByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PipelineGetPipelinesByWorkspaceCustomPager{},
		Path:       "/pipelines",
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
		Values *[]PipelineResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PipelineGetPipelinesByWorkspaceComplete retrieves all the results into a single object
func (c PipelinesClient) PipelineGetPipelinesByWorkspaceComplete(ctx context.Context) (PipelineGetPipelinesByWorkspaceCompleteResult, error) {
	return c.PipelineGetPipelinesByWorkspaceCompleteMatchingPredicate(ctx, PipelineResourceOperationPredicate{})
}

// PipelineGetPipelinesByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PipelinesClient) PipelineGetPipelinesByWorkspaceCompleteMatchingPredicate(ctx context.Context, predicate PipelineResourceOperationPredicate) (result PipelineGetPipelinesByWorkspaceCompleteResult, err error) {
	items := make([]PipelineResource, 0)

	resp, err := c.PipelineGetPipelinesByWorkspace(ctx)
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

	result = PipelineGetPipelinesByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
