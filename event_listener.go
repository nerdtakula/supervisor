package supervisor

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "reflect"
  "strconv"
  "strings"
)

type Event string

const (
  EVENT_BASE                             Event = "EVENT"
  EVENT_PROCESS_STATE                          = "PROCESS_STATE"
  EVENT_PROCESS_STATE_STARTING                 = "PROCESS_STATE_STARTING"
  EVENT_PROCESS_STATE_RUNNING                  = "PROCESS_STATE_RUNNING"
  EVENT_PROCESS_STATE_BACKOFF                  = "PROCESS_STATE_BACKOFF"
  EVENT_PROCESS_STATE_STOPPING                 = "PROCESS_STATE_STOPPING"
  EVENT_PROCESS_STATE_EXITED                   = "PROCESS_STATE_EXITED"
  EVENT_PROCESS_STATE_STOPPED                  = "PROCESS_STATE_STOPPED"
  EVENT_PROCESS_STATE_FATAL                    = "PROCESS_STATE_FATAL"
  EVENT_PROCESS_STATE_UNKNOWN                  = "PROCESS_STATE_UNKNOWN"
  EVENT_REMOTE_COMMUNICATION                   = "REMOTE_COMMUNICATION"
  EVENT_PROCESS_LOG                            = "PROCESS_LOG"
  EVENT_PROCESS_LOG_STDOUT                     = "PROCESS_LOG_STDOUT"
  EVENT_PROCESS_LOG_STDERR                     = "PROCESS_LOG_STDERR"
  EVENT_PROCESS_COMMUNICATION                  = "PROCESS_COMMUNICATION"
  EVENT_PROCESS_COMMUNICATION_STDOUT           = "PROCESS_COMMUNICATION_STDOUT"
  EVENT_PROCESS_COMMUNICATION_STDERR           = "PROCESS_COMMUNICATION_STDERR"
  EVENT_SUPERVISOR_STATE_CHANGE                = "SUPERVISOR_STATE_CHANGE"
  EVENT_SUPERVISOR_STATE_CHANGE_RUNNING        = "SUPERVISOR_STATE_CHANGE_RUNNING"
  EVENT_SUPERVISOR_STATE_CHANGE_STOPPING       = "SUPERVISOR_STATE_CHANGE_STOPPING"
  EVENT_TICK                                   = "TICK"
  EVENT_TICK_5                                 = "TICK_5"
  EVENT_TICK_60                                = "TICK_60"
  EVENT_TICK_3600                              = "TICK_3600"
  EVENT_PROCESS_GROUP                          = "PROCESS_GROUP"
  EVENT_PROCESS_GROUP_ADDED                    = "PROCESS_GROUP_ADDED "
  EVENT_PROCESS_GROUP_REMOVED                  = "PROCESS_GROUP_REMOVED"
)

type EventMessage struct {
  Headers *HeaderToken
  Body    []byte
}

func NewEventMessage(headers *HeaderToken, body []byte) EventMessage {
  return EventMessage{
    Headers: headers,
    Body:    body,
  }
}

func (e EventMessage) String() string { return string(e.Body) }

func (e EventMessage) AsMap() map[string]string {
  data := make(map[string]string)
  
  body := strings.Split(strings.TrimSpace(e.String()), " ")
  for _, s := range body {
    t := strings.Splits(s, ":")
    data[t[0]] = t[1]
  }
  return data
}

type HeaderToken struct {
  Version    string `token:"ver"` // The event system protocol version.
  Server     string `token:"server"` // The identifier of the supervisord sending the event.
  Serial     int    `token"serial"` // An integer assigned to each event.
  Pool       string `token:"pool"` // The name of the event listener pool which generated this event.
  PoolSerial int    `token:"poolserial"` // An integer assigned to each event by the eventlistener pool which is being sent from.
  EventName  Event  `token:"eventname"` // The specific event type name.
  Length     int     `token:"len"` // An integer indicating the number of bytes in the event payload.
}
