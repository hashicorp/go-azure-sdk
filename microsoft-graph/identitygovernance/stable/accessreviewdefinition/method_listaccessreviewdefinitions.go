package accessreviewdefinition

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

type ListAccessReviewDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewScheduleDefinition
}

type ListAccessReviewDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewScheduleDefinition
}

type ListAccessReviewDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitions ...
func (c AccessReviewDefinitionClient) ListAccessReviewDefinitions(ctx context.Context) (result ListAccessReviewDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewDefinitionsCustomPager{},
		Path:       "/identityGovernance/accessReviews/definitions",
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
		Values *[]stable.AccessReviewScheduleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionClient) ListAccessReviewDefinitionsComplete(ctx context.Context) (ListAccessReviewDefinitionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionsCompleteMatchingPredicate(ctx, AccessReviewScheduleDefinitionOperationPredicate{})
}

// ListAccessReviewDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionClient) ListAccessReviewDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate AccessReviewScheduleDefinitionOperationPredicate) (result ListAccessReviewDefinitionsCompleteResult, err error) {
	items := make([]stable.AccessReviewScheduleDefinition, 0)

	resp, err := c.ListAccessReviewDefinitions(ctx)
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

	result = ListAccessReviewDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
