package virtualendpointreportexportjob

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

type ListVirtualEndpointReportExportJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCExportJob
}

type ListVirtualEndpointReportExportJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCExportJob
}

type ListVirtualEndpointReportExportJobsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointReportExportJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointReportExportJobs ...
func (c VirtualEndpointReportExportJobClient) ListVirtualEndpointReportExportJobs(ctx context.Context) (result ListVirtualEndpointReportExportJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEndpointReportExportJobsCustomPager{},
		Path:       "/deviceManagement/virtualEndpoint/reports/exportJobs",
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
		Values *[]beta.CloudPCExportJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointReportExportJobsComplete retrieves all the results into a single object
func (c VirtualEndpointReportExportJobClient) ListVirtualEndpointReportExportJobsComplete(ctx context.Context) (ListVirtualEndpointReportExportJobsCompleteResult, error) {
	return c.ListVirtualEndpointReportExportJobsCompleteMatchingPredicate(ctx, CloudPCExportJobOperationPredicate{})
}

// ListVirtualEndpointReportExportJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointReportExportJobClient) ListVirtualEndpointReportExportJobsCompleteMatchingPredicate(ctx context.Context, predicate CloudPCExportJobOperationPredicate) (result ListVirtualEndpointReportExportJobsCompleteResult, err error) {
	items := make([]beta.CloudPCExportJob, 0)

	resp, err := c.ListVirtualEndpointReportExportJobs(ctx)
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

	result = ListVirtualEndpointReportExportJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
