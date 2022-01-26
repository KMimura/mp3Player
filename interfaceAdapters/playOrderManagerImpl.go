package interfaceAdapters

type PlayOrderManagerImpl struct {
	queue []string
	index int
}

func (*PlayOrderManagerImpl) GetSongPathes(path string) []string {
	return nil
}

func (*PlayOrderManagerImpl) UpdateIndex(index int) {
	return
}

func NewFileManagerImpl(path string) *PlayOrderManagerImpl {
	pom := &PlayOrderManagerImpl{queue: nil}
	var songs []string
	if path != "" {
		songs = pom.GetSongPathes(path)
	}
	pom.queue = songs
	return pom
}
