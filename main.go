package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"log"
	"time"
	"strings"
	"slices"
)

func readChatMetadata(chatDir string) (Conversation, error) {
	conversationData, err := os.ReadFile("./data/" + chatDir + "/conversation.json")
	if err != nil {
		return Conversation{}, err
	}
	var conversation Conversation
	json.Unmarshal(conversationData, &conversation)
	return conversation, nil
}

func readChatMessages(chatDir string) ([]Message, error) {
	messagesData, err := os.ReadFile("./data/" + chatDir + "/message.json")
	if err != nil {
		return []Message{}, err
	}
	var messages []Message
	json.Unmarshal(messagesData, &messages)
	slices.SortFunc(messages, func(a, b Message) int {
		return int(a.CreatedAt - b.CreatedAt)
	})
	return messages, nil
}

func listChats(w http.ResponseWriter, r *http.Request) {
	chats := make(map[string]Chat)
	chatDirs, err := os.ReadDir("./data")
	if err != nil {
		log.Println("Error reading chats directory: " + err.Error())
		http.Error(w, "Error reading chats directory", http.StatusInternalServerError)
		return
	}
	for _, chatDir := range chatDirs {
		if chatDir.IsDir() {
			conversation, err := readChatMetadata(chatDir.Name())
			if err != nil {
				log.Println("Error reading conversation.json: " + err.Error())
				http.Error(w, "Error reading conversation.json", http.StatusInternalServerError)
				return
			}
			chats[chatDir.Name()] = Chat{
				Metadata: conversation,
				Messages: []Message{},
			}
		}
	}
	var jsonData []byte
	jsonData, err = json.Marshal(chats)
	if err != nil {
		log.Println("Error marshalling chats: " + err.Error())
		http.Error(w, "Error marshalling chats", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, "chats.json", time.Now(), strings.NewReader(string(jsonData)))
}

func getChat(w http.ResponseWriter, r *http.Request) {
	chat := Chat{}
	chat.Metadata = Conversation{}
	chat.Messages = []Message{}

	var slug string
	if strings.HasPrefix(r.URL.Path, "/api/chat/") {
		slug = strings.TrimPrefix(r.URL.Path, "/api/chat/")
	} else {
		http.Error(w, "Invalid path", http.StatusNotFound)
		return
	}

	conversation, err := readChatMetadata(slug)
	if err != nil {
		log.Println("Error reading conversation.json: " + err.Error())
		http.Error(w, "Error reading conversation.json", http.StatusInternalServerError)
		return
	}
	chat.Metadata = conversation
	chat.Messages, err = readChatMessages(slug)
	if err != nil {
		log.Println("Error reading message.json: " + err.Error())
		http.Error(w, "Error reading message.json", http.StatusInternalServerError)
		return
	}
	var jsonData []byte
	jsonData, err = json.Marshal(chat)
	if err != nil {
		log.Println("Error marshalling chat: " + err.Error())
		http.Error(w, "Error marshalling chat", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, "chat.json", time.Now(), strings.NewReader(string(jsonData)))
}

func main() {
	http.HandleFunc("/api/chats/", listChats)
	http.HandleFunc("/api/chat/", getChat)
	fs := http.FileServer(http.Dir("./data"))
	http.Handle("/data/", http.StripPrefix("/data/", fs))
	http.Handle("/", http.FileServer(http.Dir("./frontend/build")))

	fmt.Println("Go to localhost:3000 in your browser")
	http.ListenAndServe(":3000", nil)
}