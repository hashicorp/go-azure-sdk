package driveitemlistitemlastmodifiedbyuserserviceprovisioningerror

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

type ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServiceProvisioningError
}

type ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServiceProvisioningError
}

type ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions() ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions {
	return ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions{}
}

func (o ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemListItemLastModifiedByUserServiceProvisioningErrors - Get serviceProvisioningErrors property value.
// Errors published by a federated service describing a nontransient, service-specific error regarding the properties or
// link from a user object.
func (c DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient) ListDriveItemListItemLastModifiedByUserServiceProvisioningErrors(ctx context.Context, id beta.UserIdDriveIdItemId, options ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions) (result ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCustomPager{},
		Path:          fmt.Sprintf("%s/listItem/lastModifiedByUser/serviceProvisioningErrors", id.ID()),
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

	temp := make([]beta.ServiceProvisioningError, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalServiceProvisioningErrorImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ServiceProvisioningError (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient) ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsComplete(ctx context.Context, id beta.UserIdDriveIdItemId, options ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions) (ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteResult, error) {
	return c.ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, options, ServiceProvisioningErrorOperationPredicate{})
}

// ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient) ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveIdItemId, options ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsOperationOptions, predicate ServiceProvisioningErrorOperationPredicate) (result ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]beta.ServiceProvisioningError, 0)

	resp, err := c.ListDriveItemListItemLastModifiedByUserServiceProvisioningErrors(ctx, id, options)
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

	result = ListDriveItemListItemLastModifiedByUserServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
