package exchangeroledefinitioninheritspermissionsfrom

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

type ListExchangeRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListExchangeRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeRoleDefinitionInheritsPermissionsFroms ...
func (c ExchangeRoleDefinitionInheritsPermissionsFromClient) ListExchangeRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id RoleManagementExchangeRoleDefinitionId) (result ListExchangeRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListExchangeRoleDefinitionInheritsPermissionsFromsCustomPager{},
		Path:       fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c ExchangeRoleDefinitionInheritsPermissionsFromClient) ListExchangeRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id RoleManagementExchangeRoleDefinitionId) (ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, UnifiedRoleDefinitionOperationPredicate{})
}

// ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeRoleDefinitionInheritsPermissionsFromClient) ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id RoleManagementExchangeRoleDefinitionId, predicate UnifiedRoleDefinitionOperationPredicate) (result ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListExchangeRoleDefinitionInheritsPermissionsFroms(ctx, id)
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

	result = ListExchangeRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
