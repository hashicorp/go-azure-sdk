package extensionproperty

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

type ListExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ExtensionProperty
}

type ListExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ExtensionProperty
}

type ListExtensionPropertiesOperationOptions struct {
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

func DefaultListExtensionPropertiesOperationOptions() ListExtensionPropertiesOperationOptions {
	return ListExtensionPropertiesOperationOptions{}
}

func (o ListExtensionPropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExtensionPropertiesOperationOptions) ToOData() *odata.Query {
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

func (o ListExtensionPropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExtensionProperties - List extensionProperties (directory extensions). Retrieve the list of directory extension
// definitions, represented by extensionProperty objects on an application.
func (c ExtensionPropertyClient) ListExtensionProperties(ctx context.Context, id beta.ApplicationId, options ListExtensionPropertiesOperationOptions) (result ListExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExtensionPropertiesCustomPager{},
		Path:          fmt.Sprintf("%s/extensionProperties", id.ID()),
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
		Values *[]beta.ExtensionProperty `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExtensionPropertiesComplete retrieves all the results into a single object
func (c ExtensionPropertyClient) ListExtensionPropertiesComplete(ctx context.Context, id beta.ApplicationId, options ListExtensionPropertiesOperationOptions) (ListExtensionPropertiesCompleteResult, error) {
	return c.ListExtensionPropertiesCompleteMatchingPredicate(ctx, id, options, ExtensionPropertyOperationPredicate{})
}

// ListExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExtensionPropertyClient) ListExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, id beta.ApplicationId, options ListExtensionPropertiesOperationOptions, predicate ExtensionPropertyOperationPredicate) (result ListExtensionPropertiesCompleteResult, err error) {
	items := make([]beta.ExtensionProperty, 0)

	resp, err := c.ListExtensionProperties(ctx, id, options)
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

	result = ListExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
