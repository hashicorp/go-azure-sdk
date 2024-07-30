package lifecycleworkflowcustomtaskextension

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

type ListLifecycleWorkflowCustomTaskExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceCustomTaskExtension
}

type ListLifecycleWorkflowCustomTaskExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceCustomTaskExtension
}

type ListLifecycleWorkflowCustomTaskExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowCustomTaskExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowCustomTaskExtensions ...
func (c LifecycleWorkflowCustomTaskExtensionClient) ListLifecycleWorkflowCustomTaskExtensions(ctx context.Context) (result ListLifecycleWorkflowCustomTaskExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLifecycleWorkflowCustomTaskExtensionsCustomPager{},
		Path:       "/identityGovernance/lifecycleWorkflows/customTaskExtensions",
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
		Values *[]stable.IdentityGovernanceCustomTaskExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowCustomTaskExtensionsComplete retrieves all the results into a single object
func (c LifecycleWorkflowCustomTaskExtensionClient) ListLifecycleWorkflowCustomTaskExtensionsComplete(ctx context.Context) (ListLifecycleWorkflowCustomTaskExtensionsCompleteResult, error) {
	return c.ListLifecycleWorkflowCustomTaskExtensionsCompleteMatchingPredicate(ctx, IdentityGovernanceCustomTaskExtensionOperationPredicate{})
}

// ListLifecycleWorkflowCustomTaskExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowCustomTaskExtensionClient) ListLifecycleWorkflowCustomTaskExtensionsCompleteMatchingPredicate(ctx context.Context, predicate IdentityGovernanceCustomTaskExtensionOperationPredicate) (result ListLifecycleWorkflowCustomTaskExtensionsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceCustomTaskExtension, 0)

	resp, err := c.ListLifecycleWorkflowCustomTaskExtensions(ctx)
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

	result = ListLifecycleWorkflowCustomTaskExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
