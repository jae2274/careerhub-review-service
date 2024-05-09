package company

import (
	"testing"

	"github.com/jae2274/careerhub-review-service/common/domain/company"
	"github.com/stretchr/testify/require"
)

func TestRefineNameForSearch(t *testing.T) {
	t.Run("return original string", func(t *testing.T) {
		originalString := "testCompany"
		result := company.RefineNameForSearch(originalString)
		require.Equal(t, originalString, result)
	})

	t.Run("return refined string when has parenthesis", func(t *testing.T) {
		originalString := "testCompany(test)"
		result := company.RefineNameForSearch(originalString)
		require.Equal(t, "testCompany", result)
	})
}
