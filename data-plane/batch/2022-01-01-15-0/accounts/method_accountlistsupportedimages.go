package accounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountListSupportedImagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ImageInformation
}

type AccountListSupportedImagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ImageInformation
}

type AccountListSupportedImagesOperationOptions struct {
	ClientRequestId       *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Timeout               *int64
}

func DefaultAccountListSupportedImagesOperationOptions() AccountListSupportedImagesOperationOptions {
	return AccountListSupportedImagesOperationOptions{}
}

func (o AccountListSupportedImagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o AccountListSupportedImagesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AccountListSupportedImagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type AccountListSupportedImagesCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *AccountListSupportedImagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AccountListSupportedImages ...
func (c AccountsClient) AccountListSupportedImages(ctx context.Context, options AccountListSupportedImagesOperationOptions) (result AccountListSupportedImagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AccountListSupportedImagesCustomPager{},
		Path:          "/supportedimages",
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
		Values *[]ImageInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AccountListSupportedImagesComplete retrieves all the results into a single object
func (c AccountsClient) AccountListSupportedImagesComplete(ctx context.Context, options AccountListSupportedImagesOperationOptions) (AccountListSupportedImagesCompleteResult, error) {
	return c.AccountListSupportedImagesCompleteMatchingPredicate(ctx, options, ImageInformationOperationPredicate{})
}

// AccountListSupportedImagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccountsClient) AccountListSupportedImagesCompleteMatchingPredicate(ctx context.Context, options AccountListSupportedImagesOperationOptions, predicate ImageInformationOperationPredicate) (result AccountListSupportedImagesCompleteResult, err error) {
	items := make([]ImageInformation, 0)

	resp, err := c.AccountListSupportedImages(ctx, options)
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

	result = AccountListSupportedImagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
