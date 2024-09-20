package virtualendpointusersettingassignment

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

type ListVirtualEndpointUserSettingAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.CloudPCUserSettingAssignment
}

type ListVirtualEndpointUserSettingAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.CloudPCUserSettingAssignment
}

type ListVirtualEndpointUserSettingAssignmentsOperationOptions struct {
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

func DefaultListVirtualEndpointUserSettingAssignmentsOperationOptions() ListVirtualEndpointUserSettingAssignmentsOperationOptions {
	return ListVirtualEndpointUserSettingAssignmentsOperationOptions{}
}

func (o ListVirtualEndpointUserSettingAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointUserSettingAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointUserSettingAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointUserSettingAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointUserSettingAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointUserSettingAssignments - Get assignments from deviceManagement. Represents the set of Microsoft
// 365 groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand.
// For an example, see Get cloudPcUserSetting.
func (c VirtualEndpointUserSettingAssignmentClient) ListVirtualEndpointUserSettingAssignments(ctx context.Context, id stable.DeviceManagementVirtualEndpointUserSettingId, options ListVirtualEndpointUserSettingAssignmentsOperationOptions) (result ListVirtualEndpointUserSettingAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointUserSettingAssignmentsCustomPager{},
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
		Values *[]stable.CloudPCUserSettingAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointUserSettingAssignmentsComplete retrieves all the results into a single object
func (c VirtualEndpointUserSettingAssignmentClient) ListVirtualEndpointUserSettingAssignmentsComplete(ctx context.Context, id stable.DeviceManagementVirtualEndpointUserSettingId, options ListVirtualEndpointUserSettingAssignmentsOperationOptions) (ListVirtualEndpointUserSettingAssignmentsCompleteResult, error) {
	return c.ListVirtualEndpointUserSettingAssignmentsCompleteMatchingPredicate(ctx, id, options, CloudPCUserSettingAssignmentOperationPredicate{})
}

// ListVirtualEndpointUserSettingAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointUserSettingAssignmentClient) ListVirtualEndpointUserSettingAssignmentsCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementVirtualEndpointUserSettingId, options ListVirtualEndpointUserSettingAssignmentsOperationOptions, predicate CloudPCUserSettingAssignmentOperationPredicate) (result ListVirtualEndpointUserSettingAssignmentsCompleteResult, err error) {
	items := make([]stable.CloudPCUserSettingAssignment, 0)

	resp, err := c.ListVirtualEndpointUserSettingAssignments(ctx, id, options)
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

	result = ListVirtualEndpointUserSettingAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
