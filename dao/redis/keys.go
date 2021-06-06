package redis

/*
	Redis Key
*/

const (
	KeyPostInfoHashPrefix = "blockchainguide:post:"
	KeyPostTimeZSet       = "blockchainguide:post:time"
	KeyPostScoreZSet      = "blockchainguide:post:score"
	//KeyPostVotedUpSetPrefix   = "bluebell:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "bluebell:post:voted:up:"
	KeyPostVotedZSetPrefix = "blockchainguide:post:voted:"

	KeyCommunityPostSetPrefix = "blockchainguide:community:"
)
