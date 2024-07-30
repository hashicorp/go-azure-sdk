package accessreviewdefinitioninstancecontactedreviewer

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

type ListAccessReviewDefinitionInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessReviewReviewer
}

type ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessReviewReviewer
}

type ListAccessReviewDefinitionInstanceContactedReviewersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceContactedReviewersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceContactedReviewers ...
func (c AccessReviewDefinitionInstanceContactedReviewerClient) ListAccessReviewDefinitionInstanceContactedReviewers(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId) (result ListAccessReviewDefinitionInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAccessReviewDefinitionInstanceContactedReviewersCustomPager{},
		Path:       fmt.Sprintf("%s/contactedReviewers", id.ID()),
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
		Values *[]stable.AccessReviewReviewer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionInstanceContactedReviewersComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceContactedReviewerClient) ListAccessReviewDefinitionInstanceContactedReviewersComplete(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId) (ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, AccessReviewReviewerOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceContactedReviewerClient) ListAccessReviewDefinitionInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAccessReviewDefinitionIdInstanceId, predicate AccessReviewReviewerOperationPredicate) (result ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult, err error) {
	items := make([]stable.AccessReviewReviewer, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceContactedReviewers(ctx, id)
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

	result = ListAccessReviewDefinitionInstanceContactedReviewersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
