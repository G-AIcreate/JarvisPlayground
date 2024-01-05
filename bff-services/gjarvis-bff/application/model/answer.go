package model

type TextAnswer string
type AudioAnswer []byte

type JarvisResponse struct {
	TextAnswer  TextAnswer
	AudioAnswer AudioAnswer
}
