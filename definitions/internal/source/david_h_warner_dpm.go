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

// https://fhir-myrecord.cerner.com/r4/_byMhQ2TvVtBFrxJSYiBZz55T5g1Rbw2/metadata
func GetSourceDavidHWarnerDpm(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/_byMhQ2TvVtBFrxJSYiBZz55T5g1Rbw2/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/_byMhQ2TvVtBFrxJSYiBZz55T5g1Rbw2/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/_byMhQ2TvVtBFrxJSYiBZz55T5g1Rbw2"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/_byMhQ2TvVtBFrxJSYiBZz55T5g1Rbw2"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeDavidHWarnerDpm]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "David H. Warner, DPM"
	sourceDef.SourceType = pkg.SourceTypeDavidHWarnerDpm
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}