package directoryobject

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type GetByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type GetByIdsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultGetByIdsOperationOptions() GetByIdsOperationOptions {
	return GetByIdsOperationOptions{}
}

func (o GetByIdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetByIdsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetByIdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetByIds - Invoke action getByIds. Return the directory objects specified in a list of IDs. Only a subset of user
// properties are returned by default in v1.0. Some common uses for this function are to:
func (c DirectoryObjectClient) GetByIds(ctx context.Context, input GetByIdsRequest, options GetByIdsOperationOptions) (result GetByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetByIdsCustomPager{},
		Path:          "/directoryObjects/getByIds",
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// GetByIdsComplete retrieves all the results into a single object
func (c DirectoryObjectClient) GetByIdsComplete(ctx context.Context, input GetByIdsRequest, options GetByIdsOperationOptions) (GetByIdsCompleteResult, error) {
	return c.GetByIdsCompleteMatchingPredicate(ctx, input, options, DirectoryObjectOperationPredicate{})
}

// GetByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryObjectClient) GetByIdsCompleteMatchingPredicate(ctx context.Context, input GetByIdsRequest, options GetByIdsOperationOptions, predicate DirectoryObjectOperationPredicate) (result GetByIdsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.GetByIds(ctx, input, options)
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

	result = GetByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
