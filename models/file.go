package models

import (
    "gorm.io/gorm"
)

type File struct {
    gorm.Model
    Name      string `json:"name"`
    Path      string `json:"path"`
    Size      int64  `json:"size"`
    FileType  string `json:"file_type"`
    Uploader  string `json:"uploader"`
}
