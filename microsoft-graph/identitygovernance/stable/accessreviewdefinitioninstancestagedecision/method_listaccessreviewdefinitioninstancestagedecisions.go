package accessreviewdefinitioninstancestagedecision

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

type ListAccessReviewDefinitionInstanceStageDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewInstanceDecisionItem
}

type ListAccessReviewDefinitionInstanceStageDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceStageDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceStageDecisions ...
func (c AccessReviewDefinitionInstanceStageDecisionClient) ListAccessReviewDefinitionInstanceStageDecisions(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId) (result ListAccessReviewDefinitionInstanceStageDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewDefinitionInstanceStageDecisionsCustomPager{},
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

// ListAccessReviewDefinitionInstanceStageDecisionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceStageDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionsComplete(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId) (ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceStageDecisionsCompleteMatchingPredicate(ctx, id, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceStageDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceStageDecisionClient) ListAccessReviewDefinitionInstanceStageDecisionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult, err error) {
	items := make([]stable.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceStageDecisions(ctx, id)
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

	result = ListAccessReviewDefinitionInstanceStageDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
