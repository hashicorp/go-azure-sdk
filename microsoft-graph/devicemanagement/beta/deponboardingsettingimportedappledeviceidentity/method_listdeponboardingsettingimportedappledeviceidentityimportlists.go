package deponboardingsettingimportedappledeviceidentity

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

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedAppleDeviceIdentityResult
}

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedAppleDeviceIdentityResult
}

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions() ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions {
	return ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions{}
}

func (o ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions) ToOData() *odata.Query {
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

func (o ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDepOnboardingSettingImportedAppleDeviceIdentityImportLists - Invoke action importAppleDeviceIdentityList
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentityImportLists(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, input ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsRequest, options ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions) (result ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCustomPager{},
		Path:          fmt.Sprintf("%s/importedAppleDeviceIdentities/importAppleDeviceIdentityList", id.ID()),
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
		Values *[]beta.ImportedAppleDeviceIdentityResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsComplete retrieves all the results into a single object
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsComplete(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, input ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsRequest, options ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions) (ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteResult, error) {
	return c.ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteMatchingPredicate(ctx, id, input, options, ImportedAppleDeviceIdentityResultOperationPredicate{})
}

// ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, input ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsRequest, options ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions, predicate ImportedAppleDeviceIdentityResultOperationPredicate) (result ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteResult, err error) {
	items := make([]beta.ImportedAppleDeviceIdentityResult, 0)

	resp, err := c.ListDepOnboardingSettingImportedAppleDeviceIdentityImportLists(ctx, id, input, options)
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

	result = ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
