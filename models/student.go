package models

// student model

type Student struct{
	StudentID string `json:"StudentID"`
	Name string `json:"Name"`
	Age int `json:"Age"`
	School string `json:"School"`
}
