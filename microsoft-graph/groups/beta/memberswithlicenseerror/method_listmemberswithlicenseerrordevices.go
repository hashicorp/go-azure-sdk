package memberswithlicenseerror

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMembersWithLicenseErrorDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Device
}

type ListMembersWithLicenseErrorDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Device
}

type ListMembersWithLicenseErrorDevicesOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListMembersWithLicenseErrorDevicesOperationOptions() ListMembersWithLicenseErrorDevicesOperationOptions {
	return ListMembersWithLicenseErrorDevicesOperationOptions{}
}

func (o ListMembersWithLicenseErrorDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMembersWithLicenseErrorDevicesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o ListMembersWithLicenseErrorDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMembersWithLicenseErrorDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMembersWithLicenseErrorDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMembersWithLicenseErrorDevices - Get the items of type microsoft.graph.device in the
// microsoft.graph.directoryObject collection
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorDevices(ctx context.Context, id beta.GroupId, options ListMembersWithLicenseErrorDevicesOperationOptions) (result ListMembersWithLicenseErrorDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMembersWithLicenseErrorDevicesCustomPager{},
		Path:          fmt.Sprintf("%s/membersWithLicenseErrors/device", id.ID()),
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
		Values *[]beta.Device `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMembersWithLicenseErrorDevicesComplete retrieves all the results into a single object
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorDevicesComplete(ctx context.Context, id beta.GroupId, options ListMembersWithLicenseErrorDevicesOperationOptions) (ListMembersWithLicenseErrorDevicesCompleteResult, error) {
	return c.ListMembersWithLicenseErrorDevicesCompleteMatchingPredicate(ctx, id, options, DeviceOperationPredicate{})
}

// ListMembersWithLicenseErrorDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorDevicesCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListMembersWithLicenseErrorDevicesOperationOptions, predicate DeviceOperationPredicate) (result ListMembersWithLicenseErrorDevicesCompleteResult, err error) {
	items := make([]beta.Device, 0)

	resp, err := c.ListMembersWithLicenseErrorDevices(ctx, id, options)
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

	result = ListMembersWithLicenseErrorDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
