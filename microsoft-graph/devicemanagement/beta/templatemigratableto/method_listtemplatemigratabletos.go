package templatemigratableto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTemplateMigratableTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementTemplate
}

type ListTemplateMigratableTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementTemplate
}

type ListTemplateMigratableTosOperationOptions struct {
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

func DefaultListTemplateMigratableTosOperationOptions() ListTemplateMigratableTosOperationOptions {
	return ListTemplateMigratableTosOperationOptions{}
}

func (o ListTemplateMigratableTosOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTemplateMigratableTosOperationOptions) ToOData() *odata.Query {
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

func (o ListTemplateMigratableTosOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTemplateMigratableTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTemplateMigratableTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTemplateMigratableTos - Get migratableTo from deviceManagement. Collection of templates this template can migrate
// to
func (c TemplateMigratableToClient) ListTemplateMigratableTos(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateMigratableTosOperationOptions) (result ListTemplateMigratableTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTemplateMigratableTosCustomPager{},
		Path:          fmt.Sprintf("%s/migratableTo", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DeviceManagementTemplate, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDeviceManagementTemplateImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DeviceManagementTemplate (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTemplateMigratableTosComplete retrieves all the results into a single object
func (c TemplateMigratableToClient) ListTemplateMigratableTosComplete(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateMigratableTosOperationOptions) (ListTemplateMigratableTosCompleteResult, error) {
	return c.ListTemplateMigratableTosCompleteMatchingPredicate(ctx, id, options, DeviceManagementTemplateOperationPredicate{})
}

// ListTemplateMigratableTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TemplateMigratableToClient) ListTemplateMigratableTosCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementTemplateId, options ListTemplateMigratableTosOperationOptions, predicate DeviceManagementTemplateOperationPredicate) (result ListTemplateMigratableTosCompleteResult, err error) {
	items := make([]beta.DeviceManagementTemplate, 0)

	resp, err := c.ListTemplateMigratableTos(ctx, id, options)
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

	result = ListTemplateMigratableTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
