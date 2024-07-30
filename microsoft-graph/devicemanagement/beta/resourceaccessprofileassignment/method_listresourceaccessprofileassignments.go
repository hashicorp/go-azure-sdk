package resourceaccessprofileassignment

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

type ListResourceAccessProfileAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementResourceAccessProfileAssignment
}

type ListResourceAccessProfileAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementResourceAccessProfileAssignment
}

type ListResourceAccessProfileAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListResourceAccessProfileAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListResourceAccessProfileAssignments ...
func (c ResourceAccessProfileAssignmentClient) ListResourceAccessProfileAssignments(ctx context.Context, id DeviceManagementResourceAccessProfileId) (result ListResourceAccessProfileAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListResourceAccessProfileAssignmentsCustomPager{},
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
		Values *[]beta.DeviceManagementResourceAccessProfileAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListResourceAccessProfileAssignmentsComplete retrieves all the results into a single object
func (c ResourceAccessProfileAssignmentClient) ListResourceAccessProfileAssignmentsComplete(ctx context.Context, id DeviceManagementResourceAccessProfileId) (ListResourceAccessProfileAssignmentsCompleteResult, error) {
	return c.ListResourceAccessProfileAssignmentsCompleteMatchingPredicate(ctx, id, DeviceManagementResourceAccessProfileAssignmentOperationPredicate{})
}

// ListResourceAccessProfileAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceAccessProfileAssignmentClient) ListResourceAccessProfileAssignmentsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementResourceAccessProfileId, predicate DeviceManagementResourceAccessProfileAssignmentOperationPredicate) (result ListResourceAccessProfileAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementResourceAccessProfileAssignment, 0)

	resp, err := c.ListResourceAccessProfileAssignments(ctx, id)
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

	result = ListResourceAccessProfileAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
