package approleassignedresource

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

type ListAppRoleAssignedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipal
}

type ListAppRoleAssignedResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipal
}

type ListAppRoleAssignedResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppRoleAssignedResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppRoleAssignedResources ...
func (c AppRoleAssignedResourceClient) ListAppRoleAssignedResources(ctx context.Context) (result ListAppRoleAssignedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppRoleAssignedResourcesCustomPager{},
		Path:       "/me/appRoleAssignedResources",
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
		Values *[]beta.ServicePrincipal `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppRoleAssignedResourcesComplete retrieves all the results into a single object
func (c AppRoleAssignedResourceClient) ListAppRoleAssignedResourcesComplete(ctx context.Context) (ListAppRoleAssignedResourcesCompleteResult, error) {
	return c.ListAppRoleAssignedResourcesCompleteMatchingPredicate(ctx, ServicePrincipalOperationPredicate{})
}

// ListAppRoleAssignedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppRoleAssignedResourceClient) ListAppRoleAssignedResourcesCompleteMatchingPredicate(ctx context.Context, predicate ServicePrincipalOperationPredicate) (result ListAppRoleAssignedResourcesCompleteResult, err error) {
	items := make([]beta.ServicePrincipal, 0)

	resp, err := c.ListAppRoleAssignedResources(ctx)
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

	result = ListAppRoleAssignedResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
