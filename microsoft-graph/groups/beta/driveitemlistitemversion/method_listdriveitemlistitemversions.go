package driveitemlistitemversion

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

type ListDriveItemListItemVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ListItemVersion
}

type ListDriveItemListItemVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ListItemVersion
}

type ListDriveItemListItemVersionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDriveItemListItemVersionsOperationOptions() ListDriveItemListItemVersionsOperationOptions {
	return ListDriveItemListItemVersionsOperationOptions{}
}

func (o ListDriveItemListItemVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemListItemVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveItemListItemVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemListItemVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemListItemVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemListItemVersions - Get versions from groups. The list of previous versions of the list item.
func (c DriveItemListItemVersionClient) ListDriveItemListItemVersions(ctx context.Context, id beta.GroupIdDriveIdItemId, options ListDriveItemListItemVersionsOperationOptions) (result ListDriveItemListItemVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveItemListItemVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/listItem/versions", id.ID()),
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

	temp := make([]beta.ListItemVersion, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalListItemVersionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ListItemVersion (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDriveItemListItemVersionsComplete retrieves all the results into a single object
func (c DriveItemListItemVersionClient) ListDriveItemListItemVersionsComplete(ctx context.Context, id beta.GroupIdDriveIdItemId, options ListDriveItemListItemVersionsOperationOptions) (ListDriveItemListItemVersionsCompleteResult, error) {
	return c.ListDriveItemListItemVersionsCompleteMatchingPredicate(ctx, id, options, ListItemVersionOperationPredicate{})
}

// ListDriveItemListItemVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemListItemVersionClient) ListDriveItemListItemVersionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdDriveIdItemId, options ListDriveItemListItemVersionsOperationOptions, predicate ListItemVersionOperationPredicate) (result ListDriveItemListItemVersionsCompleteResult, err error) {
	items := make([]beta.ListItemVersion, 0)

	resp, err := c.ListDriveItemListItemVersions(ctx, id, options)
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

	result = ListDriveItemListItemVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
