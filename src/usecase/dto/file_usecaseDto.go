package dto

type CreateFile struct {
	Name        string
	Directory   string
	Description string
	MimeType    string
}

type UpdateFile struct {
	Description string
}

type File struct {
	Id          int
	Name        string
	Directory   string
	Description string
	MimeType    string
}
