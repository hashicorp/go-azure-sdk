package accessreviewhistorydefinition

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

type ListAccessReviewHistoryDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewHistoryDefinition
}

type ListAccessReviewHistoryDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewHistoryDefinition
}

type ListAccessReviewHistoryDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewHistoryDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewHistoryDefinitions ...
func (c AccessReviewHistoryDefinitionClient) ListAccessReviewHistoryDefinitions(ctx context.Context) (result ListAccessReviewHistoryDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewHistoryDefinitionsCustomPager{},
		Path:       "/identityGovernance/accessReviews/historyDefinitions",
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
		Values *[]stable.AccessReviewHistoryDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewHistoryDefinitionsComplete retrieves all the results into a single object
func (c AccessReviewHistoryDefinitionClient) ListAccessReviewHistoryDefinitionsComplete(ctx context.Context) (ListAccessReviewHistoryDefinitionsCompleteResult, error) {
	return c.ListAccessReviewHistoryDefinitionsCompleteMatchingPredicate(ctx, AccessReviewHistoryDefinitionOperationPredicate{})
}

// ListAccessReviewHistoryDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewHistoryDefinitionClient) ListAccessReviewHistoryDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate AccessReviewHistoryDefinitionOperationPredicate) (result ListAccessReviewHistoryDefinitionsCompleteResult, err error) {
	items := make([]stable.AccessReviewHistoryDefinition, 0)

	resp, err := c.ListAccessReviewHistoryDefinitions(ctx)
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

	result = ListAccessReviewHistoryDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
