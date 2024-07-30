package elevationrequest

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

type ListElevationRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrivilegeManagementElevationRequest
}

type ListElevationRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrivilegeManagementElevationRequest
}

type ListElevationRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListElevationRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListElevationRequests ...
func (c ElevationRequestClient) ListElevationRequests(ctx context.Context) (result ListElevationRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListElevationRequestsCustomPager{},
		Path:       "/deviceManagement/elevationRequests",
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
		Values *[]beta.PrivilegeManagementElevationRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListElevationRequestsComplete retrieves all the results into a single object
func (c ElevationRequestClient) ListElevationRequestsComplete(ctx context.Context) (ListElevationRequestsCompleteResult, error) {
	return c.ListElevationRequestsCompleteMatchingPredicate(ctx, PrivilegeManagementElevationRequestOperationPredicate{})
}

// ListElevationRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ElevationRequestClient) ListElevationRequestsCompleteMatchingPredicate(ctx context.Context, predicate PrivilegeManagementElevationRequestOperationPredicate) (result ListElevationRequestsCompleteResult, err error) {
	items := make([]beta.PrivilegeManagementElevationRequest, 0)

	resp, err := c.ListElevationRequests(ctx)
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

	result = ListElevationRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
