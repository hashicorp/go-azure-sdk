package guestconfigurationassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentReportsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GuestConfigurationAssignmentReport
}

type GuestConfigurationAssignmentReportsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GuestConfigurationAssignmentReport
}

type GuestConfigurationAssignmentReportsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GuestConfigurationAssignmentReportsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GuestConfigurationAssignmentReportsList ...
func (c GuestConfigurationAssignmentsClient) GuestConfigurationAssignmentReportsList(ctx context.Context, id VirtualMachineProviders2GuestConfigurationAssignmentId) (result GuestConfigurationAssignmentReportsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GuestConfigurationAssignmentReportsListCustomPager{},
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

// GuestConfigurationAssignmentReportsListComplete retrieves all the results into a single object
func (c GuestConfigurationAssignmentsClient) GuestConfigurationAssignmentReportsListComplete(ctx context.Context, id VirtualMachineProviders2GuestConfigurationAssignmentId) (GuestConfigurationAssignmentReportsListCompleteResult, error) {
	return c.GuestConfigurationAssignmentReportsListCompleteMatchingPredicate(ctx, id, GuestConfigurationAssignmentReportOperationPredicate{})
}

// GuestConfigurationAssignmentReportsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GuestConfigurationAssignmentsClient) GuestConfigurationAssignmentReportsListCompleteMatchingPredicate(ctx context.Context, id VirtualMachineProviders2GuestConfigurationAssignmentId, predicate GuestConfigurationAssignmentReportOperationPredicate) (result GuestConfigurationAssignmentReportsListCompleteResult, err error) {
	items := make([]GuestConfigurationAssignmentReport, 0)

	resp, err := c.GuestConfigurationAssignmentReportsList(ctx, id)
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

	result = GuestConfigurationAssignmentReportsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
