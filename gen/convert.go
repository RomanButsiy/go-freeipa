// Copyright © 2022 IN2P3 Computing Centre, IN2P3, CNRS
// Copyright © 2018 Philippe Voinov
//
// Contributor(s): Remi Ferrand <remi.ferrand_at_cc.in2p3.fr>, 2021
//
// This software is governed by the CeCILL license under French law and
// abiding by the rules of distribution of free software.  You can  use,
// modify and/ or redistribute the software under the terms of the CeCILL
// license as circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and  rights to copy,
// modify and redistribute granted by the license, users are provided only
// with a limited warranty  and the software's author,  the holder of the
// economic rights,  and the successive licensors  have only  limited
// liability.
//
// In this respect, the user's attention is drawn to the risks associated
// with loading,  using,  modifying and/or developing or reproducing the
// software by the user in light of its specific status of free software,
// that may mean  that it is complicated to manipulate,  and  that  also
// therefore means  that it is reserved for developers  and  experienced
// professionals having in-depth computer knowledge. Users are therefore
// encouraged to load and test the software's suitability as regards their
// requirements in conditions enabling the security of their systems and/or
// data to be ensured and,  more generally, to use and operate it in the
// same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

package main

import (
	"strings"

	"github.com/RomanButsiy/go-freeipa/thirdparty/snaker"
)

var reservedWords = []string{
	"continue",
	"break",
	"return",
	"type",
	"map",
}

func safeName(s string) string {
	for _, r := range reservedWords {
		if r == s {
			return "_" + s
		}
	}
	return s
}

func toGoType(ipaType string) string {
	switch ipaType {
	case "":
		return "interface{}"
	case "dict":
		return "interface{}"
	case "object":
		return "interface{}"
	case "NoneType":
		return "interface{}"
	case "unicode":
		return "string"
	case "str":
		return "string"
	case "bytes":
		return "string"
	case "datetime":
		return "time.Time"
	case "DN":
		return "string"
	case "Principal":
		return "string"
	case "DNSName":
		return "string"
	case "Decimal":
		return "float64"
	case "Certificate":
		return "interface{}"
	case "CertificateSigningRequest":
		return "string"
	default:
		return ipaType
	}
}

func (t *CommandOutput) GoType(parent *Command) string {
	if t.Type == "dict_guess_receiver" {
		cls := strings.Split(parent.ObjClass, "/")[0]
		if cls != "" {
			return upperName(cls)
		}
	}
	return toGoType(t.Type)
}

func lowerName(s string) string {
	return safeName(snaker.SnakeToCamelLower(s))
}

func upperName(s string) string {
	return safeName(snaker.SnakeToCamel(s))
}

func (t *Topic) LowerName() string {
	return lowerName(t.Name)
}

func (t *Topic) UpperName() string {
	return upperName(t.Name)
}

func (t *Class) LowerName() string {
	return lowerName(t.Name)
}

func (t *Class) UpperName() string {
	return upperName(t.Name)
}

func (t *Command) LowerName() string {
	return lowerName(t.Name)
}

func (t *Command) UpperName() string {
	return upperName(t.Name)
}

func (t *Param) LowerName() string {
	return lowerName(t.Name)
}

func (t *Param) UpperName() string {
	return upperName(t.Name)
}

func (t *CommandOutput) LowerName() string {
	return lowerName(t.Name)
}

func (t *CommandOutput) UpperName() string {
	return upperName(t.Name)
}
