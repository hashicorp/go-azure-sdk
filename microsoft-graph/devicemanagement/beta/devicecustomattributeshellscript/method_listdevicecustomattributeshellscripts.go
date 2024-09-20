package devicecustomattributeshellscript

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

type ListDeviceCustomAttributeShellScriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCustomAttributeShellScript
}

type ListDeviceCustomAttributeShellScriptsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCustomAttributeShellScript
}

type ListDeviceCustomAttributeShellScriptsOperationOptions struct {
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

func DefaultListDeviceCustomAttributeShellScriptsOperationOptions() ListDeviceCustomAttributeShellScriptsOperationOptions {
	return ListDeviceCustomAttributeShellScriptsOperationOptions{}
}

func (o ListDeviceCustomAttributeShellScriptsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCustomAttributeShellScriptsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCustomAttributeShellScriptsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCustomAttributeShellScriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCustomAttributeShellScriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCustomAttributeShellScripts - Get deviceCustomAttributeShellScripts from deviceManagement. The list of
// device custom attribute shell scripts associated with the tenant.
func (c DeviceCustomAttributeShellScriptClient) ListDeviceCustomAttributeShellScripts(ctx context.Context, options ListDeviceCustomAttributeShellScriptsOperationOptions) (result ListDeviceCustomAttributeShellScriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCustomAttributeShellScriptsCustomPager{},
		Path:          "/deviceManagement/deviceCustomAttributeShellScripts",
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
		Values *[]beta.DeviceCustomAttributeShellScript `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCustomAttributeShellScriptsComplete retrieves all the results into a single object
func (c DeviceCustomAttributeShellScriptClient) ListDeviceCustomAttributeShellScriptsComplete(ctx context.Context, options ListDeviceCustomAttributeShellScriptsOperationOptions) (ListDeviceCustomAttributeShellScriptsCompleteResult, error) {
	return c.ListDeviceCustomAttributeShellScriptsCompleteMatchingPredicate(ctx, options, DeviceCustomAttributeShellScriptOperationPredicate{})
}

// ListDeviceCustomAttributeShellScriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCustomAttributeShellScriptClient) ListDeviceCustomAttributeShellScriptsCompleteMatchingPredicate(ctx context.Context, options ListDeviceCustomAttributeShellScriptsOperationOptions, predicate DeviceCustomAttributeShellScriptOperationPredicate) (result ListDeviceCustomAttributeShellScriptsCompleteResult, err error) {
	items := make([]beta.DeviceCustomAttributeShellScript, 0)

	resp, err := c.ListDeviceCustomAttributeShellScripts(ctx, options)
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

	result = ListDeviceCustomAttributeShellScriptsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
