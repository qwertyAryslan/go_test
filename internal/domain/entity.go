package domain

type Entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data Data   `json:"data"`
}

type Data struct {
	Year         int     `json:"year"`
	Price        float32 `json:"price"`
	CPUModel     string  `json:"cpu_model"`
	HardDiskSize string  `json:"hard_disk_size"`
	Color        string  `json:"color,omitempty"`
}
