package driverootlistitempermission

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

type ListDriveRootListItemPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListDriveRootListItemPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListDriveRootListItemPermissionGrantsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListDriveRootListItemPermissionGrantsOperationOptions() ListDriveRootListItemPermissionGrantsOperationOptions {
	return ListDriveRootListItemPermissionGrantsOperationOptions{}
}

func (o ListDriveRootListItemPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootListItemPermissionGrantsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootListItemPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootListItemPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootListItemPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootListItemPermissionGrants - Invoke action grant. Grant users access to a link represented by a
// permission.
func (c DriveRootListItemPermissionClient) ListDriveRootListItemPermissionGrants(ctx context.Context, id beta.MeDriveIdRootListItemPermissionId, input ListDriveRootListItemPermissionGrantsRequest, options ListDriveRootListItemPermissionGrantsOperationOptions) (result ListDriveRootListItemPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveRootListItemPermissionGrantsCustomPager{},
		Path:          fmt.Sprintf("%s/grant", id.ID()),
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

// ListDriveRootListItemPermissionGrantsComplete retrieves all the results into a single object
func (c DriveRootListItemPermissionClient) ListDriveRootListItemPermissionGrantsComplete(ctx context.Context, id beta.MeDriveIdRootListItemPermissionId, input ListDriveRootListItemPermissionGrantsRequest, options ListDriveRootListItemPermissionGrantsOperationOptions) (ListDriveRootListItemPermissionGrantsCompleteResult, error) {
	return c.ListDriveRootListItemPermissionGrantsCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveRootListItemPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootListItemPermissionClient) ListDriveRootListItemPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id beta.MeDriveIdRootListItemPermissionId, input ListDriveRootListItemPermissionGrantsRequest, options ListDriveRootListItemPermissionGrantsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveRootListItemPermissionGrantsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListDriveRootListItemPermissionGrants(ctx, id, input, options)
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

	result = ListDriveRootListItemPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
