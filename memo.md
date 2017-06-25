# 言語仕様が理解できていない箇所

## レシーバーが省略されて（いるように見える）メソッドの使用
- Parsing - Integer Literals

以下のようにParserのprefixPaseFnsからprefixParse関数を取り出す。
取り出せればその関数を実行する。
```
// parser/parser.go

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}
```
呼び出し時にはレシーバーを指定していない。

呼び出される関数は以下の通り。どちらもprefixParseFnsのmapに格納された関数。
```
// parser/parser.go

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}
```
parseIdentifier, parseIntegerLiteralともに、関数内部でレシーバーのpにアクセスしている。
どうやってこんなことができるんだろう？

未解決


## parserがexpressionを処理する箇所が難しすぎる
- できれば自分で仕組みを表現できるようになりたい
- pratt parserの論文も読んでおきたい
- 左から再帰的にparseしていく処理は単純におもしろい
	- "1 + 2 + 3" は、"((1 + 2) + 3)"になる
	- "1 + 2 * 3" は、"(1 + (2 * 3))"になる
	- これは"+"と"*"の優先順位(precedence)が異なるため
	- "2"を処理する際に次の"*"が"+"より先にparseされる

