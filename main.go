package main

import (
    "fmt"
    "github.com/atotto/clipboard"
    "os"
    "regexp"
    "strings"
)

var asciiToMorseMap map[string]string
var morseToAsciiMap map[string]string

func init() {
    asciiToMorseMap = map[string]string {
        "A": ".-",
        "B": "-...",
        "C": "-.-.",
        "D": "-..",
        "E": ".",
        "F": "..-.",
        "G": "--.",
        "H": "....",
        "I": "..",
        "J": ".---",
        "K": "-.-",
        "L": ".-..",
        "M": "--",
        "N": "-.",
        "O": "---",
        "P": ".--.",
        "Q": "--.-",
        "R": ".-.",
        "S": "...",
        "T": "-",
        "U": "..-",
        "V": "...-",
        "W": ".--",
        "X": "-..-",
        "Y": "-.--",
        "Z": "--..",
        " ": "/",
        "0": "-----",
        "1": ".----",
        "2": "..---",
        "3": "...--",
        "4": "....-",
        "5": ".....",
        "6": "-....",
        "7": "--...",
        "8": "---..",
        "9": "----.",
        ".": ".-.-.-",
        "," : "--..--",
        ":": "---...",
        "?": "..--..",
        "'": ".----.",
        "-": "-....-",
        "/": "-..-.",
        "(": "-.--.-",
        "@": ".--.-.",
        "=": "-...-",
        "\"": ".-..-.",
    }
    morseToAsciiMap = reversedMap(asciiToMorseMap)
}

func reversedMap(original_map map[string]string) map[string]string {
    reversed_map := make(map[string]string)
    for key, value := range original_map {
        reversed_map[value] = key
    }
    return reversed_map
}

func convertMorseToAscii (input string) string {
    splitted := strings.Split(input, " ")
    var result string = ""
    for i:=0; i< len(splitted); i++ {
        if val, ok := morseToAsciiMap[splitted[i]]; ok {
            result = result + string(val)
        } else {
            var errorString string = "  [ERROR] Couldn't match \"" + string(splitted[i]) + "\" to ASCII"
            return errorString
        }
    }
    return result
}

func convertAsciiToMorse (input string) string {
    var result string = ""
    for i := 0; i < len(input); i++ {
        if val, ok := asciiToMorseMap[strings.ToUpper(string(input[i]))]; ok {
            result = result + string(val) + " "
        } else {
            var errorString string = "  [ERROR] Couldn't match \"" + string(input[i]) + "\" to Morse"
            return errorString
        }
    }
    return result
}

func main() {
    if len(os.Args) == 2 {
        var inputString string = os.Args[1]
        var converted string
        if matched, _ := regexp.MatchString("^[\\/.\\-\\s]*$", inputString); matched {
            fmt.Println("Morse string detected ... ")
            fmt.Println("The corresponding ASCII string is below\n")
            converted = convertMorseToAscii(inputString)
        } else {
            fmt.Println("Ascii string detected ... ")
            fmt.Println("The corresponding Morse string is below\n")
            converted = convertAsciiToMorse(inputString)
        }
        fmt.Println(converted)
        clipboard.WriteAll(converted)
        fmt.Println("Converted string copied to your clipboard")
    } else {
        fmt.Println("Usage: mors [ascii or morse string here]")
    }
}
