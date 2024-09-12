package exportdevicemanagementvirtualendpointreportsjobs

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

type ExportVirtualEndpointReportsJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCExportJob
}

type ExportVirtualEndpointReportsJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCExportJob
}

type ExportVirtualEndpointReportsJobsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultExportVirtualEndpointReportsJobsOperationOptions() ExportVirtualEndpointReportsJobsOperationOptions {
	return ExportVirtualEndpointReportsJobsOperationOptions{}
}

func (o ExportVirtualEndpointReportsJobsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ExportVirtualEndpointReportsJobsOperationOptions) ToOData() *odata.Query {
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

func (o ExportVirtualEndpointReportsJobsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ExportVirtualEndpointReportsJobsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ExportVirtualEndpointReportsJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ExportVirtualEndpointReportsJobs - Get cloudPcExportJob. Read the properties and relationships of a cloudPcExportJob
// object. You can download a report by first creating a new cloudPcExportJob resource to initiate downloading. Use this
// GET operation to verify the exportJobStatus property of the cloudPcExportJob resource. The property becomes completed
// when the report finishes downloading in the location specified by the exportUrl property.
func (c ExportDeviceManagementVirtualEndpointReportsJobsClient) ExportVirtualEndpointReportsJobs(ctx context.Context, options ExportVirtualEndpointReportsJobsOperationOptions) (result ExportVirtualEndpointReportsJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ExportVirtualEndpointReportsJobsCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/reports/exportJobs",
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

// ExportVirtualEndpointReportsJobsComplete retrieves all the results into a single object
func (c ExportDeviceManagementVirtualEndpointReportsJobsClient) ExportVirtualEndpointReportsJobsComplete(ctx context.Context, options ExportVirtualEndpointReportsJobsOperationOptions) (ExportVirtualEndpointReportsJobsCompleteResult, error) {
	return c.ExportVirtualEndpointReportsJobsCompleteMatchingPredicate(ctx, options, CloudPCExportJobOperationPredicate{})
}

// ExportVirtualEndpointReportsJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExportDeviceManagementVirtualEndpointReportsJobsClient) ExportVirtualEndpointReportsJobsCompleteMatchingPredicate(ctx context.Context, options ExportVirtualEndpointReportsJobsOperationOptions, predicate CloudPCExportJobOperationPredicate) (result ExportVirtualEndpointReportsJobsCompleteResult, err error) {
	items := make([]beta.CloudPCExportJob, 0)

	resp, err := c.ExportVirtualEndpointReportsJobs(ctx, options)
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

	result = ExportVirtualEndpointReportsJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
