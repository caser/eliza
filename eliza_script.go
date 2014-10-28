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

// this could be better as a map of synonym to keyword instead of keyword to list
// of synonyms, but might be harder to interpret / customize
// thus processing happens in Synonymize() method
var SynonymMap = map[string][]string{
	"be":       []string{"am", "is", "are", "was"},
	"belief":   []string{"feel", "think", "believe", "wish"},
	"cannot":   []string{"can't"},
	"desire":   []string{"want", "need"},
	"everyone": []string{"everybody", "nobody", "noone"},
	"family":   []string{"mother", "mom", "father", "dad", "sister", "brother", "wife", "children", "child"},
	"happy":    []string{"elated", "glad", "better"},
	"sad":      []string{"unhappy", "depressed", "sick"},
}

type Key struct {
	Keyword string
	Rank    int
	Decomps []Decomp
}

type Decomp struct {
	DecompRule string
	AssembRule []string
}

var Keywords = map[string]Key{
	"xnone": Key{
		Keyword: "xnone",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"I'm not sure I understand you fully.",
					"Please go on.",
					"What does that suggest to you ?",
					"Do you feel strongly about discussing such things ?",
					"That is interesting.  Please continue.",
					"Tell me more about that.",
					"Does talking about this bother you ?",
				},
			},
		},
	},
	"sorry": Key{
		Keyword: "sorry",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Please don't apologise.",
					"Apologies are not necessary.",
					"I've told you that apologies are not required.",
					"It did not bother me.  Please continue.",
				},
			},
		},
	},
	"apologise": Key{
		Keyword: "apologise",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto sorry",
				},
			},
		},
	},
	"remember": Key{
		Keyword: "remember",
		Rank:    5,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* i remember *",
				AssembRule: []string{
					"Do you often think of (2) ?",
					"Does thinking of (2) bring anything else to mind ?",
					"What else do you recollect ?",
					"Why do you remember (2) just now ?",
					"What in the present situation reminds you of (2) ?",
					"What is the connection between me and (2) ?",
					"What else does (2) remind you of ?",
				},
			},
			Decomp{
				DecompRule: "* do you remember *",
				AssembRule: []string{
					"Did you think I would forget (2) ?",
					"Why do you think I should recall (2) now ?",
					"What about (2) ?",
					"goto what",
					"You mentioned (2) ?",
				},
			},
			Decomp{
				DecompRule: "* you remember *",
				AssembRule: []string{
					"How could I forget (2) ?",
					"What about (2) should I remember ?",
					"goto you",
				},
			},
		},
	},
	"forget": Key{
		Keyword: "forget",
		Rank:    5,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* i forget *",
				AssembRule: []string{
					"Can you think of why you might forget (2) ?",
					"Why can't you remember (2) ?",
					"How often do you think of (2) ?",
					"Does it bother you to forget that ?",
					"Could it be a mental block ?",
					"Are you generally forgetful ?",
					"Do you think you are suppressing (2) ?",
				},
			},
			Decomp{
				DecompRule: "* did you forget *",
				AssembRule: []string{
					"Why do you ask ?",
					"Are you sure you told me ?",
					"Would it bother you if I forgot (2) ?",
					"Why should I recall (2) just now ?",
					"goto what",
					"Tell me more about (2).",
				},
			},
		},
	},
	"if": Key{
		Keyword: "if",
		Rank:    3,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* if *",
				AssembRule: []string{
					"Do you think it's likely that (2) ?",
					"Do you wish that (2) ?",
					"What do you know about (2) ?",
					"Really, if (2) ?",
					"What would you do if (2) ?",
					"But what are the chances that (2) ?",
					"What does this speculation lead to ?",
				},
			},
		},
	},
	"dreamed": Key{
		Keyword: "dreamed",
		Rank:    4,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* i dreamed *",
				AssembRule: []string{
					"Really, (2) ?",
					"Have you ever fantasized (2) while you were awake ?",
					"Have you ever dreamed (2) before ?",
					"goto dream",
				},
			},
		},
	},
	"dream": Key{
		Keyword: "dream",
		Rank:    3,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"What does that dream suggest to you ?",
					"Do you dream often ?",
					"What persons appear in your dreams ?",
					"Do you believe that dreams have something to do with your problem ?",
				},
			},
		},
	},
	"perhaps": Key{
		Keyword: "perhaps",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"You don't seem quite certain.",
					"Why the uncertain tone ?",
					"Can't you be more positive ?",
					"You aren't sure ?",
					"Don't you know ?",
					"How likely, would you estimate ?",
				},
			},
		},
	},
	"name": Key{
		Keyword: "name",
		Rank:    15,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"I am not interested in names.",
					"I've told you before, I don't care about names -- please continue.",
				},
			},
		},
	},
	"deutsch": Key{
		Keyword: "deutsch",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto xforeign",
					"I told you before, I don't understand German.",
				},
			},
		},
	},
	"francais": Key{
		Keyword: "francais",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto xforeign",
					"I told you before, I don't understand French.",
				},
			},
		},
	},
	"italiano": Key{
		Keyword: "italiano",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto xforeign",
					"I told you before, I don't understand Italian.",
				},
			},
		},
	},
	"espanol": Key{
		Keyword: "espanol",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto xforeign",
					"I told you before, I don't understand Spanish.",
				},
			},
		},
	},
	"xforeign": Key{
		Keyword: "xforeign",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"I speak only English.",
				},
			},
		},
	},
	"hello": Key{
		Keyword: "hello",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"How do you do.  Please state your problem.",
					"Hi.  What seems to be your problem ?",
				},
			},
		},
	},
	"computer": Key{
		Keyword: "computer",
		Rank:    50,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Do computers worry you ?",
					"Why do you mention computers ?",
					"What do you think machines have to do with your problem ?",
					"Don't you think computers can help people ?",
					"What about machines worries you ?",
					"What do you think about machines ?",
					"You don't think I am a computer program, do you ?",
				},
			},
		},
	},
	"am": Key{
		Keyword: "am",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* am i *",
				AssembRule: []string{
					"Do you believe you are (2) ?",
					"Would you want to be (2) ?",
					"Do you wish I would tell you you are (2) ?",
					"What would it mean if you were (2) ?",
					"goto what",
				},
			},
			Decomp{
				DecompRule: "* i am *",
				AssembRule: []string{
					"goto i",
				},
			},
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Why do you say 'am' ?",
					"I don't understand that.",
				},
			},
		},
	},
	"are": Key{
		Keyword: "are",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* are you *",
				AssembRule: []string{
					"Why are you interested in whether I am (2) or not ?",
					"Would you prefer if I weren't (2) ?",
					"Perhaps I am (2) in your fantasies.",
					"Do you sometimes think I am (2) ?",
					"goto what",
					"Would it matter to you ?",
					"What if I were (2) ?",
				},
			},
			Decomp{
				DecompRule: "* you are *",
				AssembRule: []string{
					"goto you",
				},
			},
			Decomp{
				DecompRule: "* are *",
				AssembRule: []string{
					"Did you think they might not be (2) ?",
					"Would you like it if they were not (2) ?",
					"What if they were not (2) ?",
					"Are they always (2) ?",
					"Possibly they are (2).",
					"Are you positive they are (2) ?",
				},
			},
		},
	},
	"your": Key{
		Keyword: "your",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* your *",
				AssembRule: []string{
					"Why are you concerned over my (2) ?",
					"What about your own (2) ?",
					"Are you worried about someone else's (2) ?",
					"Really, my (2) ?",
					"What makes you think of my (2) ?",
					"Do you want my (2) ?",
				},
			},
		},
	},
	"was": Key{
		Keyword: "was",
		Rank:    2,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* was i *",
				AssembRule: []string{
					"What if you were (2) ?",
					"Do you think you were (2) ?",
					"Were you (2) ?",
					"What would it mean if you were (2) ?",
					"What does ' (2) ' suggest to you ?",
					"goto what",
				},
			},
			Decomp{
				DecompRule: "* i was *",
				AssembRule: []string{
					"Were you really ?",
					"Why do you tell me you were (2) now ?",
					"Perhaps I already know you were (2).",
				},
			},
			Decomp{
				DecompRule: "* was you *",
				AssembRule: []string{
					"Would you like to believe I was (2) ?",
					"What suggests that I was (2) ?",
					"What do you think ?",
					"Perhaps I was (2).",
					"What if I had been (2) ?",
				},
			},
		},
	},
	"i": Key{
		Keyword: "i",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* i @desire *",
				AssembRule: []string{
					"What would it mean to you if you got (3) ?",
					"Why do you want (3) ?",
					"Suppose you got (3) soon.",
					"What if you never got (3) ?",
					"What would getting (3) mean to you ?",
					"What does wanting (3) have to do with this discussion ?",
				},
			},
			Decomp{
				DecompRule: "* i am* @sad *",
				AssembRule: []string{
					"I am sorry to hear that you are (3).",
					"Do you think coming here will help you not to be (3) ?",
					"I'm sure it's not pleasant to be (3).",
					"Can you explain what made you (3) ?",
				},
			},
			Decomp{
				DecompRule: "* i am* @happy *",
				AssembRule: []string{
					"How have I helped you to be (3) ?",
					"Has your treatment made you (3) ?",
					"What makes you (3) just now ?",
					"Can you explain why you are suddenly (3) ?",
				},
			},
			Decomp{
				DecompRule: "* i was *",
				AssembRule: []string{
					"goto was",
				},
			},
			Decomp{
				DecompRule: "* i @belief i *",
				AssembRule: []string{
					"Do you really think so ?",
					"But you are not sure you (3).",
					"Do you really doubt you (3) ?",
				},
			},
			Decomp{
				DecompRule: "* i* @belief *you *",
				AssembRule: []string{
					"goto you",
				},
			},
			Decomp{
				DecompRule: "* i am *",
				AssembRule: []string{
					"Is it because you are (2) that you came to me ?",
					"How long have you been (2) ?",
					"Do you believe it is normal to be (2) ?",
					"Do you enjoy being (2) ?",
					"Do you know anyone else who is (2) ?",
				},
			},
			Decomp{
				DecompRule: "* i @cannot *",
				AssembRule: []string{
					"How do you know that you can't (3) ?",
					"Have you tried ?",
					"Perhaps you could (3) now.",
					"Do you really want to be able to (3) ?",
					"What if you could (3) ?",
				},
			},
			Decomp{
				DecompRule: "* i don't *",
				AssembRule: []string{
					"Don't you really (2) ?",
					"Why don't you (2) ?",
					"Do you wish to be able to (2) ?",
					"Does that trouble you ?",
				},
			},
			Decomp{
				DecompRule: "* i feel *",
				AssembRule: []string{
					"Tell me more about such feelings.",
					"Do you often feel (2) ?",
					"Do you enjoy feeling (2) ?",
					"Of what does feeling (2) remind you ?",
				},
			},
			Decomp{
				DecompRule: "* i * you *",
				AssembRule: []string{
					"Perhaps in your fantasies we (2) each other.",
					"Do you wish to (2) me ?",
					"You seem to need to (2) me.",
					"Do you (2) anyone else ?",
				},
			},
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"You say (1) ?",
					"Can you elaborate on that ?",
					"Do you say (1) for some special reason ?",
					"That's quite interesting.",
				},
			},
		},
	},
	"you": Key{
		Keyword: "you",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* you remind me of *",
				AssembRule: []string{
					"goto alike",
				},
			},
			Decomp{
				DecompRule: "* you are *",
				AssembRule: []string{
					"What makes you think I am (2) ?",
					"Does it please you to believe I am (2) ?",
					"Do you sometimes wish you were (2) ?",
					"Perhaps you would like to be (2).",
				},
			},
			Decomp{
				DecompRule: "* you* me *",
				AssembRule: []string{
					"Why do you think I (2) you ?",
					"You like to think I (2) you -- don't you ?",
					"What makes you think I (2) you ?",
					"Really, I (2) you ?",
					"Do you wish to believe I (2) you ?",
					"Suppose I did (2) you -- what would that mean ?",
					"Does someone else believe I (2) you ?",
				},
			},
			Decomp{
				DecompRule: "* you *",
				AssembRule: []string{
					"We were discussing you -- not me.",
					"Oh, I (2) ?",
					"You're not really talking about me -- are you ?",
					"What are your feelings now ?",
				},
			},
		},
	},
	"yes": Key{
		Keyword: "yes",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"You seem to be quite positive.",
					"You are sure.",
					"I see.",
					"I understand.",
				},
			},
		},
	},
	"no": Key{
		Keyword: "no",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* no one *",
				AssembRule: []string{
					"Are you sure, no one (2) ?",
					"Surely someone (2) .",
					"Can you think of anyone at all ?",
					"Are you thinking of a very special person ?",
					"Who, may I ask ?",
					"You have a particular person in mind, don't you ?",
					"Who do you think you are talking about ?",
				},
			},
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Are you saying no just to be negative?",
					"You are being a bit negative.",
					"Why not ?",
					"Why 'no' ?",
				},
			},
		},
	},
	"my": Key{
		Keyword: "my",
		Rank:    2,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "$ * my *",
				AssembRule: []string{
					"Does that have anything to do with the fact that your (2) ?",
					"Lets discuss further why your (2).",
					"Earlier you said your (2).",
					"But your (2).",
				},
			},
			Decomp{
				DecompRule: "* my* @family *",
				AssembRule: []string{
					"Tell me more about your family.",
					"Who else in your family (4) ?",
					"Your (3) ?",
					"What else comes to your mind when you think of your (3) ?",
				},
			},
			Decomp{
				DecompRule: "* my *",
				AssembRule: []string{
					"Your (2) ?",
					"Why do you say your (2) ?",
					"Does that suggest anything else which belongs to you ?",
					"Is it important to you that your (2) ?",
				},
			},
		},
	},
	"can": Key{
		Keyword: "can",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* can you *",
				AssembRule: []string{
					"You believe I can (2) don't you ?",
					"goto what",
					"You want me to be able to (2).",
					"Perhaps you would like to be able to (2) yourself.",
				},
			},
			Decomp{
				DecompRule: "* can i *",
				AssembRule: []string{
					"Whether or not you can (2) depends on you more than on me.",
					"Do you want to be able to (2) ?",
					"Perhaps you don't want to (2).",
					"goto what",
				},
			},
		},
	},
	"what": Key{
		Keyword: "what",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Why do you ask ?",
					"Does that question interest you ?",
					"What is it you really want to know ?",
					"Are such questions much on your mind ?",
					"What answer would please you most ?",
					"What do you think ?",
					"What comes to mind when you ask that ?",
					"Have you asked such questions before ?",
					"Have you asked anyone else ?",
				},
			},
		},
	},
	"who": Key{
		Keyword: "who",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "who *",
				AssembRule: []string{
					"goto what",
				},
			},
		},
	},
	"when": Key{
		Keyword: "when",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "when *",
				AssembRule: []string{
					"goto what",
				},
			},
		},
	},
	"where": Key{
		Keyword: "where",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "where *",
				AssembRule: []string{
					"goto what",
				},
			},
		},
	},
	"how": Key{
		Keyword: "how",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "how *",
				AssembRule: []string{
					"goto what",
				},
			},
		},
	},
	"because": Key{
		Keyword: "because",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Is that the real reason ?",
					"Don't any other reasons come to mind ?",
					"Does that reason seem to explain anything else ?",
					"What other reasons might there be ?",
				},
			},
		},
	},
	"why": Key{
		Keyword: "why",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* why don't you *",
				AssembRule: []string{
					"Do you believe I don't (2) ?",
					"Perhaps I will (2) in good time.",
					"Should you (2) yourself ?",
					"You want me to (2) ?",
					"goto what",
				},
			},
			Decomp{
				DecompRule: "* why can't i *",
				AssembRule: []string{
					"Do you think you should be able to (2) ?",
					"Do you want to be able to (2) ?",
					"Do you believe this will help you to (2) ?",
					"Have you any idea why you can't (2) ?",
					"goto what",
				},
			},
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto what",
				},
			},
		},
	},
	"everyone": Key{
		Keyword: "everyone",
		Rank:    2,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* @everyone *",
				AssembRule: []string{
					"Really, (2) ?",
					"Surely not (2).",
					"Can you think of anyone in particular ?",
					"Who, for example?",
					"Are you thinking of a very special person ?",
					"Who, may I ask ?",
					"Someone special perhaps ?",
					"You have a particular person in mind, don't you ?",
					"Who do you think you're talking about ?",
				},
			},
		},
	},
	"everybody": Key{
		Keyword: "everybody",
		Rank:    2,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto everyone",
				},
			},
		},
	},
	"nobody": Key{
		Keyword: "nobody",
		Rank:    2,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto everyone",
				},
			},
		},
	},
	"noone": Key{
		Keyword: "noone",
		Rank:    2,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"goto everyone",
				},
			},
		},
	},
	"always": Key{
		Keyword: "always",
		Rank:    1,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"Can you think of a specific example ?",
					"When ?",
					"What incident are you thinking of ?",
					"Really, always ?",
				},
			},
		},
	},
	"alike": Key{
		Keyword: "alike",
		Rank:    10,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"In what way ?",
					"What resemblence do you see ?",
					"What does that similarity suggest to you ?",
					"What other connections do you see ?",
					"What do you suppose that resemblence means ?",
					"What is the connection, do you suppose ?",
					"Could there really be some connection ?",
					"How ?",
				},
			},
		},
	},
	"like": Key{
		Keyword: "like",
		Rank:    10,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "* @be *like *",
				AssembRule: []string{
					"goto alike",
				},
			},
		},
	},
	"different": Key{
		Keyword: "different",
		Rank:    0,
		Decomps: []Decomp{
			Decomp{
				DecompRule: "*",
				AssembRule: []string{
					"How is it different ?",
					"What differences do you see ?",
					"What does that difference suggest to you ?",
					"What other distinctions do you see ?",
					"What do you suppose that disparity means ?",
					"Could there be some connection, do you suppose ?",
					"How ?",
				},
			},
		},
	},
}
