package approleassignedto

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

type ListAppRoleAssignedTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppRoleAssignment
}

type ListAppRoleAssignedTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppRoleAssignment
}

type ListAppRoleAssignedTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppRoleAssignedTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppRoleAssignedTos ...
func (c AppRoleAssignedToClient) ListAppRoleAssignedTos(ctx context.Context, id ServicePrincipalId) (result ListAppRoleAssignedTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppRoleAssignedTosCustomPager{},
		Path:       fmt.Sprintf("%s/appRoleAssignedTo", id.ID()),
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
		Values *[]beta.AppRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppRoleAssignedTosComplete retrieves all the results into a single object
func (c AppRoleAssignedToClient) ListAppRoleAssignedTosComplete(ctx context.Context, id ServicePrincipalId) (ListAppRoleAssignedTosCompleteResult, error) {
	return c.ListAppRoleAssignedTosCompleteMatchingPredicate(ctx, id, AppRoleAssignmentOperationPredicate{})
}

// ListAppRoleAssignedTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppRoleAssignedToClient) ListAppRoleAssignedTosCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate AppRoleAssignmentOperationPredicate) (result ListAppRoleAssignedTosCompleteResult, err error) {
	items := make([]beta.AppRoleAssignment, 0)

	resp, err := c.ListAppRoleAssignedTos(ctx, id)
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

	result = ListAppRoleAssignedTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
