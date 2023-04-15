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

import "strings"

func pathToURI(absolutePath string) LsDocumentURI {
	m := map[rune](string){
		' ': "%20",
		'#': "%23",
		'$': "%24",
		'&': "%26",
		'(': "%28",
		')': "%29",
		'+': "%2B",
		',': "%2C",
		':': "%3A",
		';': "%3B",
		'?': "%3F",
		'@': "%40",
	}

	result := ""
	for _, c := range absolutePath {
		if _, has := m[c]; has {
			result += m[c]
		} else {
			result += string(c)
		}
	}

	return LsDocumentURI("file://" + strings.Replace(result, "\\", "/", -1))
}


// ファイルパスの絶対値をURI（Uniform Resource Identifier）形式に変換するGoプログラミング言語コードです。この関数は、入力として絶対パスを受け取り、LsDocumentURIタイプのURIを返します。

// この関数は、まず、特殊文字とそれに対応するURLエンコード表現を対応付けるマップを定義する。次に、入力パスの文字をループして、各文字が特殊文字であるかどうかをチェックします。特殊文字であれば、その文字をURLエンコードされた表現に置き換えます。そうでない場合は、その文字をそのまま結果文字列に追加します。

// 最後に、結果の文字列の前に「file://」スキームを追加し、バックスラッシュをフォワードスラッシュに置き換えて、有効なURI形式とする。

// 全体として、この関数は、絶対ファイルパスをWebコンテキストで使用可能なURIに変換するのに便利です。