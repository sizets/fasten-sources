// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/58434/.well-known/smart-configuration
// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/58434/metadata
func GetSourceSusilaSubramanianMdsc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://open.allscripts.com/fhirroute/patientauthv2/616ef409-b1d1-464c-a53a-a8990ab1ddf4/connect/authorize"
	sourceDef.TokenEndpoint = "https://open.allscripts.com/fhirroute/patientauthv2/616ef409-b1d1-464c-a53a-a8990ab1ddf4/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/58434"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/58434"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeSusilaSubramanianMdsc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Susila Subramanian Mdsc"
	sourceDef.SourceType = pkg.SourceTypeSusilaSubramanianMdsc
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}