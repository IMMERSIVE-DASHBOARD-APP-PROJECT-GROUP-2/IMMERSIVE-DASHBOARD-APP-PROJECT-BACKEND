package handler

type MenteeResponse struct {
	Id       uint
	ClassID  uint
	Name     string
	Gender   MenteeGender
	Category MenteeCategory
	Status   MenteeStatus
}
