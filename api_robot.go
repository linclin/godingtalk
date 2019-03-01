package godingtalk

import (
	"net/url"
)

//SendRobotTextMessage can send a text message to a group chat
func (c *DingTalkClient) SendRobotTextMessage(accessToken string, msg string) error {
	var data OAPIResponse
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	err := c.httpRPC("robot/send", params, request, &data)
	return err
}

//SendRobotMarkdownMessage can send a Markdown message to a group chat
func (c *DingTalkClient) SendRobotMarkdownMessage(accessToken string, title, text string) error {
	var data OAPIResponse
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  text,
		},
	}
	err := c.httpRPC("robot/send", params, request, &data)
	return err
}
