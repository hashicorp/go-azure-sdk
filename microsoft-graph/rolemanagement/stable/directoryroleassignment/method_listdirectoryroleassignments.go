package directoryroleassignment

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

type ListDirectoryRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleAssignment
}

type ListDirectoryRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleAssignment
}

type ListDirectoryRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignments ...
func (c DirectoryRoleAssignmentClient) ListDirectoryRoleAssignments(ctx context.Context) (result ListDirectoryRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryRoleAssignmentsCustomPager{},
		Path:       "/roleManagement/directory/roleAssignments",
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
		Values *[]stable.UnifiedRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentsComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentClient) ListDirectoryRoleAssignmentsComplete(ctx context.Context) (ListDirectoryRoleAssignmentsCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentsCompleteMatchingPredicate(ctx, UnifiedRoleAssignmentOperationPredicate{})
}

// ListDirectoryRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentClient) ListDirectoryRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleAssignmentOperationPredicate) (result ListDirectoryRoleAssignmentsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleAssignment, 0)

	resp, err := c.ListDirectoryRoleAssignments(ctx)
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

	result = ListDirectoryRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
