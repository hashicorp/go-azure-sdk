package accessreviewhistorydefinitioninstance

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

type ListAccessReviewHistoryDefinitionInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewHistoryInstance
}

type ListAccessReviewHistoryDefinitionInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewHistoryInstance
}

type ListAccessReviewHistoryDefinitionInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewHistoryDefinitionInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewHistoryDefinitionInstances ...
func (c AccessReviewHistoryDefinitionInstanceClient) ListAccessReviewHistoryDefinitionInstances(ctx context.Context, id IdentityGovernanceAccessReviewHistoryDefinitionId) (result ListAccessReviewHistoryDefinitionInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewHistoryDefinitionInstancesCustomPager{},
		Path:       fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]stable.AccessReviewHistoryInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewHistoryDefinitionInstancesComplete retrieves all the results into a single object
func (c AccessReviewHistoryDefinitionInstanceClient) ListAccessReviewHistoryDefinitionInstancesComplete(ctx context.Context, id IdentityGovernanceAccessReviewHistoryDefinitionId) (ListAccessReviewHistoryDefinitionInstancesCompleteResult, error) {
	return c.ListAccessReviewHistoryDefinitionInstancesCompleteMatchingPredicate(ctx, id, AccessReviewHistoryInstanceOperationPredicate{})
}

// ListAccessReviewHistoryDefinitionInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewHistoryDefinitionInstanceClient) ListAccessReviewHistoryDefinitionInstancesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAccessReviewHistoryDefinitionId, predicate AccessReviewHistoryInstanceOperationPredicate) (result ListAccessReviewHistoryDefinitionInstancesCompleteResult, err error) {
	items := make([]stable.AccessReviewHistoryInstance, 0)

	resp, err := c.ListAccessReviewHistoryDefinitionInstances(ctx, id)
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

	result = ListAccessReviewHistoryDefinitionInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
