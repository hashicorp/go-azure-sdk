package deponboardingsettingenrollmentprofile

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

type ListDepOnboardingSettingEnrollmentProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EnrollmentProfile
}

type ListDepOnboardingSettingEnrollmentProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EnrollmentProfile
}

type ListDepOnboardingSettingEnrollmentProfilesOperationOptions struct {
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

func DefaultListDepOnboardingSettingEnrollmentProfilesOperationOptions() ListDepOnboardingSettingEnrollmentProfilesOperationOptions {
	return ListDepOnboardingSettingEnrollmentProfilesOperationOptions{}
}

func (o ListDepOnboardingSettingEnrollmentProfilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDepOnboardingSettingEnrollmentProfilesOperationOptions) ToOData() *odata.Query {
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

func (o ListDepOnboardingSettingEnrollmentProfilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDepOnboardingSettingEnrollmentProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDepOnboardingSettingEnrollmentProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDepOnboardingSettingEnrollmentProfiles - Get enrollmentProfiles from deviceManagement. The enrollment profiles.
func (c DepOnboardingSettingEnrollmentProfileClient) ListDepOnboardingSettingEnrollmentProfiles(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, options ListDepOnboardingSettingEnrollmentProfilesOperationOptions) (result ListDepOnboardingSettingEnrollmentProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDepOnboardingSettingEnrollmentProfilesCustomPager{},
		Path:          fmt.Sprintf("%s/enrollmentProfiles", id.ID()),
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

	temp := make([]beta.EnrollmentProfile, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalEnrollmentProfileImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.EnrollmentProfile (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDepOnboardingSettingEnrollmentProfilesComplete retrieves all the results into a single object
func (c DepOnboardingSettingEnrollmentProfileClient) ListDepOnboardingSettingEnrollmentProfilesComplete(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, options ListDepOnboardingSettingEnrollmentProfilesOperationOptions) (ListDepOnboardingSettingEnrollmentProfilesCompleteResult, error) {
	return c.ListDepOnboardingSettingEnrollmentProfilesCompleteMatchingPredicate(ctx, id, options, EnrollmentProfileOperationPredicate{})
}

// ListDepOnboardingSettingEnrollmentProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingEnrollmentProfileClient) ListDepOnboardingSettingEnrollmentProfilesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, options ListDepOnboardingSettingEnrollmentProfilesOperationOptions, predicate EnrollmentProfileOperationPredicate) (result ListDepOnboardingSettingEnrollmentProfilesCompleteResult, err error) {
	items := make([]beta.EnrollmentProfile, 0)

	resp, err := c.ListDepOnboardingSettingEnrollmentProfiles(ctx, id, options)
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

	result = ListDepOnboardingSettingEnrollmentProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
