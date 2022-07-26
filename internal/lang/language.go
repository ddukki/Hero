package lang

import (
	"embed"
	"encoding/xml"
	"math/rand"
	"strings"
	"time"
)

const (
	RULE_GENERATIONAL string = "generational"
	RULE_GIVEN        string = "given"
	RULE_LOCALE       string = "locale"
)

//go:embed language.xml
var f embed.FS

// The central database for accessing in-world language features.
var DB LanguageDB

// LanguageDB holds the languages database determining how words are created.
type LanguageDB struct {
	List        []*Language `xml:"language"`
	LangsByName map[string]*Language
}

// Language contains the rules that defines how words are generated.
type Language struct {
	Name        string  `xml:"name"`
	Rules       []*Rule `xml:"rule"`
	RulesByName map[string]*Rule
}

// Rule defines a specific category of words in a language, along with valid
// word parts in that category that can be summed to create a word.
type Rule struct {
	Name  string  `xml:"name"`
	Parts []*Part `xml:"part"`
}

// Part defines a set of syllables that are valid for a specific part of a word.
// Words can be thought of as having multiple "parts" or slots in which
// syllables can be placed.
type Part struct {
	Raw       string `xml:"syllables"`
	Syllables []string
}

func init() {
	rand.Seed(time.Now().UnixNano())

	data, _ := f.ReadFile("language.xml")
	err := xml.Unmarshal(data, &DB)
	if err != nil {
		panic(err.Error())
	}

	DB.LangsByName = make(map[string]*Language)

	// Parse and extract syllables from languages.
	for _, l := range DB.List {
		DB.LangsByName[l.Name] = l
		l.RulesByName = make(map[string]*Rule)

		for _, r := range l.Rules {
			l.RulesByName[r.Name] = r
			for _, p := range r.Parts {
				p.Syllables = strings.Split(p.Raw, ";")
			}
		}
	}
}

// randSyllable returns a random syllable from the available syllables in
// the part.
func (p *Part) randSyllable() string {
	n := len(p.Syllables)
	return p.Syllables[rand.Intn(n)]
}

// Generate creates a random string from the syllables defined by this rule.
func (r *Rule) Generate() string {
	word := ""
	for _, p := range r.Parts {
		word += p.randSyllable()
	}

	return word
}
