package accessreviewdefinitioninstancestage

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

type ListAccessReviewDefinitionInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewStage
}

type ListAccessReviewDefinitionInstanceStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewStage
}

type ListAccessReviewDefinitionInstanceStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceStages ...
func (c AccessReviewDefinitionInstanceStageClient) ListAccessReviewDefinitionInstanceStages(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId) (result ListAccessReviewDefinitionInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewDefinitionInstanceStagesCustomPager{},
		Path:       fmt.Sprintf("%s/stages", id.ID()),
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
		Values *[]stable.AccessReviewStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionInstanceStagesComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceStageClient) ListAccessReviewDefinitionInstanceStagesComplete(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId) (ListAccessReviewDefinitionInstanceStagesCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceStagesCompleteMatchingPredicate(ctx, id, AccessReviewStageOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceStageClient) ListAccessReviewDefinitionInstanceStagesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId, predicate AccessReviewStageOperationPredicate) (result ListAccessReviewDefinitionInstanceStagesCompleteResult, err error) {
	items := make([]stable.AccessReviewStage, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceStages(ctx, id)
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

	result = ListAccessReviewDefinitionInstanceStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
