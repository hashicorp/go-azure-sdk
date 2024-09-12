package windowsinformationprotectiondeviceregistration

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

type ListWindowsInformationProtectionDeviceRegistrationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsInformationProtectionDeviceRegistration
}

type ListWindowsInformationProtectionDeviceRegistrationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsInformationProtectionDeviceRegistration
}

type ListWindowsInformationProtectionDeviceRegistrationsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListWindowsInformationProtectionDeviceRegistrationsOperationOptions() ListWindowsInformationProtectionDeviceRegistrationsOperationOptions {
	return ListWindowsInformationProtectionDeviceRegistrationsOperationOptions{}
}

func (o ListWindowsInformationProtectionDeviceRegistrationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListWindowsInformationProtectionDeviceRegistrationsOperationOptions) ToOData() *odata.Query {
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

func (o ListWindowsInformationProtectionDeviceRegistrationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListWindowsInformationProtectionDeviceRegistrationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsInformationProtectionDeviceRegistrationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsInformationProtectionDeviceRegistrations - Get windowsInformationProtectionDeviceRegistrations from me.
// Zero or more WIP device registrations that belong to the user.
func (c WindowsInformationProtectionDeviceRegistrationClient) ListWindowsInformationProtectionDeviceRegistrations(ctx context.Context, options ListWindowsInformationProtectionDeviceRegistrationsOperationOptions) (result ListWindowsInformationProtectionDeviceRegistrationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListWindowsInformationProtectionDeviceRegistrationsCustomPager{},
		Path:          "/me/windowsInformationProtectionDeviceRegistrations",
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
		Values *[]beta.WindowsInformationProtectionDeviceRegistration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsInformationProtectionDeviceRegistrationsComplete retrieves all the results into a single object
func (c WindowsInformationProtectionDeviceRegistrationClient) ListWindowsInformationProtectionDeviceRegistrationsComplete(ctx context.Context, options ListWindowsInformationProtectionDeviceRegistrationsOperationOptions) (ListWindowsInformationProtectionDeviceRegistrationsCompleteResult, error) {
	return c.ListWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate(ctx, options, WindowsInformationProtectionDeviceRegistrationOperationPredicate{})
}

// ListWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsInformationProtectionDeviceRegistrationClient) ListWindowsInformationProtectionDeviceRegistrationsCompleteMatchingPredicate(ctx context.Context, options ListWindowsInformationProtectionDeviceRegistrationsOperationOptions, predicate WindowsInformationProtectionDeviceRegistrationOperationPredicate) (result ListWindowsInformationProtectionDeviceRegistrationsCompleteResult, err error) {
	items := make([]beta.WindowsInformationProtectionDeviceRegistration, 0)

	resp, err := c.ListWindowsInformationProtectionDeviceRegistrations(ctx, options)
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

	result = ListWindowsInformationProtectionDeviceRegistrationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
