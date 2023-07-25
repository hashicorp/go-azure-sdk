package staticsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserRolesInvitationLinkOperationResponse struct {
	HttpResponse *http.Response
	Model        *StaticSiteUserInvitationResponseResource
}

// CreateUserRolesInvitationLink ...
func (c StaticSitesClient) CreateUserRolesInvitationLink(ctx context.Context, id StaticSiteId, input StaticSiteUserInvitationRequestResource) (result CreateUserRolesInvitationLinkOperationResponse, err error) {
	req, err := c.preparerForCreateUserRolesInvitationLink(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateUserRolesInvitationLink", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateUserRolesInvitationLink", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateUserRolesInvitationLink(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateUserRolesInvitationLink", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateUserRolesInvitationLink prepares the CreateUserRolesInvitationLink request.
func (c StaticSitesClient) preparerForCreateUserRolesInvitationLink(ctx context.Context, id StaticSiteId, input StaticSiteUserInvitationRequestResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/createUserInvitation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateUserRolesInvitationLink handles the response to the CreateUserRolesInvitationLink request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForCreateUserRolesInvitationLink(resp *http.Response) (result CreateUserRolesInvitationLinkOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
