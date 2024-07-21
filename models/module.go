package models

type ModuleResponse struct {
	Modules []Module `json:"module"`
}

type Module struct {
	Name       string       `json:"module_name"`
	ID         string       `json:"module_id"`
	ImageURL   string       `json:"image_url"`
	Type       string       `json:"module_type"`
	Tier       string       `json:"module_tier"`
	SocketType string       `json:"module_socket_type"`
	Class      string       `json:"module_class"`
	Stat       []ModuleStat `json:"module_stat"`
}

type ModuleStat struct {
	Level          int64  `json:"level"`
	ModuleCapacity int64  `json:"module_capacity"`
	Value          string `json:"value"`
}
