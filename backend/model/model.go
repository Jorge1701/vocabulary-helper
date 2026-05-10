package model

type Example struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type VerbInfo struct {
	Type             string           `json:"type"`
	Infinitive       string           `json:"infinitive,omitempty"`
	PresentPaticiple string           `json:"present_participle,omitempty"`
	PastParticiple   string           `json:"past_participle,omitempty"`
	SimplePresent    VerbConjugations `json:"simple_present"`
	ImperfectPast    VerbConjugations `json:"imperfect_past"`
	SimplePast       VerbConjugations `json:"simple_past"`
	PerfectPast      VerbConjugations `json:"perfect_past"`
	SimpleFuture     VerbConjugations `json:"simple_future"`
	Conditional      VerbConjugations `json:"conditional"`
}

type VerbConjugations struct {
	FirstPersonSingular  string `json:"first_per_sin"`
	SecondPersonSingular string `json:"second_per_sin"`
	ThirdPersonSingular  string `json:"third_per_sin"`
	FirstPersonPlural    string `json:"first_per_plu"`
	SecondPersonPlural   string `json:"second_per_plu"`
	ThirdPersonPlural    string `json:"third_per_plu"`
}

type SearchResult struct {
	SearchWord  string            `json:"search_word"`
	FoundWord   string            `json:"found_word,omitempty"`
	Translation string            `json:"translation,omitempty"`
	Type        string            `json:"type,omitempty"`
	Meanings    []string          `json:"meanings,omitempty"`
	Examples    []Example         `json:"examples,omitempty"`
	Synonyms    []string          `json:"synonyms,omitempty"`
	VerbInfo    *VerbInfo         `json:"verb_info,omitempty"`
	Sources     map[string]string `json:"sources,omitempty"`
}
