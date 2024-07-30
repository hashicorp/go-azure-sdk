package cloudpcroleassignment

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

type ListCloudPCRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentMultiple
}

type ListCloudPCRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentMultiple
}

type ListCloudPCRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCRoleAssignments ...
func (c CloudPCRoleAssignmentClient) ListCloudPCRoleAssignments(ctx context.Context) (result ListCloudPCRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCloudPCRoleAssignmentsCustomPager{},
		Path:       "/roleManagement/cloudPC/roleAssignments",
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
		Values *[]beta.UnifiedRoleAssignmentMultiple `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudPCRoleAssignmentsComplete retrieves all the results into a single object
func (c CloudPCRoleAssignmentClient) ListCloudPCRoleAssignmentsComplete(ctx context.Context) (ListCloudPCRoleAssignmentsCompleteResult, error) {
	return c.ListCloudPCRoleAssignmentsCompleteMatchingPredicate(ctx, UnifiedRoleAssignmentMultipleOperationPredicate{})
}

// ListCloudPCRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCRoleAssignmentClient) ListCloudPCRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, predicate UnifiedRoleAssignmentMultipleOperationPredicate) (result ListCloudPCRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentMultiple, 0)

	resp, err := c.ListCloudPCRoleAssignments(ctx)
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

	result = ListCloudPCRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
