package drivelistcontenttypebasetype

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

type ListDriveListContentTypeBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ContentType
}

type ListDriveListContentTypeBaseTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ContentType
}

type ListDriveListContentTypeBaseTypesOperationOptions struct {
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

func DefaultListDriveListContentTypeBaseTypesOperationOptions() ListDriveListContentTypeBaseTypesOperationOptions {
	return ListDriveListContentTypeBaseTypesOperationOptions{}
}

func (o ListDriveListContentTypeBaseTypesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListContentTypeBaseTypesOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveListContentTypeBaseTypesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListContentTypeBaseTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListContentTypeBaseTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListContentTypeBaseTypes - Get baseTypes from users. The collection of content types that are ancestors of
// this content type.
func (c DriveListContentTypeBaseTypeClient) ListDriveListContentTypeBaseTypes(ctx context.Context, id beta.UserIdDriveIdListContentTypeId, options ListDriveListContentTypeBaseTypesOperationOptions) (result ListDriveListContentTypeBaseTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveListContentTypeBaseTypesCustomPager{},
		Path:          fmt.Sprintf("%s/baseTypes", id.ID()),
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
		Values *[]beta.ContentType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveListContentTypeBaseTypesComplete retrieves all the results into a single object
func (c DriveListContentTypeBaseTypeClient) ListDriveListContentTypeBaseTypesComplete(ctx context.Context, id beta.UserIdDriveIdListContentTypeId, options ListDriveListContentTypeBaseTypesOperationOptions) (ListDriveListContentTypeBaseTypesCompleteResult, error) {
	return c.ListDriveListContentTypeBaseTypesCompleteMatchingPredicate(ctx, id, options, ContentTypeOperationPredicate{})
}

// ListDriveListContentTypeBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListContentTypeBaseTypeClient) ListDriveListContentTypeBaseTypesCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveIdListContentTypeId, options ListDriveListContentTypeBaseTypesOperationOptions, predicate ContentTypeOperationPredicate) (result ListDriveListContentTypeBaseTypesCompleteResult, err error) {
	items := make([]beta.ContentType, 0)

	resp, err := c.ListDriveListContentTypeBaseTypes(ctx, id, options)
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

	result = ListDriveListContentTypeBaseTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
