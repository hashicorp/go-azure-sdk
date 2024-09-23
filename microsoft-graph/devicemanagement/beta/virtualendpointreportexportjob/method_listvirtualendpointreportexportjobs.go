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

type ListVirtualEndpointReportExportJobsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListVirtualEndpointReportExportJobsOperationOptions() ListVirtualEndpointReportExportJobsOperationOptions {
	return ListVirtualEndpointReportExportJobsOperationOptions{}
}

func (o ListVirtualEndpointReportExportJobsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointReportExportJobsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListVirtualEndpointReportExportJobsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListVirtualEndpointReportExportJobs - Get cloudPcExportJob. Read the properties and relationships of a
// cloudPcExportJob object. You can download a report by first creating a new cloudPcExportJob resource to initiate
// downloading. Use this GET operation to verify the exportJobStatus property of the cloudPcExportJob resource. The
// property becomes completed when the report finishes downloading in the location specified by the exportUrl property.
func (c VirtualEndpointReportExportJobClient) ListVirtualEndpointReportExportJobs(ctx context.Context, options ListVirtualEndpointReportExportJobsOperationOptions) (result ListVirtualEndpointReportExportJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointReportExportJobsCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/reports/exportJobs",
		RetryFunc:     options.RetryFunc,
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
func (c VirtualEndpointReportExportJobClient) ListVirtualEndpointReportExportJobsComplete(ctx context.Context, options ListVirtualEndpointReportExportJobsOperationOptions) (ListVirtualEndpointReportExportJobsCompleteResult, error) {
	return c.ListVirtualEndpointReportExportJobsCompleteMatchingPredicate(ctx, options, CloudPCExportJobOperationPredicate{})
}

// ListVirtualEndpointReportExportJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointReportExportJobClient) ListVirtualEndpointReportExportJobsCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointReportExportJobsOperationOptions, predicate CloudPCExportJobOperationPredicate) (result ListVirtualEndpointReportExportJobsCompleteResult, err error) {
	items := make([]beta.CloudPCExportJob, 0)

	resp, err := c.ListVirtualEndpointReportExportJobs(ctx, options)
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
