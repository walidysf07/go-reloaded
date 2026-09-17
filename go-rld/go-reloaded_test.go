package main

import "testing"

func TestProcessText(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		//  README sample cases
		{
			name:  "readme sample: mixed low/up/cap with counts",
			input: "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			want:  "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:  "readme sample: hex and bin",
			input: "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			want:  "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:  "readme sample: a/an before consonant-looking untold",
			input: "There is no greater agony than bearing a untold story inside you.",
			want:  "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:  "readme sample: punctuation grouping and spacing",
			input: "Punctuation tests are ... kinda boring ,what do you think ?",
			want:  "Punctuation tests are... kinda boring, what do you think?",
		},

		//  audit test cases
		{
			name:  "audit: low/cap/up with counts and question mark",
			input: "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			want:  "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			name:  "audit: bin and hex conversions",
			input: "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			want:  "I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			name:  "audit: comma and period spacing, no cap/up/low markers",
			input: "Don not be sad ,because sad backwards is das . And das not good",
			want:  "Don not be sad, because sad backwards is das. And das not good",
		},
		{
			name:  "audit: cap count, quotes, and repeated a/an conversions",
			input: "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			want:  "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},

		// test for the a/an capitalization
		{
			name:  "a/an: capitalized word after 'A' should still convert",
			input: "There it was. A amazing rock!",
			want:  "There it was. An amazing rock!",
		},
		{
			name:  "a/an: lowercase a before capitalized vowel word",
			input: "I need a hour break and a Umbrella.",
			want:  "I need an hour break and an Umbrella.",
		},
		{
			name:  "a/an: leaves ordinary a/A alone before consonants",
			input: "A cat sat. A dog ran. A elephant walked.",
			want:  "A cat sat. A dog ran. An elephant walked.",
		},

		// Quote handling
		{
			name:  "quotes: multiple quoted spans in one text",
			input: "He said ' hi ' and then ' see you later ' okay?",
			want:  "He said 'hi' and then 'see you later' okay?",
		},
		{
			name:  "quotes: single word between quotes",
			input: "I am exactly how they describe me: ' awesome '",
			want:  "I am exactly how they describe me: 'awesome'",
		},

		// Punctuation groups
		{
			name:  "punctuation: ellipsis and combined marks stay grouped",
			input: "I was thinking ... You were right !?",
			want:  "I was thinking... You were right!?",
		},

		// malformed-input ( not panic)
		{
			name:  "marker with nothing before it is left as a literal word",
			input: "(up) hello",
			want:  "(up) hello",
		},
		{
			name:  "cap count larger than available preceding words does not panic",
			input: "hello (cap, 2)",
			want:  "Hello",
		},
		{
			name:  "malformed hex operand is left untouched, no panic",
			input: "zz (hex) test",
			want:  "zz (hex) test",
		},
		{
			name:  "trailing newline in input does not leak into output",
			input: "Hello world.\n",
			want:  "Hello world.",
		},
		{
			name:  "empty input produces empty output",
			input: "",
			want:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := processText(tc.input)
			if got != tc.want {
				t.Errorf("processText(%q)\n  got:  %q\n  want: %q", tc.input, got, tc.want)
			}
		})
	}
}
