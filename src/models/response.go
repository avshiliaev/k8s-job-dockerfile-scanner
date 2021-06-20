package models

type JsonResponse struct {
	Data map[string]map[string][]string `json:"data,omitempty"`
}
