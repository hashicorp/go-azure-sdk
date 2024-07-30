package lifecycleworkflowworkflowversion

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

type ListLifecycleWorkflowWorkflowVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceWorkflowVersion
}

type ListLifecycleWorkflowWorkflowVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceWorkflowVersion
}

type ListLifecycleWorkflowWorkflowVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowVersions ...
func (c LifecycleWorkflowWorkflowVersionClient) ListLifecycleWorkflowWorkflowVersions(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId) (result ListLifecycleWorkflowWorkflowVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowVersionsCustomPager{},
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

// ListLifecycleWorkflowWorkflowVersionsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowVersionClient) ListLifecycleWorkflowWorkflowVersionsComplete(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId) (ListLifecycleWorkflowWorkflowVersionsCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowVersionsCompleteMatchingPredicate(ctx, id, IdentityGovernanceWorkflowVersionOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowVersionClient) ListLifecycleWorkflowWorkflowVersionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceLifecycleWorkflowWorkflowId, predicate IdentityGovernanceWorkflowVersionOperationPredicate) (result ListLifecycleWorkflowWorkflowVersionsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceWorkflowVersion, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowVersions(ctx, id)
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

	result = ListLifecycleWorkflowWorkflowVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
