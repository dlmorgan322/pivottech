package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv/cmd/godotenv"
)

type characterClient struct {
	public     string
	private    string
	httpClient *http.Client
}

const baseURL = "https://gateway.marvel.com"

type MarvelCharacters []RosterRepsonse

func (m MarvelCharacters) String() string {
	var str string
	for _, i := range m {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	public := os.Getenv("PUB_KEY")
	private := os.Getenv("PRIV_KEY")

	server := characterClient{
		public:     public,
		private:    private,
		httpClient: &http.Client{},
	}

	characters, err := server.marvelRoster()
	if err != nil {
		log.Fatal(err)
	}

	result := MarvelCharacters(characters)
	fmt.Print(result)
}

func (c *characterClient) md5Hash(timeStamp int64) string {
	tsmd5 := strconv.Itoa(int(timeStamp))
	hash := md5.Sum([]byte(tsmd5 + c.private + c.public))
	return hex.EncodeToString(hash[:])
}

func (c *characterClient) marvelRoster() ([]RosterRepsonse, error) {
	url := (baseURL + "/v1/public/characters?limit=25")

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var roster RosterRepsonse
	if err := json.NewDecoder(res.Body).Decode(&roster); err != nil {
		return nil, err
	}

	var final25 []string

	roster.Data.Results.Name = append(Name, final25)

	return final25, nil
}

func (c *characterClient) verifyKey(url string) string {
	timeStamp := time.Now().Unix()
	hash := c.md5Hash(timeStamp)
	return fmt.Sprintf("%sts=%d&apikey=%s&hash=%s", url, timeStamp, c.public, hash)
}
