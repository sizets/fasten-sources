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

// https://fhir-myrecord.cerner.com/r4/235ba15f-6bef-42d8-831c-5f3f60c7755f/metadata
func GetSourceTregoCountyLemkeMemorialHospital(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/235ba15f-6bef-42d8-831c-5f3f60c7755f/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/235ba15f-6bef-42d8-831c-5f3f60c7755f/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/235ba15f-6bef-42d8-831c-5f3f60c7755f"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/235ba15f-6bef-42d8-831c-5f3f60c7755f"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTregoCountyLemkeMemorialHospital]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Trego County-Lemke Memorial Hospital"
	sourceDef.SourceType = pkg.SourceTypeTregoCountyLemkeMemorialHospital
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}