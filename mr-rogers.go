package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Add your favorite Mr. Rogers quotes here
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
	"If you could only sense how important you are to the lives of those you meet; how important you can be to the people you may never even dream of. There is something of yourself that you leave at every meeting with another person.",
	"In times of stress, the best thing we can do for each other is to listen with our ears and our hearts and to be assured that our questions are just as important as our answers.",
	"I hope you're proud of yourself for the times you've said \"yes,\" when all it meant was extra work for you and was seemingly helpful only to someone else.",
	"Whether we're a preschooler or a young teen, a graduating college senior or a retired person, we human beings all want to know that we're acceptable, that our being alive somehow makes a difference in the lives of others.",
	"We need to help people to discover the true meaning of love. Love is generally confused with dependence. Those of us who have grown in true love know that we can love only in proportion to our capacity for independence.",
	"Who we are in the present includes who we were in the past.",
	"It's very dramatic when two people come together to work something out. It's easy to take a gun and annihilate your opposition, but what is really exciting to me is to see people with differing views come together and finally respect each other.",
	"In the external scheme of things, shining moments are as brief as the twinkling of an eye, yet such twinklings are what eternity is made of -- moments when we human beings can say \"I love you,\" \"I'm proud of you,\" \"I forgive you,\" \"I'm grateful for you.\" That's what eternity is made of: invisible imperishable good stuff.",
	"The thing I remember best about successful people I've met all through the years is their obvious delight in what they're doing and it seems to have very little to do with worldly success. They just love what they're doing, and they love it in front of others.",
	"At the center of the Universe is a loving heart that continues to beat and that wants the best for every person.",
	"Anything that we can do to help foster the intellect and spirit and emotional growth of our fellow human beings, that is our job.",
	"Those of us who have this particular vision must continue against all odds.",
	"Life is for service.",
	"The kingdom of God is for the broken hearted",
	"Little by little we human beings are confronted with situations that give us more and more clues that we are not perfect.",
	"What's been important in my understanding of myself and others is the fact that each one of us is so much more than any one thing. A sick child is much more than his or her sickness. A person with a disability is much, much more than a handicap. A pediatrician is more than a medical doctor. You're MUCH more than your job description or your age or your income or your output.",
	"When we love a person, we accept him or her exactly as is: the lovely with the unlovely, the strong with the fearful, the true mixed in with the façade, and of course, the only way we can do it is by accepting ourselves that way.",
	"Confronting our feelings and giving them appropriate expression always takes strength, not weakness. It takes strength to acknowledge our anger, and sometimes more strength yet to curb the aggressive urges anger may bring and to channel them into nonviolent outlets. It takes strength to face our sadness and to grieve and to let our grief and our anger flow in tears when they need to. It takes strength to talk about our feelings and to reach out for help and comfort when we need it.",
	"It's not the honors and the prizes and the fancy outsides of life which ultimately nourish our souls. It's the knowing that we can be trusted, that we never have to fear the truth, that the bedrock of our very being is good stuff.",
	"I believe that appreciation is a holy thing--that when we look for what's best in a person we happen to be with at the moment, we're doing what God does all the time. So in loving and appreciating our neighbor, we're participating in something sacred.",
	"Love is like infinity: You can't have more or less infinity, and you can't compare two things to see if they're \"equally infinite.\" Infinity just is, and that's the way I think love is, too.",
	"I'm proud of you for the times you came in second, or third, or fourth, but what you did was the best you have ever done",
	"Peace means far more than the opposite of war.",
	"There's no \"should\" or \"should not\" when it comes to having feelings. They're part of who we are and their origins are beyond our control. When we can believe that, we may find it easier to make constructive choices about what to do with those feelings.",
	"The only thing evil can’t stand is forgiveness.",
	"Solitude is different from loneliness, and it doesn't have to be a lonely kind of thing.",
	"Some days, doing \"the best we can\" may still fall short of what we would like to be able to do, but life isn't perfect on any front-and doing what we can with what we have is the most we should expect of ourselves or anyone else.",
	"The world needs a sense of worth, and it will achieve it only by its people feeling that they are worthwhile.",
	"Everyone longs to be loved. And the greatest thing we can do is to let people know that they are loved and capable of loving.",
	"It's a mistake to think that we have to be lovely to be loved by human beings or by God",
	"It's really easy to fall into the trap of believing that what we do is more important than what we are. Of course, it's the opposite that's true: What we are ultimately determines what we do!",
	"Whatever we choose to imagine can be as private as we want it to be. Nobody knows what you're thinking or feeling unless you share it.",
	"You are a very special person. There is only one like you in the whole world. There's never been anyone exactly like you before, and there will never be again. Only you. And people can like you exactly as you are.",
	"Life is deep and simple, and what our society gives us is shallow and complicated.",
	"Often out of periods of losing come the greatest strivings toward a new winning streak.",
	"I like to compare the holiday season with the way a child listens to a favorite story. The pleasure is in the familiar way the story begins, the anticipation of familiar turns it takes, the familiar moments of suspense, and the familiar climax and ending.",
	"We speak with more than our mouths. We listen with more than our ears.",
	"Our society is much more interested in information than wonder, in noise rather than silence...And I feel that we need a lot more wonder and a lot more silence in our lives",
	"It's good to be curious about many things.",
	"How many times have you noticed that it’s the little quiet moments in the midst of life that seem to give the rest extra-special meaning?",
	"Being able to resolve conflicts peacefully is one of the greatest strengths we can give our children.",
	"I realize that it isn't very fashionable to talk about some things being holy; nevertheless, if we ever want to rid ourselves of personal and corporate emptiness, brokenness, loneliness, and fear, we have to allow ourselves room for that which we can not see, hear, touch , or control.",
	"There’s a part of all of us that longs to know that even what’s weakest about us is still redeemable and can ultimately count for something good.",
	"There are times when explanations, no matter how reasonable, just don't seem to help.",
	"We all have different gifts, so we all have different ways of saying to the world who we are.",
	"Listening is where love begins: listening to ourselves and then to our neighbors.",
	"Love is at the root of everything. All learning, all parenting, all relationships. Love or the lack of it.",
	"I think that those who would try to make you feel less than who you are, I think that's the greatest evil.",
	"The media shows the tiniest percentage of what people do. There are millions and millions of people doing wonderful things all over the world, and they’re generally not the ones being touted in the news.",
	"You can think about things and make believe. All you have to do is think and they'll grow.",
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	// Pick a random quote
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Println("Mr. Rogers says:")
	fmt.Printf("\"%s\"\n", quote)
}
