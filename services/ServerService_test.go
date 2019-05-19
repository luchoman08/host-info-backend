package services

import (
	"../models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerService_GetMinorSSLGrade(t *testing.T) {
	serverService := ServerService{}
	minorGrade := "C-"
	servers := []models.ServerModel{{SslGrade: "C+"}, {SslGrade: "A"}, {SslGrade: "C"}, {SslGrade: minorGrade}}
	minorGradeResult := serverService.GetMinorSSLGrade(servers)
	assert.Equal(t, minorGrade, minorGradeResult, fmt.Sprintf("Minor grade: %s, given: %s", minorGrade, minorGradeResult))
}
func TestServerService_EqualSetOfServers(t *testing.T) {
	serverService := ServerService{}
	set1 := []models.ServerModel{
		{IPAddress: "192.168.117.40", Country: "US", SslGrade: "A", Owner: "Google"},
		{IPAddress: "192.0.0.4", Country: "EN", SslGrade: "A+", Owner: "Amazon"},
	}
	set2 := []models.ServerModel{
		{IPAddress: "192.168.117.40", Country: "US", SslGrade: "A", Owner: "Google"},
		{IPAddress: "192.0.0.4", Country: "EN", SslGrade: "A", Owner: "Amazon"},
	}
	result1 := serverService.EqualSetOfServers(set1, set1)
	assert.True(t, result1, fmt.Sprintf("Servers are equal, why false is returned?"))
	result2 :=  serverService.EqualSetOfServers(set1, set2)
	assert.False(t, result2, fmt.Sprintf("Servers are equal, why false is returned?"))

}
