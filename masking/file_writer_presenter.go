package masking

import (
	"os"
)

type FileWriterPresenter struct {
	filePath string
}

func NewFileWriterPresenter(filePath string) *FileWriterPresenter {
	return &FileWriterPresenter{
		filePath: filePath,
	}
}

func (fw *FileWriterPresenter) present(data []string) error {
	file, err := os.Create(fw.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		len_line := len([]rune(line))

		// len_line-2 cause we need remove the ' ' in the end of line
		_, err := file.WriteString(line[:len_line-1] + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
