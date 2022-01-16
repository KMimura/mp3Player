package interfaceAdapters

type FileManagerImpl struct {
	queue []string
	index int
}

func (*FileManagerImpl) GetSongsFromPath(path string) []string {
	return nil
}

func (*FileManagerImpl) UpdateIndex(index int) {
	return
}

func NewFileManagerImpl(path string) *FileManagerImpl {
	fm := &FileManagerImpl{queue: nil}
	var songs []string
	if path != "" {
		songs = fm.GetSongsFromPath(path)
	}
	fm.queue = songs
	return fm
}
