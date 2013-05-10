// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"github.com/fiam/gounidecode/unidecode"
	"regexp"
	"strings"
)

var (
	// Custom substitution map
	CustomSub map[string]string
	// Custom rune substitution map
	CustomRuneSub map[rune]string
)

//=============================================================================

// Make returns slug generated from provided string. Will use "en" as language
// substitution.
func Make(s string) (slug string) {
	return MakeLang(s, "en")
}

// MakeLang returns slug generated from provided string and will use provided
// language for chars substitution.
func MakeLang(s string, lang string) (slug string) {
	slug = strings.TrimSpace(s)

	// Select substitution language
	switch lang {
	case "de":
		slug = SubstituteRune(slug, deSub)
	case "en":
		slug = SubstituteRune(slug, enSub)
	case "pl":
		slug = SubstituteRune(slug, plSub)
	default: // fallback to "en" if lang not found
		slug = SubstituteRune(slug, enSub)
	}

	slug = SubstituteRune(slug, defaultSub)

	// Custom substitutions
	slug = SubstituteRune(slug, CustomRuneSub)
	slug = Substitute(slug, CustomSub)

	slug = unidecode.Unidecode(slug)

	slug = strings.ToLower(slug)

	slug = regexp.MustCompile("[^a-z0-9-_]").ReplaceAllString(slug, "-")
	slug = regexp.MustCompile("-+").ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}

// Substitute returns string with superseded all substrings from
// provided substitution map.
func Substitute(s string, sub map[string]string) (buf string) {
	buf = s
	for key, val := range sub {
		buf = strings.Replace(s, key, val, -1)
	}
	return
}

// SubstituteRune substitutes string chars with provided rune
// substitution map.
func SubstituteRune(s string, sub map[rune]string) (buf string) {
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf += d
		} else {
			buf += string(c)
		}
	}
	return
}
