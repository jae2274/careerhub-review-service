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
		t.Run("case 1", func(t *testing.T) {
			originalString := "(주)엔터인"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "엔터인", result)
		})
		t.Run("case 2", func(t *testing.T) {
			originalString := "(주)리버티랩스"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "리버티랩스", result)
		})
		t.Run("case 3", func(t *testing.T) {
			originalString := "(주)더블티"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "더블티", result)
		})
		t.Run("case 4", func(t *testing.T) {
			originalString := "(주)에이아이세스"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "에이아이세스", result)
		})
		t.Run("case 5", func(t *testing.T) {
			originalString := "(주)사람인 TEST"
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "사람인 TEST", result)
		})
		t.Run("case 6", func(t *testing.T) {
			originalString := " (주) 라츠온 "
			result := company.RefineNameForSearch(originalString)
			require.Equal(t, "라츠온", result)
		})

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
