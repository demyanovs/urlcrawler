package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var HTMLBodyWikiFyodorDostoevsky = "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Fyodor Dostoevsky - Wikipedia</title>\n    <meta property=\"og:title\" content=\"Fyodor Dostoevsky - Wikipedia\">\n    <meta name=\"description\" content=\"Russian novelist, short story writer, essayist and journalist\" />\n    <meta name=\"keywords\" content=\"Fyodor Dostoevsky, novelist, essayist, journalist\" />\n</head>\n<body>\n<p><b>Fyodor Mikhailovich Dostoevsky</b><sup id=\"cite_ref-1\" class=\"reference\"><a href=\"#cite_note-1\">[a]</a></sup> (<span class=\"rt-commentedText nowrap\"><style data-mw-deduplicate=\"TemplateStyles:r1177148991\">.mw-parser-output .IPA-label-small{font-size:85%}.mw-parser-output .references .IPA-label-small,.mw-parser-output .infobox .IPA-label-small,.mw-parser-output .navbox .IPA-label-small{font-size:100%}</style><span class=\"IPA-label IPA-label-small\"><a href=\"/wiki/British_English\" title=\"British English\">UK</a>: </span><span class=\"IPA nopopups noexcerpt\" lang=\"en-fonipa\"><a href=\"/wiki/Help:IPA/English\" title=\"Help:IPA/English\">/<span style=\"border-bottom:1px dotted\"><span title=\"/ˌ/: secondary stress follows\">ˌ</span><span title=\"'d' in 'dye'\">d</span><span title=\"/ɒ/: 'o' in 'body'\">ɒ</span><span title=\"'s' in 'sigh'\">s</span><span title=\"'t' in 'tie'\">t</span><span title=\"/ɔɪ/: 'oi' in 'choice'\">ɔɪ</span><span title=\"/ˈ/: primary stress follows\">ˈ</span><span title=\"/ɛ/: 'e' in 'dress'\">ɛ</span><span title=\"'f' in 'find'\">f</span><span title=\"'s' in 'sigh'\">s</span><span title=\"'k' in 'kind'\">k</span><span title=\"/i/: 'y' in 'happy'\">i</span></span>/</a></span></span>,<sup id=\"cite_ref-2\" class=\"reference\"><a href=\"#cite_note-2\">[1]</a></sup> <span class=\"rt-commentedText nowrap\"><link rel=\"mw-deduplicated-inline-style\" href=\"mw-data:TemplateStyles:r1177148991\"><span class=\"IPA-label IPA-label-small\"><a href=\"/wiki/American_English\" title=\"American English\">US</a>: </span><span class=\"IPA nopopups noexcerpt\" lang=\"en-fonipa\"><a href=\"/wiki/Help:IPA/English\" title=\"Help:IPA/English\">/<span style=\"border-bottom:1px dotted\"><span title=\"/ˌ/: secondary stress follows\">ˌ</span><span title=\"'d' in 'dye'\">d</span><span title=\"/ɒ/: 'o' in 'body'\">ɒ</span><span title=\"'s' in 'sigh'\">s</span><span title=\"'t' in 'tie'\">t</span><span title=\"/ə/: 'a' in 'about'\">ə</span><span title=\"/ˈ/: primary stress follows\">ˈ</span><span title=\"/j/: 'y' in 'yes'\">j</span><span title=\"/ɛ/: 'e' in 'dress'\">ɛ</span><span title=\"'f' in 'find'\">f</span><span title=\"'s' in 'sigh'\">s</span><span title=\"'k' in 'kind'\">k</span><span title=\"/i/: 'y' in 'happy'\">i</span></span>,<span class=\"wrap\"> </span><span style=\"border-bottom:1px dotted\"><span title=\"/ˌ/: secondary stress follows\">ˌ</span><span title=\"'d' in 'dye'\">d</span><span title=\"/ʌ/: 'u' in 'cut'\">ʌ</span><span title=\"'s' in 'sigh'\">s</span></span>-/</a></span></span>;<sup id=\"cite_ref-3\" class=\"reference\"><a href=\"#cite_note-3\">[2]</a></sup> <a href=\"/wiki/Russian_language\" title=\"Russian language\">Russian</a>: <span title=\"Russian-language text\"><span lang=\"ru\">Фёдор Михайлович Достоевский<sup id=\"cite_ref-4\" class=\"reference\"><a href=\"#cite_note-4\">[b]</a></sup></span></span>, <small><a href=\"/wiki/Romanization_of_Russian\" title=\"Romanization of Russian\">romanized</a>:</small> <span title=\"Russian-language romanization\"><i lang=\"ru-Latn\">Fyodor Mikhaylovich Dostoyevskiy</i></span>, <link rel=\"mw-deduplicated-inline-style\" href=\"mw-data:TemplateStyles:r1177148991\"><span class=\"IPA-label IPA-label-small\">IPA:</span> <span class=\"IPA nowrap\" lang=\"ru-Latn-fonipa\"><a href=\"/wiki/Help:IPA/Russian\" title=\"Help:IPA/Russian\">[ˈfʲɵdər<span class=\"wrap\"> </span>mʲɪˈxajləvʲɪdʑ<span class=\"wrap\"> </span>dəstɐˈjefskʲɪj]</a></span> <span class=\"noprint\"><span class=\"ext-phonos\"><span data-nosnippet=\"\" id=\"ooui-php-1\" class=\"ext-phonos-PhonosButton noexcerpt ext-phonos-PhonosButton-emptylabel oo-ui-widget oo-ui-widget-enabled oo-ui-buttonElement oo-ui-buttonElement-frameless oo-ui-iconElement oo-ui-buttonWidget\" data-ooui=\"{&quot;_&quot;:&quot;mw.Phonos.PhonosButton&quot;,&quot;href&quot;:&quot;\\/\\/upload.wikimedia.org\\/wikipedia\\/commons\\/transcoded\\/6\\/64\\/Ru-Dostoevsky.ogg\\/Ru-Dostoevsky.ogg.mp3&quot;,&quot;rel&quot;:[&quot;nofollow&quot;],&quot;framed&quot;:false,&quot;icon&quot;:&quot;volumeUp&quot;,&quot;data&quot;:{&quot;ipa&quot;:&quot;&quot;,&quot;text&quot;:&quot;&quot;,&quot;lang&quot;:&quot;en&quot;,&quot;wikibase&quot;:&quot;&quot;,&quot;file&quot;:&quot;Ru-Dostoevsky.ogg&quot;},&quot;classes&quot;:[&quot;ext-phonos-PhonosButton&quot;,&quot;noexcerpt&quot;,&quot;ext-phonos-PhonosButton-emptylabel&quot;]}\"><a role=\"button\" tabindex=\"0\" href=\"//upload.wikimedia.org/wikipedia/commons/transcoded/6/64/Ru-Dostoevsky.ogg/Ru-Dostoevsky.ogg.mp3\" rel=\"nofollow\" aria-label=\"Play audio\" title=\"Play audio\" class=\"oo-ui-buttonElement-button\"><span class=\"oo-ui-iconElement-icon oo-ui-icon-volumeUp\"></span><span class=\"oo-ui-labelElement-label\"></span><span class=\"oo-ui-indicatorElement-indicator oo-ui-indicatorElement-noIndicator\"></span></a></span><sup class=\"ext-phonos-attribution noexcerpt navigation-not-searchable\"><a href=\"/wiki/File:Ru-Dostoevsky.ogg\" title=\"File:Ru-Dostoevsky.ogg\">ⓘ</a></sup></span></span>; 11 November 1821&nbsp;– 9 February 1881<sup id=\"cite_ref-Morson_Britannica_5-0\" class=\"reference\"><a href=\"#cite_note-Morson_Britannica-5\">[3]</a></sup><sup id=\"cite_ref-6\" class=\"reference\"><a href=\"#cite_note-6\">[c]</a></sup>), sometimes transliterated as <b>Dostoyevsky</b>, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of <a href=\"/wiki/World_literature\" title=\"World literature\">world literature</a>, as many of his works are considered highly influential masterpieces.<sup id=\"cite_ref-FOOTNOTEScanlan2002_7-0\" class=\"reference\"><a href=\"#cite_note-FOOTNOTEScanlan2002-7\">[4]</a></sup><sup class=\"noprint Inline-Template\" style=\"white-space:nowrap;\">[<i><a href=\"/wiki/Wikipedia:Citing_sources\" title=\"Wikipedia:Citing sources\"><span title=\"This citation requires a reference to the specific page or range of pages in which the material appears. (November 2023)\">page&nbsp;needed</span></a></i>]</sup></p>\n<p>Dostoevsky's literary works explore the <a href=\"/wiki/Human_condition\" title=\"Human condition\">human condition</a> in the troubled political, social, and spiritual atmospheres of <a href=\"/wiki/Russian_Empire\" title=\"Russian Empire\">19th-century Russia</a>, and engage with a variety of philosophical and religious themes. His most acclaimed novels include <i><a href=\"/wiki/Crime_and_Punishment\" title=\"Crime and Punishment\">Crime and Punishment</a></i> (1866), <i><a href=\"/wiki/The_Idiot\" title=\"The Idiot\">The Idiot</a></i> (1869), <a href=\"/wiki/Demons_(Dostoevsky_novel)\" title=\"Demons (Dostoevsky novel)\"><i>Demons</i></a> (1872), and <i><a href=\"/wiki/The_Brothers_Karamazov\" title=\"The Brothers Karamazov\">The Brothers Karamazov</a></i> (1880). His 1864 <a href=\"/wiki/Novella\" title=\"Novella\">novella</a> <i><a href=\"/wiki/Notes_from_Underground\" title=\"Notes from Underground\">Notes from Underground</a></i> is considered to be one of the first works of <a href=\"/wiki/Existentialism\" title=\"Existentialism\">existentialist</a> literature.<sup id=\"cite_ref-8\" class=\"reference\"><a href=\"#cite_note-8\">[5]</a></sup></p>\n<p>Born in Moscow in 1821, Dostoevsky was introduced to literature at an early age through <a href=\"/wiki/Fairy_tales\" class=\"mw-redirect\" title=\"Fairy tales\">fairy tales</a> and <a href=\"/wiki/Legend\" title=\"Legend\">legends</a>, and through books by Russian and foreign authors. His mother died in 1837 when he was 15, and around the same time, he left school to enter the <a href=\"/wiki/Military_Engineering-Technical_University\" title=\"Military Engineering-Technical University\">Nikolayev Military Engineering Institute</a>. After graduating, he worked as an engineer and briefly enjoyed a lavish lifestyle, translating books to earn extra money. In the mid-1840s he wrote his first novel, <i><a href=\"/wiki/Poor_Folk\" title=\"Poor Folk\">Poor Folk</a></i>, which gained him entry into <a href=\"/wiki/Saint_Petersburg\" title=\"Saint Petersburg\">Saint Petersburg</a>'s literary circles. However, he was arrested in 1849 for belonging to a literary group, the <a href=\"/wiki/Petrashevsky_Circle\" title=\"Petrashevsky Circle\">Petrashevsky Circle</a>, that discussed banned books critical of <a href=\"/wiki/Russian_Empire\" title=\"Russian Empire\">Tsarist Russia</a>. Dostoevsky was sentenced to death but the sentence was <a href=\"/wiki/Mock_execution\" title=\"Mock execution\">commuted at the last moment</a>. He spent four years in a <a href=\"/wiki/Siberia\" title=\"Siberia\">Siberian</a> prison camp, followed by six years of compulsory military service in exile. In the following years, Dostoevsky worked as a journalist, publishing and editing several magazines of his own and later <i><a href=\"/wiki/A_Writer%27s_Diary\" title=\"A Writer's Diary\">A Writer's Diary</a></i>, a collection of his writings. He began to travel around western Europe and developed a <a href=\"/wiki/Gambling_addiction\" class=\"mw-redirect\" title=\"Gambling addiction\">gambling addiction</a>, which led to financial hardship. For a time, he had to beg for money, but he eventually became one of the most widely read and highly regarded Russian writers.</p>\n<p>Dostoevsky's <a href=\"/wiki/Fyodor_Dostoevsky_bibliography\" title=\"Fyodor Dostoevsky bibliography\">body of work</a> consists of thirteen novels, three novellas, seventeen short stories, and numerous other works. His writings were widely read both within and beyond his native Russia and influenced an equally great number of later writers including Russians such as <a href=\"/wiki/Aleksandr_Solzhenitsyn\" title=\"Aleksandr Solzhenitsyn\">Aleksandr Solzhenitsyn</a> and <a href=\"/wiki/Anton_Chekhov\" title=\"Anton Chekhov\">Anton Chekhov</a>, philosophers <a href=\"/wiki/Friedrich_Nietzsche\" title=\"Friedrich Nietzsche\">Friedrich Nietzsche</a> and <a href=\"/wiki/Jean-Paul_Sartre\" title=\"Jean-Paul Sartre\">Jean-Paul Sartre</a>, and the emergence of <a href=\"/wiki/Existentialism\" title=\"Existentialism\">Existentialism</a> and <a href=\"/wiki/Freudianism\" class=\"mw-redirect\" title=\"Freudianism\">Freudianism</a>.<sup id=\"cite_ref-Morson_Britannica_5-1\" class=\"reference\"><a href=\"#cite_note-Morson_Britannica-5\">[3]</a></sup> His books have been translated into more than 170 languages, and served as the inspiration for many films.</p>\n</body>\n</html>"

