package lifecycleworkflowworkflowtemplate

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

type ListLifecycleWorkflowWorkflowTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceWorkflowTemplate
}

type ListLifecycleWorkflowWorkflowTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceWorkflowTemplate
}

type ListLifecycleWorkflowWorkflowTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowWorkflowTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowWorkflowTemplates ...
func (c LifecycleWorkflowWorkflowTemplateClient) ListLifecycleWorkflowWorkflowTemplates(ctx context.Context) (result ListLifecycleWorkflowWorkflowTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowWorkflowTemplatesCustomPager{},
		Path:       "/identityGovernance/lifecycleWorkflows/workflowTemplates",
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
		Values *[]stable.IdentityGovernanceWorkflowTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowWorkflowTemplatesComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowTemplateClient) ListLifecycleWorkflowWorkflowTemplatesComplete(ctx context.Context) (ListLifecycleWorkflowWorkflowTemplatesCompleteResult, error) {
	return c.ListLifecycleWorkflowWorkflowTemplatesCompleteMatchingPredicate(ctx, IdentityGovernanceWorkflowTemplateOperationPredicate{})
}

// ListLifecycleWorkflowWorkflowTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowTemplateClient) ListLifecycleWorkflowWorkflowTemplatesCompleteMatchingPredicate(ctx context.Context, predicate IdentityGovernanceWorkflowTemplateOperationPredicate) (result ListLifecycleWorkflowWorkflowTemplatesCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceWorkflowTemplate, 0)

	resp, err := c.ListLifecycleWorkflowWorkflowTemplates(ctx)
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

	result = ListLifecycleWorkflowWorkflowTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
