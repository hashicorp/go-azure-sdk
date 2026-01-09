package remotedesktopsecurityconfigurationtargetdevicegroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TargetDeviceGroup
}

type ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TargetDeviceGroup
}

type ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions struct {
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

func DefaultListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions() ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions {
	return ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions{}
}

func (o ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRemoteDesktopSecurityConfigurationTargetDeviceGroups - List targetDeviceGroups. Get a list of the
// targetDeviceGroup objects and their properties on the remoteDesktopSecurityConfiguration resource on the
// servicePrincipal. Any user authenticating using the Microsoft Entra ID Remote Desktop Services (RDS) authentication
// protocol to a Microsoft Entra joined or Microsoft Entra hybrid joined device that belongs to the targetDeviceGroup
// will get SSO.
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) ListRemoteDesktopSecurityConfigurationTargetDeviceGroups(ctx context.Context, id beta.ServicePrincipalId, options ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions) (result ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/remoteDesktopSecurityConfiguration/targetDeviceGroups", id.ID()),
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
		Values *[]beta.TargetDeviceGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsComplete retrieves all the results into a single object
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsComplete(ctx context.Context, id beta.ServicePrincipalId, options ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions) (ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteResult, error) {
	return c.ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteMatchingPredicate(ctx, id, options, TargetDeviceGroupOperationPredicate{})
}

// ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteMatchingPredicate(ctx context.Context, id beta.ServicePrincipalId, options ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsOperationOptions, predicate TargetDeviceGroupOperationPredicate) (result ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteResult, err error) {
	items := make([]beta.TargetDeviceGroup, 0)

	resp, err := c.ListRemoteDesktopSecurityConfigurationTargetDeviceGroups(ctx, id, options)
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

	result = ListRemoteDesktopSecurityConfigurationTargetDeviceGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