func TestParseURL_WrongStatusCodeError(t *testing.T) {
	resp := http.Response{
		StatusCode: http.StatusBadRequest,
		Status:     "Bad Request",
		Body:       io.NopCloser(strings.NewReader(HTMLBodyWikiFyodorDostoevsky)),
		Request: &http.Request{
			URL: &url.URL{
				Scheme: "https",
				Host:   "en.wikipedia.org",
				Path:   "/wiki/Fyodor_Dostoevsky",
			},
		},
	}

	parser := NewParser()
	_, _, err := parser.ParseResponse(&resp)

	require.Error(t, err)
}

func TestParseURL_Success(t *testing.T) {
	resp := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(HTMLBodyWikiFyodorDostoevsky)),
		Request: &http.Request{
			URL: &url.URL{
				Scheme: "https",
				Host:   "en.wikipedia.org",
				Path:   "/wiki/Fyodor_Dostoevsky",
			},
		},
	}

	parser := NewParser()
	pageData, linksOnPage, err := parser.ParseResponse(&resp)

	require.NoError(t, err)
	require.Equal(t, 35, len(linksOnPage))
	require.Equal(t, PageData{
		URL:        "https://en.wikipedia.org/wiki/Fyodor_Dostoevsky",
		StatusCode: 200,
		Title:      "Fyodor Dostoevsky - Wikipedia",
		Desc:       "Russian novelist, short story writer, essayist and journalist",
		Keywords:   "Fyodor Dostoevsky, novelist, essayist, journalist",
	}, pageData)
}
