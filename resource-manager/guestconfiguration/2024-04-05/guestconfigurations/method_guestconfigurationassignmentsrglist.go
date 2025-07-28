package guestconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentsRGListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GuestConfigurationAssignment
}

type GuestConfigurationAssignmentsRGListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GuestConfigurationAssignment
}

type GuestConfigurationAssignmentsRGListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GuestConfigurationAssignmentsRGListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GuestConfigurationAssignmentsRGList ...
func (c GuestconfigurationsClient) GuestConfigurationAssignmentsRGList(ctx context.Context, id commonids.ResourceGroupId) (result GuestConfigurationAssignmentsRGListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GuestConfigurationAssignmentsRGListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments", id.ID()),
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
		Values *[]GuestConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GuestConfigurationAssignmentsRGListComplete retrieves all the results into a single object
func (c GuestconfigurationsClient) GuestConfigurationAssignmentsRGListComplete(ctx context.Context, id commonids.ResourceGroupId) (GuestConfigurationAssignmentsRGListCompleteResult, error) {
	return c.GuestConfigurationAssignmentsRGListCompleteMatchingPredicate(ctx, id, GuestConfigurationAssignmentOperationPredicate{})
}

// GuestConfigurationAssignmentsRGListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GuestconfigurationsClient) GuestConfigurationAssignmentsRGListCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate GuestConfigurationAssignmentOperationPredicate) (result GuestConfigurationAssignmentsRGListCompleteResult, err error) {
	items := make([]GuestConfigurationAssignment, 0)

	resp, err := c.GuestConfigurationAssignmentsRGList(ctx, id)
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

	result = GuestConfigurationAssignmentsRGListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
