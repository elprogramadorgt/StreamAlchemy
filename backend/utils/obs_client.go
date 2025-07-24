package utils

import (
	"encoding/json"
	"net/url"

	model_obs "github.com/elprogramadorgt/StreamAlchemy/models/obs"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type OBSClient struct {
	conn *websocket.Conn
}

func NewOBSClient(wsURL string) (*OBSClient, error) {

	u, err := url.Parse(wsURL)

	if err != nil {
		logrus.Error("Failed to parse WebSocket URL:", err)
		return nil, err
	}

	websocket.DefaultDialer.EnableCompression = true
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		return nil, err
	}

	client := &OBSClient{conn: conn}

	_, _, err = client.conn.ReadMessage()

	if err != nil {
		logrus.Error("Failed to read hello message:", err)
		client.conn.Close()
		return nil, err
	}

	identify := model_obs.OBSIdentify{
		Op: 1,
		D: model_obs.IdentifyOp{
			RpcVersion: 1,
		},
	}

	if err := client.send(identify); err != nil {
		logrus.Error("Failed to send identify message:", err)
		client.conn.Close()
		return nil, err
	}

	_, _, err = client.conn.ReadMessage()

	if err != nil {
		logrus.Error("Failed to read identify confirmation:", err)
		client.conn.Close()
		return nil, err
	}

	return client, nil
}

func (c *OBSClient) SendRequest(payload interface{}) ([]byte, error) {

	if err := c.send(payload); err != nil {
		return nil, err
	}

	_, response, err := c.conn.ReadMessage()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *OBSClient) send(payload interface{}) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, bytes)
}

func (c *OBSClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
