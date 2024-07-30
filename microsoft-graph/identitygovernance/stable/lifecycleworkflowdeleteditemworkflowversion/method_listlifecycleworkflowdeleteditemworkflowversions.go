package lifecycleworkflowdeleteditemworkflowversion

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

type ListLifecycleWorkflowDeletedItemWorkflowVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceWorkflowVersion
}

type ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceWorkflowVersion
}

type ListLifecycleWorkflowDeletedItemWorkflowVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowDeletedItemWorkflowVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowDeletedItemWorkflowVersions ...
func (c LifecycleWorkflowDeletedItemWorkflowVersionClient) ListLifecycleWorkflowDeletedItemWorkflowVersions(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (result ListLifecycleWorkflowDeletedItemWorkflowVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowDeletedItemWorkflowVersionsCustomPager{},
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]stable.IdentityGovernanceWorkflowVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionsComplete retrieves all the results into a single object
func (c LifecycleWorkflowDeletedItemWorkflowVersionClient) ListLifecycleWorkflowDeletedItemWorkflowVersionsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteResult, error) {
	return c.ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteMatchingPredicate(ctx, id, IdentityGovernanceWorkflowVersionOperationPredicate{})
}

// ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowDeletedItemWorkflowVersionClient) ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, predicate IdentityGovernanceWorkflowVersionOperationPredicate) (result ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceWorkflowVersion, 0)

	resp, err := c.ListLifecycleWorkflowDeletedItemWorkflowVersions(ctx, id)
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

	result = ListLifecycleWorkflowDeletedItemWorkflowVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
