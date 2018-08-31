// Tags parse and unparse this JSON data, add this code to your project and do:
//
//    traceSegments, err := UnmarshalTraceSegments(bytes)
//    bytes, err = traceSegments.Marshal()

package main

import "encoding/json"

// TraceSegments should have a comment
type TraceSegments []TraceSegment

// UnmarshalTraceSegments should have a comment
func UnmarshalTraceSegments(data []byte) (TraceSegments, error) {
	var r TraceSegments
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal should have a comment
func (r *TraceSegments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// TraceSegment should have a comment
type TraceSegment struct {
	Gt                 [][]int64          `json:"gt"`
	TraceSegmentObject TraceSegmentObject `json:"sg"`
}

// TraceSegmentObject should have a comment
type TraceSegmentObject struct {
	ApplicationID         int64        `json:"ai"`
	ApplicationInstanceID int64        `json:"ii"`
	SpanObject            []SpanObject `json:"ss"`
	Ts                    []int64      `json:"ts"`
}

// SpanObject should have a comment
type SpanObject struct {
	ComponentID            int64                   `json:"ci"`
	Component              string                  `json:"cn"`
	EndTime                int64                   `json:"et"`
	IsError                bool                    `json:"ie"`
	Logs                   []Log                   `json:"lo"`
	SpanLayer              int64                   `json:"lv"`
	OperationNameID        int64                   `json:"oi"`
	OperationName          string                  `json:"on"`
	PeerID                 int64                   `json:"pi"`
	Peer                   string                  `json:"pn"`
	ParentSpanID           int64                   `json:"ps"`
	TraceSegmentReferences []TraceSegmentReference `json:"rs"`
	SpanID                 int64                   `json:"si"`
	StartTime              int64                   `json:"st"`
	Tags                   []Tag                   `json:"to"`
	SpanType               int64                   `json:"tv"`
}

// Log should have a comment
type Log struct {
	LogTags   []Tag `json:"ld"`
	Timestamp int64 `json:"ti"`
}

// Tag should have a comment
type Tag struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

// TraceSegmentReference should have a comment
type TraceSegmentReference struct {
	EntryApplicationInstanceID  int64   `json:"eii"`
	EntryServiceID              int64   `json:"esi"`
	EntryServiceName            string  `json:"esn"`
	NetworkAddressID            int64   `json:"ni"`
	NetworkAddress              string  `json:"nn"`
	ParentApplicationInstanceID int64   `json:"pii"`
	ParentSpanID                int64   `json:"psi"`
	ParentServiceName           string  `json:"psn"`
	ParentServiceID             int64   `json:"psp"`
	ParentTraceSegmentIDs       []int64 `json:"pts"`
	RefTypeValue                int64   `json:"rv"`
}

