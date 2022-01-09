package interfaceAdapters

type FileManagerImpl struct {
	queue []string
}

func (*FileManagerImpl) GetSongsFromPath(path string) []string {
	return nil
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
