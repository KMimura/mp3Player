package interfaceAdapters

type TestFileManagerImpl struct {
	queue []string
}

func (*TestFileManagerImpl) GetSongsFromPath(path string) []string {
	sampleOne := "../sampleSounds/sample_song.mp3"
	sampleSongs := make([]string, 1)
	sampleSongs[0] = sampleOne
	return sampleSongs
}
