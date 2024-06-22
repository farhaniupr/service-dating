package model

type RabbitQueue struct {
	Arguments struct {
	} `json:"arguments"`
	AutoDelete bool   `json:"auto_delete"`
	Durable    bool   `json:"durable"`
	Exclusive  bool   `json:"exclusive"`
	Name       string `json:"name"`
	Node       string `json:"node"`
	State      string `json:"state"`
	Type       string `json:"type"`
	Vhost      string `json:"vhost"`
}
