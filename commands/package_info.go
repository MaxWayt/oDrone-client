package commands

type packageInfo struct {
	Name         string `json:"name"`
	AuthorEmail  string `json:"author_email"`
	Dependencies string `json:"dependencies,omitempty"`
	Revision     int64  `json:"revision"`
	Summary      string `json:"summary,omitempty"`
	FileUrl      string `json:"file_url"`
}
