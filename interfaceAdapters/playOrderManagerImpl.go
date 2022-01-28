package interfaceAdapters

type PlayOrderManagerImpl struct {
	queue []string
	index int
}

func (pom *PlayOrderManagerImpl) GetSongPathes(path string) []string {
	pom.index = 0
	return nil
}

func (pom *PlayOrderManagerImpl) UpdateIndex(index int) {
	pom.index++
	return
}

func (*PlayOrderManagerImpl) GetCurrentSong() string {
	return ""
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
