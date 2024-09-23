package deviceregisteredowner

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

type ListDeviceRegisteredOwnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceRegisteredOwnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceRegisteredOwnersOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListDeviceRegisteredOwnersOperationOptions() ListDeviceRegisteredOwnersOperationOptions {
	return ListDeviceRegisteredOwnersOperationOptions{}
}

func (o ListDeviceRegisteredOwnersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceRegisteredOwnersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o ListDeviceRegisteredOwnersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceRegisteredOwnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceRegisteredOwnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceRegisteredOwners - Get registeredOwners from me. The user that cloud joined the device or registered their
// personal device. The registered owner is set at the time of registration. Read-only. Nullable. Supports $expand.
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwners(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredOwnersOperationOptions) (result ListDeviceRegisteredOwnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceRegisteredOwnersCustomPager{},
		Path:          fmt.Sprintf("%s/registeredOwners", id.ID()),
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

// ListDeviceRegisteredOwnersComplete retrieves all the results into a single object
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnersComplete(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredOwnersOperationOptions) (ListDeviceRegisteredOwnersCompleteResult, error) {
	return c.ListDeviceRegisteredOwnersCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDeviceRegisteredOwnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnersCompleteMatchingPredicate(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredOwnersOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDeviceRegisteredOwnersCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceRegisteredOwners(ctx, id, options)
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

	result = ListDeviceRegisteredOwnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
