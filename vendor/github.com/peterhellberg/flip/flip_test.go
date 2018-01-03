package flip_test

import (
	"fmt"
	"testing"

	"github.com/peterhellberg/flip"
)

var funcTests = []struct {
	emoticon string
	in       string
	out      string
}{
	{"foo", "bar", "fooɹɐq"},
	{flip.DefaultEmoticon, "baz", "(╯°□°）╯︵zɐq"},
	{flip.GopherEmoticon, "qux", "ʕ╯◔ϖ◔ʔ╯︵xnb"},
	{flip.AngryEmoticon, "quux", "(ノಠ益ಠ)ノ︵xnnb"},
	{flip.SparklyEmoticon, "corge", "(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧*:･ﾟ✧ ǝƃɹoɔ"},
}

func TestFunc(t *testing.T) {
	for _, tt := range funcTests {
		if got := flip.Func(tt.emoticon)(tt.in); got != tt.out {
			t.Errorf("Func(%#v)(%#v) = %#v, want %#v", tt.emoticon, tt.in, got, tt.out)
		}
	}
}

func ExampleFunc() {
	fmt.Println(flip.Func("(O_o/︵")("example"))
	// Output: (O_o/︵ǝןdɯɐxǝ
}

var flippersTests = []struct {
	name string
	in   string
	out  string
}{
	{"table", "foo", "(╯°□°）╯︵ooɟ"},
	{"gopher", "bar", "ʕ╯◔ϖ◔ʔ╯︵ɹɐq"},
	{"angry", "baz", "(ノಠ益ಠ)ノ︵zɐq"},
	{"sparkly", "qux", "(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧*:･ﾟ✧ xnb"},
}

func TestFlippers(t *testing.T) {
	for _, tt := range flippersTests {
		if got := flip.Flippers[tt.name](tt.in); got != tt.out {
			t.Errorf("Flippers[%#v](%#v) = %v, want %v", tt.name, tt.in, got, tt.out)
		}
	}
}

func ExampleFlipper() {
	fmt.Println(flip.Flippers["angry"]("example"))
	// Output: (ノಠ益ಠ)ノ︵ǝןdɯɐxǝ
}

var upsideDownTests = []struct {
	in  string
	out string
}{
	{"@", ""},
	{"ab", "qɐ"},
	{"abcde", "ǝpɔqɐ"},
	{"abcdefghijk", "ʞɾᴉɥƃɟǝpɔqɐ"},
	{"abcdefghijklmnopqrstuvwxyz", "zʎxʍʌnʇsɹbdouɯןʞɾᴉɥƃɟǝpɔqɐ"},
	{"this", "sᴉɥʇ"},
}

func TestUpsideDown(t *testing.T) {
	for _, tt := range upsideDownTests {
		if got := flip.UpsideDown(tt.in); got != tt.out {
			t.Errorf("UpsideDown(in) = %v, want %v", got, tt.out)
		}
	}
}

func ExampleUpsideDown() {
	fmt.Println(flip.UpsideDown("example"))
	// Output: ǝןdɯɐxǝ
}

var tableTests = []struct {
	out string
	in  string
}{
	{"(╯°□°）╯︵q∀", "AB"},
	{"(╯°□°）╯︵ɾǝɥ", "hej"},
	{"(╯°□°）╯︵ʇxǝʇ", "text"},
	{"(╯°□°）╯︵ƃuᴉɹʇs ɹǝƃuoן ∀", "A longer string"},
}

func TestTable(t *testing.T) {
	for _, tt := range tableTests {
		if got := flip.Table(tt.in); got != tt.out {
			t.Errorf("Table(in) = %v, want %v", got, tt.out)
		}
	}
}

func ExampleTable() {
	fmt.Println(flip.Table("example"))
	// Output: (╯°□°）╯︵ǝןdɯɐxǝ
}

var gopherTests = []struct {
	out string
	in  string
}{
	{"ʕ╯◔ϖ◔ʔ╯︵q∀", "AB"},
	{"ʕ╯◔ϖ◔ʔ╯︵ɾǝɥ", "hej"},
	{"ʕ╯◔ϖ◔ʔ╯︵ʇxǝʇ", "text"},
	{"ʕ╯◔ϖ◔ʔ╯︵ƃuᴉɹʇs ɹǝƃuoן ∀", "A longer string"},
}

func TestGopher(t *testing.T) {
	for _, tt := range gopherTests {
		if got := flip.Gopher(tt.in); got != tt.out {
			t.Errorf("Gopher(in) = %v, want %v", got, tt.out)
		}
	}
}

func ExampleGopher() {
	fmt.Println(flip.Gopher("example"))
	// Output: ʕ╯◔ϖ◔ʔ╯︵ǝןdɯɐxǝ
}

var reverseTests = []struct {
	in  string
	out string
}{
	{"let", "tel"},
	{"open", "nepo"},
	{"enough", "hguone"},
	{"side", "edis"},
	{"case", "esac"},
	{"days", "syad"},
	{"yet", "tey"},
	{"better", "retteb"},
	{"nothing", "gnihton"},
	{"tell", "llet"},
	{"problem", "melborp"},
	{"toward", "drawot"},
	{"given", "nevig"},
	{"why", "yhw"},
	{"national", "lanoitan"},
	{"room", "moor"},
	{"young", "gnuoy"},
	{"social", "laicos"},
	{"light", "thgil"},
	{"business", "ssenisub"},
	{"president", "tnediserp"},
	{"help", "pleh"},
	{"power", "rewop"},
	{"country", "yrtnuoc"},
	{"next", "txen"},
	{"things", "sgniht"},
	{"word", "drow"},
	{"looked", "dekool"},
	{"real", "laer"},
	{"John", "nhoJ"},
}

func TestReverse(t *testing.T) {
	for _, tt := range reverseTests {
		if got := flip.Reverse(tt.in); got != tt.out {
			t.Errorf("Reverse(in) = %v, want %v", got, tt.out)
		}
	}
}

func ExampleReverse() {
	fmt.Println(flip.Reverse("foo bar"))
	// Output: rab oof
}
