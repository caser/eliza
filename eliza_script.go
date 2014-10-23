package eliza

var Pre = map[string]string{
	"dont":       "don't",
	"cant":       "can't",
	"wont":       "won't",
	"recollect":  "remember",
	"recall":     "remember",
	"dreamt":     "dreamed",
	"dreams":     "dream",
	"maybe":      "perhaps",
	"certainly":  "yes",
	"machine":    "computer",
	"machines":   "computer",
	"computers":  "computer",
	"were":       "was",
	"you're":     "you are",
	"i'm":        "i am",
	"same":       "alike",
	"identical":  "alike",
	"equivalent": "alike",
}

var Post = map[string]string{
	"am":       "are",
	"your":     "my",
	"me":       "you",
	"myself":   "yourself",
	"yourself": "myself",
	"i":        "you",
	"you":      "I",
	"my":       "your",
	"i'm":      "you are",
}
