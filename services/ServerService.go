package services

import (
	"../app"
	"../interfaces"
	"../models"
	"github.com/golang/glog"
	"github.com/luchoman08/ssllabs"
	"sort"
)

// ServerService implements the methods for access the server data repository
type ServerService struct {
	interfaces.ServerRepository
}

// GetServersOfDomain returns the servers related to the given domain, the servers are finded
// locally
func (service *ServerService) GetServersOfDomain(domain *models.DomainModel) []models.ServerModel {
	return service.GetServersForDomain(domain)
}

// ServerCollectionsDistinct compare two arrays of servers and return if their are distinct
// collections of servers or if both collections are equal.
// The comparison is made in a profound way.
func (service *ServerService) ServerCollectionsDistinct([]models.ServerModel, []models.ServerModel) bool {
	return false
}

// ServiceCreateServer create a server in local storage
func (service *ServerService) ServiceCreateServer(server *models.ServerModel) {
	service.CreateServer(server)
}

// GenerateServer constructs a server based in a ssllabs endpoint and in a existing domain,
func (service *ServerService) GenerateServer(domain models.DomainModel, endpoint ssllabs.Endpoint) (server models.ServerModel, err error) {
	exists := service.ExistsByIPAdress(endpoint.IPAddress)
	server, err = service.GetServerFromExtern(endpoint)
	if !exists {
		if err != nil {
			glog.Warning("Get server From extern has been failed on who is info get. Error: ", err)
		}
		server.DomainID = domain.ID
	}
	return
}
func (service *ServerService) GetServerByIP(ip  string) models.ServerModel {
	server, _ :=  service.GetServerFromLocal(ip)
	return server
}
// UpdateLocalServersIfChanged check if
func (service *ServerService) UpdateLocalServersIfChanged(servers []models.ServerModel) {
	for _, server := range servers {
		localServer, found := service.GetServerFromLocal(server.IPAddress)
		if !found {
			service.CreateServer(&server)
		} else if !service.EqualServers(server, localServer) {
			server.ID = localServer.ID
			server.DomainID = localServer.DomainID
			service.UpdateServer(&server)
		}
	}
}

// EqualServers check if the servers have the same attributes and return a response
// Does not check the attributes updated_at, ID and all attributes than are not related
// with main server info.
func (service *ServerService) EqualServers(s1, s2 models.ServerModel) bool {
	return s1.SslGrade == s2.SslGrade && s1.Owner == s2.Owner && s1.Country == s2.Country && s1.IPAddress == s2.IPAddress
}

// EqualSetOfServers compare two sets of servers and return if are equal or not
// Does not check the attributes like ID or updated_at than are no related directly
// with main server info.
func (service *ServerService) EqualSetOfServers(set1, set2 []models.ServerModel) bool {
	if len(set1) != len(set2) {
		return false
	}
	minnorServer := func(s1, s2 models.ServerModel) bool {
		return s1.IPAddress < s2.IPAddress
	}
	sort.Slice(set1, func(i, j int) bool { return minnorServer(set1[i], set1[j]) })
	sort.Slice(set2, func(i, j int) bool { return minnorServer(set2[i], set2[j]) })
	for i := 0; i < len(set1); i++ {
		if !service.EqualServers(set1[i], set2[i]) {
			return false
		}
	}
	return true
}

// GetMinorSSLGrade returns the minor ssl grade from a collection of servers, if empty collection is given
// return empty string
func (service *ServerService) GetMinorSSLGrade(servers []models.ServerModel) (ssllGrade string) {
	ssllGrade = ""
	if len(servers) == 0 {
		return
	}
	for _, server := range servers {
		ssllGrade = app.GetMinorSSLGrade(ssllGrade, server.SslGrade)
	}
	return
}
