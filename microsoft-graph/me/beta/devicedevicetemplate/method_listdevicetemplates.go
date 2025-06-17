package devicedevicetemplate

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

type ListDeviceTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceTemplate
}

type ListDeviceTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceTemplate
}

type ListDeviceTemplatesOperationOptions struct {
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

func DefaultListDeviceTemplatesOperationOptions() ListDeviceTemplatesOperationOptions {
	return ListDeviceTemplatesOperationOptions{}
}

func (o ListDeviceTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceTemplates - Get deviceTemplate from me. Device template used to instantiate this device. Nullable.
// Read-only.
func (c DeviceDeviceTemplateClient) ListDeviceTemplates(ctx context.Context, id beta.MeDeviceId, options ListDeviceTemplatesOperationOptions) (result ListDeviceTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceTemplatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceTemplate", id.ID()),
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
		Values *[]beta.DeviceTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceTemplatesComplete retrieves all the results into a single object
func (c DeviceDeviceTemplateClient) ListDeviceTemplatesComplete(ctx context.Context, id beta.MeDeviceId, options ListDeviceTemplatesOperationOptions) (ListDeviceTemplatesCompleteResult, error) {
	return c.ListDeviceTemplatesCompleteMatchingPredicate(ctx, id, options, DeviceTemplateOperationPredicate{})
}

// ListDeviceTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceDeviceTemplateClient) ListDeviceTemplatesCompleteMatchingPredicate(ctx context.Context, id beta.MeDeviceId, options ListDeviceTemplatesOperationOptions, predicate DeviceTemplateOperationPredicate) (result ListDeviceTemplatesCompleteResult, err error) {
	items := make([]beta.DeviceTemplate, 0)

	resp, err := c.ListDeviceTemplates(ctx, id, options)
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

	result = ListDeviceTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
