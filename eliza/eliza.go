//Damian Gavin
//adpted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe

package eliza

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var reflections map[string]string //map of strings of type string
//a struct called response, contains a compiled regular expression and a string array
//of answers as from regexp package
type response struct {
	rex     *regexp.Regexp
	answers []string
} //struct

func newResponse(pattern string, answers []string) response {
	response := response{}
	rex := regexp.MustCompile(pattern)
	response.rex = rex
	response.answers = answers
	return response
} //newResponse

//buildResponseList reads an array of Responses from a text file.
//It takes no arguments
func buildResponseList() []response {

	allResponses := []response{}
	//File takes data from my patterns.dat. If anything goes wrong itwill exit
	file, err := os.Open("./data/patterns.dat") //data file from static
	if err != nil {                             // an error
		panic(err) // escape
	} //if err

	// The file exists!
	defer file.Close() // this will be called AFTER this function.

	//read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		patternStr := scanner.Text()
		scanner.Scan() // move onto the next line which holds the answers
		answersAsStr := scanner.Text()

		answerList := strings.Split(answersAsStr, ";")     //In my patterns the various eliza responses to one input are seperated by ";"
		resp := newResponse("(?i)"+patternStr, answerList) //this regex will allow for any case (upper&lower) entered by the user
		allResponses = append(allResponses, resp)
	} //scanner

	//return the allResponses array
	return allResponses
} //buildResponse

func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano()) // seed to make it return different values.
	index := rand.Intn(len(answers)) // Intn generates a number between 0 and num - 1
	return answers[index]            // can be any element
}

func subWords(original string) string {
	//reflections from https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

	if reflections == nil { // map hasn't been made yet
		reflections = map[string]string{ // will only happen once.
			"am":     "are",
			"was":    "were",
			"i":      "you",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"you":    "me",
			"me":     "you",
		}
	}
	// If I get to here reflections map is populated.

	words := strings.Split(original, " ")

	for index, word := range words {
		// we want to change the word if it's in the map
		val, ok := reflections[word]
		if ok { // value WAS in the map
			// we want to swap with the value
			words[index] = val // eg. you -> me
		}
	}

	return strings.Join(words, " ")
}

func Ask(userInput string) string {

	// My name is bob
	responses := buildResponseList()
	//fmt.Println(responses)
	for _, resp := range responses { // look at every single response/pattern/answers
		//fmt.Println("User input: " + userInput)
		if resp.rex.MatchString(userInput) {
			match := resp.rex.FindStringSubmatch(userInput)
			//match[0] is full match, match[1] is the capture group
			captured := match[1]

			captured = subWords(captured)

			formatAnswer := getRandomAnswer(resp.answers) // get random element.

			if strings.Contains(formatAnswer, "%s") { // string needs to be formatted, %s will be sub target
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer

		} // if

	} // for

	// if we're down here, it means there were no matches;
	return "Sorry I was busy." // catch all.

}
