package guestconfigurationconnectedvmwarevsphereassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GuestConfigurationAssignmentReport
}

type ReportsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GuestConfigurationAssignmentReport
}

type ReportsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReportsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReportsList ...
func (c GuestConfigurationConnectedVMwarevSphereAssignmentsClient) ReportsList(ctx context.Context, id ProviderVirtualMachineProviders2GuestConfigurationAssignmentId) (result ReportsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ReportsListCustomPager{},
		Path:       fmt.Sprintf("%s/reports", id.ID()),
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
		Values *[]GuestConfigurationAssignmentReport `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReportsListComplete retrieves all the results into a single object
func (c GuestConfigurationConnectedVMwarevSphereAssignmentsClient) ReportsListComplete(ctx context.Context, id ProviderVirtualMachineProviders2GuestConfigurationAssignmentId) (ReportsListCompleteResult, error) {
	return c.ReportsListCompleteMatchingPredicate(ctx, id, GuestConfigurationAssignmentReportOperationPredicate{})
}

// ReportsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GuestConfigurationConnectedVMwarevSphereAssignmentsClient) ReportsListCompleteMatchingPredicate(ctx context.Context, id ProviderVirtualMachineProviders2GuestConfigurationAssignmentId, predicate GuestConfigurationAssignmentReportOperationPredicate) (result ReportsListCompleteResult, err error) {
	items := make([]GuestConfigurationAssignmentReport, 0)

	resp, err := c.ReportsList(ctx, id)
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

	result = ReportsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
