package profileproject

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListProfileProjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ProjectParticipation
}

type ListProfileProjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ProjectParticipation
}

type ListProfileProjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileProjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileProjects ...
func (c ProfileProjectClient) ListProfileProjects(ctx context.Context) (result ListProfileProjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileProjectsCustomPager{},
		Path:       "/me/profile/projects",
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
		Values *[]beta.ProjectParticipation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileProjectsComplete retrieves all the results into a single object
func (c ProfileProjectClient) ListProfileProjectsComplete(ctx context.Context) (ListProfileProjectsCompleteResult, error) {
	return c.ListProfileProjectsCompleteMatchingPredicate(ctx, ProjectParticipationOperationPredicate{})
}

// ListProfileProjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileProjectClient) ListProfileProjectsCompleteMatchingPredicate(ctx context.Context, predicate ProjectParticipationOperationPredicate) (result ListProfileProjectsCompleteResult, err error) {
	items := make([]beta.ProjectParticipation, 0)

	resp, err := c.ListProfileProjects(ctx)
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

	result = ListProfileProjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
