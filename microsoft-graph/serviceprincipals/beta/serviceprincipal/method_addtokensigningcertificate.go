package serviceprincipal

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

type AddTokenSigningCertificateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SelfSignedCertificate
}

// AddTokenSigningCertificate - Invoke action addTokenSigningCertificate. Creates a self-signed signing certificate and
// returns a selfSignedCertificate object, which is the public part of the generated certificate. The self-signed
// signing certificate is composed of the following objects which are added to the servicePrincipal: + The
// keyCredentials object with the following objects: + A private key object with usage set to Sign. + A public key
// object with usage set to Verify. + The passwordCredentials object. All the objects have the same value of
// customKeyIdentifier. The passwordCredential is used to open the PFX file (private key). It and the associated private
// key object have the same value of keyId. Once set during creation through the displayName property, the subject of
// the certificate cannot be updated. The startDateTime is set to the same time the certificate is created using the
// action. The endDateTime can be up to three years after the certificate is created.
func (c ServicePrincipalClient) AddTokenSigningCertificate(ctx context.Context, id beta.ServicePrincipalId, input AddTokenSigningCertificateRequest) (result AddTokenSigningCertificateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/addTokenSigningCertificate", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.SelfSignedCertificate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
