package tfd

type LanguageCode string

var LanguageCodes struct {
	Korean            LanguageCode
	English           LanguageCode
	Danish            LanguageCode
	French            LanguageCode
	Japanese          LanguageCode
	ChineseSimplified LanguageCode
	Italian           LanguageCode
	Polish            LanguageCode
	Portuguese        LanguageCode
	Russian           LanguageCode
	Spanish           LanguageCode
}

func (l *LanguageCode) String() string {
	return string(*l)
}
