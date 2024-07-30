package accessreviewdefinitioninstancedecision

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

type ListAccessReviewDefinitionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceDecisions ...
func (c AccessReviewDefinitionInstanceDecisionClient) ListAccessReviewDefinitionInstanceDecisions(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId) (result ListAccessReviewDefinitionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewDefinitionInstanceDecisionsCustomPager{},
		Path:       fmt.Sprintf("%s/decisions", id.ID()),
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
		Values *[]stable.AccessReviewInstanceDecisionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionInstanceDecisionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceDecisionClient) ListAccessReviewDefinitionInstanceDecisionsComplete(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId) (ListAccessReviewDefinitionInstanceDecisionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceDecisionsCompleteMatchingPredicate(ctx, id, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceDecisionClient) ListAccessReviewDefinitionInstanceDecisionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListAccessReviewDefinitionInstanceDecisionsCompleteResult, err error) {
	items := make([]stable.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceDecisions(ctx, id)
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

	result = ListAccessReviewDefinitionInstanceDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
