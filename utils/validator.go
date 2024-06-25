package utils

import "strings"

var validFileTypes = []string{".jpg", ".jpeg", ".png", ".pdf", ".doc", ".docx"}

func IsValidFileType(fileType string) bool {
    for _, validType := range validFileTypes {
        if strings.EqualFold(validType, fileType) {
            return true
        }
    }
    return false
}
