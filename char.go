package decexpr

// === Массив для быстрой классификации символов ===
type CharType int

const (
	CharInvalid CharType = iota + 1
	CharWhitespace
	CharDigit
	CharLetter
	CharDot
	CharOperator
	CharLeftParen
	CharRightParen
	CharComma
	CharEOF
)

// Массив размером 256 для быстрого определения типа символа
var charMap [256]CharType

// Инициализация массива классификации символов
func init() {
	// Белые пробелы
	charMap[0] = CharEOF
	charMap[' '] = CharWhitespace
	charMap['\t'] = CharWhitespace
	charMap['\n'] = CharWhitespace
	charMap['\r'] = CharWhitespace

	// Цифры
	for ch := '0'; ch <= '9'; ch++ {
		charMap[ch] = CharDigit
	}

	// Буквы (латиница)
	for ch := 'a'; ch <= 'z'; ch++ {
		charMap[ch] = CharLetter
	}

	for ch := 'A'; ch <= 'Z'; ch++ {
		charMap[ch] = CharLetter
	}

	charMap['_'] = CharLetter // Подчеркивание тоже часть идентификатора

	// Специальные символы
	charMap['.'] = CharDot
	charMap['('] = CharLeftParen
	charMap[')'] = CharRightParen
	charMap[','] = CharComma

	// Операторы
	charMap['+'] = CharOperator
	charMap['-'] = CharOperator
	charMap['*'] = CharOperator
	charMap['/'] = CharOperator
	charMap['%'] = CharOperator
	charMap['^'] = CharOperator
}

// === Быстрые проверки через массив ===
func isWhitespace(ch byte) bool {
	return charMap[ch] == CharWhitespace
}

func isDigit(ch byte) bool {
	return charMap[ch] == CharDigit
}

func isLetter(ch byte) bool {
	return charMap[ch] == CharLetter
}

func isDot(ch byte) bool {
	return charMap[ch] == CharDot
}

func isOperatorChar(ch byte) bool {
	return charMap[ch] == CharOperator
}

func isLeftParen(ch byte) bool {
	return charMap[ch] == CharLeftParen
}

func isRightParen(ch byte) bool {
	return charMap[ch] == CharRightParen
}

func isComma(ch byte) bool {
	return charMap[ch] == CharComma
}

func isValidChar(ch byte) bool {
	return charMap[ch] != CharInvalid
}
