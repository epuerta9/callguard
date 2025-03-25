package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CheckSpam(phone string) bool {
	fmt.Println("Checking spam for number:", phone)

	url := "https://lookups.twilio.com/v1/PhoneNumbers/" + phone + "/?AddOns=nomorobo_spamscore"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.SetBasicAuth(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
		return false
	}

	var result TwillioAntiSpam
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return false
	}

	// Check the spam score
	spamScore := result.AddOns.Results.NomoroboSpamscore.Result.Score
	fmt.Println("Spam score:", spamScore)
	return spamScore > 0
}
