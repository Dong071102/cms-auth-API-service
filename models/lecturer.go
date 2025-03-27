package models

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	pgvector "github.com/pgvector/pgvector-go"
)

func GenerateLecturerCode() string {
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}

type Lecturer struct {
	LecturerID    uuid.UUID       `json:"lecturer_id" gorm:"type:uuid;primaryKey"`
	LecturerCode  string          `json:"lecturer_code" gorm:"type:varchar(100);unique"`
	FaceEmbedding pgvector.Vector `json:"face_embedding" gorm:"type:vector(512)"`
	User          User            `gorm:"foreignKey:LecturerID;references:UserID"`
}
