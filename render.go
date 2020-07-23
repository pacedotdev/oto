package main

import (
	"bytes"
	"encoding/json"
	"go/doc"
	"html/template"
	"strings"

	"github.com/fatih/structtag"
	"github.com/gobuffalo/plush"
	"github.com/markbates/inflect"
	"github.com/pkg/errors"
)

var defaultRuleset = inflect.NewDefaultRuleset()

// render renders the template using the Definition.
func render(template string, def Definition, params map[string]interface{}) (string, error) {
	ctx := plush.NewContext()
	ctx.Set("camelize_down", camelizeDown)
	ctx.Set("def", def)
	ctx.Set("params", params)
	ctx.Set("json", toJSONHelper)
	ctx.Set("format_comment_text", formatCommentText)
	ctx.Set("format_comment_html", formatCommentHTML)
	ctx.Set("format_tags", formatTags)
	s, err := plush.Render(string(template), ctx)
	if err != nil {
		return "", err
	}
	return s, nil
}

func toJSONHelper(v interface{}) (template.HTML, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}

func formatCommentText(s string) string {
	var buf bytes.Buffer
	doc.ToText(&buf, s, "// ", "", 80)
	return buf.String()
}

func formatCommentHTML(s string) template.HTML {
	var buf bytes.Buffer
	doc.ToHTML(&buf, s, nil)
	return template.HTML(buf.String())
}

// camelizeDown converts a name or other string into a camel case
// version with the first letter lowercase. "ModelID" becomes "modelID".
func camelizeDown(word string) string {
	if isAcronym(word) {
		// entire word is an acronym
		return strings.ToLower(word)
	}
	words := Split(word)
	for i := range words {
		if isAcronym(words[i]) {
			if i == 0 {
				words[i] = strings.ToLower(words[i])
			} else {
				words[i] = strings.ToUpper(words[i])
			}
		}
	}
	word = strings.Join(words, "")
	return strings.ToLower(word[:1]) + word[1:]
}

// formatTags formats a list of struct tag strings into one.
// Will return an error if any of the tag strings are invalid.
func formatTags(tags ...string) (template.HTML, error) {
	alltags := &structtag.Tags{}
	for _, tag := range tags {
		theseTags, err := structtag.Parse(tag)
		if err != nil {
			return "", errors.Wrapf(err, "parse tags: `%s`", tag)
		}
		for _, t := range theseTags.Tags() {
			alltags.Set(t)
		}
	}
	tagsStr := alltags.String()
	if tagsStr == "" {
		return "", nil
	}
	tagsStr = "`" + tagsStr + "`"
	return template.HTML(tagsStr), nil
}

func isAcronym(word string) bool {
	for _, ac := range baseAcronyms {
		if strings.ToUpper(ac) == strings.ToUpper(word) {
			return true
		}
	}
	return false
}

var baseAcronyms = strings.Split(`HTML,JSON,JWT,ID,UUID,SQL,ACK,ACL,ADSL,AES,ANSI,API,ARP,ATM,BGP,BSS,CAT,CCITT,CHAP,CIDR,CIR,CLI,CPE,CPU,CRC,CRT,CSMA,CMOS,DCE,DEC,DES,DHCP,DNS,DRAM,DSL,DSLAM,DTE,DMI,EHA,EIA,EIGRP,EOF,ESS,FCC,FCS,FDDI,FTP,GBIC,gbps,GEPOF,HDLC,HTTP,HTTPS,IANA,ICMP,IDF,IDS,IEEE,IETF,IMAP,IP,IPS,ISDN,ISP,kbps,LACP,LAN,LAPB,LAPF,LLC,MAC,MAN,Mbps,MC,MDF,MIB,MoCA,MPLS,MTU,NAC,NAT,NBMA,NIC,NRZ,NRZI,NVRAM,OSI,OSPF,OUI,PAP,PAT,PC,PIM,PIM,PCM,PDU,POP3,POP,POTS,PPP,PPTP,PTT,PVST,RADIUS,RAM,RARP,RFC,RIP,RLL,ROM,RSTP,RTP,RCP,SDLC,SFD,SFP,SLARP,SLIP,SMTP,SNA,SNAP,SNMP,SOF,SRAM,SSH,SSID,STP,SYN,TDM,TFTP,TIA,TOFU,UDP,URL,URI,USB,UTP,VC,VLAN,VLSM,VPN,W3C,WAN,WEP,WiFi,WPA,WWW`, ",")
