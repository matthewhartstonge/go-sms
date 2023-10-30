package gsm7

var (
	EnglishBasic = map[rune]struct{}{
		'@':  keyExists,
		'£':  keyExists,
		'$':  keyExists,
		'¥':  keyExists,
		'è':  keyExists,
		'é':  keyExists,
		'ù':  keyExists,
		'ì':  keyExists,
		'ò':  keyExists,
		'Ç':  keyExists,
		'\n': keyExists, // LF: Line Feed control.
		'Ø':  keyExists,
		'ø':  keyExists,
		'\r': keyExists, // CR: Carriage Return
		'Å':  keyExists,
		'å':  keyExists,

		'Δ': keyExists,
		'_': keyExists,
		'Φ': keyExists,
		'Γ': keyExists,
		'Λ': keyExists,
		'Ω': keyExists,
		'Π': keyExists,
		'Ψ': keyExists,
		'Σ': keyExists,
		'Θ': keyExists,
		'Ξ': keyExists,
		' ': keyExists, // ESC: Control Character. maps to non-breaking space.
		'Æ': keyExists,
		'æ': keyExists,
		'ß': keyExists,
		'É': keyExists,

		' ':  keyExists,
		'!':  keyExists,
		'"':  keyExists,
		'#':  keyExists,
		'¤':  keyExists,
		'%':  keyExists,
		'&':  keyExists,
		'\'': keyExists,
		'(':  keyExists,
		')':  keyExists,
		'*':  keyExists,
		'+':  keyExists,
		',':  keyExists,
		'-':  keyExists,
		'.':  keyExists,
		'/':  keyExists,

		'0': keyExists,
		'1': keyExists,
		'2': keyExists,
		'3': keyExists,
		'4': keyExists,
		'5': keyExists,
		'6': keyExists,
		'7': keyExists,
		'8': keyExists,
		'9': keyExists,
		':': keyExists,
		';': keyExists,
		'<': keyExists,
		'=': keyExists,
		'>': keyExists,
		'?': keyExists,

		'¡': keyExists,
		'A': keyExists,
		'B': keyExists,
		'C': keyExists,
		'D': keyExists,
		'E': keyExists,
		'F': keyExists,
		'G': keyExists,
		'H': keyExists,
		'I': keyExists,
		'J': keyExists,
		'K': keyExists,
		'L': keyExists,
		'M': keyExists,
		'N': keyExists,
		'O': keyExists,

		'P': keyExists,
		'Q': keyExists,
		'R': keyExists,
		'S': keyExists,
		'T': keyExists,
		'U': keyExists,
		'V': keyExists,
		'W': keyExists,
		'X': keyExists,
		'Y': keyExists,
		'Z': keyExists,
		'Ä': keyExists,
		'Ö': keyExists,
		'Ñ': keyExists,
		'Ü': keyExists,
		'§': keyExists,

		'¿': keyExists,
		'a': keyExists,
		'b': keyExists,
		'c': keyExists,
		'd': keyExists,
		'e': keyExists,
		'f': keyExists,
		'g': keyExists,
		'h': keyExists,
		'i': keyExists,
		'j': keyExists,
		'k': keyExists,
		'l': keyExists,
		'm': keyExists,
		'n': keyExists,
		'o': keyExists,

		'p': keyExists,
		'q': keyExists,
		'r': keyExists,
		's': keyExists,
		't': keyExists,
		'u': keyExists,
		'v': keyExists,
		'w': keyExists,
		'x': keyExists,
		'y': keyExists,
		'z': keyExists,
		'ä': keyExists,
		'ö': keyExists,
		'ñ': keyExists,
		'ü': keyExists,
		'à': keyExists,
	}
	EnglishExtended = map[rune]struct{}{
		'|':      keyExists,
		'^':      keyExists,
		'€':      keyExists,
		'{':      keyExists,
		'}':      keyExists,
		'\f':     keyExists, // FF: is a Page Break control.
		'\u008E': keyExists, // SS2: is a second Single Shift Escape control reserved for future extensions.
		'[':      keyExists,
		// CR2: is a control character. No language specific character shall be encoded at this position.
		'~':      keyExists,
		']':      keyExists,
		'\u005C': keyExists, // Backslash \
	}
	EnglishReplacements = map[rune]string{
		'\u0000': ` `,   // Null (NUL) �
		'\u0003': ` `,   // End of Text (ETX)
		'\u0004': ` `,   // End of Transmission (EOT)
		'\u0009': ` `,   // Character Tabulation (HT, TAB)
		'\u0010': ` `,   // Data Link Escape (DLE)
		'\u0011': ` `,   // Device Control One (DC1)
		'\u0012': ` `,   // Device Control Two (DC2)
		'\u0013': ` `,   // Device Control Three (DC3)
		'\u0014': ` `,   // Device Control Four (DC4)
		'\u0017': ` `,   // End of Transmission Block (ETB)
		'\u0019': ` `,   // End of Medium (EOM)
		'\u0060': `'`,   // `
		'\u0080': ` `,   // Padding Character (PAD)
		'\u008D': ` `,   // Reverse Index (RI) 
		'\u0090': ` `,   // Device Control String (DCS) 
		'\u009B': ` `,   // Control Sequence Introducer (CSI)
		'\u009F': ` `,   // Application Program Command (APC)
		'\u00A0': ` `,   // No-Break Space (NBSP)
		'\u00AB': `"`,   // «
		'\u00B4': `'`,   // ´
		'\u00BB': `"`,   // »
		'\u00BC': "1/4", // ¼
		'\u00BD': "1/2", // ½
		'\u00BE': "3/4", // ¾
		'\u00F7': `/`,   // ÷
		'\u0141': `L`,   // Ł
		'\u0142': `l`,   // ł
		'\u01C3': `!`,   // ǃ
		'\u0262': `G`,   // ɢ
		'\u026A': `I`,   // ɪ
		'\u0274': `N`,   // ɴ
		'\u0280': `R`,   // ʀ
		'\u028F': `Y`,   // ʏ
		'\u0299': `B`,   // ʙ
		'\u029C': `H`,   // ʜ
		'\u029F': `L`,   // ʟ
		'\u02B9': `'`,   // ʹ
		'\u02BA': `"`,   // ʺ
		'\u02BB': `'`,   // ʻ
		'\u02BC': `'`,   // ʼ
		'\u02BD': `'`,   // ʽ
		'\u02C6': `^`,   // ˆ
		'\u02C8': `'`,   // ˈ
		'\u02CA': `'`,   // ˊ
		'\u02CB': `'`,   // ˋ
		'\u02D0': `:`,   // ː
		'\u02D6': `+`,   // ˖
		'\u02DC': `~`,   // ˜
		'\u02EE': `"`,   // ˮ
		'\u02F7': `~`,   // ˷
		'\u02F8': `:`,   // ˸
		'\u0302': `^`,   // ̂
		'\u0303': `~`,   // ̃
		'\u0313': `'`,   // ̓
		'\u0314': `\`,   // ⃥
		'\u0326': `,`,   // ̦
		'\u0330': `~`,   // ̰
		'\u0332': `_`,   // ̲
		'\u0334': `~`,   // ̴
		'\u0337': `/`,   // ̷
		'\u0338': `/`,   // ̸
		'\u0347': `=`,   // ͇
		'\u1D00': `A`,   // ᴀ
		'\u1D04': `C`,   // ᴄ
		'\u1D05': `D`,   // ᴅ
		'\u1D07': `E`,   // ᴇ
		'\u1D0A': `J`,   // ᴊ
		'\u1D0B': `K`,   // ᴋ
		'\u1D0D': `M`,   // ᴍ
		'\u1D0F': `O`,   // ᴏ
		'\u1D18': `P`,   // ᴘ
		'\u1D1B': `T`,   // ᴛ
		'\u1D1C': `U`,   // ᴜ
		'\u1D20': `V`,   // ᴠ
		'\u1D21': `W`,   // ᴡ
		'\u1D22': `Z`,   // ᴢ
		'\u1DCD': `^`,   // ᷍
		'\u2000': ` `,   // En Quad
		'\u2001': ` `,   // Em Quad
		'\u2002': ` `,   // En Space
		'\u2003': ` `,   // Em Space
		'\u2004': ` `,   // Three-Per-Em Space
		'\u2005': ` `,   // Four-Per-Em Space
		'\u2006': ` `,   // Six-Per-Em Space
		'\u2007': ` `,   // Figure Space
		'\u2008': ` `,   // Punctuation Space
		'\u2009': ` `,   // Thin Space
		'\u200A': ` `,   // Hair Space
		'\u200B': ` `,   // Zero Width Space (ZWSP) ​
		'\u2010': `-`,   // ‐
		'\u2013': `-`,   // –
		'\u2014': `-`,   // —
		'\u2015': `-`,   // ―
		'\u2017': `_`,   // ‗
		'\u2018': `'`,   // ‘
		'\u2019': `'`,   // ’
		'\u201A': `'`,   // ‚
		'\u201B': `'`,   // ‛
		'\u201C': `"`,   // “
		'\u201D': `"`,   // ”
		'\u201E': `"`,   // „
		'\u201F': `"`,   // ‟
		'\u2022': `-`,   // •
		'\u2026': `...`, // …
		'\u2028': ` `,   // Line Separator
		'\u2029': ` `,   // Paragraph Separator
		'\u202F': ` `,   // Narrow No-Break Space (NNBSP)
		'\u2039': `<`,   // ‹
		'\u203A': `>`,   // ›
		'\u203C': `!!`,  // ‼
		'\u2043': `-`,   // ⁃
		'\u2044': `/`,   // ⁄
		'\u204E': `*`,   // ⁎
		'\u204F': `\`,   // ⧵
		'\u205F': ` `,   // Medium Mathematical Space (MMSP)
		'\u2060': ` `,   // Word Joiner (WJ) ⁠
		'\u20D2': `|`,   // ⃒
		'\u20D3': `|`,   // ⃓
		'\u20E5': `'`,   // ̔
		'\u2215': `/`,   // ∕
		'\u2217': `*`,   // ∗
		'\u2223': `|`,   // ∣
		'\u223C': `~`,   // ∼
		'\u229B': `*`,   // ⊛
		'\u239C': `|`,   // ⎜
		'\u239F': `|`,   // ⎟
		'\u23B8': `|`,   // ⎸
		'\u23B9': `|`,   // ⎹
		'\u23BC': `-`,   // ⎼
		'\u23BD': `-`,   // ⎽
		'\u23D0': `|`,   // ⏐
		'\u2722': `*`,   // ✢
		'\u2723': `*`,   // ✣
		'\u2724': `*`,   // ✤
		'\u2725': `*`,   // ✥
		'\u2731': `*`,   // ✱
		'\u2732': `*`,   // ✲
		'\u2733': `*`,   // ✳
		'\u273A': `*`,   // ✺
		'\u273B': `*`,   // ✻
		'\u273C': `*`,   // ✼
		'\u273D': `*`,   // ✽
		'\u2743': `*`,   // ❃
		'\u2749': `*`,   // ❉
		'\u274A': `*`,   // ❊
		'\u274B': `*`,   // ❋
		'\u275B': `'`,   // ❛
		'\u275C': `'`,   // ❜
		'\u275D': `"`,   // ❝
		'\u275E': `"`,   // ❞
		'\u2768': `(`,   // ❨
		'\u2769': `)`,   // ❩
		'\u276A': `(`,   // ❪
		'\u276B': `)`,   // ❫
		'\u2774': `{`,   // ❴
		'\u2775': `}`,   // ❵
		'\u27EE': `(`,   // ⟮
		'\u27EF': `)`,   // ⟯
		'\u2982': `:`,   // ⦂
		'\u2985': `(`,   // ⦅
		'\u2986': `)`,   // ⦆
		'\u29C6': `*`,   // ⧆
		'\u29F5': `;`,   // ⁏
		'\u29F8': `/`,   // ⧸
		'\u29F9': `\`,   // ⧹
		'\u3000': ` `,   // Ideographic Space
		'\u3001': `,`,   // 、
		'\u3002': `.`,   // 。
		'\u301D': `"`,   // 〝
		'\u301E': `"`,   // 〞
		'\uA730': `F`,   // ꜰ
		'\uA731': `S`,   // ꜱ
		'\uA789': `:`,   // ꞉
		'\uA78A': `=`,   // ꞊
		'\uFE10': `'`,   // ︐
		'\uFE11': `'`,   // ︑
		'\uFE13': `:`,   // ︓
		'\uFE14': `;`,   // ︔
		'\uFE15': `!`,   // ︕
		'\uFE16': `?`,   // ︖
		'\uFE50': `,`,   // ﹐
		'\uFE51': `,`,   // ﹑
		'\uFE52': `.`,   // ﹒
		'\uFE54': `;`,   // ﹔
		'\uFE56': `?`,   // ﹖
		'\uFE57': `!`,   // ﹗
		'\uFE59': `(`,   // ﹙
		'\uFE5A': `)`,   // ﹚
		'\uFE5B': `{`,   // ﹛
		'\uFE5C': `}`,   // ﹜
		'\uFE5F': `#`,   // ﹟
		'\uFE60': `&`,   // ﹠
		'\uFE61': `*`,   // ﹡
		'\uFE62': `+`,   // ﹢
		'\uFE63': `-`,   // ﹣
		'\uFE64': `<`,   // ﹤
		'\uFE65': `>`,   // ﹥
		'\uFE66': `=`,   // ﹦
		'\uFE68': `\`,   // ﹨
		'\uFE69': `$`,   // ﹩
		'\uFE6A': `%`,   // ﹪
		'\uFE6B': `@`,   // ﹫
		'\uFEFF': ` `,   // Zero Width No-Break Space (BOM, ZWNBSP)
		'\uFF01': `!`,   // ！
		'\uFF02': `"`,   // ＂
		'\uFF03': `#`,   // ＃
		'\uFF04': `$`,   // ＄
		'\uFF05': `%`,   // ％
		'\uFF06': `&`,   // ＆
		'\uFF07': `'`,   // ＇
		'\uFF08': `(`,   // （
		'\uFF09': `)`,   // ）
		'\uFF0A': `*`,   // ＊
		'\uFF0B': `+`,   // ＋
		'\uFF0C': `,`,   // ，
		'\uFF0D': `-`,   // －
		'\uFF0E': `.`,   // ．
		'\uFF0F': `/`,   // ／
		'\uFF10': `0`,   // ０
		'\uFF11': `1`,   // １
		'\uFF12': `2`,   // ２
		'\uFF13': `3`,   // ３
		'\uFF14': `4`,   // ４
		'\uFF15': `5`,   // ５
		'\uFF16': `6`,   // ６
		'\uFF17': `7`,   // ７
		'\uFF18': `8`,   // ８
		'\uFF19': `9`,   // ９
		'\uFF1A': `:`,   // ：
		'\uFF1B': `;`,   // ；
		'\uFF1C': `<`,   // ＜
		'\uFF1D': `=`,   // ＝
		'\uFF1E': `>`,   // ＞
		'\uFF1F': `?`,   // ？
		'\uFF20': `@`,   // ＠
		'\uFF21': `A`,   // Ａ
		'\uFF22': `B`,   // Ｂ
		'\uFF23': `C`,   // Ｃ
		'\uFF24': `D`,   // Ｄ
		'\uFF25': `E`,   // Ｅ
		'\uFF26': `F`,   // Ｆ
		'\uFF27': `G`,   // Ｇ
		'\uFF28': `H`,   // Ｈ
		'\uFF29': `I`,   // Ｉ
		'\uFF2A': `J`,   // Ｊ
		'\uFF2B': `K`,   // Ｋ
		'\uFF2C': `L`,   // Ｌ
		'\uFF2D': `M`,   // Ｍ
		'\uFF2E': `N`,   // Ｎ
		'\uFF2F': `O`,   // Ｏ
		'\uFF30': `P`,   // Ｐ
		'\uFF31': `Q`,   // Ｑ
		'\uFF32': `R`,   // Ｒ
		'\uFF33': `S`,   // Ｓ
		'\uFF34': `T`,   // Ｔ
		'\uFF35': `U`,   // Ｕ
		'\uFF36': `V`,   // Ｖ
		'\uFF37': `W`,   // Ｗ
		'\uFF38': `X`,   // Ｘ
		'\uFF39': `Y`,   // Ｙ
		'\uFF3A': `Z`,   // Ｚ
		'\uFF3B': `[`,   // ［
		'\uFF3C': `\`,   // ＼
		'\uFF3D': `]`,   // ］
		'\uFF3E': `^`,   // ＾
		'\uFF3F': `_`,   // ＿
		'\uFF5B': `{`,   // ｛
		'\uFF5C': `|`,   // ｜
		'\uFF5D': `}`,   // ｝
		'\uFF5E': `~`,   // ～
		'\uFF61': `.`,   // ｡
		'\uFF64': `,`,   // ､
	}
)
