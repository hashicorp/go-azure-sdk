package administrativeunit

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

type ListGetsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListGetsByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListGetsByIdsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListGetsByIdsOperationOptions() ListGetsByIdsOperationOptions {
	return ListGetsByIdsOperationOptions{}
}

func (o ListGetsByIdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGetsByIdsOperationOptions) ToOData() *odata.Query {
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

func (o ListGetsByIdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGetsByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGetsByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGetsByIds - Invoke action getByIds. Return the directory objects specified in a list of IDs. Some common uses for
// this function are to
func (c AdministrativeUnitClient) ListGetsByIds(ctx context.Context, input ListGetsByIdsRequest, options ListGetsByIdsOperationOptions) (result ListGetsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListGetsByIdsCustomPager{},
		Path:          "/administrativeUnits/getByIds",
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

// ListGetsByIdsComplete retrieves all the results into a single object
func (c AdministrativeUnitClient) ListGetsByIdsComplete(ctx context.Context, input ListGetsByIdsRequest, options ListGetsByIdsOperationOptions) (ListGetsByIdsCompleteResult, error) {
	return c.ListGetsByIdsCompleteMatchingPredicate(ctx, input, options, DirectoryObjectOperationPredicate{})
}

// ListGetsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitClient) ListGetsByIdsCompleteMatchingPredicate(ctx context.Context, input ListGetsByIdsRequest, options ListGetsByIdsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListGetsByIdsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListGetsByIds(ctx, input, options)
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

	result = ListGetsByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
