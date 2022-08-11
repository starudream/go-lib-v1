package testx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	AssertNew = assert.New

	AssertEmptyf    = assert.Emptyf
	AssertNotEmptyf = assert.NotEmptyf
	AssertNilf      = assert.Nilf
	AssertNotNilf   = assert.NotNilf

	AssertEqualf          = assert.Equalf
	AssertNotEqualf       = assert.NotEqualf
	AssertEqualValuesf    = assert.EqualValuesf
	AssertNotEqualValuesf = assert.NotEqualValuesf

	AssertErrorf      = assert.Errorf
	AssertNoErrorf    = assert.NoErrorf
	AssertErrorIsf    = assert.ErrorIsf
	AssertNotErrorIsf = assert.NotErrorIsf

	RequireNew = require.New

	RequireEmptyf    = require.Emptyf
	RequireNotEmptyf = require.NotEmptyf
	RequireNilf      = require.Nilf
	RequireNotNilf   = require.NotNilf

	RequireEqualf          = require.Equalf
	RequireNotEqualf       = require.NotEqualf
	RequireEqualValuesf    = require.EqualValuesf
	RequireNotEqualValuesf = require.NotEqualValuesf

	RequireErrorf      = require.Errorf
	RequireNoErrorf    = require.NoErrorf
	RequireErrorIsf    = require.ErrorIsf
	RequireNotErrorIsf = require.NotErrorIsf
)
