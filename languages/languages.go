package languages

type Lanuage struct {
	StopWords     map[string]uint8
	Contractions  map[string]string
	Abbreviations map[string]string
}

func NewLanguage(stopwrds map[string]uint8, contractions map[string]string, abbreviations map[string]string) Lanuage {
	return Lanuage{
		StopWords:     stopwrds,
		Contractions:  contractions,
		Abbreviations: abbreviations,
	}
}
