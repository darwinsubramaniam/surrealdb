// Code generated by https://github.com/abcum/tmpl
// Source file: auth.gen.go.tmpl
// DO NOT EDIT!

// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

func (s *UseStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *InfoStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *LetStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *ReturnStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *LiveStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *SelectStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *CreateStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *UpdateStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DeleteStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RelateStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineNamespaceStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveNamespaceStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineDatabaseStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveDatabaseStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineLoginStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveLoginStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineTokenStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveTokenStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineScopeStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveScopeStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineTableStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveTableStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineFieldStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveFieldStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineIndexStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveIndexStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *DefineViewStatement) Auth() (string, string) {
	return s.NS, s.DB
}

func (s *RemoveViewStatement) Auth() (string, string) {
	return s.NS, s.DB
}