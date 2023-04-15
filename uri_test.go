// Copyright 2018 Jacob Dufault
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURINix(t *testing.T) {
	assert.Equal(t, LsDocumentURI("file:///a/b/c"), pathToURI("/a/b/c"))
	assert.Equal(t, LsDocumentURI("file:///a/b/c/"), pathToURI("/a/b/c/"))
	assert.Equal(t, LsDocumentURI("file:///a/b/c/"), pathToURI("\\a/b\\c\\"))
	assert.Equal(t, LsDocumentURI("file:///"), pathToURI("/"))
}

func TestURISlashConversion(t *testing.T) {
	assert.Equal(t, LsDocumentURI("file:///a%20b"), pathToURI("/a b"))
}

// TODO: add windows-style path tests

// これらは、pathToURI関数の正しさを検証するためにGoで書かれたユニットテストです。TestURINix関数は、さまざまな入力パスで関数をテストし、関数がUnixライクなファイルシステムに対して期待されるURIを生成することを確認するためのものです。TestURISlashConversion関数は、スペースを含む入力パスで関数をテストし、結果のURIでスペースが"%20 "に適切に変換されるかどうかを確認します。

// テストの定義にはtestingパッケージを使用し、関数呼び出しの結果についてのアサーションにはgithub.com/stretchr/testify/assertパッケージを使用しています。各テストケースでは、assert.Equal関数が2つの引数で呼び出されます：期待される結果と、特定の入力でpathToURIを呼び出した場合の実際の結果です。2つの値が等しくない場合、テストは失敗し、エラーメッセージが出力されます。すべてのテストに合格した場合、関数は何も出力せずに返します。
