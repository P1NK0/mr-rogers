package main

import (
	"fmt"
	"math/rand"
	"time"
)


var quotes = []string{
	// Original Quotes
	"Often when you think you're at the end of something, you're at the beginning of something else.",
	"Anyone who does anything to help a child in his life is a hero to me.",
	"All of us, at some time or other, need help.",
	"Mutual caring relationships require kindness and patience, tolerance, optimism, joy in the other’s achievements, confidence in oneself, and the ability to give without undue thought of gain.",
	"There are three ways to ultimate success: The first way is to be kind. The second way is to be kind. The third way is to be kind.",
	"Love isn’t a state of perfect caring. It is an active noun like struggle.",
	"Forgiveness is a strange thing. It can sometimes be easier to forgive our enemies than our friends. It can be hardest of all to forgive people we love. Like all of life’s important coping skills, the ability to forgive and the capacity to let go of resentments most likely take root very early in our lives.",
	"Love and trust, in the space between what’s said and what’s heard in our life, can make all the difference in the world.",
	"Discovering the truth about ourselves is a lifetime’s work, but it’s worth the effort.",
	"Knowing that we can be loved exactly as we are is what gives us all the best opportunity for growing into the healthiest of people.",
	"I don’t think anyone can grow unless he’s loved exactly as he is now, appreciated for what he is rather than what he will be.",
	"Feeling good about ourselves is essential in our being able to love others.",
	"Try your best to make goodness attractive. That’s one of the toughest assignments you’ll ever be given.",
	"The connections we make in the course of a life – maybe that’s what heaven is.",

	// New Additions
	"When I was a boy and I would see scary things in the news, my mother would say to me, 'Look for the helpers. You will always find people who are helping.'",
	"It's not so much what we have in this life that matters. It's what we do with what we have.",
	"The greatest gift you ever give is your honest self.",
	"There is no normal life that is free of pain. It's the very wrestling with our problems that can be the impetus for our growth.",
	"We live in a world in which we need to share responsibility. It's easy to say 'It's not my child, not my community, not my world, not my problem.' Then there are those who see the need and respond. I consider those people my heroes.",
	"Play is often talked about as if it were a relief from serious learning. But for children play is serious learning.",
	"You rarely have time for everything you want in this life, so you need to make choices. And hopefully your choices can come from a deep sense of who you are.",
	"The real issue in life is not how many blessings we have, but what we do with our blessings. Some people have many blessings and hoard them. Some have few and share everything they have.",
	"Real strength has to do with helping others.",
	"Part of the problem with the word 'disabilities' is that it immediately suggests an inability to see or hear or walk or do other things that many of us take for granted. But what of people who can't feel? Or talk about their feelings? Or manage their feelings in constructive ways? What of people who aren't able to form loving, trusting relationships? And people who get stuck in anger or hate or revenge? Those are disabilities, too.",
	"Imagining something may be the first step in making it happen, but it takes the real time and real efforts of real people to learn things, make things, turn thoughts into deeds or visions into inventions.",
	"We need to do our best to work together, to be good neighbors, to be good friends.",
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	// Pick a random quote
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Println("Mr. Rogers says:")
	fmt.Printf("\"%s\"\n", quote)
}
