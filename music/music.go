package music

import "bytes"

func GetMusicFilePath() string {
	// ファイルを一時ファイルに保存しているなら、このあたりに書いていただけると

	return "/path/to/file"
}

func GetMusicData() []byte {
	// byte列で音楽データを保存しているなら、このあたりに

	// 返り値の例
	return bytes.NewBuffer(nil).Bytes()
}
