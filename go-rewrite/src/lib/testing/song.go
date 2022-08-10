package testlib

import (
	"embed"
	"encoding/json"
	. "github.com/onsi/gomega"
)

// compare most of the fields, except last saved at
func ExpectJSONEqualExceptLastSavedAt(a map[string]interface{}, b map[string]interface{}) {
	a["lastSavedAt"] = nil
	b["lastSavedAt"] = nil
	ExpectWithOffset(1, a).To(Equal(b))
}

//go:embed demo_song_test.json
var demoSongFS embed.FS

func LoadDemoSong() map[string]interface{} {
	file := ExpectSuccess(demoSongFS.Open("demo_song_test.json"))

	output := map[string]interface{}{}
	err := json.NewDecoder(file).Decode(&output)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	output["id"] = ""
	output["owner"] = PrimaryUser.ID

	return output
}
