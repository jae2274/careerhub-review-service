package company

import (
	"testing"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
	"github.com/stretchr/testify/require"
)

func TestRefineNameForSearch(t *testing.T) {

	t.Run("return original string", func(t *testing.T) {
		originalString := "testCompany"
		result := company.RefineNameForSearch(originalString)
		require.Equal(t, originalString, result)
	})

	t.Run("return refined string when has parenthesis lately", func(t *testing.T) {
		originalString := "testCompany(test)"
		result := company.RefineNameForSearch(originalString)
		require.Equal(t, "testCompany", result)
	})

	t.Run("return refined string when has parenthesis early", func(t *testing.T) {
		originalString := "(주)엔터인"
		result := company.RefineNameForSearch(originalString)
		require.Equal(t, "엔터인", result)
	})

	t.Run("return refined string when has parenthesis with whitespace", func(t *testing.T) {
		t.Run("case 1", func(t *testing.T) {
			originalString := "testCompany (test)"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "testCompany", result)
		})

		t.Run("case 2", func(t *testing.T) {
			originalString := "(주) 엔터인"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "엔터인", result)
		})
	})
}
