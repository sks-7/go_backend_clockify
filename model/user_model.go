package model

import "time"

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Tag       string    `json:"tag"`
	Billable  string    `json:"billable"`
	Useremail string    `json:"useremail"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UreatedAt time.Time `json:"updated_at"`
}

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Tag         string    `json:"tag"`
	Billable    string    `json:"billable"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
	ProjectName string    `json:"project_name"`
	TotalTime   string    `json:"totalTime"`
	CreatedAt   time.Time `json:"created_at"`
	UreatedAt   time.Time `json:"updated_at"`
}
