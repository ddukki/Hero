package lang

import (
	"embed"
	"encoding/xml"
	"strings"

	"github.com/ddukki/Hero/src/util"
)

const (
	RULE_GENERATIONAL string = "generational"
	RULE_GIVEN        string = "given"
	RULE_LOCALE       string = "locale"
)

//go:embed language.xml
var f embed.FS

// The central database for accessing in-world language features.
var DB *db

// db holds the languages database determining how words are created.
type db struct {
	Languages map[string]*Language
}

// Language contains the rules that defines how words are generated.
type Language struct {
	Name  string
	Rules map[string]*Rule
}

// Rule defines a specific category of words in a language, along with valid
// word parts in that category that can be summed to create a word.
type Rule struct {
	Name  string
	Parts []*Part
}

// Part defines a set of syllables that are valid for a specific part of a word.
// Words can be thought of as having multiple "parts" or slots in which
// syllables can be placed.
type Part struct {
	Syllables []string
}

func init() {

	// Create a holder for the XML data.
	langDBStruct := struct {
		XList []*struct {
			XName  string `xml:"name"`
			XRules []*struct {
				XName  string `xml:"name"`
				XParts []*struct {
					XRaw string `xml:"syllables"`
				} `xml:"part"`
			} `xml:"rule"`
		} `xml:"language"`
	}{}

	data, _ := f.ReadFile("language.xml")
	err := xml.Unmarshal(data, &langDBStruct)
	if err != nil {
		panic(err.Error())
	}

	DB = &db{}
	DB.Languages = make(map[string]*Language)

	// Parse and extract syllables from languages.
	for _, l := range langDBStruct.XList {
		lang := &Language{
			Name:  l.XName,
			Rules: make(map[string]*Rule),
		}
		DB.Languages[l.XName] = lang

		for _, r := range l.XRules {
			rule := &Rule{
				Name:  r.XName,
				Parts: make([]*Part, 0, len(r.XParts)),
			}
			lang.Rules[r.XName] = rule
			for _, p := range r.XParts {
				part := &Part{
					Syllables: strings.Split(p.XRaw, ";"),
				}
				rule.Parts = append(rule.Parts, part)
			}
		}
	}
}

// randSyllable returns a random syllable from the available syllables in
// the part.
func (p *Part) randSyllable() string {
	return p.Syllables[util.RandIdx(p.Syllables)]
}

// Generate creates a random string from the syllables defined by this rule.
func (r *Rule) Generate() string {
	word := ""
	for _, p := range r.Parts {
		word += p.randSyllable()
	}

	return word
}
