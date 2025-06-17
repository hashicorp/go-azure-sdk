package authenticationhardwareoathmethod

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

type ListAuthenticationHardwareOathMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwareOathAuthenticationMethod
}

type ListAuthenticationHardwareOathMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwareOathAuthenticationMethod
}

type ListAuthenticationHardwareOathMethodsOperationOptions struct {
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

func DefaultListAuthenticationHardwareOathMethodsOperationOptions() ListAuthenticationHardwareOathMethodsOperationOptions {
	return ListAuthenticationHardwareOathMethodsOperationOptions{}
}

func (o ListAuthenticationHardwareOathMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationHardwareOathMethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationHardwareOathMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationHardwareOathMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationHardwareOathMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationHardwareOathMethods - Get hardwareOathMethods from users. The hardware OATH time-based one-time
// password (TOTP) devices assigned to a user for authentication.
func (c AuthenticationHardwareOathMethodClient) ListAuthenticationHardwareOathMethods(ctx context.Context, id beta.UserId, options ListAuthenticationHardwareOathMethodsOperationOptions) (result ListAuthenticationHardwareOathMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationHardwareOathMethodsCustomPager{},
		Path:          fmt.Sprintf("%s/authentication/hardwareOathMethods", id.ID()),
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
		Values *[]beta.HardwareOathAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationHardwareOathMethodsComplete retrieves all the results into a single object
func (c AuthenticationHardwareOathMethodClient) ListAuthenticationHardwareOathMethodsComplete(ctx context.Context, id beta.UserId, options ListAuthenticationHardwareOathMethodsOperationOptions) (ListAuthenticationHardwareOathMethodsCompleteResult, error) {
	return c.ListAuthenticationHardwareOathMethodsCompleteMatchingPredicate(ctx, id, options, HardwareOathAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationHardwareOathMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationHardwareOathMethodClient) ListAuthenticationHardwareOathMethodsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListAuthenticationHardwareOathMethodsOperationOptions, predicate HardwareOathAuthenticationMethodOperationPredicate) (result ListAuthenticationHardwareOathMethodsCompleteResult, err error) {
	items := make([]beta.HardwareOathAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationHardwareOathMethods(ctx, id, options)
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

	result = ListAuthenticationHardwareOathMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
