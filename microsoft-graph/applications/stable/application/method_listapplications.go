package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListApplicationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Application
}

type ListApplicationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Application
}

type ListApplicationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListApplicationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListApplications ...
func (c ApplicationClient) ListApplications(ctx context.Context) (result ListApplicationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListApplicationsCustomPager{},
		Path:       "/applications",
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
		Values *[]stable.Application `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationsComplete retrieves all the results into a single object
func (c ApplicationClient) ListApplicationsComplete(ctx context.Context) (ListApplicationsCompleteResult, error) {
	return c.ListApplicationsCompleteMatchingPredicate(ctx, ApplicationOperationPredicate{})
}

// ListApplicationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationClient) ListApplicationsCompleteMatchingPredicate(ctx context.Context, predicate ApplicationOperationPredicate) (result ListApplicationsCompleteResult, err error) {
	items := make([]stable.Application, 0)

	resp, err := c.ListApplications(ctx)
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

	result = ListApplicationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
