package types

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupList struct {
	Value []Group `json:"value"`
}
