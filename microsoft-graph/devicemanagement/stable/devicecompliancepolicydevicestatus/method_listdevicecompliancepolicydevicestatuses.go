package devicecompliancepolicydevicestatus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeviceCompliancePolicyDeviceStatusesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceComplianceDeviceStatus
}

type ListDeviceCompliancePolicyDeviceStatusesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceComplianceDeviceStatus
}

type ListDeviceCompliancePolicyDeviceStatusesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDeviceCompliancePolicyDeviceStatusesOperationOptions() ListDeviceCompliancePolicyDeviceStatusesOperationOptions {
	return ListDeviceCompliancePolicyDeviceStatusesOperationOptions{}
}

func (o ListDeviceCompliancePolicyDeviceStatusesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePolicyDeviceStatusesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCompliancePolicyDeviceStatusesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePolicyDeviceStatusesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyDeviceStatusesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyDeviceStatuses - List deviceComplianceDeviceStatuses. List properties and relationships of
// the deviceComplianceDeviceStatus objects.
func (c DeviceCompliancePolicyDeviceStatusClient) ListDeviceCompliancePolicyDeviceStatuses(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyDeviceStatusesOperationOptions) (result ListDeviceCompliancePolicyDeviceStatusesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePolicyDeviceStatusesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceStatuses", id.ID()),
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
		Values *[]stable.DeviceComplianceDeviceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyDeviceStatusesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyDeviceStatusClient) ListDeviceCompliancePolicyDeviceStatusesComplete(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyDeviceStatusesOperationOptions) (ListDeviceCompliancePolicyDeviceStatusesCompleteResult, error) {
	return c.ListDeviceCompliancePolicyDeviceStatusesCompleteMatchingPredicate(ctx, id, options, DeviceComplianceDeviceStatusOperationPredicate{})
}

// ListDeviceCompliancePolicyDeviceStatusesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyDeviceStatusClient) ListDeviceCompliancePolicyDeviceStatusesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyDeviceStatusesOperationOptions, predicate DeviceComplianceDeviceStatusOperationPredicate) (result ListDeviceCompliancePolicyDeviceStatusesCompleteResult, err error) {
	items := make([]stable.DeviceComplianceDeviceStatus, 0)

	resp, err := c.ListDeviceCompliancePolicyDeviceStatuses(ctx, id, options)
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

	result = ListDeviceCompliancePolicyDeviceStatusesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
