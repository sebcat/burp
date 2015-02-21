package burp

import (
	"strings"
	"testing"
)

var xmlData = `
<items>
<item>
<time>Sat Feb 21 06:59:58 CET 2015</time>
<url><![CDATA[http://www.reddit.com/api/request_promo]]></url>
<host ip="198.41.209.136">www.reddit.com</host>
<port>80</port>
<protocol>http</protocol>
<method>POST</method>
<path><![CDATA[/api/request_promo]]></path>
<extension>null</extension>
<request base64="true"><![CDATA[AAAA]]>
</request>
<status>200</status>
<responselength>383</responselength>
<mimetype></mimetype>
<response base64="true">
BBBB
</response>
<comment></comment>
</item>

  <item>
    <time>Sat Feb 21 06:59:59 CET 2015</time>
    <url><![CDATA[http://www.google-analytics.com/plugins/ga/inpage_lin
kid.js]]></url>
    <host ip="216.58.209.142">www.google-analytics.com</host>
    <port>80</port>
    <protocol>http</protocol>
    <method>GET</method>
    <path><![CDATA[/plugins/ga/inpage_linkid.js]]></path>
    <extension>js</extension>
    <request base64="true">
</request>
    <status>304</status>
    <responselength>170</responselength>
    <mimetype></mimetype>
    <response base64="true">
</response>
    <comment></comment>
  </item>


</items>
`

// A demo usage example, not a testcase proper
func TestDecoder(t *testing.T) {
	r := strings.NewReader(xmlData)
	x := NewDecoder(r)

	for x.Next() {
		if item := x.Item(); item != nil {
			t.Log(item)
		}
	}

	if err := x.Error(); err != nil {
		t.Fatal(err)
	}

}
