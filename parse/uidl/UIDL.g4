grammar UIDL;

options {
    language=Go;
}

@parser::header {
import "strings"
}

uidl
    : WhiteSpace? version commandSeparator commands WhiteSpace? EOF
    ;

version
    : ('version' | 'v') WhiteSpace Natural
    ;

commands
    : command (commandSeparator command)* commandSeparator?
    ;

command
    : Identifier WhiteSpace Identifier WhiteSpace? attributes WhiteSpace? commandBody?
    ;

commandSeparator
    : Semicolon
    | ws=WhiteSpace {strings.ContainsRune($ws.text, '\n')}?<fail='expected semicolon or new line'>
    ;

commandBody
    : '{' WhiteSpace? commands WhiteSpace? '}'
    ;

attributes
    : '(' WhiteSpace? (attribute (Comma attribute)* )? Comma? ')'
    ;

attribute
    : Identifier WhiteSpace? '=' WhiteSpace? value WhiteSpace?
    ;

value
    : DoubleQuotedString
    | BackQuotedString
    | Float
    | (Natural | Int)
    | Bool
    ;

Bool
    : ('true' | 'false')
    ;

DoubleQuotedString
    : '"' (EscapedChar | SafeCodepoint)* '"'
    ;

BackQuotedString
    : '`' ~[`]* '`'
    ;

fragment EscapedChar
    : '\\' (["\\bfnrt] | UnicodeChar)
    ;

fragment UnicodeChar
    : 'u' HexDigit HexDigit HexDigit HexDigit
    ;

fragment HexDigit
    : [0-9a-fA-F]
    ;

fragment SafeCodepoint
    : ~["\\\u0000-\u001F]
    ;

Identifier
    : [\p{Alpha}_] [\p{Alnum}_]*
    ;

Natural
    // integer part forbids leading 0s (e.g. `01`)
    : [1-9] [0-9]*
    ;

Float
    : Int '.' [0-9]+ Exponent?
    ;

Int
    : ('+' | '-')? ('0' | Natural)
    ;

fragment Exponent
    // exponent number permits leading 0s (e.g. `1e01`)
    : [Ee] [+-]? [0-9]+
    ;

Semicolon
    : WhiteSpace? ';' WhiteSpace?
    ;

Comma
    : WhiteSpace? ',' WhiteSpace?
    ;

WhiteSpace
    : (Space | Comment)+
    ;

fragment Space
    : [\p{White_Space}]+
    ;

fragment Comment
    : '#' ~[\n]* ('\n' | EOF)
    ;
