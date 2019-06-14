package creates

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
} 

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemeoryStorage
)

type DiskStore struct {

}
func (s *DiskStore)Open(file string)(io.ReadWriteCloser, error) {
	return nil,nil
}

type TempStore struct {

}
func (s *TempStore)Open(file string)(io.ReadWriteCloser, error) {
	return nil,nil
}

type MemeoryStore struct {

}
func (s *MemeoryStore)Open(file string)(io.ReadWriteCloser, error) {
	return nil,nil
}

func NewStorage(storageType StorageType) Store  {
	switch storageType {
	case DiskStorage:
		return &DiskStore{}
	case TempStorage:
		return &TempStore{}
	case MemeoryStorage:
		return &MemeoryStore{}
	default:
		return &DiskStore{}
	}
}
