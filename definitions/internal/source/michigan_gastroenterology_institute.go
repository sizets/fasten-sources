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

// https://fhir.fhirpoint.open.allscripts.com/fhirroute/fhir/10055650/metadata
func GetSourceMichiganGastroenterologyInstitute(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/authorization/10055650/connect/authorize"
	sourceDef.TokenEndpoint = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/authorization/10055650/connect/token"

	sourceDef.Audience = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/fhir/10055650"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.fhirpoint.open.allscripts.com/fhirroute/fhir/10055650"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeMichiganGastroenterologyInstitute]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Michigan Gastroenterology Institute"
	sourceDef.SourceType = pkg.SourceTypeMichiganGastroenterologyInstitute
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}