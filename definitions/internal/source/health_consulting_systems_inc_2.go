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

// https://fhir-myrecord.cerner.com/r4/yM1TVbA1pYVBqAL5r6hrPXGVr5nFBJbP/metadata
func GetSourceHealthConsultingSystemsInc2(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/yM1TVbA1pYVBqAL5r6hrPXGVr5nFBJbP/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/yM1TVbA1pYVBqAL5r6hrPXGVr5nFBJbP/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/yM1TVbA1pYVBqAL5r6hrPXGVr5nFBJbP"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/yM1TVbA1pYVBqAL5r6hrPXGVr5nFBJbP"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeHealthConsultingSystemsInc2]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Health Consulting Systems Inc"
	sourceDef.SourceType = pkg.SourceTypeHealthConsultingSystemsInc2
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}