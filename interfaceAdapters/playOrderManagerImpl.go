package interfaceAdapters

var acceptedExtentions = []string{"mp3"}

type PlayOrderManagerImpl struct {
	queue []string
	index int
}

func (pom *PlayOrderManagerImpl) UpdateIndex(index int) {
	pom.index++
	return
}

func (*PlayOrderManagerImpl) GetCurrentSong() string {
	return ""
}

func NewPlayOrderManagerImpl() *PlayOrderManagerImpl {
	pom := &PlayOrderManagerImpl{queue: nil}
	return pom
}
