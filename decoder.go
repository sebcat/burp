package burp

import (
	"encoding/xml"
	"io"
)

type Host struct {
	IP   string `xml:"ip,attr"`
	Host string `xml:",chardata"`
}

type HTTPMessage struct {
	Base64 bool   `xml:"base64,attr"`
	Data   []byte `xml:",chardata"`
}

type Item struct {
	Time           string      `xml:"time"`
	URL            string      `xml:"url"`
	Host           Host        `xml:"host"`
	Port           int         `xml:"port"`
	Protocol       string      `xml:"protocol"`
	Method         string      `xml:"method"`
	Path           string      `xml:"path"`
	Request        HTTPMessage `xml:"request"`
	Status         int         `xml:"status"`
	ResponseLength int         `xml:"responselength"`
	MIMEType       string      `xml:"mimetype"`
	Response       HTTPMessage `xml:"response"`
	Comment        string      `xml:"comment"`
}

type Decoder struct {
	decoder *xml.Decoder
	done    bool
	err     error
	item    *Item
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{decoder: xml.NewDecoder(reader)}
}

func (d *Decoder) finalize() {
	d.done = true
}

func (d *Decoder) Next() bool {
OUTER:
	for {
		tok, err := d.decoder.Token()
		if err != nil {
			if err != io.EOF {
				d.err = err
			}
			d.finalize()
			break
		}

		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "item" {
				el := &Item{}
				d.decoder.DecodeElement(el, &se)
				d.item = el
				break OUTER
			}
		}
	}

	return !d.done
}

func (d *Decoder) Error() error {
	return d.err
}

func (d *Decoder) Item() *Item {
	return d.item
}
