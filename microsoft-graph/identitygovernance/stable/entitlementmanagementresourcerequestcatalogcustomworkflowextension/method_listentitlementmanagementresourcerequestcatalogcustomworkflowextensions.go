package entitlementmanagementresourcerequestcatalogcustomworkflowextension

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

type ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CustomCalloutExtension
}

type ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CustomCalloutExtension
}

type ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensions ...
func (c EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient) ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensions(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId) (result ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCustomPager{},
		Path:       fmt.Sprintf("%s/catalog/customWorkflowExtensions", id.ID()),
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
		Values *[]stable.CustomCalloutExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient) ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsComplete(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId) (ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteResult, error) {
	return c.ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteMatchingPredicate(ctx, id, CustomCalloutExtensionOperationPredicate{})
}

// ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient) ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceEntitlementManagementResourceRequestId, predicate CustomCalloutExtensionOperationPredicate) (result ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteResult, err error) {
	items := make([]stable.CustomCalloutExtension, 0)

	resp, err := c.ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensions(ctx, id)
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

	result = ListEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
