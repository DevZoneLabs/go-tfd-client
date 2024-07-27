package tfd

type LanguageCode string

var LanguageCodes = struct {
	Korean             LanguageCode
	English            LanguageCode
	Danish             LanguageCode
	French             LanguageCode
	Japanese           LanguageCode
	ChineseSimplified  LanguageCode
	ChineseTraditional LanguageCode
	Italian            LanguageCode
	Polish             LanguageCode
	Portuguese         LanguageCode
	Russian            LanguageCode
	Spanish            LanguageCode
}{
	"ko", "en", "de", "fr", "ja", "zh-CN", "zh-TW", "it", "pl", "pt", "ru", "es",
}
