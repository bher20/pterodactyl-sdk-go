package pterodactyl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ApiEndpointBase    string = "api"
	ApiEndpointServers string = "client"
	ApiEndpointServer  string = "client/servers"
	ApiEndpointBackups string = "backups"
)

func buildApiUrl(pterodactylServer PterodactylServer, endpoint string, subPaths []string) string {
	url := fmt.Sprintf("%s/%s/%s", pterodactylServer.Url, ApiEndpointBase, endpoint)

	for _, path := range subPaths {
		url = fmt.Sprintf("%s/%s", url, path)
	}
	return url
}

func callApi[T any](apiObject *T, pterodactylServer PterodactylServer, method string, endpoint string, subPaths []string, data map[string]string) error {
	apiUrl := buildApiUrl(pterodactylServer, endpoint, subPaths)

	dataToSend := url.Values{}

	for k, v := range data {
		dataToSend.Set(k, v)
	}

	req, _ := http.NewRequest(method, apiUrl, strings.NewReader(dataToSend.Encode()))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", pterodactylServer.ApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		var apiErrors ApiErrors

		err = json.Unmarshal(body, &apiErrors)
		if err != nil {
			return err
		}

		return fmt.Errorf("api call failed with errors: %s", apiErrors)
	}

	err = json.Unmarshal(body, &apiObject)
	return err
}

func GetServers(pterodactylServer PterodactylServer) ([]Server, error) {
	var servers Servers
	err := callApi(&servers, pterodactylServer, http.MethodGet, ApiEndpointServers, nil, nil)
	if err != nil {
		return nil, err
	}

	if servers.Servers == nil {
		return nil, errors.New("no servers returned")
	}

	return servers.Servers, nil
}

func GetServer(pterodactylServer PterodactylServer, serverId string) (Server, error) {
	var server Server
	err := callApi(&server, pterodactylServer, http.MethodGet, ApiEndpointServer, []string{serverId}, nil)
	if err != nil {
		return server, err
	}

	return server, nil
}

func GetServerBackups(pterodactylServer PterodactylServer, server Server) ([]Backup, error) {
	var backups Backups
	err := callApi(&backups, pterodactylServer, http.MethodGet, ApiEndpointServer, []string{server.Attributes.UUID, ApiEndpointBackups}, nil)
	if err != nil {
		return nil, err
	}

	if backups.Backups == nil {
		return nil, errors.New("no backups returned")
	}

	return backups.Backups, nil
}

func GetServerBackup(pterodactylServer PterodactylServer, server Server, backupId string) (Backup, error) {
	var backup Backup
	err := callApi(&backup, pterodactylServer, http.MethodGet, ApiEndpointServer, []string{server.Attributes.UUID, ApiEndpointBackups, backupId}, nil)
	if err != nil {
		return backup, err
	}

	return backup, nil
}

func DownloadServerBackup(pterodactylServer PterodactylServer, server Server, backupId string) ([]byte, error) {
	var backupUrl BackupUrl
	err := callApi(&backupUrl, pterodactylServer, http.MethodGet, ApiEndpointServer, []string{server.Attributes.UUID, ApiEndpointBackups, backupId, "download"}, nil)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodGet, backupUrl.Attributes.URL, nil)
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		var apiErrors ApiErrors

		err = json.Unmarshal(body, &apiErrors)
		if err != nil {
			return nil, err
		}
	}

	return body, nil
}

func BackupServer(pterodactylServer PterodactylServer, server Server) (Backup, error) {
	var backup Backup

	err := callApi(&backup, pterodactylServer, http.MethodPost, fmt.Sprintf("%s/%s/%s", ApiEndpointServer, server.Attributes.UUID, ApiEndpointBackups), nil, nil)
	if err != nil {
		return backup, err
	}

	return backup, nil
}
