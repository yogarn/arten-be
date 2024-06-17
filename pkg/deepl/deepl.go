package deepl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const deeplAPI_URL = "https://api-free.deepl.com/v2/translate"

type Translation struct {
	Text string `json:"text"`
}

type TranslationRequest struct {
	Text       []string `json:"text"`
	SourceLang string   `json:"source_lang"`
	TargetLang string   `json:"target_lang"`
}

type TranslationResponse struct {
	Translations []Translation `json:"translations"`
}

func Translate(text, sourceLang, targetLang string) (*TranslationResponse, error) {
	reqBody := TranslationRequest{
		Text:       []string{text},
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("DEEPL_API_KEY")
	req, err := http.NewRequest("POST", deeplAPI_URL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

	req.Header.Set("Access-Control-Allow-Origin", deeplAPI_URL)
	req.Header.Set("Access-Control-Allow-Methods", "POST")
	req.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("failed to translate text: %s", resp.Status)
		return nil, errors.New(errMsg)
	}

	var translationResp TranslationResponse
	if err := json.NewDecoder(resp.Body).Decode(&translationResp); err != nil {
		return nil, err
	}

	if len(translationResp.Translations) == 0 {
		return nil, errors.New("no translation found")
	}

	return &translationResp, nil
}
