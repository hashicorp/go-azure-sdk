package deviceregistereduser

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

type ListDeviceRegisteredUserRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceRegisteredUserRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceRegisteredUserRefsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	Search           *string
	Skip             *int64
	Top              *int64
}

func DefaultListDeviceRegisteredUserRefsOperationOptions() ListDeviceRegisteredUserRefsOperationOptions {
	return ListDeviceRegisteredUserRefsOperationOptions{}
}

func (o ListDeviceRegisteredUserRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceRegisteredUserRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDeviceRegisteredUserRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceRegisteredUserRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceRegisteredUserRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceRegisteredUserRefs - Get ref of registeredUsers from me. Collection of registered users of the device. For
// cloud joined devices and registered personal devices, registered users are set to the same value as registered owners
// at the time of registration. Read-only. Nullable. Supports $expand.
func (c DeviceRegisteredUserClient) ListDeviceRegisteredUserRefs(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredUserRefsOperationOptions) (result ListDeviceRegisteredUserRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceRegisteredUserRefsCustomPager{},
		Path:          fmt.Sprintf("%s/registeredUsers/$ref", id.ID()),
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

// ListDeviceRegisteredUserRefsComplete retrieves all the results into a single object
func (c DeviceRegisteredUserClient) ListDeviceRegisteredUserRefsComplete(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredUserRefsOperationOptions) (ListDeviceRegisteredUserRefsCompleteResult, error) {
	return c.ListDeviceRegisteredUserRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDeviceRegisteredUserRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceRegisteredUserClient) ListDeviceRegisteredUserRefsCompleteMatchingPredicate(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredUserRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDeviceRegisteredUserRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceRegisteredUserRefs(ctx, id, options)
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

	result = ListDeviceRegisteredUserRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
