package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func sendMailHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	mailInfo, err := unmarshalRequestBody(r)
	if err != nil {
		response := Response{
			Message: "Invalid Request Body ",
		}
		writeResponse(r.Context(), w, response, http.StatusBadRequest)
		return
	}

	mailer := GetMailer()
	if err := mailer.Send(ctx, mailInfo); err != nil {
		response := Response{
			Message: "Failed to send Email",
		}
		writeResponse(r.Context(), w, response, http.StatusInternalServerError)
		return
	}

	response := Response{
		Message: "Mail Sent succesfully",
	}

	writeResponse(r.Context(), w, response, http.StatusCreated)
}

func unmarshalRequestBody(r *http.Request) (info MailInfo, err error) {
	body, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Error("Error while unmarshalling data.", err.Error())
		return
	}
	return
}
