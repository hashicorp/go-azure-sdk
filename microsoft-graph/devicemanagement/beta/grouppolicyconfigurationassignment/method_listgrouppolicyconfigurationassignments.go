package grouppolicyconfigurationassignment

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

type ListGroupPolicyConfigurationAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyConfigurationAssignment
}

type ListGroupPolicyConfigurationAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyConfigurationAssignment
}

type ListGroupPolicyConfigurationAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyConfigurationAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyConfigurationAssignments ...
func (c GroupPolicyConfigurationAssignmentClient) ListGroupPolicyConfigurationAssignments(ctx context.Context, id DeviceManagementGroupPolicyConfigurationId) (result ListGroupPolicyConfigurationAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyConfigurationAssignmentsCustomPager{},
		Path:       fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.GroupPolicyConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyConfigurationAssignmentsComplete retrieves all the results into a single object
func (c GroupPolicyConfigurationAssignmentClient) ListGroupPolicyConfigurationAssignmentsComplete(ctx context.Context, id DeviceManagementGroupPolicyConfigurationId) (ListGroupPolicyConfigurationAssignmentsCompleteResult, error) {
	return c.ListGroupPolicyConfigurationAssignmentsCompleteMatchingPredicate(ctx, id, GroupPolicyConfigurationAssignmentOperationPredicate{})
}

// ListGroupPolicyConfigurationAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyConfigurationAssignmentClient) ListGroupPolicyConfigurationAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyConfigurationId, predicate GroupPolicyConfigurationAssignmentOperationPredicate) (result ListGroupPolicyConfigurationAssignmentsCompleteResult, err error) {
	items := make([]beta.GroupPolicyConfigurationAssignment, 0)

	resp, err := c.ListGroupPolicyConfigurationAssignments(ctx, id)
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

	result = ListGroupPolicyConfigurationAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
