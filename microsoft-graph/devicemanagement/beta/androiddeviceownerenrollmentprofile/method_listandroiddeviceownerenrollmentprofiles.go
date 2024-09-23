package androiddeviceownerenrollmentprofile

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

type ListAndroidDeviceOwnerEnrollmentProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AndroidDeviceOwnerEnrollmentProfile
}

type ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AndroidDeviceOwnerEnrollmentProfile
}

type ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListAndroidDeviceOwnerEnrollmentProfilesOperationOptions() ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions {
	return ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions{}
}

func (o ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions) ToOData() *odata.Query {
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

func (o ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAndroidDeviceOwnerEnrollmentProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAndroidDeviceOwnerEnrollmentProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAndroidDeviceOwnerEnrollmentProfiles - Get androidDeviceOwnerEnrollmentProfiles from deviceManagement. Android
// device owner enrollment profile entities.
func (c AndroidDeviceOwnerEnrollmentProfileClient) ListAndroidDeviceOwnerEnrollmentProfiles(ctx context.Context, options ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions) (result ListAndroidDeviceOwnerEnrollmentProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAndroidDeviceOwnerEnrollmentProfilesCustomPager{},
		Path:          "/deviceManagement/androidDeviceOwnerEnrollmentProfiles",
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
		Values *[]beta.AndroidDeviceOwnerEnrollmentProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAndroidDeviceOwnerEnrollmentProfilesComplete retrieves all the results into a single object
func (c AndroidDeviceOwnerEnrollmentProfileClient) ListAndroidDeviceOwnerEnrollmentProfilesComplete(ctx context.Context, options ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions) (ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult, error) {
	return c.ListAndroidDeviceOwnerEnrollmentProfilesCompleteMatchingPredicate(ctx, options, AndroidDeviceOwnerEnrollmentProfileOperationPredicate{})
}

// ListAndroidDeviceOwnerEnrollmentProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AndroidDeviceOwnerEnrollmentProfileClient) ListAndroidDeviceOwnerEnrollmentProfilesCompleteMatchingPredicate(ctx context.Context, options ListAndroidDeviceOwnerEnrollmentProfilesOperationOptions, predicate AndroidDeviceOwnerEnrollmentProfileOperationPredicate) (result ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult, err error) {
	items := make([]beta.AndroidDeviceOwnerEnrollmentProfile, 0)

	resp, err := c.ListAndroidDeviceOwnerEnrollmentProfiles(ctx, options)
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

	result = ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
