package pendingaccessreviewinstancedecisioninstancecontactedreviewer

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

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewReviewer
}

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewReviewer
}

type ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceDecisionInstanceContactedReviewers ...
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) ListPendingAccessReviewInstanceDecisionInstanceContactedReviewers(ctx context.Context, id MePendingAccessReviewInstanceIdDecisionId) (result ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCustomPager{},
		Path:       fmt.Sprintf("%s/instance/contactedReviewers", id.ID()),
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
		Values *[]beta.AccessReviewReviewer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersComplete(ctx context.Context, id MePendingAccessReviewInstanceIdDecisionId) (ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteMatchingPredicate(ctx, id, AccessReviewReviewerOperationPredicate{})
}

// ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceIdDecisionId, predicate AccessReviewReviewerOperationPredicate) (result ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult, err error) {
	items := make([]beta.AccessReviewReviewer, 0)

	resp, err := c.ListPendingAccessReviewInstanceDecisionInstanceContactedReviewers(ctx, id)
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

	result = ListPendingAccessReviewInstanceDecisionInstanceContactedReviewersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
