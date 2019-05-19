package services

import (
	"testing"
	"../models"
)

func TestGetMinorSSLGrade(t *testing.T) {
	serverService := ServerService{}
	minorGrade := "C-"
	servers := []models.ServerModel{{SslGrade:"C+"}, {SslGrade:"A"}, {SslGrade:"C"}, {SslGrade:minorGrade}}
	minorGradeResult := serverService.GetMinorSSLGrade(servers)
	if minorGradeResult != minorGrade {
		t.Errorf("Minor grade: %s, given: %s", minorGrade, minorGradeResult)
	}
}
