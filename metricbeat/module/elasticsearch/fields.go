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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded zlib format compressed contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "eJzsXV+P3Dhyf/enIPx0B6wF3KsfcgEud4kD7GKR9eUlCHQcqbqbtiTKFNU7k08fiJTUpMS/EtXT9o2fdmdGv/pV8V+xWCx+QF/h5SOCCnecFB1gVlzeIcQJr+Ajev9X9efv3yFUQlcw0nJCm4/oX94hhJD2N6imZV/BO4QYVIA7+IjO+B1CHXBOmnP3Ef3P+66r3v+E3l84b9//7/C7C2U8L2hzIueP6ISrbvj+RKAqu49CxAfU4Bo+ItKU8JwzKOgV2Iv4FUL8pR2kMNq340/UT9XPuwtmZZd1HDOec1JDTpq8JlVFuvlvJzxcEaz+tMX8srBTJuhkEx0FN6s7u3DaHiJ7hJ1Ez2I5Lr7mHce8i7YXbuvsRPum3MSwqPqOA8uE7EzwyNaIk6zndvh9UbAMGvxUQTqZduS1bHzFpBr+6ADpOvYkuyIFNB3E92WOeb+t6+g0RwLZAnCSM8AmlDLDacMiWvuWkRrPM0AcMSExWyKodt2msMTVv9cmrh3jfEBZgTa03MZ0+DAj63GgtsUW3Zu+fgKmNe/YCzZOQKQpSQHrXq5+a/peY0D7hmu/sSlmU07vySOnjFOOK6PEcaZf/0EawSO8rpfaJxLYKzl5wStbt8Uk9cu1NkpbMrexV7Fq/Jz3rXWN9asTo9KXa53dBKoL/4oW1NkFcJv3HZQDs6cXDgcTg5qyFyE1G6RmZpErhoNCdydY42eFn2kCiV8lhaT8grtLmgWdQ2aAnKRdgXWENslELfFuHVyYZPP8b5JlwtTWxLzvSSKnjEtvYwGZ3LNRgGZvhtTQcVy372zQEvb9v85/+d7YHRXmNgwztfGzWVnaswJUs4d37s0NYlv/NS8jGnD+ep7T6ZNc97PhvyLtVVfDV0tz3SBPlEGBO96N/68uWFES7ED6rnOzB6M5foEuntn4uvdLdi70HacMso78H9jm+rgFX6oxc8t8+BOPkhYmz2CveAvsrD2ca2j4EZId0JN0BicG3UV2Nns4YD+XYEE3H4CdJ+82P65zBIrRhhFpzsl8RDmmTV6uW7dQ/SbCmUvQgkx6h9XHyi1xXjQvjHJewV0Z+oTO5BaWjZ8Hv/XAXvICFxcY/dHk/V6QzKIETewE8xJzfCy3CDG3OexbDx2/h+UiRd1lLpPMIuexg+f9yVqRc/4xnoAkE+oFfCczvFTqkWZ3M6NHmdkX7EIE3jy08TBqzWNLf5Cz7zH9YfyRS8SCxrH21vmEmFp6rMmMzYHVXT5O1EkjSbqao6MdJm7ePlHS8DuyC5TnClgnZGOGVzZj+RVXPdzRPhEyb3G9u/avMHHaclrmcozcj2Sc2Jt/+Qxl/kR43gG/H9k4seq0kl+h4JTdeXYJlrqIBec1bu/HNEao7p78zggHdj+mUVKV84n7sFsK0sKBRcFy3HOan2hV0d83BgblWWlOT/kJk2oYtxLNduQZFP4uWKYwyyRypiPbDg6XfBjUlEOunb/k40YsKT2nIC/bri8K6LpTXx1hwRE9zITyj4BlswVxud9isyVmc+FSIaD2ym09UZwn51vPdgqWzXqvsgCQtvVVm3mHOYwwczoN4BJYvj3fYpAhQTIdZNnKO2XMRjNLGfU4V/QJV3lxgeKrcCP36mQHXEiu8XPewbe8oXtFGpBWtkyn52xXv6az9AS6zmId2g4fTfMBlLut60SbZNKedxw3JWnOqecjBXo5KdkYiPX+IAoC28JByn3qT6dh0WiBYT64ScutUBwLFTSbQUMY2KJhO+QPkIvsDkM3b9uhFXalKeh93Qy4kjxl/6YTbUVcyRbQkFC0DVA7kZCDTYRediT5DnJvUVw5ygypxYrEcXmEZygOka7gG5Ocb95Y4tnmhuyabO7pealy5+HfCdn7O9kNMFDsMBtwSNC9HZiTWDHP7FPUAKH3Zjm7Je7Cco1w99vUa5TScV1L1NhzU0sfu64qWds2NLSEjfuG05pezLGrPcN0S8B5Tq53RQesqUtjStHQWHWN2cstV9+SMemONOg9OQkhGaCIJXPr3vUiXT6Rta2MXsXUCdjsszOh2ljab2j3cZUJ0YWKjEuH8c9CbDZpOxpNWbJsRxjIG6JT4wW4zNOzFKtpYqpy0k7PVc7aO8imvEchjheSdWxXsoFNm2BLzjcl5twDT+K740/uQSVxJqRBvCcZ8qjkC98pdyB7Q3KD6aJF2ErhzgvZyEgM1y0Ht/GZGLEMNeQ4s61y0JL1j4BssQhdl8lrm9dvJUPvQVVd5xBuVlZLqntQdU2Jf9u94EMSihKN11XKzvYZzp3ptInPpvntsOwiG5Fg1ZypMGFGjkqUiSV2A9/a59GOFI9Ythp+CsKRKQmxfFX4FHTTM0xAKjRDKpaawE1BMDgLL5ahBE5BMTaRK5aphp+CcGR+VCxfFT4V3aN4JiEYl8QVS1NB30J2vhOpXZ2P39ifi4QLc1XJ7nFMzKoqncEPG7YP36CC6Xh4+c/T3l+udXYusptNMlqV2Q3fGc1BAeEnC22vR5qKv9FBDSU/EX+h/SruoGM8eqsKDb7rdl1psKllFXco2YQSUo0iQN2psMWixkSYXis+vvodMYQcVTliGbXACtiyL1oTaovQbZF2jqjlEIYUHnBj0n0B66Ltk/XDiuIyx1dg+LwMlbiBXeCqgD8tx8z0z9N2tMuKts9GfufMiuMbtYWJ/Q5PoO1x4ehFe2zVd/gMeYMbuvGkZTCaIJCNNDMBmVlPbkIG4rq7pdG2OHX5t55ynNekYElUzopTlwnMrN+iMtL2SNh9IpVi+b6le0OF23GyI7TcsQrqBhl+li2wk67jNw1EYZl8isrbHNptGiywk2og3I8Z2jn8tpNXuFsHo4+4HpY5ZkzShjNa5a6+HWyAcesXghk6KCtSE+5yUbYQFKBWXyWGnpzAE9OTU3gsvTkaxWgB3eM4HJuduVERMazi3Th+EWkXLaX7amo89dXXZLb41kNv8ro8llB0yQY+mcDZFfRn8AUKbpy0Y8lMUJtPVc5gzhJ4JQufgT+MgQcuu+27vN7z6haWNxUfxcZT/emdVk5/JrvXzLeDz4ew8/jb3YYWx12PZGeZ3PYoZpZsoq28SLDbmEmd42rfcmsqCGhHsqEhf0aYG9QFvAK3JiigoI2J/S73lMprK3vm2BVFpq8dztSUB3JM/rxxHnZDumBRcDoKSmTDdeGbgHYOTOE5kKOvaPQPMA5NBaoMF9g1iOgpdCqv25iqqH6Fl9+pVtne8IzJ9E9/zmTEFVIyq1RDSDuBTFLaJcoqymnlCkyjVGOl6ZgFiiyXZxtZD+Hh3y+0BPTp34xyFs2fQpLe8qowWTLbKO6J0gpwEyfuU4f4BYSxxX9IfPH/fzYTqGjxVfcd9lOYQNH4WgqizUzrz+v+WKxLOSw7hkPmXxjtug9Th2fQVqQQdx3Q8h6N/pzQ9M/V56wFJ5CzVzgvOd4+rehiUp5jb55b/gEQhooct69Iw+Gs6GN3DsTClsxDWF4+9mqz+thydzgKaHWD0/m1pV/gMofnAlrTZR2J0ojGs3y+upuJ9rle8wXRQxZ5w/Vhkwij1dHK8huB5lRUcR31x9dTKUzx4ylrHBZqNYp0Q8N/L8vdDJaqHPHqKUWhkml34GBw1qvyWsCEGFOlK1jAYgp8QFOYl3hR2ChZPzBFmpHHc0V+X2749wuuAdHTyNgi6ebPGsogOWwTxeRn/Ezqvkbd0GWaAsYT8YHcPEonV3Nku3yOTGfrKhnlIG1s0Kn4xffUpBNnT6Maa7I5LRTHZm7FoeGEMPQ74RciW9LNzVm+JD3DmzjJC0r0h2nLAeUfB8+aCtazaaU+J0br8H4pokkdaQrIx63ABr85SLPPpIafEGlQ3f2EhESd/SAenYAXF1gpkXpYRRH/dyED3WQgcdlpGP666R9/qgrmayu9FUD6FqO1VtGKQLHXw4oAcZcnckKZIln+vAtXDGGMHuyPGJhN4g8VLN8MtathY4AWwbbVm3V+MmgVmVo+EBeH4esfvu+Nj8b5IdTPV8mJwbE8tU/0HfrDmQE0P6EXGEbrT4hB+UdzTG/5fiZyNqUm85fhUyGRiKBqFtTuswvSOWvPWMek8r2vko4XwzyKHdOBpv/ngb4yVwpTDsvTNNItUo1h1HCxt/VdAn2AipzJUwXBBAyVDQKMJc+Q4761Uh9ggvmu30p175PsuyR1ObU/YOpRaZlx6b9psQHOeU/CG1Bd179B4ePaECWXcJvHurcujbfzmc/sQvudye1Zjt7RZAHdUb5HHNUjvWx+E4633b7+VVRUvHGXz3J6OBEeTYReaztL6znM7M/yOILW+Dy6n968/oTVYNnTRX8WwGiYCtCJMkWkcdDrL+vrBML9svl4+rkl7CUvB6fGnisQtGoY7ly4HJt5zOsv8fs+VCQW6XJxcVtnJ9o3dqfOfO54Q3hucfFVFtqf/JUEWONhZTDS3KwNI8X6ceiI7cdfBUKC3Qc8Q9EPG7y8pRUp0lVrNByKo2A3nuNu2XvcbFyMHAkBKqzHL7BQ8kfX5k4sbrnugihwU0Bl672+/qvM6ZhBw/NBpfWBbxwlw2mvCmC/V+K+haXMI4y7bgZ7DhKUo4S+aUhzzhpsDN840bQEVOFCGYeJgYUv1+O2IstsmOlsA42hTCnRMXLFYf/wKZgdj/TM+AXzcQabLrBQ1qELvsLMaQwMiqQS0YR9a1ZiPOBJfrxd9IyZb7Ls9Vn+IpHV8Nzt2HXUZzaW3a3a4CTGBp7ttMLayXasEL1YLQHiV6mCAV5nW3tG64WU5Sq4b58inTuL7T1xx25gg+s9fWqIgCWPfh2T/vZp6CyO/Lf0WecPkIXsPQiPzhn2JEbZizCmUTywIGOEBeBKCled40CYC+HOYiyBMDXpuq04ept+T23wQMbzvx4dCGQtfxzzfQkVmG8DBZ2wGStVoiSdIEVzeWuQRYCFVl+LgIwtmBfJNqK0XQRyaKXECMjg2oYRmFEFRyNwI4ttRiDHVc+LAI4taOqBvq3FwY/zRyLCMwfWDKDJoGtg52mj6V1YAvAe4qJRYGfd63ypiOGF1x2grh3Gdm9463q4KeBf0qIXax9SU7fsu+bti+02elJSNM0w5+REWY35eOfjAG0mfQYeU66cID5oIYTaNTjGN432TIMis+6dQRCE27ENOzd2O7UBk+HyMCvFBHvDdFnpbcvognmgXU/q9fe7cj4O3SclcRX9czoKnbrXi1GtHEQ/vcjM59EijgkdHbv3OmLPsWOsbTauXBfnItmvsic6YEt44NY48XYI3WXXfdjmMO1+9p83Mv22IdwI+s7IMWNQ0CtoBSlf4fgubVWzE6kc/s6eDmKvf4gCUiJUJ0hY3VpnNbBTMBgW+50da1tvn2egda4DSmLq2USJJmBpq1QLkO8JXQfWplNRXbA5Jdojr+O03folZnzbp5ZD5vh0E5l6O81WWjkY5M5G3HHUvBA6ILnOnM25cZsqnnxmPSCySHg1y+74uvr7Dp3/a9JW4FqszHDTVXTZhNtnbPu8GpgOuD2HQo4o2uSil4dj3NL+2Bl4JkzSUsZzXJZsfWs8YGhLoJQ1iT4LSJm9Yx0xUuyFdubCrbsFD8hoNAr6Q0H7qkRPgD79Ov+QMvFHAx/L/aeRZNpEEpWknk5iHme0ZwUkaOgRKGVD/yYg3Q09ik3b0KrgFA09kkzb0CpJe97QFRg5DZuylK6ouEychz0QHDBNRUHojv76sf1X8PPT52a9nUalPI3aUtXGH9MMVmXLQVTgs6nprOyLBK6IBT2id7Rp94aJ7xd/SVEWyF6nNgjqXuE1J7M0D8yFvLTiiJYrC5DTmB6wCcTwAANKYk6l8kOC7uOcCiNwrHdfkrfx5pZZouHrOSmWy4oOuLfUlcPpvTkLb86Cl/8/hbPgf/fMFcnTwbRzmk3Ox1rHNz/mzY+J0ve79GMewPOYc32q7At92hMqqatdcZKUscG/N+RbD6iu0Bf6ZI8OWouFbRL6n/RJQpqlnSiDAnd8fJom5mLy3Ea0BJnilsw9nBIMzcUwQkK7JeZYpvClfPPmiitSyhofO+4hjVexocwZFJSVW7AW7f7rBCkrRcF17eSolslUVYzz29bTwc8XtRyi1K8TFT8REH4BhrBIryTNeWAC0uiIDj8X/y/urMuQdUM5egLUYtZBaTgTWM0WQc9QOBRYfH98gcCwFyZGVHNzrot92Tu4b6r475/Rp+ZEw4pV+bT2aR5AaCJlNABazReyEhhpHK+QHu21/wfgFg0MNEd90MHvo4dWR7uLDjV+3q5CQ5vXb4pfaPMhQXNMurxmi8yqhLfKYq3JDngK5jSBowEbSlFpWBrNODPvr+T621zvDeEn2nMEuLiMZ60NwuZnkva5f5GFAR8q7gXPg7GGhXaKLD1uzOuBIkXfW4xoCr9FhIP2hg43hQaF0ABy80WoQ/b1nibzqo68TXfPBOlkoS2Za50i1LE/7qbnqx99je0eXeFeV/LuocvxNx1SdGrl0sTux2eOvMx2u6HkW0wClA+51Ga8FHRn2eoVnDuLFje37ixT3ry7s1DtPtmdZau3wl5B9L1lKlfVDpasBH0GSSJWkGxvctAOd4K3XJDZP4XeaxQX7jWkpL2t2OtNjCm9ftmoLaVVuh0nrY6xurt+fxLLB18R/hk/h1yqbgF/fRjOvwL+Gko6fyRjC+J1mMU9rzLcl/jfZeDMufd8of1BG7ZHasK38fI2XtKMl65nV3Kl9pcD34bM63F+GzL3Jm4bMqqLdy6yglaV3B2ldPMmWNdR7GvXJ90RDRRL8w+mo7GTnNL1Cvdt3D2G2nPvfok1v/OSog7L6nJdmM4BA/xvpALUvXQcaoeYYOPdLdDHYFfqYYws3wuDSQTOUUSaL0+UdQm7+7eTf4pMy8AkUxSXcyleXD0GWsRsE2DPBkg31xVtf0hPqCguc3w9Z39av6qoyrjg6pSfKorX5pgp2ngkWLjaHhcFz/oOnyHbW0fRbEs/Ux9bTcapy771lOPMmOIZyBgtMme9SC7qIfRVgVDhtoMyb4ERWvoHQ6A+aJHwrdyKOEqEIsHadwLh9RPJ47sRbTijVe5rV18qoo5akdqRuqZjegwdsr1QtzPD+D1Y8DI5rGj7bB3edoS1Q8LZ/DKsQ3lLabprg4fd9pCPdyXw/hl8gYJ7BmqAH3sGe6rTD6WoqcrHD6rqsZkTj6WrcBR/TFXfLYGml29bEE+45Yv3J6PTXP9hAvyHWOkwaTqE0fgL8dKliqRGnbYku3bAeE5ZuXpufet1kE8CEq0hb0sQoYxwc8G2eHm/muDm8SeqDxkl7ahmlKG/UYbgGddtNSjU8w81bttl4p/mbZEml904tIqf/94NqUVupYBd9VBRtG5PlxQAY+fZ1ccOq8/HL6RDpBOZpgG1+mS6byrja5eeBBN3mcCUF/w+i9RazCFENoOKFlg8IixS8pu0ddwuY20vke4ruszvuJuEQolOjNZhxJIW3wuihT5xdMGyA8EzLjjqcA1IJOchfsGN0XjiofGC1i3m5IlUhL+gtmct7WxHAHISyhdXzdAut9jQij6TKTuO9fuT64//PwAA//+L5Sb/"
}
