// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Command struct {
	DeviceID string `json:"deviceId"`
	Command  string `json:"command"`
}

type CommandOutput struct {
	Code   int    `json:"code"`
	Output string `json:"output"`
}

type Device struct {
	ID string `json:"_id"`
}
