package lifecycleworkflowtaskdefinition

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

type ListLifecycleWorkflowTaskDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskDefinition
}

type ListLifecycleWorkflowTaskDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskDefinition
}

type ListLifecycleWorkflowTaskDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowTaskDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowTaskDefinitions ...
func (c LifecycleWorkflowTaskDefinitionClient) ListLifecycleWorkflowTaskDefinitions(ctx context.Context) (result ListLifecycleWorkflowTaskDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowTaskDefinitionsCustomPager{},
		Path:       "/identityGovernance/lifecycleWorkflows/taskDefinitions",
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
		Values *[]stable.IdentityGovernanceTaskDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowTaskDefinitionsComplete retrieves all the results into a single object
func (c LifecycleWorkflowTaskDefinitionClient) ListLifecycleWorkflowTaskDefinitionsComplete(ctx context.Context) (ListLifecycleWorkflowTaskDefinitionsCompleteResult, error) {
	return c.ListLifecycleWorkflowTaskDefinitionsCompleteMatchingPredicate(ctx, IdentityGovernanceTaskDefinitionOperationPredicate{})
}

// ListLifecycleWorkflowTaskDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowTaskDefinitionClient) ListLifecycleWorkflowTaskDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate IdentityGovernanceTaskDefinitionOperationPredicate) (result ListLifecycleWorkflowTaskDefinitionsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskDefinition, 0)

	resp, err := c.ListLifecycleWorkflowTaskDefinitions(ctx)
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

	result = ListLifecycleWorkflowTaskDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
