package instagram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

// GetInstagramMessage Gets the message to send to
func GetInstagramMessage() string {
	const instagramUser = "theartidote"
	const instagramURL = "https://www.instagram.com/" + instagramUser + "/?__a=1"

	instagramResponse := getInstagramResponse(instagramURL)
	instagramTextPosts := formatInstagramResponse(instagramResponse)

	if len(instagramTextPosts) == 0 {
		log.Fatalln("No Posts Retrieved. Is the user's recent timeline empty?")
	}
	itr := determineItrToSend(len(instagramTextPosts), instagramTextPosts[0].Timestamp)
	text := formatTextPost(instagramTextPosts, itr)
	return text
}

// get raw instagram response from source
func getInstagramResponse(instagramURL string) Response {
	client := http.Client{}

	request, err := http.NewRequest("GET", instagramURL, bytes.NewBuffer(nil))

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		instagramResponse := *new(Response)
		json.NewDecoder(resp.Body).Decode(&instagramResponse)
		return instagramResponse
	} else {
		fmt.Println("resp.StatusCode")
		fmt.Println(resp.StatusCode)
		instagramResponse := *new(Response)
		return instagramResponse
	}
}

// Reduces the complexity of the raw instagram response to a simpler format
func formatInstagramResponse(instagramResponse Response) []TextPosts {

	textPosts := []TextPosts{}
	for _, post := range instagramResponse.GraphQLResponse.User.RecentTimeline.Edges {
		if post.InstagramPost.Typename == "GraphImage" {
			// Assign instagram text to post
			var postText string
			if len(post.InstagramPost.EdgeMediaToCaption.Edges) >= 1 {
				postText = post.InstagramPost.EdgeMediaToCaption.Edges[0].Node.Text
			}

			// Get post thumbnail resource
			postThumbnailResources := getThumbnailResources(post.InstagramPost)

			textPosts = append(textPosts, TextPosts{
				Text:               postText,
				ThumbnailResources: postThumbnailResources,
				Timestamp:          time.Unix(int64(post.InstagramPost.Timestamp), 0),
			})
		}
		// log.Println(string(i) + " " + post.InstagramPost.EdgeMediaToCaption.Edges[0].Node.Text)
	}
	return textPosts
}

// reassigns thumbnail resources from instagram post to simpler format
func getThumbnailResources(post UnformattedInstagramPost) []ThumbnailResource {
	thumbnailResources := make([]ThumbnailResource, len(post.ThumbnailResources))

	for i, resource := range post.ThumbnailResources {
		thumbnailResources[i] = ThumbnailResource{
			Source:       resource.Source,
			ConfigHeight: resource.ConfigHeight,
			ConfigWidth:  resource.ConfigWidth,
		}
	}

	return thumbnailResources
}

// simple algorithm to determine which the iterate number of which item to send based on days elapsed
func determineItrToSend(lengthOfArray int, timeOfLastPost time.Time) int {
	today := time.Now()
	daysElapsed := int(math.Floor(today.Sub(timeOfLastPost).Hours() / 24))

	itrInArrayToSend := daysElapsed % lengthOfArray
	return itrInArrayToSend
}

func formatTextPost(instagramTextPosts []TextPosts, itr int) string {
	var imageURL string
	var text string
	post := instagramTextPosts[itr]
	seperator := "\nart"

	splitText := strings.Split(post.Text, seperator)
	if len(post.ThumbnailResources) > 1 {
		// 320 x 320 picture link
		imageURL = post.ThumbnailResources[2].Source
	}

	if len(splitText) >= 2 {
		text = splitText[0] + "\n\n" + imageURL + "\n\nart"

		for i := 1; i < len(splitText); i++ {
			text += splitText[i]
		}
	} else {
		text = splitText[0]
	}

	if itr != 0 {
		// This is to show that there has been repeats due to lack of daily posts
		text = "**" + text
	}
	return text
}

// TextPosts Instagram Posts shape
type TextPosts struct {
	Text               string
	ThumbnailResources []ThumbnailResource
	Timestamp          time.Time
}

type ThumbnailResource struct {
	Source       string
	ConfigWidth  int
	ConfigHeight int
}

// Response Instagram response shape
type Response struct {
	GraphQLResponse struct {
		User struct {
			RecentTimeline struct {
				Edges []struct {
					InstagramPost UnformattedInstagramPost `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
		} `json:"user"`
	} `json:"graphql"`
}

// UnformattedInstagramPost Instagram post response shape
type UnformattedInstagramPost struct {
	Typename           string `json:"__typename"`
	ID                 string `json:"id"`
	ThumbnailResources []struct {
		Source       string `json:"src"`
		ConfigWidth  int    `json:"config_width"`
		ConfigHeight int    `json:"config_height"`
	} `json:"thumbnail_resources"`
	EdgeMediaToCaption struct {
		Edges []struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`
	Timestamp int `json:"taken_at_timestamp"`
}
