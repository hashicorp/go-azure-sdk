package driveitemlistitempermission

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

type ListDriveItemListItemPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListDriveItemListItemPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListDriveItemListItemPermissionGrantsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListDriveItemListItemPermissionGrantsOperationOptions() ListDriveItemListItemPermissionGrantsOperationOptions {
	return ListDriveItemListItemPermissionGrantsOperationOptions{}
}

func (o ListDriveItemListItemPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemListItemPermissionGrantsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveItemListItemPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemListItemPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemListItemPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemListItemPermissionGrants - Invoke action grant. Grant users access to a link represented by a
// permission.
func (c DriveItemListItemPermissionClient) ListDriveItemListItemPermissionGrants(ctx context.Context, id beta.UserIdDriveIdItemIdListItemPermissionId, input ListDriveItemListItemPermissionGrantsRequest, options ListDriveItemListItemPermissionGrantsOperationOptions) (result ListDriveItemListItemPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveItemListItemPermissionGrantsCustomPager{},
		Path:          fmt.Sprintf("%s/grant", id.ID()),
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
		Values *[]beta.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveItemListItemPermissionGrantsComplete retrieves all the results into a single object
func (c DriveItemListItemPermissionClient) ListDriveItemListItemPermissionGrantsComplete(ctx context.Context, id beta.UserIdDriveIdItemIdListItemPermissionId, input ListDriveItemListItemPermissionGrantsRequest, options ListDriveItemListItemPermissionGrantsOperationOptions) (ListDriveItemListItemPermissionGrantsCompleteResult, error) {
	return c.ListDriveItemListItemPermissionGrantsCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveItemListItemPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemListItemPermissionClient) ListDriveItemListItemPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveIdItemIdListItemPermissionId, input ListDriveItemListItemPermissionGrantsRequest, options ListDriveItemListItemPermissionGrantsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveItemListItemPermissionGrantsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListDriveItemListItemPermissionGrants(ctx, id, input, options)
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

	result = ListDriveItemListItemPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
