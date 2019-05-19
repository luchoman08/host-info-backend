package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinorSSLGrade(t *testing.T) {
	minorGrade := "C"
	minorGradeResult := GetMinorSSLGrade(minorGrade, "A")
	assert.Equal(t, minorGradeResult, minorGrade, fmt.Sprintf("Minor grade incorrect: %s, want: %s.", minorGradeResult, minorGrade))
}
func TestGetMinorSSLGradeWithEmptyInput(t *testing.T) {
	minorGrade := "C"
	minorGradeResult := GetMinorSSLGrade(minorGrade, "")
	assert.Equal(t, minorGradeResult, minorGrade, fmt.Sprintf("Minor grade incorrect: %s, want: %s.", minorGradeResult, minorGrade))
}
func TestGetGetMinorSSLGradeWithModifier(t *testing.T) {
	minorGrade := "A-"
	minorGradeResult := GetMinorSSLGrade("A+", minorGrade)
	assert.Equal(t, minorGrade, minorGradeResult, fmt.Sprintf("Minor grade incorrect: %s, want: %s.", minorGradeResult, minorGrade))
}
func TestGetMinorSSLGradeFromList(t *testing.T) {
	minorGrade := "C-"
	grades := []string{"A+", "A-", "C", minorGrade}
	minorGradeResult := GetMinorSSLGradeFromList(grades)
	assert.Equal(t, minorGradeResult, minorGrade, fmt.Sprintf("Minor grade incorrect: %s, want: %s", minorGradeResult, minorGrade))
}
