package accesspolicyassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicyAssignmentListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AccessPolicyAssignment
}

type AccessPolicyAssignmentListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AccessPolicyAssignment
}

type AccessPolicyAssignmentListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AccessPolicyAssignmentListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AccessPolicyAssignmentList ...
func (c AccessPolicyAssignmentsClient) AccessPolicyAssignmentList(ctx context.Context, id DatabaseId) (result AccessPolicyAssignmentListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AccessPolicyAssignmentListCustomPager{},
		Path:       fmt.Sprintf("%s/accessPolicyAssignments", id.ID()),
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
		Values *[]AccessPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AccessPolicyAssignmentListComplete retrieves all the results into a single object
func (c AccessPolicyAssignmentsClient) AccessPolicyAssignmentListComplete(ctx context.Context, id DatabaseId) (AccessPolicyAssignmentListCompleteResult, error) {
	return c.AccessPolicyAssignmentListCompleteMatchingPredicate(ctx, id, AccessPolicyAssignmentOperationPredicate{})
}

// AccessPolicyAssignmentListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessPolicyAssignmentsClient) AccessPolicyAssignmentListCompleteMatchingPredicate(ctx context.Context, id DatabaseId, predicate AccessPolicyAssignmentOperationPredicate) (result AccessPolicyAssignmentListCompleteResult, err error) {
	items := make([]AccessPolicyAssignment, 0)

	resp, err := c.AccessPolicyAssignmentList(ctx, id)
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

	result = AccessPolicyAssignmentListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
