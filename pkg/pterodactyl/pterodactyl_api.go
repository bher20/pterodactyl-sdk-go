package pterodactyl

import (
	"time"
)

type PterodactylServer struct {
	ApiKey string `json:"apiKey"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

type Servers struct {
	Object  string      `json:"object"`
	Servers []Server    `json:"data"`
	Meta    ApiMetaData `json:"meta"`
}

type ApiMetaData struct {
	Pagination ApiPagination `json:"pagination"`
}
type ApiLinks struct {
}
type ApiPagination struct {
	Total       int      `json:"total"`
	Count       int      `json:"count"`
	PerPage     int      `json:"per_page"`
	CurrentPage int      `json:"current_page"`
	TotalPages  int      `json:"total_pages"`
	Links       ApiLinks `json:"links"`
}

type ApiErrors struct {
	Errors []ApiError `json:"errors"`
}
type ApiError struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}

type Server struct {
	Object     string `json:"object"`
	Attributes struct {
		ServerOwner            bool   `json:"server_owner"`
		Identifier             string `json:"identifier"`
		InternalID             int    `json:"internal_id"`
		UUID                   string `json:"uuid"`
		Name                   string `json:"name"`
		Node                   string `json:"node"`
		IsNodeUnderMaintenance bool   `json:"is_node_under_maintenance"`
		SftpDetails            struct {
			IP   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"sftp_details"`
		Description string `json:"description"`
		Limits      struct {
			Memory      int  `json:"memory"`
			Swap        int  `json:"swap"`
			Disk        int  `json:"disk"`
			Io          int  `json:"io"`
			CPU         int  `json:"cpu"`
			Threads     any  `json:"threads"`
			OomDisabled bool `json:"oom_disabled"`
		} `json:"limits"`
		Invocation    string   `json:"invocation"`
		DockerImage   string   `json:"docker_image"`
		EggFeatures   []string `json:"egg_features"`
		FeatureLimits struct {
			Databases   int `json:"databases"`
			Allocations int `json:"allocations"`
			Backups     int `json:"backups"`
		} `json:"feature_limits"`
		Status         any  `json:"status"`
		IsSuspended    bool `json:"is_suspended"`
		IsInstalling   bool `json:"is_installing"`
		IsTransferring bool `json:"is_transferring"`
		Renewable      bool `json:"renewable"`
		Renewal        int  `json:"renewal"`
		Bg             any  `json:"bg"`
		Relationships  struct {
			Allocations struct {
				Object string `json:"object"`
				Data   []struct {
					Object     string `json:"object"`
					Attributes struct {
						ID        int    `json:"id"`
						IP        string `json:"ip"`
						IPAlias   string `json:"ip_alias"`
						Port      int    `json:"port"`
						Notes     any    `json:"notes"`
						IsDefault bool   `json:"is_default"`
					} `json:"attributes"`
				} `json:"data"`
			} `json:"allocations"`
			Variables struct {
				Object string `json:"object"`
				Data   []struct {
					Object     string `json:"object"`
					Attributes struct {
						Name         string `json:"name"`
						Description  string `json:"description"`
						EnvVariable  string `json:"env_variable"`
						DefaultValue string `json:"default_value"`
						ServerValue  string `json:"server_value"`
						IsEditable   bool   `json:"is_editable"`
						Rules        string `json:"rules"`
					} `json:"attributes"`
				} `json:"data"`
			} `json:"variables"`
		} `json:"relationships"`
	} `json:"attributes"`
	Meta struct {
		IsServerOwner   bool     `json:"is_server_owner"`
		UserPermissions []string `json:"user_permissions"`
	} `json:"meta"`
}

type Backups struct {
	Object  string      `json:"object"`
	Backups []Backup    `json:"data"`
	Meta    ApiMetaData `json:"meta"`
}
type Backup struct {
	Object     string `json:"object"`
	Attributes struct {
		UUID         string    `json:"uuid"`
		Name         string    `json:"name"`
		IgnoredFiles []any     `json:"ignored_files"`
		Sha256Hash   string    `json:"sha256_hash"`
		Bytes        int       `json:"bytes"`
		CreatedAt    time.Time `json:"created_at"`
		CompletedAt  time.Time `json:"completed_at,omitempty"`
	} `json:"attributes"`
}

type BackupUrl struct {
	Object     string `json:"object"`
	Attributes struct {
		URL string `json:"url"`
	} `json:"attributes"`
}
