package interfaceAdapters

type TestFileManagerImpl struct {
	queue []string
}

func (*TestFileManagerImpl) GetSongsFromPath(path string) []string {
	return nil
}
