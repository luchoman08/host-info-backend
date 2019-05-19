package app

import "testing"

func TestGetMinorSSLGrade(t *testing.T) {
	minorGrade := "C"
	minorGradeResult:= GetMinorSSLGrade(minorGrade, "C")
	if minorGradeResult != minorGrade  {
		t.Errorf("Minor grade incorrect: %s, want: %s.", minorGradeResult,  minorGrade )
	}
}
func TestGetMinorSSLGradeWithEmptyInput(t *testing.T) {
	minorGrade := "C"
	minorGradeResult:= GetMinorSSLGrade(minorGrade, "")
	if minorGradeResult != minorGrade  {
		t.Errorf("Minor grade incorrect: %s, want: %s.", minorGradeResult,  minorGrade )
	}
}
func TestGetGetMinorSSLGradeWithModifier(t *testing.T) {
	minorGrade := "A-"
	minorGradeResult := GetMinorSSLGrade("A+", minorGrade)
	if minorGradeResult != minorGrade  {
		t.Errorf("Minor grade incorrect: %s, want: %s.", minorGradeResult, minorGrade)
	}
}
func TestGetMinorSSLGradeFromList(t *testing.T) {
	minorGrade := "C-"
	grades := []string{"A+", "A-", "C", minorGrade}
	minorGradeResult := GetMinorSSLGradeFromList(grades)
	if minorGradeResult != minorGrade {
		t.Errorf("Minor grade incorrect: %s, want: %s", minorGradeResult, minorGrade)
	}
}