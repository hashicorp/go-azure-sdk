package deleteditem

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeletedItemsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type GetDeletedItemsByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type GetDeletedItemsByIdsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetDeletedItemsByIdsOperationOptions() GetDeletedItemsByIdsOperationOptions {
	return GetDeletedItemsByIdsOperationOptions{}
}

func (o GetDeletedItemsByIdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeletedItemsByIdsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetDeletedItemsByIdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetDeletedItemsByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDeletedItemsByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDeletedItemsByIds - Invoke action getByIds. Return the directory objects specified in a list of IDs. Some common
// uses for this function are to
func (c DeletedItemClient) GetDeletedItemsByIds(ctx context.Context, input GetDeletedItemsByIdsRequest, options GetDeletedItemsByIdsOperationOptions) (result GetDeletedItemsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetDeletedItemsByIdsCustomPager{},
		Path:          "/directory/deletedItems/getByIds",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// GetDeletedItemsByIdsComplete retrieves all the results into a single object
func (c DeletedItemClient) GetDeletedItemsByIdsComplete(ctx context.Context, input GetDeletedItemsByIdsRequest, options GetDeletedItemsByIdsOperationOptions) (GetDeletedItemsByIdsCompleteResult, error) {
	return c.GetDeletedItemsByIdsCompleteMatchingPredicate(ctx, input, options, DirectoryObjectOperationPredicate{})
}

// GetDeletedItemsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedItemClient) GetDeletedItemsByIdsCompleteMatchingPredicate(ctx context.Context, input GetDeletedItemsByIdsRequest, options GetDeletedItemsByIdsOperationOptions, predicate DirectoryObjectOperationPredicate) (result GetDeletedItemsByIdsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.GetDeletedItemsByIds(ctx, input, options)
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

	result = GetDeletedItemsByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
