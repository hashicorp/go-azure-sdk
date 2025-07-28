package guestconfigurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentsVMSSListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GuestConfigurationAssignment
}

type GuestConfigurationAssignmentsVMSSListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GuestConfigurationAssignment
}

type GuestConfigurationAssignmentsVMSSListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GuestConfigurationAssignmentsVMSSListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GuestConfigurationAssignmentsVMSSList ...
func (c GuestconfigurationsClient) GuestConfigurationAssignmentsVMSSList(ctx context.Context, id VirtualMachineScaleSetId) (result GuestConfigurationAssignmentsVMSSListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GuestConfigurationAssignmentsVMSSListCustomPager{},
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

// GuestConfigurationAssignmentsVMSSListComplete retrieves all the results into a single object
func (c GuestconfigurationsClient) GuestConfigurationAssignmentsVMSSListComplete(ctx context.Context, id VirtualMachineScaleSetId) (GuestConfigurationAssignmentsVMSSListCompleteResult, error) {
	return c.GuestConfigurationAssignmentsVMSSListCompleteMatchingPredicate(ctx, id, GuestConfigurationAssignmentOperationPredicate{})
}

// GuestConfigurationAssignmentsVMSSListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GuestconfigurationsClient) GuestConfigurationAssignmentsVMSSListCompleteMatchingPredicate(ctx context.Context, id VirtualMachineScaleSetId, predicate GuestConfigurationAssignmentOperationPredicate) (result GuestConfigurationAssignmentsVMSSListCompleteResult, err error) {
	items := make([]GuestConfigurationAssignment, 0)

	resp, err := c.GuestConfigurationAssignmentsVMSSList(ctx, id)
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

	result = GuestConfigurationAssignmentsVMSSListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
