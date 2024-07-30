package applicationsignindetailedsummary

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

type ListApplicationSignInDetailedSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApplicationSignInDetailedSummary
}

type ListApplicationSignInDetailedSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApplicationSignInDetailedSummary
}

type ListApplicationSignInDetailedSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListApplicationSignInDetailedSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListApplicationSignInDetailedSummaries ...
func (c ApplicationSignInDetailedSummaryClient) ListApplicationSignInDetailedSummaries(ctx context.Context) (result ListApplicationSignInDetailedSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListApplicationSignInDetailedSummariesCustomPager{},
		Path:       "/reports/applicationSignInDetailedSummary",
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
		Values *[]beta.ApplicationSignInDetailedSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationSignInDetailedSummariesComplete retrieves all the results into a single object
func (c ApplicationSignInDetailedSummaryClient) ListApplicationSignInDetailedSummariesComplete(ctx context.Context) (ListApplicationSignInDetailedSummariesCompleteResult, error) {
	return c.ListApplicationSignInDetailedSummariesCompleteMatchingPredicate(ctx, ApplicationSignInDetailedSummaryOperationPredicate{})
}

// ListApplicationSignInDetailedSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationSignInDetailedSummaryClient) ListApplicationSignInDetailedSummariesCompleteMatchingPredicate(ctx context.Context, predicate ApplicationSignInDetailedSummaryOperationPredicate) (result ListApplicationSignInDetailedSummariesCompleteResult, err error) {
	items := make([]beta.ApplicationSignInDetailedSummary, 0)

	resp, err := c.ListApplicationSignInDetailedSummaries(ctx)
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

	result = ListApplicationSignInDetailedSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
