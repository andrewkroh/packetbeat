// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package file_integrity

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mimeSamples = map[string][]byte{
	"docx": []byte("PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x12\x00\x00\x00word/numbering.xml\xa5\x93MN\xc30\x10\x85O\xc0\x1d\"\xef\xdb$\x15 \x145\xed\x82\n6\xec\x80\x03\xb8\x8e\x93X\xb5=\xd6\xd8I\xe8\xedq\x9b\xbfR$\x94\x86U\xe4\x8c\xdf\xf7\xc6\xe3\xe7\xf5\xf6K\u0260\xe6h\x05\xe8\x94\xc4\u02c8\x04\\3\u0204.R\xf2\xf9\xf1\xb2x\"\x81uTgT\x82\xe6)9rK\xb6\x9b\xbbu\x93\xe8J\xed9\xfa}\x81Gh\x9b(\x96\x92\xd29\x93\x84\xa1e%W\xd4.\xc1p\xed\x8b9\xa0\xa2\xce/\xb1\b\x15\xc5Ce\x16\f\x94\xa1N\xec\x85\x14\xee\x18\xae\xa2\xe8\x91t\x18HI\x85:\xe9\x10\v%\x18\x82\x85\u071d$\t\xe4\xb9`\xbc\xfb\xf4\n\x9c\xe2\xdbJv\xc0*\u0175;;\x86\u0225\xef\x01\xb4-\x85\xb1=M\u0365\xf9b\xd9C\xea\xbf\x0eQ+\xd9\xefk\xcc\x14\xb7\fi\xe3\xe7\xacdk\xd4\x00f\x06\x81qk\xfd\xdf][\x1c\x88q4a\x80'\u0120\x98\xd2\xc2O\u03fe\x13E\x85\x1e0\xa7t\\\x81\x06\xef\xa5\xf7\xee\x86vF\x8d\a\x19ga\xe5\x94F\xda\u049b\xd8#\xc5\xe3\xef.\xe8\x8cy^\ua358\x94\xe2+\x82W\xb9\n\x87@\xceA\xb0\x92\xa2\xeb\x01r\x0eA\x02;\xf0\xec\x99\xea\x9a\x0ea\u038aIq\xbe\"e\x82\x16H\xd5\x18R{\xd3\xcd\xc6\xd1U\\\xdeKj\xf8H+\xfeG{E\xa8\xcc\x18\xf7\xfb9\xb4\x8b\x17\x18?\xdc\x06X\xf5\x80p\xf3\rPK\a\bI\x13C\u007fh\x01\x00\x00=\x05\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00\x00\x00word/settings.xml\xa5\x95\xcdn\xdb0\f\u01df`\xef\x10\xe8\x9e\xf8\xa3I6\x18uzX\xb1\xed\xb0\x9e\xd2=\x00#\u0276\x10}A\x92\xe3\xe5\xed'\u01d6\u0564@\xe1f\xa7H\u007f\x92?\xd2\fM?>\xfd\x15|q\xa2\xc62%K\x94\xadR\xb4\xa0\x12+\xc2d]\xa2?\xaf?\x96\xdf\xd0\xc2:\x90\x04\xb8\x92\xb4Dgj\xd1\xd3\xee\xcbcWX\xea\x9c\xf7\xb2\vO\x90\xb6\x10\xb8D\x8ds\xbaH\x12\x8b\x1b*\xc0\xae\x94\xa6\xd2\x1b+e\x048\u007f5u\"\xc0\x1c[\xbd\xc4Jhp\xec\xc08s\xe7$O\xd3-\x1a1\xaaD\xad\x91\u0148X\n\x86\x8d\xb2\xaar}H\xa1\xaa\x8aa:\xfe\x84\b3'\xef\x10\xf2\xacp+\xa8t\x97\x8c\x89\xa1\xdc\u05e0\xa4m\x98\xb6\x81&\xee\xa5yc\x13 \xa7\x8f\x1e\xe2$x\xf0\xeb\xf4\x9cl\xc4@\xe7\x1b-\xf8\x90\xa8S\x86h\xa30\xb5\u05ab\u03c3q\"f\xe9\x8c\x06\xf6\x88)bN\t\xd79C%\x02\x98\x9c0\xfdp\u0700\xa6\xdc+\x9f{l\xda\x05\x15\x1f$\xf6\xc2\xf29\x85\f\xa6\xdf\xec`\xc0\x9c\xdfW\x01w\xf4\xf3m\xbcf\xb3\xa6\xf8\x86\xe0\xa3\\k\xa6\x81\xbc\a\x81\x1b0.\x00\xf8=\x04\xae\xf0\x91\x92\xef O0\r3\xa9g\x8d\xf3\r\x890\xa8\r\x888\xa4\xf6S\xffl\x96\u078c\u02fe\x01M#\xad\xfe?\xdaO\xa3Z\x1d\xc7}}\x0f\xed\xcd\x1b\x98m>\a\xc8\x03`\xe7W \xa1\x15\xb4\u073d\xc2a\xef\x94^t\xc5\t\xfc\x14\u007f\xcdS\x94\xf4\xe6a\xcb\xc5\xd3~\u0618\xc1/\xdb \u007f\x94 \xfc\x9bs\xb5\x10_\x14\xa1\xbd\xa95l~q}\xca\xe4*'7\xfb>\x88\xbe\x80\xd6C\xdaC\x9d\x95\x88\xb3\xbaqY\xcfw\xfeF\xfcB\xbe\\\x0eu>\xda\xf2\x8b-\x1fl\x97\v`\xec\xf7\x9c\xf7\x1e\x0fQ\u02c3\xf6\xc6\xef!h\x0fQ[\am\x1d\xb5M\xd06Q\xdb\x06m\xdbk\xcdYS\u00d9<\xfa6\x84c\xafW\x8as\xd5Q\xf2+\xda\xdfIc?\xc2Wj\xf7\x0fPK\a\b\x8e\xb3\u00e4\x05\x02\x00\x00\xea\x06\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x12\x00\x00\x00word/fontTable.xml\xa5\x94MN\xc30\x10\x85O\xc0\x1d\"\xef\u06e4\b\x10\x8a\x9aT\b\x04\x1bv\xc0\x01\x06\xc7I\xac\xda\x1ek\xec4\xf4\xf6\xb84?P$\x94\x86U\x94\x8c\xdf\xf7\xc6\xe3\x17\xaf7\x1fZE;AN\xa2\xc9\xd8j\x99\xb0H\x18\x8e\x854U\xc6\xde^\x1f\x17\xb7,r\x1eL\x01\n\x8d\xc8\xd8^8\xb6\xc9/\xd6mZ\xa2\xf1.\nr\xe3R\xcd3V{o\xd38v\xbc\x16\x1a\xdc\x12\xad0\xa1X\"i\xf0\u156aX\x03m\x1b\xbb\xe0\xa8-x\xf9.\x95\xf4\xfb\xf82InX\x87\xc1\x8c5d\xd2\x0e\xb1\u0412\x13:,\xfdA\x92bYJ.\xbaG\xaf\xa0)\xbeG\xc9\x03\xf2F\v\xe3\xbf\x1cc\x12*\xf4\x80\xc6\xd5\u04ba\x9e\xa6\xe7\xd2B\xb1\xee!\xbb\xbf6\xb1\u04ea_\xd7\xda)n\x05A\x1b\xceB\xab\xa3Q\x8bTXB.\x9c\v_\x1f\x8e\u0141\xb8J&\f\xf0\x80\x18\x14SZ\xf8\xe9\xd9w\xa2A\x9a\x01sH\xc6\th\xf0^\x06\xefnh_\xa8q#\xe3,\x9c\x9a\xd2\u0231\xf4,\xdf\th\xff\xbb\v\x981\xcf\xefz+'\xa5\xf8\x84\x10T\xbe\xa1!\x90s\x10\xbc\x06\xf2=@\xcd!(\xe4[Q\u0703\xd9\xc1\x10\u689a\x14\xe7\x13R!\xa1\"\xd0cH\xddY'\xbbJN\xe2\xf2R\x83\x15#\xad\xfa\x1f\ud270\xb1c\u072f\xe6\u043e\xfd\x81\xab\xeb\xf3\x00\x97= \xef\uefe8M\r\xe8\x10\xfe;\x92\xa0X\x9c\xaf\xe3\xeeb\xcc?\x01PK\a\b\xad\x87m\x00y\x01\x00\x00Z\x05\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00\x00\x00word/styles.xml\u0556\xedn\xda0\x14\x86\xaf`\xf7\x80\xf2\xbfMH\x02CQiU\xb5\xea6\xa9\ua9b5\xbb\x80\x83c\x88U\u01f6l\a\u02ae~\xce7$\xa1J\x03\x12\x1d\xfc\x00\x1f\xfb\xbc\xc7~\xfc:\xce\xd5\xcd[LGk,\x15\xe1ln\x8d/\x1dk\x84\x19\xe2!a\xab\xb9\xf5\xe7\xe5\xe1bf\x8d\x94\x06\x16\x02\xe5\f\u03ed-V\xd6\xcd\xf5\x97\xabM\xa0\xf4\x96b52\xf9L\x051\x9a[\x91\xd6\"\xb0m\x85\"\x1c\x83\xba\xe4\x023\u04f9\xe42\x06m\x9are\xc7 _\x13q\x81x,@\x93\x05\xa1Dom\xd7q\xa6V!\xc3\xe7V\"YPH\\\xc4\x04I\xae\xf8R\xa7)\x01_.\t\xc2\xc5O\x99!\xfb\xd4\xcdS\xee9Jb\xcctV\u0456\x98\x9a9p\xa6\"\"T\xa9\x16\x0fU3\x9dQ)\xb2~o\x11\ub616\xe36\xa2O\xb5P\xc2\xc6lFL\xf3B\x1b.C!9\xc2J\x99\xe8}\xdeY)\x8e\x9d\x1e\x00S\x89*\xa3\xcf\x14\xf6k\x963\x89\x81\xb0J&\xb5FC\xa8\xaa}ij\x17\xd02\xa9z!5\vE\xfbL$\xefz$\v\tr\u06de\x05\f\u0e5b/H/\x177\x14L\x96Nde\xc8!\x12(\x02\xa9K\x01:D\x81r\xf4\x8a\xc3;`k\xa8\xcc\x1c\xaez\u0679\xa1\x14\x12XI\x88k\x93\xaa\x0f\xed\xec\xd8i\xd8\xe59\x02\x81k\xb5\xd5qj\xdf$ODmw\u007f\x88\xda\xce\t\x1cO>&\xe0\x96\x02\xd7\xe6\x01\x18rt\x8f\x97\x90P\xad\u04a6\xfc%\x8bf\xd1\xca~\x1e8\xd3j\xb4\t@!B\xe6\u05ad$`\xcao\x02\xa4v\x1a\x18\x94\xbeU\x04vB\xd1-S\xd5x;\x95R\u007fMx\r\u6838n\x19\xb9S\xcd\x18\x05\xb6*c\x98\xa51\xbb\x98\x8c\u075c\xa2h\xb62M\x01\x88d\x12\x94\xa4\x87\xda\xfd:\xb5\x8a\xc6\uf11a\x00$\x9a\x17\xb2\xa2\x90\xdd\x15\xb2[\\\xb2{\xc2H\xe8\xad0\xe9\x02d\xea/\x11\xa5\xaaY\u05cfpn=\xa5~\xcc\xd6\x1d\xe6\x99\xe6*\xca\x183\x88q\xb9\x1c\x96\x0f\xcakg\xa9my\r\v\x8a\xf7\xa4_\xd2H/\xfdl\xe4\xe8\xa9G\x95\xeeE|\u01d0^\x9bm\xe1(\xef\x18\x8d\xf3-Z\x80\xc2\xe1OV\xf6\xd6\x05M\x16~\xd3]\xf1bs^1\x16O;C\n\xc14\xfch6H5\xe2\xf5^\xc2RcsS\x8e]'\x9d\xf1\x02\x9b\xf3o\x96\xe1;\xce\xfb{[\u0678\xf6\x9e\ufd3d\x97\xc7v|6\x04\x9b{\x10\x9b\xfb\u0270y\u04fe\xd8\x16\xa5\xb2\xd3<\xc2^\xc7\x11\xcecGb\xf4\x0eb\xf4\u038dq\xb6O\xd1\x1dJ\x11q\xcae\xe5=/\xfd\xb6\x9e\x90\xb3\x8e'\xe4\xec\x04x\xfd\x83x\xfd\u03c5\u05dd\xf5\u017b\x87s\x9a}Z8\xfd\x0e\x9c\xfe\tpN\x0e\xe2\x9c|2\x9c\xfe)q\x1e\xbc\xbf\x8f\xc49=\x88s\xfa\xbf\xe2$\r\xe1\xb3\xe0}!\u06bcU\xb4\xde\x17\xb2\u8679N\xf7\xb8~\xfc>\x9ft\xc0\x9a\x1c\x05\xeb9Y\xe8N^U\u01d9\x91y\xee f'|\x95\xafL\xddu\xa3u\x9b\xda\xebx\xef\xf2\x0e\xbcw\x95\xff\xd4\xf5?PK\a\b\x88\xceE\f\x1d\x03\x00\x00\xdf\x11\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00\x00\x00word/document.xml\xa5\x95Yn\xdb0\x10\x86O\xd0;\x18|\xb7%\x19N\x1b\b\x96\xf3P\xa3E\x81\xb60\xb2\x1c\x80&)\x89\b7\f)\xab\xee\xe9Kj\xf5\x12\x04\x8a\xeb\x17b\xb6o~\x91cr\xfd\xf0G\x8a\u0641\x81\xe5Ze(Y\xc4h\xc6\x14\u0454\xab\"C/\xcf\xdf\xe6\xf7hf\x1dV\x14\v\xadX\x86\x8e\u0322\x87\u0367u\x9dRM*\u0254\x9by\x82\xb2\xa9$\x19*\x9d3i\x14YR2\x89\xedB\x1b\xa6|0\xd7 \xb1\xf3&\x14\x91\xc4\xf0Z\x999\xd1\xd2`\xc7\xf7\\pw\x8c\x96q\xfc\x19u\x18\x9d\xa1\nT\xda!\xe6\x92\x13\xd0V\xe7.\x94\xa4:\xcf9a\xdd\xd2W\xc0\x94\xbem\u0276\x93\xdct\x8c\x80\t\xafA+[rc{\x9a\xbc\x95\xe6\x83e\x0f9\xbc\xf7\x11\a)\xfa\xbc\xdaL\xe9F\x01\xd7\xfe8\xa4h\x1b\xd5\x1a\xa8\x01M\x98\xb5\u07bbm\x83\x031\x89'l`@\f\x15S$\x9c\xf7\xec\x95H\xcc\u0540\t\xc3q\x01\x1az/|\xefn\xd3\x1a\xd4\xf8!\xe3^X1EH\x1b\xfa\xc9\xf7\x80\xe1x\xad\x02\u07f0\x9f\xa7\xf5\x86O\x9a\xe2\v\x82\xafr\x15\f\x03y\v\x82\x94\x18\\\x0f\x10\xb7\x10\x84&\xaf\x8c~\xc5\ua007a\xa6\u0164q\xbe Q\x8e\v\xc0r\x1cR\xfb\xa1\x93M\xe2\x8bqy*\xb1a#\xad\xf8?\xdaw\u0415\x19\xc7}u\v\xed\xe4\x1f\x98\xdc}\f\xb0\xec\x01\x1b\u007f\x05\xee5=\x86\xd5\xcc\xea\xd4\u07e0\xf41Cq\xf7C\x9dk\xcb\u0135sw\xedz\u0732\x1cW\u00bd\x11\xd9\xc1\x993Y\xa5\x06\x03\xfeA\ao\u04881;\b\v\xec \u06ac\xa3\xd1~O\xc8\x1b\x82\xcf\xdbu\xc4fq\u00a7\x1cp\xc0\xa0\xb6E\x13\tk\xdb0dYF\\\x9bo\x8a\xa7\xbf\xbe\xa0\xf4\xaf\xca\xdd\xfd\xaa\xe1\xfb\xbb&Y.WMyH\xf8\x85\x83\xba\xbdvN\xfbAMVm\x96\xd3f4\x04\xcb\xddh\x01/\xca\x13\xb3d\x982/\xf7\u02f21s\xad]ov\x1d~W\xf2\xf9h\x98\x0f\xfaG\fBi'\xbd\xd7\x19\xf5\xa7\x18\x8d/\xda\xe6\x1fPK\a\b@Z0\xeb\x17\x02\x00\x00\x16\a\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1c\x00\x00\x00word/_rels/document.xml.rels\xad\x92Mj\xc30\x10\x85O\xd0;\x88\xd9\u05f2\xd3\x1fJ\x89\x9cM\bd[\xdc\x03(\xf2\xf8\x87Z#!MJ}\xfb\x8a\x94$\x0e\x04\u04c5\x97\xef\x89y\xf3\u034c\u059b\x1f;\x88o\f\xb1w\xa4\xa0\xc8r\x10H\xc6\xd5=\xb5\n>\xab\xdd\xe3\x1b\x88\u021aj=8B\x05#F\u0614\x0f\xeb\x0f\x1c4\xa7\x9a\xd8\xf5>\x8a\x14BQA\xc7\xec\u07e5\x8c\xa6C\xabc\xe6<Rzi\\\xb0\x9a\x93\f\xad\xf4\xda|\xe9\x16\xe5*\xcf_e\x98f@y\x93)\xf6\xb5\x82\xb0\xaf\v\x10\xd5\xe8\xf1?\u066eiz\x83[g\x8e\x16\x89\ufd10\x9cj1\x05\xea\xd0\"+8\xc9?\xb3\xc8R\x18\xc8\xfb\f\xab%\x19\"2\xa7\xe5\xc6+\xc6\u0659CxZ\x12\xa1q\u0115>\f\x93U\\\xac9\x88\xe7%!\xe8h\x0f\x18\xd2\xdcW\x88\x8b5\a\xf1\xb2\xe81x\x1cpz\x8a\x93>\xb7\x977\x9f\xbc\xfc\x05PK\a\b\x90\x00\xab\xeb\xf1\x00\x00\x00,\x03\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\v\x00\x00\x00_rels/.rels\x8d\xcf;\x0e\xc20\f\x06\xe0\x13p\x87\xc8;M\u02c0\x10j\xd2\x05!uE\xe5\x00Q\xe2\xa6\x11\xcdCIx\xf4\xf6d`\x00\xc4\xc0h\xfb\xf7g\xb9\xed\x1ev&7\x8c\xc9x\u01e0\xa9j \xe8\xa4W\xc6i\x06\xe7\xe1\xb8\xde\x01IY8%f\xef\x90\xc1\x82\t:\xbejO8\x8b\\v\xd2dB\"\x05q\x89\xc1\x94s\xd8S\x9a\xe4\x84V\xa4\xca\ate2\xfahE.e\xd44\by\x11\x1a\u99ae\xb74\xbe\x1b\xc0?L\xd2+\x06\xb1W\r\x90a\t\xf8\x8f\xed\xc7\xd1H<xy\xb5\xe8\xf2\x8f\x13_\x89\"\x8b\xa813\xb8\xfb\xa8\xa8z\xb5\xab\xc2\x02\xe5-\xfdx\x91?\x01PK\a\b-h\xcf\"\xb1\x00\x00\x00*\x01\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x15\x00\x00\x00word/theme/theme1.xml\xedYKo\xdb6\x1c\xbf\x0f\xd8w toe\xd9V\xea\x04u\x8a\u0631\u06edM\x1b$n\x87\x1ei\x89\x96\xd8P\xa2@\xd2I|\x1b\xda\xe3\x80\x01\u00faa\x87\x15\xd8m\x87a[\x81\x16\u0625\xfb4\xd9:l\x1d\u042f\xb0\xbf\x1e\x96)\x9b\u03a3M\xb7\x0e\xad\x0f6I\xfd\xfe\xef\aI\xf9\xf2\x95\u00c8\xa1}\"$\xe5q\xdbr.\xd6,Db\x8f\xfb4\x0e\xda\xd6\xedA\xffB\xcbBR\xe1\xd8\u01cc\u01e4mM\x88\xb4\xae\xac\u007f\xf8\xc1e\xbc\xa6B\x12\x11\x04\xf4\xb1\\\xc3m+T*Y\xb3m\xe9\xc12\x96\x17yBbx6\xe2\"\xc2\n\xa6\"\xb0}\x81\x0f\x80o\xc4\xecz\xad\xb6bG\x98\xc6\x16\x8aq\x04lo\x8dF\xd4#h\x90\xb2\xb4\u05a7\xcc{\f\xbeb%\xd3\x05\x8f\x89]/\x93\xa8SdX\u007f\xcfI\u007f\xe4Dv\x99@\xfb\x98\xb5-\x90\xe3\xf3\x83\x019T\x16bX*x\u0436j\xd9\u01f2\xd7/\xdb%\x11SKh5\xba~\xf6)\xe8\n\x02\u007f\xaf\x9e\u0449`X\x12:\xfd\xe6\xea\xa5\u0352\u007f=\u7fc8\xeb\xf5z\u075eS\xf2\xcb\x00\xd8\xf3\xc0Rg\x01\xdb\ucddc\u0394\xa7\x06\u0287\x8b\xbc\xbb5\xb7\u05ac\xe25\xfe\x8d\x05\xfcj\xa7\xd3qW+\xf8\xc6\f\xdf\\\xc0\xb7j+\u034dz\x05\u07dc\xe1\xddE\xfd;\x1b\xdd\xeeJ\x05\xef\xce\xf0+\v\xf8\xfe\xa5\u0555f\x15\x9f\x81BF\xe3\xbd\x05t\x1a\xcf22%d\xc4\xd95#\xbc\x05\xf0\xd64\x01f([\u02ee\x9c>V\xcbr-\xc2\xf7\xb8\xe8\x03 \v.V4Fj\x92\x90\x11\xf6\x00\xd7\u014c\x0e\x05M\x05\xe05\x82\xb5'\xf9\x92'\x17\x96RYHz\x82&\xaam}\x9c`\xa8\x88\x19\xe4\xe5\xb3\x1f_>{\x82\x8e\xee?=\xba\xff\xcb\u0443\aG\xf7\u007f6P]\xc3q\xa0S\xbd\xf8\xfe\x8b\xbf\x1f}\x8a\xfez\xf2\u074b\x87_\x99\xf1R\xc7\xff\xfe\xd3g\xbf\xfd\xfa\xa5\x19\xa8t\xe0\xf3\xaf\x1f\xff\xf1\xf4\xf1\xf3o>\xff\xf3\x87\x87\x06\xf8\x86\xc0C\x1d>\xa0\x11\x91\xe8&9@;<\x02\xc3\f\x02\xc8P\x9c\x8db\x10b\xaaSl\u0101\xc41Ni\f\xe8\x9e\n+\xe8\x9b\x13\u0330\x01\xd7!U\x0f\xde\x11\xd0\x02L\xc0\xab\xe3{\x15\x85wC1V\xd4\x00\xbc\x1eF\x15\xe0\x16\xe7\xac\u00c5\u0466\xeb\xa9,\xdd\v\xe380\v\x17c\x1d\xb7\x83\xf1\xbeIvw.\xbe\xbdq\x02\xb9LM,\xbb!\xa9\xa8\xb9\xcd \xe48 1Q(}\xc6\xf7\b1\x90\u0765\xb4\xe2\xd7-\xea\t.\xf9H\xa1\xbb\x14u05\xbad@\x87\xcaLt\x8dF\x10\x97\x89IA\x88w\xc57[wP\x873\x13\xfbM\xb2_EBU`fbIX\u014dW\xf1X\xe1\u02281\x8e\x98\x8e\xbc\x81UhRrw\"\xbc\x8a\u00e5\x82H\a\x84q\xd4\xf3\x89\x94&\x9a[bRQ\xf7:\xb4\x0es\u0637\xd8$\xaa\"\x85\xa2{&\xe4\r\u0339\x8e\xdc\xe4{\xdd\x10G\x89Qg\x1a\x87:\xf6#\xb9\a)\x8a\xd16WF%x\xb5B\xd29\xc4\x01\xc7K\xc3}\x87\x12u\xb6\u06beM\x83\u041c \u94f10\x95\x04\xe1\xd5z\x9c\xb0\x11&q\xd1\xe1+\xbd:\xa2\xf1q\x8d;\x82\xbe\x8d\u03fbqC\xab|\xfe\xed\xa3\xffQ\xcb\xde\x00'\x98jf\xbeQ/\xc3\u0377\xe7.\x17>}\xfb\xbb\xf3&\x1e\xc7\xdb\x04\n\xe2}s~\u07dc\xdf\xc5\u6f2c\x9e\u03ff%\u03fa\xb0\xad\x1f\xb436\xd1\xd2S\xf7\x882\xb6\xab&\x8c\u0710Y\xff\x96`\x9e\u07c7\xc5l\x92\x11\x95\x87\xfc$\x84a!\xae\x82\v\x04\xce\xc6Hp\xf5\tU\xe1n\x88\x13\x10\xe3d\x12\x02Y\xb0\x0e$J\xb8\x84\xab\x85\xb5\x94wv?\xa5`s\xb6\xe6N/\x95\x80\xc6j\x8b\xfb\xf9rC\xbfl\x96l\xb2Y uA\x8d\x94\xc1i\x855.\xbd\x9e0'\a\x9eR\x9a\u36a5\xb9\xc7J\xb35oB\xdd \x9c\xbeJpV\xea\xb9hH\x14\u0308\x9f\xfa=g0\r\xcb\x1b\f\x91S\xd3b\x14b\x9f\x18\x965\xfb\x9c\xc6\x1b\xf1\xa6{&%\xce\xc7\u0275\x05'\u06cb\xd5\xc4\xe2\xea\f\x1d\xb4\xadU\xb7\xeeZ\xc8\xc3I\xdb\x1a\xc1i\t\x86Q\x02\xfcd\xdai0\v\xe2\xb6\xe5\xa9\xdc\xc0\x93kq\xce\xe2UsV95w\x99\xc1\x15\x11\x89\x90j\x13\xcb0\xa7\xca\x1eM_\xa5\xc43\xfd\xebn3\xf5\xc3\xf9\x18`h&\xa7\u04e2\xd1r\xfeC-\xec\xf9\u0412\u0448xj\xc9\xcalZ<\xe3cE\xc4n\xe8\x1f\xa0!\x1b\x8b\x1d\fz7\xf3\xec\xf2\xa9\x84N_\x9fN\x04\xe4v\xb3H\xbcj\xe1\x16\xb51\xff\u02a6\xa8\x19\u0312\x10\x17\xd9\xde\xd2b\x9f\u00f3q\xa9C6\xd3\u0533\x97\xe8\xfe\x8a\xa64\xce\xd1\x14\xf7\xdd5%\xcd\\8\x9f6\xfc\xec\xd2\x04\xbb\xb8\xc0(\xcd\u0476\u0145\n9t\xa1$\xa4^_\xc0\xbe\x9f\xc9\x02\xbd\x10\x94E\xaa\x12b\xe9\v\xe8TW\xb2?\xeb[9\x8f\xbc\xc9\x05\xa1\u06a1\x01\x12\x14:\x9d\n\x05!\u06ea\xb0\xf3\x04fN]\xdf\x1e\xa7\x8c\x8a>S\xaa+\x93\xfcwH\xf6\t\x1b\xa4\u057b\x92\xdao\xa1p\xdaM\nGd\xb8\xf9\xa0\u0666\xea\x1a\x06\xfd\xb7\xf8\xe0\xd2|\xa5\x8dg&\xa8y\x96\u036f\xa95}m+X}=\x15N\xb3\x01k\xe2\xeaf\x8b\xeb\xee\u049dg~\xabM\xe0\x96\x81\xd2/h\xdcTxlv<\x1d\xf0\x1d\x88>*\xf7y\x04\x89x\xa1U\x94_\xb98\x04\x9d[\x9aq)\xab\u007f\xeb\x14\xd4Z\x12\xef\xf3<;j\xcen,q\xf6\xf1\xe2^\xdd\u066e\xc1\xd7\xee\xf1\xae\xb6\x17K\xd4\xd6\xee!\xd9l\xe1\x8f(>\xbc\a\xb27\xe1z3f\xf9\x8aL`\x96\x0f\xb6Ef\xf0\x90\xfb\x93b\xc8d\xde\x12rGL[:\x8bw\xc8\bQ\xffp\x1a\xd69\x8f\x16\xff\xf4\x94\x9b\xf9N. \xb5\xbd$l\x9cLX\xe0g\x9bHI\\?\x99\xb8\xa4\x98\xde\xf1J\xe2\xec\x16gb\xc0f\x92s|\x1e\xe5\xb2E\x96\x9eb\xf1\xeb\xb8\xec\x14\u029b]f\xcc\xde\u04fa\xec\x14\x81z\x05\x97\xa9\xc3\xe3]Vx\xca6%\x1e9T\x02w\xa7\u007f]A\xfe\u06b3\x94]\xff\aPK\a\b!Z\xa2\x84,\x06\x00\x00\xdb\x1d\x00\x00PK\x03\x04\x14\x00\b\b\b\x00\x13b\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x13\x00\x00\x00[Content_Types].xml\xb5\x93Mn\xc20\x10\x85O\xd0;D\xdeV\xc4\xd0EUU\x04\x16\xfdY\xb6]\xd0\x03\f\xce\x04\xac\xfaO\x9e\x81\xc2\xed;\t\x90\x05\x02\xa9\x95\x9a\x8de\xfb\u037c\xf7y$O\xe7;\xef\x8a-f\xb21TjR\x8eU\x81\xc1\xc4\u0686U\xa5>\x17\xaf\xa3\aU\x10C\xa8\xc1\u0140\x95\xda#\xa9\xf9\xecf\xba\xd8'\xa4B\x9a\x03Uj\u035c\x1e\xb5&\xb3F\x0fT\u0184A\x94&f\x0f,\u01fc\xd2\t\xcc\x17\xacP\u07cd\xc7\xf7\xda\xc4\xc0\x18x\u012d\x87\x9aM\x9f\xb1\x81\x8d\xe3\xe2\xe9p\xdfZW\nRr\xd6\x00\v\x97\x163U\xbc\xecD<`\xb6g\xfd\x8b\xbem\xa8\xcf`FG\x902\xa3\xebjhm\x13\u075e\a\x88Jm\u00bbL&\xdb\x1a\xff\x14\x11\x9b\xc6\x1a\xac\xa3\xd9xi)\xbfc\xaeS\x8e\x06\x89d\xa8\u0795\x84\u0332;\xa6~@\xe67\xf0b\xab\xdbJ}R\xcb\xe3#\x87A\xe0\xbd\xc3k\x00\x9d6h|#^\vX:\xbcL\xd0\u02c3B\x84\x8d_b\x96\xfde\x88^\x1e\x14\xa2W<\xd8p\x19\xa4/\xf9G\x0e\x96\x8fze\xf8\x9dtX'\xa7H\xdd\xfd\xf6\xd9\x0fPK\a\b3\xaf\x0f\xb7,\x01\x00\x00-\x04\x00\x00PK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96PI\x13C\u007fh\x01\x00\x00=\x05\x00\x00\x12\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00word/numbering.xmlPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P\x8e\xb3\u00e4\x05\x02\x00\x00\xea\x06\x00\x00\x11\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa8\x01\x00\x00word/settings.xmlPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P\xad\x87m\x00y\x01\x00\x00Z\x05\x00\x00\x12\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xec\x03\x00\x00word/fontTable.xmlPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P\x88\xceE\f\x1d\x03\x00\x00\xdf\x11\x00\x00\x0f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa5\x05\x00\x00word/styles.xmlPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P@Z0\xeb\x17\x02\x00\x00\x16\a\x00\x00\x11\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\b\x00\x00word/document.xmlPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P\x90\x00\xab\xeb\xf1\x00\x00\x00,\x03\x00\x00\x1c\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00U\v\x00\x00word/_rels/document.xml.relsPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P-h\xcf\"\xb1\x00\x00\x00*\x01\x00\x00\v\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x90\f\x00\x00_rels/.relsPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P!Z\xa2\x84,\x06\x00\x00\xdb\x1d\x00\x00\x15\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00z\r\x00\x00word/theme/theme1.xmlPK\x01\x02\x14\x00\x14\x00\b\b\b\x00\x13b\x96P3\xaf\x0f\xb7,\x01\x00\x00-\x04\x00\x00\x13\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xe9\x13\x00\x00[Content_Types].xmlPK\x05\x06\x00\x00\x00\x00\t\x00\t\x00B\x02\x00\x00V\x15\x00\x00\x00\x00"),
	"elf":  []byte("\u007fELF\x01\x01\x01\x00\x00\xb3*1\xc0@\u0340\x02\x00\x03\x00\x01\x00\x00\x00\t\x00 \x00 \x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00 \x00\x01\x00\x00\x00@\x00\x00\x00@\x00\x00\x00\x05\x00\x00\x00\x00\x10\x00\x00"),
	"jpg":  []byte("\xff\xd8\xff\xe0\x00\x10JFIF\x00\x01\x01\x01\x00`\x00`\x00\x00\xff\xdb\x00C\x00\b\x06\x06\a\x06\x05\b\a\a\a\t\t\b\n\f\x14\r\f\v\v\f\x19\x12\x13\x0f\x14\x1d\x1a\x1f\x1e\x1d\x1a\x1c\x1c $.' \",#\x1c\x1c(7),01444\x1f'9=82<.342\xff\xdb\x00C\x01\t\t\t\f\v\f\x18\r\r\x182!\x1c!22222222222222222222222222222222222222222222222222\xff\xc0\x00\x11\b\x00\x01\x00\x01\x03\x01\"\x00\x02\x11\x01\x03\x11\x01\xff\xc4\x00\x15\x00\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\a\xff\xc4\x00\x14\x10\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xc4\x00\x14\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xc4\x00\x14\x11\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xda\x00\f\x03\x01\x00\x02\x11\x03\x11\x00?\x00\xbf\x80\x0f\xff\xd9"),
	"gif":  []byte("GIF89a\x01\x00\x01\x00\x80\x00\x00\x00\x00\x00\xff\xff\xff!\xf9\x04\x01\x00\x00\x00\x00,\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x01D\x00;"),
	"png":  []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x01\x00\x00\x00\x01\b\x04\x00\x00\x00\xb5\x1c\f\x02\x00\x00\x00\vIDATx\xdacd`\x00\x00\x00\x06\x00\x020\x81\xd0/\x00\x00\x00\x00IEND\xaeB`\x82"),
}

func TestGetMimeType(t *testing.T) {
	tests := []struct {
		extension string
		expected  string
	}{
		{"docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
		{"elf", "application/x-executable"},
		{"jpg", "image/jpeg"},
		{"gif", "image/gif"},
		{"png", "image/png"},
	}

	dir := t.TempDir()

	for extension, sample := range mimeSamples {
		samplePath := filepath.Join(dir, "sample."+extension)
		if err := os.WriteFile(samplePath, sample, 0o700); err != nil {
			t.Fatal(err)
		}
	}

	for _, test := range tests {
		t.Run(test.extension, func(t *testing.T) {
			samplePath := filepath.Join(dir, "sample."+test.extension)
			assert.Equal(t, test.expected, getMimeType(samplePath))
		})
	}
}
