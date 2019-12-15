package Day201914

var Entry dayEntry

type dayEntry bool

func (td dayEntry) Describe() (int, int, string) {
	return 2019, 14, "Space Stoichiometry"
}

func (td dayEntry) PuzzleInput() string {
	return `3 CFGBR, 9 PFMFC, 2 FQFPN => 2 PKPWN
9 XQHK => 3 KDXDC
9 MPQFZ, 6 SGHLD => 6 DXPTR
6 QDBK, 2 SGHLD => 3 WDJKB
1 RXJS, 9 XQHK => 4 GTXPN
2 XJHR, 3 WNFC => 4 NRLM
3 RHWD => 7 NDQK
3 PZXG, 3 TNCBS, 1 GNSW => 7 CFGBR
1 VRMZK => 2 TVDH
3 JQFP => 8 VRMZK
124 ORE => 6 NKRXN
2 KRDMT, 11 MPQFZ => 2 WDWNX
3 ZPCP => 9 WLMB
2 MPQFZ => 1 DQRQW
13 KHXVX => 8 RHWD
5 ZPVWS => 8 JQFP
1 NDQK, 1 JZQN, 1 GNSW => 6 XHRQW
4 KRDMT => 7 HCVLB
3 NRLM => 9 WHWK
172 ORE => 5 ZPCP
104 ORE => 1 TJHD
1 LFPG => 6 TNCBS
3 XJHR => 9 TVBNZ
3 JQFP => 4 DSJK
3 ZPVWS => 3 SGHLD
15 NRLM, 5 KDXDC, 1 DQRQW, 5 WDWNX, 12 RXJS, 3 GTXPN => 5 QTSK
1 WDWNX => 1 GNSW
1 QDBK => 1 LBTRH
3 FQFPN, 13 WDWNX => 4 RXJS
1 QDBK => 7 MPQFZ
6 LBTRH, 6 TVDH => 6 JDKMB
4 KWXF, 8 XJHR => 9 JZQN
8 MPQFZ, 8 VRMZK => 7 WNFC
16 QGZSZ, 9 XHRQW, 17 MRBFL, 10 WHWK, 36 JDKMB, 82 LFNZ, 11 TDRWG, 7 QTSK, 7 MNWVT, 6 CDNHC, 3 NDQK, 4 TNCBS => 1 FUEL
1 DQRQW, 1 MRBFL, 1 GTXPN, 1 CFGBR, 2 HCVLB, 1 DGXBN, 3 GZQSX => 8 QGZSZ
13 SGHLD, 11 XQHK, 17 PKPWN, 1 RXJS, 1 FQFPN, 11 JZQN => 1 CDNHC
21 NKRXN, 9 TJHD, 2 ZXJCJ => 5 KHXVX
2 WLMB => 8 XJKTS
2 WDJKB => 6 KRDMT
2 MGXB, 1 KWXF => 8 LFNZ
1 TVBNZ, 5 VRMZK => 8 CSDWQ
7 LFPG => 8 TDRWG
1 RHWD, 8 XJKTS => 2 QDBK
182 ORE => 4 ZXJCJ
3 ZXJCJ => 8 ZPVWS
1 WNFC, 2 CSDWQ, 2 NRLM => 6 GZQSX
4 TVDH, 2 DGXBN => 6 MRBFL
3 DSJK => 4 FQFPN
9 NDQK, 7 WLMB => 2 KWXF
4 CSDWQ => 2 XQHK
1 NKRXN => 5 PZXG
2 LFPG => 2 DGXBN
7 MGXB => 7 XJHR
2 WLMB => 7 LFPG
8 DXPTR, 7 WNFC, 5 MPQFZ => 9 PFMFC
5 PFMFC, 4 NRLM => 9 MNWVT
7 ZPVWS, 14 ZPCP, 11 TJHD => 2 MGXB`
}
