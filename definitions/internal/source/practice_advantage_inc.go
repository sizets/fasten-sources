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

// https://fhir-myrecord.cerner.com/r4/e2feebdd-7feb-4075-9a95-20e06d5a42e5/metadata
func GetSourcePracticeAdvantageInc(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/e2feebdd-7feb-4075-9a95-20e06d5a42e5/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/e2feebdd-7feb-4075-9a95-20e06d5a42e5/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/e2feebdd-7feb-4075-9a95-20e06d5a42e5"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/e2feebdd-7feb-4075-9a95-20e06d5a42e5"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePracticeAdvantageInc]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Practice Advantage, Inc"
	sourceDef.SourceType = pkg.SourceTypePracticeAdvantageInc
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}