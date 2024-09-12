package devicecommand

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

type ListDeviceCommandsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Command
}

type ListDeviceCommandsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Command
}

type ListDeviceCommandsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDeviceCommandsOperationOptions() ListDeviceCommandsOperationOptions {
	return ListDeviceCommandsOperationOptions{}
}

func (o ListDeviceCommandsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCommandsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCommandsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCommandsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCommandsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCommands - Get commands from users. Set of commands sent to this device.
func (c DeviceCommandClient) ListDeviceCommands(ctx context.Context, id beta.UserIdDeviceId, options ListDeviceCommandsOperationOptions) (result ListDeviceCommandsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCommandsCustomPager{},
		Path:          fmt.Sprintf("%s/commands", id.ID()),
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
		Values *[]beta.Command `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCommandsComplete retrieves all the results into a single object
func (c DeviceCommandClient) ListDeviceCommandsComplete(ctx context.Context, id beta.UserIdDeviceId, options ListDeviceCommandsOperationOptions) (ListDeviceCommandsCompleteResult, error) {
	return c.ListDeviceCommandsCompleteMatchingPredicate(ctx, id, options, CommandOperationPredicate{})
}

// ListDeviceCommandsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCommandClient) ListDeviceCommandsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDeviceId, options ListDeviceCommandsOperationOptions, predicate CommandOperationPredicate) (result ListDeviceCommandsCompleteResult, err error) {
	items := make([]beta.Command, 0)

	resp, err := c.ListDeviceCommands(ctx, id, options)
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

	result = ListDeviceCommandsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
