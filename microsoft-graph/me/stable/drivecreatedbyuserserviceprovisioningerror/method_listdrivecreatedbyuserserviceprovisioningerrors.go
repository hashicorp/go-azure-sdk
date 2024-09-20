package drivecreatedbyuserserviceprovisioningerror

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

type ListDriveCreatedByUserServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServiceProvisioningError
}

type ListDriveCreatedByUserServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServiceProvisioningError
}

type ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions struct {
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

func DefaultListDriveCreatedByUserServiceProvisioningErrorsOperationOptions() ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions {
	return ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions{}
}

func (o ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveCreatedByUserServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveCreatedByUserServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveCreatedByUserServiceProvisioningErrors - Get serviceProvisioningErrors property value. Errors published by a
// federated service describing a nontransient, service-specific error regarding the properties or link from a user
// object. Supports $filter (eq, not, for isResolved and serviceInstance).
func (c DriveCreatedByUserServiceProvisioningErrorClient) ListDriveCreatedByUserServiceProvisioningErrors(ctx context.Context, id stable.MeDriveId, options ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions) (result ListDriveCreatedByUserServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveCreatedByUserServiceProvisioningErrorsCustomPager{},
		Path:          fmt.Sprintf("%s/createdByUser/serviceProvisioningErrors", id.ID()),
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

	temp := make([]stable.ServiceProvisioningError, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalServiceProvisioningErrorImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ServiceProvisioningError (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDriveCreatedByUserServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c DriveCreatedByUserServiceProvisioningErrorClient) ListDriveCreatedByUserServiceProvisioningErrorsComplete(ctx context.Context, id stable.MeDriveId, options ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions) (ListDriveCreatedByUserServiceProvisioningErrorsCompleteResult, error) {
	return c.ListDriveCreatedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, options, ServiceProvisioningErrorOperationPredicate{})
}

// ListDriveCreatedByUserServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveCreatedByUserServiceProvisioningErrorClient) ListDriveCreatedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveId, options ListDriveCreatedByUserServiceProvisioningErrorsOperationOptions, predicate ServiceProvisioningErrorOperationPredicate) (result ListDriveCreatedByUserServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]stable.ServiceProvisioningError, 0)

	resp, err := c.ListDriveCreatedByUserServiceProvisioningErrors(ctx, id, options)
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

	result = ListDriveCreatedByUserServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
