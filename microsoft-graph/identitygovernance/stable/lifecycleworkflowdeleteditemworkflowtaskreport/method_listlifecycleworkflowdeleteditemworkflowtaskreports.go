package lifecycleworkflowdeleteditemworkflowtaskreport

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

type ListLifecycleWorkflowDeletedItemWorkflowTaskReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskReport
}

type ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskReport
}

type ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowTaskReports ...
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportClient) ListLifecycleWorkflowDeletedItemWorkflowTaskReports(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (result ListLifecycleWorkflowDeletedItemWorkflowTaskReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCustomPager{},
		Path:       fmt.Sprintf("%s/taskReports", id.ID()),
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
		Values *[]stable.IdentityGovernanceTaskReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowTaskReportsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportClient) ListLifecycleWorkflowDeletedItemWorkflowTaskReportsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteMatchingPredicate(ctx, id, IdentityGovernanceTaskReportOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportClient) ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, predicate IdentityGovernanceTaskReportOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskReport, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowTaskReports(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowTaskReportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
