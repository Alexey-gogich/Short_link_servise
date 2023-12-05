package entity

type Link struct {
	Link string
}

type ShortLink struct {
	Link string
}

type LinkCreate struct {
	Link      string
	ShortLink string
}
