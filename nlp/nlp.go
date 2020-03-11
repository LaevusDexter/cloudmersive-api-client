package nlp

// ExtractEntities - Extract the named entitites from an input string.
func (client Client) ExtractEntities(input string) (resp EntitiesResponse, err error) {
	err = client.apiCall("/nlp-v2/extract-entities", inputString{input}, &resp)

	return
}

// DetectLanguage - Automatically determine which language a text string is written in.
// Supports Danish (DAN), German (DEU), English (ENG), French (FRA), Italian (ITA), Japanese (JPN),
// Korean (KOR), Dutch (NLD), Norwegian (NOR), Portuguese (POR), Russian (RUS), Spanish (SPA),
// Swedish (SWE), Chinese (ZHO).
func (client Client) DetectLanguage(text string) (resp DetectLanguageResponse, err error) {
	err = client.apiCall("/nlp-v2/language/detect", textToDetect{text}, &resp)

	return
}

// ParseToTree - Parses the input string into a Penn Treebank syntax tree
func (client Client) ParseToTree(input string) (resp ParseTree, err error) {
	err = client.apiCall("/nlp-v2/parse/tree", inputString{input}, &resp)

	return
}

// PoSTagSentence - Part-of-speech (POS) tag a text
func (client Client) PoSTagSentences(text string) (resp TaggedSentencesResponse, err error) {
	err = client.apiCall("/nlp-v2/pos/tag/sentence", inputText{text}, &resp)

	return
}

// PoSTagVerbs - Part-of-speech (POS) tag a text, find the verbs
func (client Client) PoSTagVerbs(text string) (resp TaggedSentencesResponse, err error) {
	err = client.apiCall("/nlp-v2/pos/tag/verbs", inputText{text}, &resp)

	return
}

// PoSTagNouns - Part-of-speech (POS) tag a text, find the nouns
func (client Client) PoSTagNouns(text string) (resp TaggedSentencesResponse, err error) {
	err = client.apiCall("/nlp-v2/pos/tag/nouns", inputText{text}, &resp)

	return
}

// PoSTagAdjectives - Part-of-speech (POS) tag a text, find the adjectives
func (client Client) PoSTagAdjectives(text string) (resp TaggedSentencesResponse, err error) {
	err = client.apiCall("/nlp-v2/pos/tag/adjectives", inputText{text}, &resp)

	return
}

// PoSTagAdverbs - Part-of-speech (POS) tag a text, find the adverbs
func (client Client) PoSTagAdverbs(text string) (resp TaggedSentencesResponse, err error) {
	err = client.apiCall("/nlp-v2/pos/tag/adverbs", inputText{text}, &resp)

	return
}

// PoSTagPronouns - Part-of-speech (POS) tag a text, find the pronouns
func (client Client) PoSTagPronouns(text string) (resp TaggedSentencesResponse, err error) {
	err = client.apiCall("/nlp-v2/pos/tag/pronouns", inputText{text}, &resp)

	return
}

// Segment - Segment an input string into separate sentences
func (client Client) Segment(text string) (resp SegmentResponse, err error) {
	err = client.apiCall("/nlp-v2/segmentation/sentences", inputString{text}, &resp)

	return
}

// Words - Get the component words in an input string
func (client Client) Words(text string) (resp WordsResponse, err error) {
	err = client.apiCall("/nlp-v2/segmentation/words", inputText{text}, &resp)

	return
}

// SpellCheckWord - Find spelling correction suggestions
func (client Client) SpellCheckWord(word string) (resp SpellCheckWordResponse, err error) {
	err = client.apiCall("/nlp-v2/spellcheck/check/word", inputWord{word}, &resp)

	return
}

// SpellCheckSentence - Checks whether the sentence is spelled correctly
func (client Client) SpellCheckSentence(sentence string) (resp SpellCheckSentenceResponse, err error) {
	err = client.apiCall("/nlp-v2/spellcheck/check/sentence", inputSentence{sentence}, &resp)

	return
}

type inputString struct {
	InputString string `json:"InputString"`
}

type inputText struct {
	InputText string `json:"InputText"`
}

type inputWord struct {
	Word string `json:"Word"`
}

type inputSentence struct {
	Sentence string `json:"Sentence"`
}

type ParseTree struct {
	ParseTree string `json:"ParseTree"`
}

type SpellCheckWordResponse struct {
	Correct     bool     `json:"Correct"`
	Suggestions []string `json:"Suggestions"`
}

type SpellCheckSentenceResponse struct {
	IncorrectCount int                `json:"IncorrectCount"`
	Words          []SpellcheckedWord `json:"Words"`
}

type SpellcheckedWord struct {
	Word        Word     `json:"Word"`
	Correct     bool     `json:"Correct"`
	Suggestions []string `json:"Suggestions"`
}

type WordsResponse struct {
	Words []Word `json:"Words"`
}

type Word struct {
	Word          string `json:"Word"`
	WordIndex     int    `json:"WordIndex"`
	StartPosition int    `json:"StartPosition"`
	EndPosition   int    `json:"EndPosition"`
}

type SegmentResponse struct {
	Successful bool     `json:"Successful"`
	Sentences  []string `json:"Sentences"`
}

type EntitiesResponse struct {
	Successful bool     `json:"Successful"`
	Entities   []Entity `json:"Entities"`
}

type Entity struct {
	Type string `json:"EntityType"`
	Text string `json:"EntityText"`
}

type textToDetect struct {
	TextToDetect string `json:"textToDetect"`
}

type DetectLanguageResponse struct {
	Successful bool   `json:"Successful"`
	Code       string `json:"DetectedLanguage_ThreeLetterCode"`
	Name       string `json:"DetectedLanguage_FullName"`
}

type TaggedSentencesResponse struct {
	TaggedSentences []TaggedSentence `json:"TaggedSentences"`
}

type TaggedSentence struct {
	Words []TaggedWord `json:"Words"`
}

type TaggedWord struct {
	Word string `json:"Word"`
	Tag  string `json:"Tag"`
}
