package hardwareconfigurationassignment

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

type ListHardwareConfigurationAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwareConfigurationAssignment
}

type ListHardwareConfigurationAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwareConfigurationAssignment
}

type ListHardwareConfigurationAssignmentsOperationOptions struct {
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

func DefaultListHardwareConfigurationAssignmentsOperationOptions() ListHardwareConfigurationAssignmentsOperationOptions {
	return ListHardwareConfigurationAssignmentsOperationOptions{}
}

func (o ListHardwareConfigurationAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHardwareConfigurationAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListHardwareConfigurationAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHardwareConfigurationAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHardwareConfigurationAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHardwareConfigurationAssignments - Get assignments from deviceManagement. A list of the Entra user group ids that
// hardware configuration will be applied to. Only security groups and Office 365 Groups are supported. Optional.
func (c HardwareConfigurationAssignmentClient) ListHardwareConfigurationAssignments(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, options ListHardwareConfigurationAssignmentsOperationOptions) (result ListHardwareConfigurationAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHardwareConfigurationAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.HardwareConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHardwareConfigurationAssignmentsComplete retrieves all the results into a single object
func (c HardwareConfigurationAssignmentClient) ListHardwareConfigurationAssignmentsComplete(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, options ListHardwareConfigurationAssignmentsOperationOptions) (ListHardwareConfigurationAssignmentsCompleteResult, error) {
	return c.ListHardwareConfigurationAssignmentsCompleteMatchingPredicate(ctx, id, options, HardwareConfigurationAssignmentOperationPredicate{})
}

// ListHardwareConfigurationAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HardwareConfigurationAssignmentClient) ListHardwareConfigurationAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, options ListHardwareConfigurationAssignmentsOperationOptions, predicate HardwareConfigurationAssignmentOperationPredicate) (result ListHardwareConfigurationAssignmentsCompleteResult, err error) {
	items := make([]beta.HardwareConfigurationAssignment, 0)

	resp, err := c.ListHardwareConfigurationAssignments(ctx, id, options)
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

	result = ListHardwareConfigurationAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
