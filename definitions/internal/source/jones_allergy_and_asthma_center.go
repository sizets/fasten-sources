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

// https://fhir-myrecord.cerner.com/r4/4b193ed0-9700-43d2-bc5c-c454e535cd06/metadata
func GetSourceJonesAllergyAndAsthmaCenter(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/4b193ed0-9700-43d2-bc5c-c454e535cd06/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/4b193ed0-9700-43d2-bc5c-c454e535cd06/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/4b193ed0-9700-43d2-bc5c-c454e535cd06"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/4b193ed0-9700-43d2-bc5c-c454e535cd06"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeJonesAllergyAndAsthmaCenter]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Jones Allergy & Asthma Center"
	sourceDef.SourceType = pkg.SourceTypeJonesAllergyAndAsthmaCenter
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}