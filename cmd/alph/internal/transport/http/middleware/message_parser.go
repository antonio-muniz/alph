package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/antonio-muniz/alph/pkg/respond"
	"github.com/antonio-muniz/alph/pkg/validator"
)

type ContextKey string

const Message ContextKey = "message"

type MessageType interface {
	NewPointer() interface{}
	Dereference(pointer interface{}) interface{}
}

type messageParser struct {
	messageType MessageType
	nextHandler http.Handler
}

func MessageParser(messageType MessageType, nextHandler http.Handler) http.Handler {
	return messageParser{
		messageType: messageType,
		nextHandler: nextHandler,
	}
}

func (m messageParser) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	messagePointer := m.messageType.NewPointer()
	err := json.NewDecoder(request.Body).Decode(messagePointer)
	if err != nil {
		respond.MalformedRequest(response)
		return
	}
	message := m.messageType.Dereference(messagePointer)
	validator := validator.New(validator.ErrorFieldFromJSONTag())
	validationResult, err := validator.Validate(message)
	if err != nil {
		respond.InternalServerError(response)
		return
	}
	if validationResult.Invalid() {
		respond.InvalidRequestParameters(response, validationResult)
		return
	}
	originalContext := request.Context()
	contextWithMessage := context.WithValue(originalContext, Message, message)
	requestWithMessage := request.WithContext(contextWithMessage)
	m.nextHandler.ServeHTTP(response, requestWithMessage)
}
