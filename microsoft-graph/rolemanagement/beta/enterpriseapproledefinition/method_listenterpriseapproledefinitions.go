package enterpriseapproledefinition

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

type ListEnterpriseAppRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListEnterpriseAppRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListEnterpriseAppRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleDefinitions ...
func (c EnterpriseAppRoleDefinitionClient) ListEnterpriseAppRoleDefinitions(ctx context.Context, id RoleManagementEnterpriseAppId) (result ListEnterpriseAppRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEnterpriseAppRoleDefinitionsCustomPager{},
		Path:       fmt.Sprintf("%s/roleDefinitions", id.ID()),
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

// ListEnterpriseAppRoleDefinitionsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleDefinitionClient) ListEnterpriseAppRoleDefinitionsComplete(ctx context.Context, id RoleManagementEnterpriseAppId) (ListEnterpriseAppRoleDefinitionsCompleteResult, error) {
	return c.ListEnterpriseAppRoleDefinitionsCompleteMatchingPredicate(ctx, id, UnifiedRoleDefinitionOperationPredicate{})
}

// ListEnterpriseAppRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleDefinitionClient) ListEnterpriseAppRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, id RoleManagementEnterpriseAppId, predicate UnifiedRoleDefinitionOperationPredicate) (result ListEnterpriseAppRoleDefinitionsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListEnterpriseAppRoleDefinitions(ctx, id)
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

	result = ListEnterpriseAppRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
