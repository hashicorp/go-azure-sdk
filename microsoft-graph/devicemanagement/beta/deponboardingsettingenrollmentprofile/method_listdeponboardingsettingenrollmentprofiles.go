package deponboardingsettingenrollmentprofile

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

type ListDepOnboardingSettingEnrollmentProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EnrollmentProfile
}

type ListDepOnboardingSettingEnrollmentProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EnrollmentProfile
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

// ListDepOnboardingSettingEnrollmentProfiles ...
func (c DepOnboardingSettingEnrollmentProfileClient) ListDepOnboardingSettingEnrollmentProfiles(ctx context.Context, id DeviceManagementDepOnboardingSettingId) (result ListDepOnboardingSettingEnrollmentProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDepOnboardingSettingEnrollmentProfilesCustomPager{},
		Path:       fmt.Sprintf("%s/enrollmentProfiles", id.ID()),
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
		Values *[]beta.EnrollmentProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDepOnboardingSettingEnrollmentProfilesComplete retrieves all the results into a single object
func (c DepOnboardingSettingEnrollmentProfileClient) ListDepOnboardingSettingEnrollmentProfilesComplete(ctx context.Context, id DeviceManagementDepOnboardingSettingId) (ListDepOnboardingSettingEnrollmentProfilesCompleteResult, error) {
	return c.ListDepOnboardingSettingEnrollmentProfilesCompleteMatchingPredicate(ctx, id, EnrollmentProfileOperationPredicate{})
}

// ListDepOnboardingSettingEnrollmentProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingEnrollmentProfileClient) ListDepOnboardingSettingEnrollmentProfilesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDepOnboardingSettingId, predicate EnrollmentProfileOperationPredicate) (result ListDepOnboardingSettingEnrollmentProfilesCompleteResult, err error) {
	items := make([]beta.EnrollmentProfile, 0)

	resp, err := c.ListDepOnboardingSettingEnrollmentProfiles(ctx, id)
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
