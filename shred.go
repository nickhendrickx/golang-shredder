package shred

import (
		"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
)


func Shred(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	} else if fileInfo.IsDir() {
		return errors.New("Given filepath is a directory not a file")
	} else {
		for i = 0; i < 3; i++ {
			fileSize := fileInfo.Size() //depended on system
			err = shredFile(path, fileSize)
			if err != nil {
				return err
			}
		}
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func shredFile(path string, size int64) error{
	f, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return err
	} 
	
	buff := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		file.Close()
		return err
	}
	
	_, err = file.WriteAt(buff, 0)
	if err != nil {
		file.Close()
		return err
	}
	err = file.Sync()
	if err != nil {
		file.Close()
		return err
	}
	
	err = file.Close()
	if err != nil {
		return err
	}
	
	return nil
}