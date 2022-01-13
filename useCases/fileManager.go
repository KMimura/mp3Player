package useCases

type FileManager interface {
	GetSongsFromPath(string) []string
}
