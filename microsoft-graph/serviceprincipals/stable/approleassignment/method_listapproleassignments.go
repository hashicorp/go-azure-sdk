package approleassignment

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

type ListAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AppRoleAssignment
}

type ListAppRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AppRoleAssignment
}

type ListAppRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppRoleAssignments ...
func (c AppRoleAssignmentClient) ListAppRoleAssignments(ctx context.Context, id ServicePrincipalId) (result ListAppRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppRoleAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/appRoleAssignments", id.ID()),
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
		Values *[]stable.AppRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppRoleAssignmentsComplete retrieves all the results into a single object
func (c AppRoleAssignmentClient) ListAppRoleAssignmentsComplete(ctx context.Context, id ServicePrincipalId) (ListAppRoleAssignmentsCompleteResult, error) {
	return c.ListAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, AppRoleAssignmentOperationPredicate{})
}

// ListAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppRoleAssignmentClient) ListAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate AppRoleAssignmentOperationPredicate) (result ListAppRoleAssignmentsCompleteResult, err error) {
	items := make([]stable.AppRoleAssignment, 0)

	resp, err := c.ListAppRoleAssignments(ctx, id)
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

	result = ListAppRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
