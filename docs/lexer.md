# lexer.go
字句解析器。
入力された文字列をトークンに変換するのが役割。


## Lexer
字句解析器を表現する構造体。
この構造体のメンバに必要な値をセットし、parserに渡すことで字句解析を行う。

## New(input string) *Lexer
`input`の先頭から１つ文字を読んだ状態で、`Lexer`のポインタを返却する。

## (l *Lexer) readChar()
`input`を1文字読み進める。

## (l *Lexer) peekChar() byte
`input`の現在位置から次の文字を返却する。

## (l *Lexer) NextToken() token.Token
`input`から次のトークンを判定しそれを返却する。
空白や改行を読み飛ばして文字を読み進めトークンを取得する。