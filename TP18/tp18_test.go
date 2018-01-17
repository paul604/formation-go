package main

import "testing"

func TestNoAnagram(t *testing.T) {
        if areAnagrams("this is not", "an anagram") {
                t.Error("this is not, an anagram")
        }
}

func TestAnagram(t *testing.T) {
        if !areAnagrams("mary", "army") {
                t.Error("mary", "army")
        }
        if !areAnagrams("silent", "listen") {
                t.Error("silent", "listen")
        }
        if !areAnagrams("Madam Curie", "Radium came") {
                t.Error("Madam Curie", "Radium came")
        }
        if !areAnagrams("william Shakespeare", "I am a weakish speller") {
                t.Error("william Shakespeare", "I am a weakish speller")
        }
}