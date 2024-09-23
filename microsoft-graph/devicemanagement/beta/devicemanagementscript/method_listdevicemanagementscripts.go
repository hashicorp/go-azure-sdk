package devicemanagementscript

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

type ListDeviceManagementScriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScript
}

type ListDeviceManagementScriptsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScript
}

type ListDeviceManagementScriptsOperationOptions struct {
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

func DefaultListDeviceManagementScriptsOperationOptions() ListDeviceManagementScriptsOperationOptions {
	return ListDeviceManagementScriptsOperationOptions{}
}

func (o ListDeviceManagementScriptsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementScriptsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementScriptsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementScriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScripts - Get deviceManagementScripts from deviceManagement. The list of device management
// scripts associated with the tenant.
func (c DeviceManagementScriptClient) ListDeviceManagementScripts(ctx context.Context, options ListDeviceManagementScriptsOperationOptions) (result ListDeviceManagementScriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementScriptsCustomPager{},
		Path:          "/deviceManagement/deviceManagementScripts",
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
		Values *[]beta.DeviceManagementScript `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementScriptsComplete retrieves all the results into a single object
func (c DeviceManagementScriptClient) ListDeviceManagementScriptsComplete(ctx context.Context, options ListDeviceManagementScriptsOperationOptions) (ListDeviceManagementScriptsCompleteResult, error) {
	return c.ListDeviceManagementScriptsCompleteMatchingPredicate(ctx, options, DeviceManagementScriptOperationPredicate{})
}

// ListDeviceManagementScriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptClient) ListDeviceManagementScriptsCompleteMatchingPredicate(ctx context.Context, options ListDeviceManagementScriptsOperationOptions, predicate DeviceManagementScriptOperationPredicate) (result ListDeviceManagementScriptsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScript, 0)

	resp, err := c.ListDeviceManagementScripts(ctx, options)
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

	result = ListDeviceManagementScriptsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
