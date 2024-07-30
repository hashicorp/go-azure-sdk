package approvalstep

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

type ListApprovalStepsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ApprovalStep
}

type ListApprovalStepsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ApprovalStep
}

type ListApprovalStepsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListApprovalStepsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListApprovalSteps ...
func (c ApprovalStepClient) ListApprovalSteps(ctx context.Context, id UserIdApprovalId) (result ListApprovalStepsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListApprovalStepsCustomPager{},
		Path:       fmt.Sprintf("%s/steps", id.ID()),
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
		Values *[]beta.ApprovalStep `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApprovalStepsComplete retrieves all the results into a single object
func (c ApprovalStepClient) ListApprovalStepsComplete(ctx context.Context, id UserIdApprovalId) (ListApprovalStepsCompleteResult, error) {
	return c.ListApprovalStepsCompleteMatchingPredicate(ctx, id, ApprovalStepOperationPredicate{})
}

// ListApprovalStepsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApprovalStepClient) ListApprovalStepsCompleteMatchingPredicate(ctx context.Context, id UserIdApprovalId, predicate ApprovalStepOperationPredicate) (result ListApprovalStepsCompleteResult, err error) {
	items := make([]beta.ApprovalStep, 0)

	resp, err := c.ListApprovalSteps(ctx, id)
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

	result = ListApprovalStepsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
