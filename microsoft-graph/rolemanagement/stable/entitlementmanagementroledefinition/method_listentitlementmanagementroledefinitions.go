package entitlementmanagementroledefinition

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

type ListEntitlementManagementRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleDefinition
}

type ListEntitlementManagementRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleDefinition
}

type ListEntitlementManagementRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleDefinitions ...
func (c EntitlementManagementRoleDefinitionClient) ListEntitlementManagementRoleDefinitions(ctx context.Context) (result ListEntitlementManagementRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEntitlementManagementRoleDefinitionsCustomPager{},
		Path:       "/roleManagement/entitlementManagement/roleDefinitions",
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
		Values *[]stable.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleDefinitionsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleDefinitionClient) ListEntitlementManagementRoleDefinitionsComplete(ctx context.Context) (ListEntitlementManagementRoleDefinitionsCompleteResult, error) {
	return c.ListEntitlementManagementRoleDefinitionsCompleteMatchingPredicate(ctx, UnifiedRoleDefinitionOperationPredicate{})
}

// ListEntitlementManagementRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleDefinitionClient) ListEntitlementManagementRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleDefinitionOperationPredicate) (result ListEntitlementManagementRoleDefinitionsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleDefinition, 0)

	resp, err := c.ListEntitlementManagementRoleDefinitions(ctx)
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

	result = ListEntitlementManagementRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
