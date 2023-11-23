package masking

import (
	"os"
	"strings"
)

type FileProducer struct {
	filePath string
}

func NewFileProducer(filePath string) *FileProducer {
	return &FileProducer{
		filePath: filePath,
	}
}

func (fp *FileProducer) produce() ([]string, error) {
	content, err := os.ReadFile(fp.filePath)
	if err != nil {
		return nil, err
	}

	// we added this cause we can't change link masker function
	// that's why we need add ' ' to the end of every line
	// and after masking we will delete it
	text := strings.Split(string(content), "\n")

	for i := 0; i < len(text); i++ {
		text[i] = string(append([]rune(text[i]), rune(' ')))
	}

	return text, nil
}
