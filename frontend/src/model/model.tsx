export interface Meaning {
  text: string;
  translation: string;
}

export interface Example {
  source: string;
  target: string;
}

export interface VerbConjugations {
  first_per_sin: string;
  second_per_sin: string;
  third_per_sin: string;
  first_per_plu: string;
  second_per_plu: string;
  third_per_plu: string;
}

export interface VerbInfo {
  type: string;
  infinitive?: string;
  present_participle?: string;
  past_participle?: string;
  simple_present: VerbConjugations;
  imperfect_past: VerbConjugations;
  simple_past: VerbConjugations;
  perfect_past: VerbConjugations;
  simple_future: VerbConjugations;
  conditional: VerbConjugations;
}

export interface SearchResult {
  search_word: string;
  found_word?: string;
  translation?: string;
  meanings?: Meaning[];
  examples?: Example[];
  synonyms?: string[];
  verb_info?: VerbInfo;
  sources?: Record<string, string>;
}
